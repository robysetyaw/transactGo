package middleware

import (
	"net/http"
	"strings"
	"transactgo/app/model/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil, nil, "Unauthorized"))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid token", nil, nil, "Invalid token"))
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Return secret key directly
			return []byte("secret.puppey"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil, nil, "Unauthorized"))
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil, nil, "Unauthorized"))
			c.Abort()
			return
		}

		// Token is valid, extract the username
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username, ok := claims["username"].(string)
			if !ok {
				c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Internal server error", nil, nil, "Failed to get username"))
				c.Abort()
				return
			}
			// Add username to the context
			c.Set("username", username)
		}

		// Token is valid
		c.Next()
	}
}
