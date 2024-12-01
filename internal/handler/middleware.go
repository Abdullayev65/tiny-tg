package handler

import (
	"github.com/gin-gonic/gin"
	"tiny-tg/internal/pkg/app_errors"
)

const (
	AuthorizationHeader = "Authorization"
	UserIdCtx           = "user_id"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(AuthorizationHeader)
	if header == "" {
		failErr(c, app_errors.AuthMwMissingToken)
		c.Abort()
		return
	}

	userId, err := h.jwtManager.Parse(header)
	if err != nil {
		failErr(c, err)
		c.Abort()
		return
	}

	c.Set(UserIdCtx, userId)
	c.Next()
}
