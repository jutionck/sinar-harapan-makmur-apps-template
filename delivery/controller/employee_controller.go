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

type EmployeeController struct {
	router  *gin.Engine
	useCase usecase.EmployeeUseCase
	api.BaseApi
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var payload model.Employee
	if err := e.ParseRequestBody(c, &payload); err != nil {
		e.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := e.useCase.SaveData(&payload); err != nil {
		e.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	e.NewSuccessSingleResponse(c, payload, "OK")
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		e.NewFailedResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		e.NewFailedResponse(c, http.StatusBadRequest, "invalid limit number")
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
	brands, paging, err := e.useCase.Pagination(requestQueryParams)
	if err != nil {
		e.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var brandInterface []interface{}
	for _, b := range brands {
		brandInterface = append(brandInterface, b)
	}
	e.NewSuccessPagedResponse(c, brandInterface, "OK", paging)
}

func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := e.useCase.FindById(id)
	if err != nil {
		e.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	e.NewSuccessSingleResponse(c, brand, "OK")
}

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var payload model.Employee
	if err := e.ParseRequestBody(c, &payload); err != nil {
		e.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := e.useCase.SaveData(&payload); err != nil {
		e.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	e.NewSuccessSingleResponse(c, payload, "OK")
}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := e.useCase.DeleteData(id)
	if err != nil {
		e.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewEmployeeController(r *gin.Engine, useCase usecase.EmployeeUseCase) *EmployeeController {
	controller := &EmployeeController{
		router:  r,
		useCase: useCase,
	}
	r.GET("/employees", controller.listHandler)
	r.GET("/employees/:id", controller.getHandler)
	r.POST("/employees", controller.createHandler)
	r.PUT("/employees", controller.updateHandler)
	r.DELETE("/employees/:id", controller.deleteHandler)
	return controller
}
