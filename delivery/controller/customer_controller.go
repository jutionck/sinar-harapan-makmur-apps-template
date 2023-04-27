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

type CustomerController struct {
	router  *gin.Engine
	useCase usecase.CustomerUseCase
	api.BaseApi
}

func (cc *CustomerController) createHandler(c *gin.Context) {
	var payload model.Customer
	if err := cc.ParseRequestBody(c, &payload); err != nil {
		cc.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := cc.useCase.SaveData(&payload); err != nil {
		cc.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	cc.NewSuccessSingleResponse(c, payload, "OK")
}

func (cc *CustomerController) listHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		cc.NewFailedResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		cc.NewFailedResponse(c, http.StatusBadRequest, "invalid limit number")
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
	brands, paging, err := cc.useCase.Pagination(requestQueryParams)
	if err != nil {
		cc.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var brandInterface []interface{}
	for _, b := range brands {
		brandInterface = append(brandInterface, b)
	}
	cc.NewSuccessPagedResponse(c, brandInterface, "OK", paging)
}

func (cc *CustomerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := cc.useCase.FindById(id)
	if err != nil {
		cc.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	cc.NewSuccessSingleResponse(c, brand, "OK")
}

func (cc *CustomerController) searchHandler(c *gin.Context) {
	//name := c.Query("name")
	name := c.DefaultQuery("name", "Honda") // memberikan default query -> Honda (case sensitive)
	filter := map[string]interface{}{"name": name}
	brands, err := cc.useCase.SearchBy(filter)
	if err != nil {
		cc.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	cc.NewSuccessSingleResponse(c, brands, "OK")
}

func (cc *CustomerController) updateHandler(c *gin.Context) {
	var payload model.Customer
	if err := cc.ParseRequestBody(c, &payload); err != nil {
		cc.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := cc.useCase.SaveData(&payload); err != nil {
		cc.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cc.NewSuccessSingleResponse(c, payload, "OK")
}

func (cc *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := cc.useCase.DeleteData(id)
	if err != nil {
		cc.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewCustomerController(r *gin.Engine, useCase usecase.CustomerUseCase) *CustomerController {
	controller := &CustomerController{
		router:  r,
		useCase: useCase,
	}
	r.GET("/customers", controller.listHandler)
	r.GET("/customers/:id", controller.getHandler)
	r.GET("/customers/search", controller.searchHandler)
	r.POST("/customers", controller.createHandler)
	r.PUT("/customers", controller.updateHandler)
	r.DELETE("/customers/:id", controller.deleteHandler)
	return controller
}
