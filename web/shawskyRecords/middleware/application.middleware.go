package middleware

import (
	"log"
	"net/http"
	"shawskyRecords/settings"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApplicationHeader(appSettings settings.ApplicationSettings) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if appSettings.ApplicationId == "" {
			log.Fatal("please configure application id environment variable")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server is misconfigured for this application"})
			return
		}

		xApplicationId := strings.TrimSpace(ctx.Request.Header.Get("X-ApplicationId"))
		if (len(xApplicationId) == 0 || !strings.EqualFold(appSettings.ApplicationId, xApplicationId)) &&
			!strings.HasPrefix(ctx.Request.URL.Path, "/swagger/") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect application id"})
			return
		}

		ctx.Next()
	}
}
