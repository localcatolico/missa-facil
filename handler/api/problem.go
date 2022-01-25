package api

import (
	"github.com/gin-gonic/gin"
)

type Problem struct {
	StatusCode int    `json:"status"`
	Problem    string `json:"problem"`
	Details    string `json:"details,omitempty"`
}

func NewProblem(status int, problem string, details string) *Problem {
	return &Problem{
		StatusCode: status,
		Problem:    problem,
		Details:    details,
	}
}

func (me *Problem) JSON(c *gin.Context) {
	c.JSON(me.StatusCode, me)
}
