package routes

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterSubmissionRoutes(r *gin.RouterGroup) {
	r.GET("/submissions", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:8082/submissions")
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Submission service unavailable"})
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, "application/json", body)
	})
}
