package handlers

import (
	"github.com/artur-karunas/pop-up-museum/internal/services"
	"github.com/artur-karunas/pop-up-museum/pkg/emailhandling"
	"github.com/artur-karunas/pop-up-museum/pkg/imagehandling"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services          *services.Service
	imageHandler      *imagehandling.ImageHandler
	emailHandler      *emailhandling.EmailService
	passUpdateSubject string
	passUpdateMessage string
}

func NewHandler(services *services.Service, imageHandler *imagehandling.ImageHandler, emailHandler *emailhandling.EmailService, passUpdateSubject, passUpdateMessage string) *Handler {
	return &Handler{
		services:          services,
		imageHandler:      imageHandler,
		emailHandler:      emailHandler,
		passUpdateSubject: passUpdateSubject,
		passUpdateMessage: passUpdateMessage,
	}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()

	// Public: accessable for everyone, token is not necessary.
	// User: accessable only for users, needs token with role = 0.
	// Admin: accessable only for admins, needs token with role = 1.
	// Moderator: accessable only for moderators, needs token with role = 2.

	router.MaxMultipartMemory = 1 << 20

	router.Static("/uploads", "./uploads")

	auth := router.Group("")
	{
		auth.POST("/sign-up", h.signup)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/ws", h.wsRestorePassword)
	}

	exhibit := router.Group("/exhibit")
	{
		public := exhibit.Group("/")
		{
			public.GET("/all", h.getAllExhibits)
			public.GET("/:exhibitId", h.getExhibitById)
			public.GET("/statuses", h.getExhibitStatuses)
		}
		admin := exhibit.Group("/", h.adminIdentity)
		{
			admin.POST("/all", h.createExhibit)
			admin.POST("/:exhibitId/upload-image", h.updateExhibitImage)
			admin.PUT("/:exhibitId", h.updateExhibit)
			admin.DELETE("/:exhibitId", h.deleteExhibit)
		}
	}

	item := router.Group("/item")
	{
		public := item.Group("/")
		{
			public.GET("/:itemId", h.getItemById)
			public.GET("/all", h.getAllItems)
			public.GET("/statuses", h.getItemStatuses)
		}
		admin := item.Group("/", h.adminIdentity)
		{
			admin.POST("/all", h.createItem)
			admin.POST("/:itemId/upload-image", h.updateItemImage)
			admin.PUT("/:itemId", h.updateItem)
			admin.DELETE("/:itemId", h.deleteItem)
		}
	}

	author := router.Group("/author")
	{
		public := author.Group("/")
		{
			public.GET("/:authorId", h.getAuthorById)
			public.GET("/all", h.getAllAuthors)
		}
		admin := author.Group("/", h.adminIdentity)
		{
			admin.POST("/all", h.createAuthor)
			admin.POST("/:authorId/upload-image", h.updateAuthorImage)
			admin.PUT("/:authorId", h.updateAuthor)
			admin.DELETE("/:authorId", h.deleteAuthor)
		}
	}

	profile := router.Group("/profile")
	{
		user := profile.Group("/", h.userIdentity)
		{
			user.GET("/", h.getUserById)
			user.POST("/upload-image", h.updateUserImage)
			user.DELETE("/", h.deleteUser)
		}
	}

	collection := router.Group("/collection")
	{
		public := collection.Group("/")
		{
			public.GET("/all", h.getAllCollection)
		}
		user := collection.Group("/", h.userIdentity)
		{
			user.GET("/", h.getCollectionById)
			user.POST("/", h.addToUserCollection)
			user.DELETE("/", h.deleteFromUserCollection)
		}
	}

	appeal := router.Group("/appeal")
	{
		public := appeal.Group("/", h.optionalIdentity)
		{
			public.POST("/", h.createAppeal)
			public.GET("/statuses", h.getAppealStatuses)
		}
		admin := appeal.Group("/", h.moderatorIdentity)
		{
			admin.GET("/all", h.getAllAppeals)
			admin.PUT("/:appealId", h.confirmAppeal)
		}
	}

	reservation := router.Group("/reservation")
	{
		public := reservation.Group("/", h.optionalIdentity)
		{
			public.POST("/", h.createReservation)
			public.GET("/statuses", h.getReservationStatuses)
		}
		admin := reservation.Group("/", h.moderatorIdentity)
		{
			admin.GET("/all", h.getAllReservations)
			admin.PUT("/:reservationId", h.confirmReservation)
		}
	}

	info := router.Group("")
	{
		info.GET("/info", h.getInfo)
		info.GET("/faq", h.getFAQ)
	}

	return router
}
