package response

import (
	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}

	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)

func Json(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Data: data, Meta: Meta{
			Message: "OK",
			Code:    code,
		}})
}
