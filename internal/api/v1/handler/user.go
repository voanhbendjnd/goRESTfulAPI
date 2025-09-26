package v1hanlder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHanlder struct {
}

func NewUserHanlder() *UserHanlder {
	return &UserHanlder{}
}
func (u *UserHanlder) GetUserAPI(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User hello",
	})
}