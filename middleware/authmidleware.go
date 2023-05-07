package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/celpostgress-api/common"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
		authHeader := c.GetHeader(authorizationHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ResponseUnAuthorized(err.Error()))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ResponseUnAuthorized(err.Error()))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ResponseUnAuthorized(err.Error()))
			return
		}

		accessToken := fields[1]
		marker, err := NewJwt()
		payload, err := marker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ResponseUnAuthorized(err.Error()))
			return
		}

		c.Set(authorizationPayloadKey, payload)

		c.Next()
	}
}
