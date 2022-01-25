package api

import (
	"github.com/rafaelbmateus/slides-gospel/handler/middleware"
)

func (me *API) Routes() {
	me.Router.GET("/songs", me.Songs)
	me.Router.GET("/songs/:id", me.Song)
	me.Router.GET("/prayers", me.Prayers)
	me.Router.GET("/prayers/:id", me.Prayer)

	private := me.Router.Group("/api")
	private.Use(middleware.AuthRequired)
	private.GET("/", Home)
	private.POST("/presentation", me.CreatePresentation)
}
