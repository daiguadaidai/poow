package views

import (
	"github.com/daiguadaidai/poow/pili/controllers"
	"github.com/daiguadaidai/poow/pili/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(UtilHandler)
	AddHandlerV1("/pili/utils", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *UtilHandler) RegisterV1(group *gin.RouterGroup) {
	group.GET("/encrypt", this.Encrypt)
	group.GET("/decrypt", this.Decrypt)
}

type UtilHandler struct{}

// 加密
func (this *UtilHandler) Encrypt(c *gin.Context) {
	f, err := form.NewForm(c, &form.UtilEncreptForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewUtilController().Encrypt(f.(*form.UtilEncreptForm))
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	returnSuccess(c, data)
}

// 解密
func (this *UtilHandler) Decrypt(c *gin.Context) {
	f, err := form.NewForm(c, &form.UtilDecryptForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewUtilController().Decrypt(f.(*form.UtilDecryptForm))
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	returnSuccess(c, data)
}
