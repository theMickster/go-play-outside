package middleware

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"shawskyRecords/settings"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
)

func JwtAuthorization(authSettings settings.AuthSettings) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/swagger/") {
			ctx.Next()
			return
		}

		header := AuthHeader{}
		err := ctx.ShouldBindHeader(&header)
		if err != nil || header.Authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			ctx.Abort()
			return
		}

		const bearer = "Bearer"
		keySet, err := jwk.Fetch(ctx, "https://login.microsoftonline.com/common/discovery/v2.0/keys")
		tokenString := strings.TrimSpace(strings.Replace(header.Authorization, bearer, "", 1))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, err
			}

			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, fmt.Errorf("kid header not found")
			}

			keys, ok := keySet.LookupKeyID(kid)
			if !ok {
				return nil, fmt.Errorf("key %v not found", kid)
			}

			publicKey := &rsa.PublicKey{}
			err = keys.Raw(publicKey)
			if err != nil {
				return nil, fmt.Errorf("unable to not parse public key")
			}

			return publicKey, nil
		})

		if err != nil {
			log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization error AUTH-001"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			result, ok := err.(*jwt.ValidationError)
			if ok {
				if result.Errors&jwt.ValidationErrorMalformed != 0 {
					log.Println("The token string provided is not even a token.")
				} else if result.Errors&jwt.ValidationErrorExpired != 0 {
					log.Println("The token is expired")
				} else if result.Errors&jwt.ValidationErrorNotValidYet != 0 {
					log.Println("The token is not yet valid")
				} else {
					log.Println("Couldn't handle this token:", err)
				}
			} else {
				log.Println("Something awful occurred with token validation. ")
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization error AUTH-002"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			ctx.Set("claims", claims)
		}

		err = claims.Valid()
		if err != nil {
			log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization error AUTH-003"})
			ctx.Abort()
			return
		}

		if !strings.Contains(claims["scp"].(string), authSettings.Scope) {
			log.Printf("Scope '%s' is missing from token scope '%s'", authSettings.Scope, claims["scp"])
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization error AUTH-004"})
			ctx.Abort()
			return
		}

		if !claims.VerifyAudience(authSettings.Audience, true) {
			log.Printf("Invalid audience claim. Audience claim should be %v, but got %v", authSettings.Audience, claims["aud"])
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization error AUTH-005"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
