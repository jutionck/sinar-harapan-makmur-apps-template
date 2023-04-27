package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api/response"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/mitchellh/mapstructure"
)

type BaseApi struct{}

func (b *BaseApi) ParseRequestBody(c *gin.Context, payload interface{}) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) ParseRequestFormData(c *gin.Context, requestModel interface{}, postFormKey ...string) (map[string]interface{}, error) {
	mapRes := make(map[string]interface{})
	for _, v := range postFormKey {
		mapRes[v] = c.PostForm(v)
	}
	err := mapstructure.Decode(mapRes, &requestModel)
	if err != nil {
		return nil, err
	}
	return mapRes, nil
}

func (b *BaseApi) NewSuccessSingleResponse(c *gin.Context, data interface{}, description string) {
	response.SendSingleResponse(c, data, description)
}

func (b *BaseApi) NewSuccessPagedResponse(c *gin.Context, data []interface{}, description string, paging dto.Paging) {
	response.SendPagedResponse(c, data, description, paging)
}

func (b *BaseApi) NewFailedResponse(c *gin.Context, code int, description string) {
	response.SendErrorResponse(c, code, description)
}

func (b *BaseApi) NewSuccessFileResponse(c *gin.Context, fileName string, description string) {
	response.SendFileResponse(c, fileName, description)
}
