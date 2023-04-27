package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type VehicleController struct {
	router  *gin.Engine
	useCase usecase.VehicleUseCase
	api.BaseApi
}

func (v *VehicleController) createHandler(c *gin.Context) {
	var payload model.Vehicle
	vehicle := c.PostForm("vehicle")
	vehicleImage, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		v.NewFailedResponse(c, http.StatusBadRequest, "Failed get file")
	}
	err = json.Unmarshal([]byte(vehicle), &payload)
	if err != nil {
		log.Fatalln(err)
	}
	// 1-min.png
	// joko.anwar.png
	fileExtension := strings.Split(fileHeader.Filename, ".")
	// .png, jpg
	if fileExtension[1] == "png" || fileExtension[1] == "jpg" {

	}

	if err := v.useCase.UploadImage(&payload, vehicleImage, fileExtension[1]); err != nil {
		v.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	v.NewSuccessSingleResponse(c, payload, "OK")
}

func (v *VehicleController) listHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		v.NewFailedResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		v.NewFailedResponse(c, http.StatusBadRequest, "invalid limit number")
		return
	}
	order := c.DefaultQuery("order", "created_at")
	sort := c.DefaultQuery("sort", "desc")
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
	vehicles, paging, err := v.useCase.Pagination(requestQueryParams)
	if err != nil {
		v.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var vehicleInterface []interface{}
	for _, b := range vehicles {
		vehicleInterface = append(vehicleInterface, b)
	}
	v.NewSuccessPagedResponse(c, vehicleInterface, "OK", paging)
}

func (v *VehicleController) getHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := v.useCase.FindById(id)
	if err != nil {
		v.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	v.NewSuccessSingleResponse(c, vehicle, "OK")
}

func (v *VehicleController) searchHandler(c *gin.Context) {
	name := c.Query("model")
	filter := map[string]interface{}{"model": name}
	vehicles, err := v.useCase.SearchBy(filter)
	if err != nil {
		v.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	v.NewSuccessSingleResponse(c, vehicles, "OK")
}

func (v *VehicleController) updateHandler(c *gin.Context) {
	var payload model.Vehicle
	if err := v.ParseRequestBody(c, &payload); err != nil {
		v.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := v.useCase.SaveData(&payload); err != nil {
		v.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	v.NewSuccessSingleResponse(c, payload, "OK")
}

func (v *VehicleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := v.useCase.DeleteData(id)
	if err != nil {
		v.NewFailedResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewVehicleController(r *gin.Engine, useCase usecase.VehicleUseCase) *VehicleController {
	controller := &VehicleController{
		router:  r,
		useCase: useCase,
	}
	r.GET("/vehicles", controller.listHandler)
	r.GET("/vehicles/:id", controller.getHandler)
	r.GET("/vehicles/search", controller.searchHandler)
	r.POST("/vehicles", controller.createHandler)
	r.PUT("/vehicles", controller.updateHandler)
	r.DELETE("/vehicles/:id", controller.deleteHandler)
	return controller
}
