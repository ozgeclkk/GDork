package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/sorgula", func(ctx *gin.Context) {
		hedefSite := ctx.Query("hedef")

		if hedefSite == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"hata": "Lütfen bir hedef domain girin!"})
			return
		}

		uretilenDorklar := SorguHazirla(hedefSite)

		ctx.JSON(http.StatusOK, uretilenDorklar)

	})

	r.Run(":5050")

}
