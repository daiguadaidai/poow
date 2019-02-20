package controllers

import (
	"github.com/daiguadaidai/peep"
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/views/form"
)

type UtilController struct {
	SC *config.ServerConfig
}

func NewUtilController() *UtilController {
	return &UtilController{
		SC: config.GetServerConfig(),
	}
}

// 加密
func (this *UtilController) Encrypt(f *form.UtilEncreptForm) (string, error) {
	return peep.Encrypt(f.Data)
}

// 解密
func (this *UtilController) Decrypt(f *form.UtilDecryptForm) (string, error) {
	return peep.Decrypt(f.Data)
}
