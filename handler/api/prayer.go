package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (me *API) Prayers(c *gin.Context) {
	c.JSON(http.StatusOK, me.Usecase.GetPrayers())
}

func (me *API) Prayer(c *gin.Context) {
	prayer, err := me.Usecase.GetPrayer(c.Param("id"))
	if err != nil {
		NewProblem(http.StatusInternalServerError,
			"error on search prayer", err.Error()).JSON(c)
		return
	}

	c.JSON(http.StatusOK, prayer)
}
