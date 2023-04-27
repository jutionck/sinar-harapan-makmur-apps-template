package response

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"net/http"
)

func SendSingleResponse(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, &SingleResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendPagedResponse(c *gin.Context, data []interface{}, description string, paging dto.Paging) {
	c.JSON(http.StatusOK, &PagedResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendErrorResponse(c *gin.Context, code int, description string) {
	c.AbortWithStatusJSON(code, &Status{
		Code:        code,
		Description: description,
	})
}

func SendFileResponse(c *gin.Context, fileName string, description string) {
	c.JSON(http.StatusOK, &FileResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		FileName: fileName,
	})
}
