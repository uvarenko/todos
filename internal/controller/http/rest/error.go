package rest

import "github.com/gin-gonic/gin"

type response struct {
	Error string `json:"error"`
}

func errorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, response{message})
}
