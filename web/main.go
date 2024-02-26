package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/llmpros/lark2md/utils"
	"github.com/gin-gonic/gin"
)

//go:embed templ/*
var f embed.FS

func main() {
	if mode := os.Getenv("GIN_MODE"); mode != "release" {
		utils.LoadEnv()
	}

	router := gin.New()
	templ := template.Must(template.New("").ParseFS(f, "templ/*.templ.html"))
	router.SetHTMLTemplate(templ)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.templ.html", nil)
	})
	router.GET("/download", downloadHandler)

	if err := router.Run(); err != nil {
		log.Panicf("error: %s", err)
	}
}
