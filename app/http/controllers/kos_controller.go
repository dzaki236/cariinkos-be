package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	serviceContracts "cariinkos/app/services/contracts"
	"math"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

type KosController struct {
	service serviceContracts.KosService
}

func NewKosController(service serviceContracts.KosService) *KosController {
	return &KosController{service: service}
}

func (c *KosController) Index(ctx http.Context) http.Response {
	return c.getAllData(ctx, false)
}

func (c *KosController) PublicIndex(ctx http.Context) http.Response {
	return c.getAllData(ctx, true)
}

func (c *KosController) getAllData(ctx http.Context, publishedOnly bool) http.Response {
	page, err := strconv.Atoi(ctx.Request().Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err := strconv.Atoi(ctx.Request().Query("per_page", "10"))
	if err != nil || perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}
	var kos []models.Kos
	var total int64
	if publishedOnly {
		kos, total, err = c.service.GetPublishedData(page, perPage)
	} else {
		kos, total, err = c.service.GetAllData(page, perPage)
	}
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve kos"})
	}
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	return ctx.Response().Success().Json(http.Json{
		"data": kos,
		"meta": http.Json{"page": page, "per_page": perPage, "total": total, "last_page": lastPage},
	})
}

func (c *KosController) Show(ctx http.Context) http.Response {
	return c.show(ctx, false)
}

func (c *KosController) PublicShow(ctx http.Context) http.Response {
	return c.show(ctx, true)
}

func (c *KosController) show(ctx http.Context, publishedOnly bool) http.Response {
	id := ctx.Request().Route("id")
	var kos models.Kos
	var err error
	if publishedOnly {
		kos, err = c.service.GetPublishedByID(id)
	} else {
		kos, err = c.service.GetByID(id)
	}
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": kos})
}

func (c *KosController) Store(ctx http.Context) http.Response {
	var request dto.CreateKosRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{
		"name":          "required|string|max:255",
		"address":       "required|string",
		"province":      "required|string|max:255",
		"city":          "required|string|max:255",
		"district":      "required|string|max:255",
		"price_monthly": "required|numeric|min:1",
		"type":          "in:mixed,man,woman",
		"status":        "in:draft,published,archived",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}

	kos, err := c.service.Create(userID, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": kos})
}

func (c *KosController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var request dto.UpdateKosRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusNotFound).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{
		"name":          "required|string|max:255",
		"address":       "required|string",
		"province":      "required|string|max:255",
		"city":          "required|string|max:255",
		"district":      "required|string|max:255",
		"price_monthly": "required|numeric|min:1",
		"type":          "required|in:mixed,man,woman",
		"status":        "required|in:draft,published",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	kos, err := c.service.Update(id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": kos})
}

func (c *KosController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	if err := c.service.Delete(id); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Kos deleted successfully"})
}
