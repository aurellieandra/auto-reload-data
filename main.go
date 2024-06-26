package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Element struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Data map[string]Element

func main() {
	g := gin.Default()

	elementData := make(Data)

	go func() {
		for {
			elementData["water"] = Element{
				Value: rand.Intn(100),
				Unit:  "m",
			}
			elementData["wind"] = Element{
				Value: rand.Intn(100),
				Unit:  "m/s",
			}
			time.Sleep(15 * time.Second)
		}
	}()

	g.GET("/json", func(ctx *gin.Context) {

		var waterStatus string
		var windStatus string

		water := elementData["water"].Value
		wind := elementData["wind"].Value

		switch {
		case water <= 50:
			waterStatus = "Aman"
		case water >= 80:
			waterStatus = "Bahaya"
		default:
			waterStatus = "Siaga"
		}

		switch {
		case wind <= 60:
			windStatus = "Aman"
		case wind >= 90:
			windStatus = "Bahaya"
		default:
			windStatus = "Siaga"
		}

		response := map[string]interface{}{
			"water":       fmt.Sprintf("%d %s", water, elementData["water"].Unit),
			"waterStatus": waterStatus,
			"wind":        fmt.Sprintf("%d %s", wind, elementData["wind"].Unit),
			"windStatus":  windStatus,
		}

		ctx.JSON(http.StatusOK, response)
	})

	g.LoadHTMLGlob("template/*")
	g.GET("/index", func(ctx *gin.Context) {
		var waterStatus string
		var windStatus string

		water := elementData["water"].Value
		wind := elementData["wind"].Value

		switch {
		case water <= 50:
			waterStatus = "Aman"
		case water >= 80:
			waterStatus = "Bahaya"
		default:
			waterStatus = "Siaga"
		}

		switch {
		case wind <= 60:
			windStatus = "Aman"
		case wind >= 90:
			windStatus = "Bahaya"
		default:
			windStatus = "Siaga"
		}

		ctx.HTML(http.StatusOK, "index.html", map[string]any{
			"water":       fmt.Sprintf("%d %s", water, elementData["water"].Unit),
			"waterStatus": waterStatus,
			"wind":        fmt.Sprintf("%d %s", wind, elementData["wind"].Unit),
			"windStatus":  windStatus,
		})
	})
	g.Run(":3000")
}
