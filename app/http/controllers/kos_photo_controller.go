package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"
	"errors"

	"github.com/goravel/framework/contracts/http"
)

type KosPhotoController struct {
	service serviceContracts.KosPhotoService
}

func NewKosPhotoController(service serviceContracts.KosPhotoService) *KosPhotoController {
	return &KosPhotoController{service: service}
}

func (c *KosPhotoController) Index(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	photos, err := c.service.GetAllByKosID(id)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve kos photos"})
	}
	return ctx.Response().Success().Json(http.Json{"data": photos})
}

func (c *KosPhotoController) Store(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	file, err := ctx.Request().File("photo")
	if err != nil {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Photo file is required"})
	}
	if err := validatePhotoFile(file); err != nil {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": err.Error()})
	}
	photo, err := c.service.Create(id, dto.CreateKosPhotoRequest{File: file})
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": photo})
}

func (c *KosPhotoController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	file, err := ctx.Request().File("photo")
	if err != nil {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Photo file is required"})
	}
	if err := validatePhotoFile(file); err != nil {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": err.Error()})
	}
	photo, err := c.service.Update(id, ctx.Request().Route("photo_id"), dto.UpdateKosPhotoRequest{File: file})
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": photo})
}

func (c *KosPhotoController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	if err := c.service.Delete(id, ctx.Request().Route("photo_id")); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Kos photo deleted successfully"})
}

func validatePhotoFile(file interface {
	Size() (int64, error)
	MimeType() (string, error)
}) error {
	size, err := file.Size()
	if err != nil {
		return err
	}
	if size <= 0 || size > 10*1024*1024 {
		return errors.New("Photo must be between 1 byte and 10 MB")
	}
	mimeType, err := file.MimeType()
	if err != nil {
		return err
	}
	if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/webp" {
		return errors.New("Photo must be a JPEG, PNG, or WebP image")
	}
	return nil
}
