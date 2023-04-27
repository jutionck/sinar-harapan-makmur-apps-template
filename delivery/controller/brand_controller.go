package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"net/http"
	"strconv"
)

type BrandController struct {
	router  *gin.Engine
	useCase usecase.BrandUseCase
	api.BaseApi
}

func (b *BrandController) createHandler(c *gin.Context) {
	var payload model.Brand
	if err := b.ParseRequestBody(c, &payload); err != nil {
		b.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := b.useCase.SaveData(&payload); err != nil {
		b.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	b.NewSuccessSingleResponse(c, payload, "OK")
}

func (b *BrandController) listHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		b.NewFailedResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		b.NewFailedResponse(c, http.StatusBadRequest, "invalid limit number")
		return
	}
	order := c.DefaultQuery("order", "id")
	sort := c.DefaultQuery("sort", "asc")
	requestQueryParams := dto.RequestQueryParams{
		QueryParams: dto.QueryParams{
			Sort:  sort,
			Order: order,
		},
		PaginationParam: dto.PaginationParam{
			Page:  page,
			Limit: limit,
		},
	}
	brands, paging, err := b.useCase.Pagination(requestQueryParams)
	if err != nil {
		b.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var brandInterface []interface{}
	for _, b := range brands {
		brandInterface = append(brandInterface, b)
	}
	b.NewSuccessPagedResponse(c, brandInterface, "OK", paging)
}

func (b *BrandController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := b.useCase.FindById(id)
	if err != nil {
		b.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	b.NewSuccessSingleResponse(c, brand, "OK")
}

func (b *BrandController) searchHandler(c *gin.Context) {
	//name := c.Query("name")
	name := c.DefaultQuery("name", "Honda") // memberikan default query -> Honda (case sensitive)
	filter := map[string]interface{}{"name": name}
	brands, err := b.useCase.SearchBy(filter)
	if err != nil {
		b.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	b.NewSuccessSingleResponse(c, brands, "OK")
}

func (b *BrandController) updateHandler(c *gin.Context) {
	var payload model.Brand
	if err := b.ParseRequestBody(c, &payload); err != nil {
		b.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := b.useCase.SaveData(&payload); err != nil {
		b.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	b.NewSuccessSingleResponse(c, payload, "OK")
}

func (b *BrandController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := b.useCase.DeleteData(id)
	if err != nil {
		b.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewBrandController(r *gin.Engine, useCase usecase.BrandUseCase) *BrandController {
	controller := &BrandController{
		router:  r,
		useCase: useCase,
	}
	r.GET("/brands", controller.listHandler)
	r.GET("/brands/:id", controller.getHandler)
	r.GET("/brands/search", controller.searchHandler)
	r.POST("/brands", controller.createHandler)
	r.PUT("/brands", controller.updateHandler)
	r.DELETE("/brands/:id", controller.deleteHandler)
	return controller
}
