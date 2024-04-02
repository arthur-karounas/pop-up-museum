package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(ctx *gin.Context) {
	var input entities.CreateUser

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	userId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	_, token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.Writer.Header().Set("Authorization", token)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId": userId,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input entities.LoginUser

	if err := ctx.ShouldBind(&input); err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	userId, token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.Writer.Header().Set("Authorization", token)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId": userId,
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Request struct {
	UUID      string `json:"uuid"`
	EventType string `json:"event_type"`
	Data      string `json:"data"`
}

type Response struct {
	UUID      string `json:"uuid"`
	EventType string `json:"event_type"`
}

type RestoreSession struct {
	uuid           string
	email          string
	nextEventType  string
	recoveryCode   string
	attemptsUsed   int
	sessionTimeout int64
}

func (h *Handler) wsRestorePassword(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	var session RestoreSession = RestoreSession{
		uuid:           "",
		email:          "",
		nextEventType:  "START_RECOVERY",
		recoveryCode:   "",
		attemptsUsed:   0,
		sessionTimeout: time.Now().Unix() + 5*60,
	}

	go func() {
		for {
			// Checks the session lifetime and closes the connection when it expires.
			if time.Now().Unix() > session.sessionTimeout && conn != nil {
				var response Response = Response{
					UUID:      session.uuid,
					EventType: "RECOVERY_TIMEOUT",
				}

				if err := conn.WriteJSON(response); err != nil {
					return
				}

				conn.Close()
				return
			}

			time.Sleep(time.Second)
		}
	}()

	for {
		var request Request
		if err := conn.ReadJSON(&request); err != nil {
			break
		}

		var response Response = Response{
			UUID:      session.uuid,
			EventType: "",
		}

		switch {
		case request.EventType == "START_RECOVERY" && session.nextEventType == "START_RECOVERY":
			// Save the info of the client.
			session.email = request.Data
			session.uuid = request.UUID

			// Prepare response to the client.
			response.UUID = session.uuid
			response.EventType = "RECOVERY_STARTED"

			// Send the response to the client.
			if err := conn.WriteJSON(response); err != nil {
				return
			}

			// Send recovery code via email handler.
			session.recoveryCode = generateCode()

			if err := h.emailHandler.Send(session.email, h.passUpdateSubject, fmt.Sprintf(h.passUpdateMessage, session.recoveryCode)); err != nil {
				fmt.Println(err)
				return
			}

			// Allow user to move to the next event.
			session.nextEventType = "SEND_RECOVERY_CODE"

			// Prepare response to the client.
			response.EventType = "RECOVERY_CODE_SENDED"

			// Send the response to the client.
			if err := conn.WriteJSON(response); err != nil {
				return
			}

		case request.EventType == "SEND_RECOVERY_CODE" && session.nextEventType == "SEND_RECOVERY_CODE":
			// Validate the recovery code that was received from the client.
			if request.Data == session.recoveryCode {
				response.EventType = "RECOVERY_CODE_SUCCESS"
				session.nextEventType = "SEND_NEW_PASSWORD"
			} else {
				response.EventType = "RECOVERY_CODE_FAILURE"
				session.attemptsUsed += 1
			}

			if session.attemptsUsed >= 3 {
				response.EventType = "RECOVERY_CODE_NO_ATTEMPTS"

				if err := conn.WriteJSON(response); err != nil {
					return
				}

				conn.Close()
			}

			if err := conn.WriteJSON(response); err != nil {
				return
			}

			// Allow user to move to the next event.
			response.EventType = "RECEIVE_NEW_PASSWORD_STARTED"

			if err := conn.WriteJSON(response); err != nil {
				return
			}

		case request.EventType == "SEND_NEW_PASSWORD" && session.nextEventType == "SEND_NEW_PASSWORD":
			err := h.services.Authorization.ChangePassword(session.email, request.Data)
			if err != nil {
				response.EventType = "RECOVERY_FAILURE"

				if err := conn.WriteJSON(response); err != nil {
					return
				}

				conn.Close()
			}

			response.EventType = "RECOVERY_SUCCESS"

			if err := conn.WriteJSON(response); err != nil {
				return
			}

			conn.Close()
		}
	}
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(999999) + 1
	formattedNumber := fmt.Sprintf("%06d", randomNumber) // Ensures 6 digits with leading zeros.

	return formattedNumber
}
