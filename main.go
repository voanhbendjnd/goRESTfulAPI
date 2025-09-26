package main

import (
	v1hanlder "my-go-api/internal/api/v1/handler"
	// "net/http"

	"github.com/gin-gonic/gin"
)

//	func GetUserAPI(ctx*gin.Context){
//		ctx.JSON(http.StatusOK, gin.H{
//			"message": "hehe",
//		})
//	}
func main() {
	r := gin.Default()       // create router
	v1 := r.Group("/api/v1") // requestMapping
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1hanlder.NewUserHanlder() // service
			user.GET("", userHandlerV1.GetUserAPI)
		}
	}

	r.Run(":8080") // port
}
