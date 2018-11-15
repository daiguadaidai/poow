package views

import (
	"net/http"
	"strconv"

	"github.com/daiguadaidai/poow/utils"
	"github.com/gin-gonic/gin"
)

const (
	defaultStrPage    = "1"
	defaultStrPerPage = "100"
	defaultPage       = 1
	defaultPerPage    = 100
	maxPerPage        = 1000
)

func parsePaginator(c *gin.Context) *utils.Paginator {
	strPage := c.DefaultQuery("page", defaultStrPage)
	page, err := strconv.ParseUint(strPage, 10, 64)
	if err != nil {
		page, _ = strconv.ParseUint(defaultStrPage, 10, 64)
	}
	if page <= 0 {
		page = defaultPage
	}
	strPerPage := c.DefaultQuery("per_page", defaultStrPerPage)
	perPage, err := strconv.ParseUint(strPerPage, 10, 64)
	if err != nil {
		perPage, _ = strconv.ParseUint(defaultStrPerPage, 10, 64)
	}
	if perPage > maxPerPage {
		perPage = maxPerPage
	}
	return utils.NewPaginator(uint64((page-1)*perPage), uint64(perPage))
}

type response struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 返回错误
func returnError(c *gin.Context, code int, err error) {
	resp := &response{Data: make([]interface{}, 0)}
	resp.Status = false
	resp.Message = err.Error()
	c.JSON(code, resp)
}

// 成功并放回数据
func returnSuccess(c *gin.Context, obj interface{}) {
	resp := &response{Data: obj}
	resp.Status = true
	resp.Message = "success"
	c.JSON(http.StatusOK, resp)
}

// 返回文件信息
func returnFile(c *gin.Context, fileName string, path string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}
