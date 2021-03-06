package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseUnauthorized wraps c.AbortWithStatusJSON
func ResponseUnauthorized(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": msg})
}

// ResponseBadRequest wraps c.AbortWithStatusJSON
func ResponseBadRequest(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": msg})
}

// ResponseNotFound wraps c.AbortWithStatusJSON
func ResponseNotFound(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": msg})
}

// ResponseInternalServerError wraps c.AbortWithStatusJSON
func ResponseInternalServerError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": msg})
}

// ResponseSuccess wraps c.AbortWithStatusJSON
func ResponseSuccess(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": msg})
}

// ResponseSuccessPayload wraps c.AbortWithStatusJSON
func ResponseSuccessPayload(c *gin.Context, msg, payload interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": msg, "payload": payload})
}
