package form

import (
	"github.com/gin-gonic/gin"
)

func NewForm(c *gin.Context, obj interface{}) (interface{}, error) {
	if err := c.ShouldBind(obj); err != nil {
		return nil, err
	}
	return obj, nil
}
