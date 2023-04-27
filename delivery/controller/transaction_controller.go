package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"net/http"
)

type TransactionController struct {
	router  *gin.Engine
	useCase usecase.TransactionUseCase
	api.BaseApi
}

func (b *TransactionController) createHandler(c *gin.Context) {
	var payload model.Transaction
	if err := b.ParseRequestBody(c, &payload); err != nil {
		b.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := b.useCase.RegisterNewTransaction(&payload); err != nil {
		b.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	b.NewSuccessSingleResponse(c, payload, "OK")
}

func (b *TransactionController) listHandler(c *gin.Context) {
	brands, err := b.useCase.FindAllTransaction()
	if err != nil {
		b.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var brandInterface []interface{}
	for _, b := range brands {
		brandInterface = append(brandInterface, b)
	}
	b.NewSuccessPagedResponse(c, brandInterface, "OK", dto.Paging{
		Page:        0,
		RowsPerPage: 0,
		TotalRows:   0,
		TotalPages:  0,
	})
}

func (b *TransactionController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := b.useCase.FindByTransaction(id)
	if err != nil {
		b.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	b.NewSuccessSingleResponse(c, brand, "OK")
}

func NewTransactionController(r *gin.Engine, useCase usecase.TransactionUseCase) *TransactionController {
	controller := &TransactionController{
		router:  r,
		useCase: useCase,
	}
	r.GET("/transactions", controller.listHandler)
	r.GET("/transactions/:id", controller.getHandler)
	r.POST("/transactions", controller.createHandler)
	return controller
}
