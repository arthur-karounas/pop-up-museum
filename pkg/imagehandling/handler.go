package imagehandling

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageHandler struct {
	directoryPath string
}

func NewImageHandler(directoryPath string) *ImageHandler {
	return &ImageHandler{
		directoryPath: directoryPath,
	}
}

type ContextImageHandler struct {
	directoryPath string
	ctx           *gin.Context
}

func (handler *ImageHandler) NewContextImageHandler(ctx *gin.Context) *ContextImageHandler {
	return &ContextImageHandler{
		directoryPath: handler.directoryPath,
		ctx:           ctx,
	}
}

type Input struct {
	Image *multipart.FileHeader `form:"file"`
}

func (handler *ContextImageHandler) Delete(pathToImage string) (string, error) {
	if pathToImage == "" {
		return "", nil
	}

	if _, err := os.Stat(pathToImage); os.IsNotExist(err) {
		return "", errors.New("file do not exist")
	}

	if err := os.Remove(pathToImage); err != nil {
		return "", errors.New("error occured while removing file")
	}

	return "", nil
}

func (handler *ContextImageHandler) Update(pathToImage string) (string, error) {
	var input Input
	if err := handler.ctx.ShouldBind(&input); err != nil {
		return "", err
	}

	if input.Image == nil {
		return handler.Delete(pathToImage)
	}

	if input.Image.Size > 1*1024*1024 {
		return "", errors.New("too large image")
	}

	newPathToImage := generateImagePath(handler.directoryPath)
	if err := saveFile(handler.ctx, input.Image, newPathToImage); err != nil {
		return "", err
	}

	_, err := handler.Delete(pathToImage)
	if err != nil {
		_, err = handler.Delete(newPathToImage)
		return pathToImage, err
	}

	return newPathToImage, nil
}

func generateImagePath(directoryPath string) string {
	imageName := uuid.New().String() + ".jpg"
	pathToImage := fmt.Sprintf("%s%s", directoryPath, imageName)
	return pathToImage
}

func saveFile(ctx *gin.Context, image *multipart.FileHeader, pathToImage string) error {
	if err := ctx.SaveUploadedFile(image, pathToImage); err != nil {
		return err
	}

	return nil
}
