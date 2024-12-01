package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) WS(c *gin.Context) {
	info, err := getUserInfo(c)
	if hasErr(c, err) {
		return
	}

	h.wsHub.Handle(info.Id, c.Writer, c.Request)

}
