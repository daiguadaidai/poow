package form

type UtilEncreptForm struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilDecryptForm struct {
	Data string `json:"data" form:"data" binding:"required"`
}
