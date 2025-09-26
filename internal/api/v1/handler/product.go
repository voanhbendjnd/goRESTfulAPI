package v1hanlder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHanlder struct {
}

func NewProductHanlder() *ProductHanlder {
	return &ProductHanlder{}
}
func (u *ProductHanlder) GetProductAPI(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product hello",
	})
}