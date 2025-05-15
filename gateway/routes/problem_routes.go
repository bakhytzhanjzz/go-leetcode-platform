package routes

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProblemRoutes(r *gin.RouterGroup) {
	r.GET("/problems", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:8081/products") // your problem-service
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Problem service unavailable"})
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, "application/json", body)
	})
}
