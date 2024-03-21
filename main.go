package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default

	data := map[string]any{}

	go func()  {
		for {
			// read json file
			// update json file
			data["water"] = rand.Int31n(100)
			data["wind"] = rand.Int31n(100)
			data["fire"] = rand.Int31n(100)
			data["earth"] = rand.Int31n(100)
			time.Sleep(15 * time.Second)
		}
	}()

	g.GET("/", func(ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": data,
		})
	})

	g.Run(":3000")
}