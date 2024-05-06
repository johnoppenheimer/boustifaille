package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	"github.com/johnoppenheimer/boustifaille/database"
	"github.com/johnoppenheimer/boustifaille/database/services"
	"github.com/johnoppenheimer/boustifaille/web/gintemplrenderer"
	"github.com/johnoppenheimer/boustifaille/web/pages"
)

func main() {
	db := database.InitDB()

	log.SetLevel(log.DebugLevel)

	engine := gin.Default()
	ginHtmlRenderer := engine.HTMLRender
	engine.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	engine.SetTrustedProxies(nil)

	placesService := services.NewPlaceService(db)

	engine.GET("/", func(c *gin.Context) {
		places, err := placesService.All()
		log.Debugf("places: %v", places)
		if err != nil {
			log.Errorf("Error getting all places: %v", err)
		}

		c.HTML(http.StatusOK, "", pages.Index(places))
	})

	log.Info("Listening on port 8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Error listening %v", err)
	}
}
