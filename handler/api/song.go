package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (me *API) Songs(c *gin.Context) {
	c.JSON(http.StatusOK, me.Usecase.GetSongs())
}

func (me *API) Song(c *gin.Context) {
	song, err := me.Usecase.GetPrayer(c.Param("id"))
	if err != nil {
		NewProblem(http.StatusInternalServerError,
			"error on search song", err.Error()).JSON(c)
		return
	}

	c.JSON(http.StatusOK, song)
}
