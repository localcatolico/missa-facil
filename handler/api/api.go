package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"golang.org/x/oauth2"
)

type API struct {
	Router  *gin.Engine
	Usecase *usecase.Usecase
	OAuth2  *oauth2.Config
}

func NewAPI(router *gin.Engine, usecase *usecase.Usecase, oauth2 *oauth2.Config) *API {
	return &API{
		Router:  router,
		Usecase: usecase,
		OAuth2:  oauth2,
	}
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "slides-gospel",
		"endpoints": []string{
			fmt.Sprintf("http://%s/songs", c.Request.Host),
			fmt.Sprintf("http://%s/songs/tu-es-o-centro-frei-gilson", c.Request.Host),
			fmt.Sprintf("http://%s/prayers", c.Request.Host),
			fmt.Sprintf("http://%s/prayers/1", c.Request.Host),
		},
	})
}
