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

	restaurantsService := services.NewRestaurantService(db)

	engine.GET("/", func(c *gin.Context) {
		restaurants, err := restaurantsService.All()
		log.Debugf("restaurants: %v", restaurants)
		if err != nil {
			log.Errorf("Error getting all restaurants: %v", err)
		}

		c.HTML(http.StatusOK, "", pages.Index(restaurants))
	})

	log.Info("Listening on port 8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Error listening %v", err)
	}
}
