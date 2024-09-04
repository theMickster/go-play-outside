package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ApplicationHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		applicationId := strings.TrimSpace(os.Getenv("AdventureWorksApplicationId"))
		if applicationId == "" {
			log.Fatal("please configure application id environment variable")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server is misconfigured for this application"})
			return
		}

		xApplicationId := strings.TrimSpace(ctx.Request.Header.Get("X-ApplicationId"))
		if (len(xApplicationId) == 0 || !strings.EqualFold(applicationId, xApplicationId)) &&
			!strings.HasPrefix(ctx.Request.URL.Path, "/swagger/") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect application id"})
			return
		}

		ctx.Next()
	}
}

func RequestIdHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-RequestId", uuid.NewString())
		ctx.Next()
	}
}
