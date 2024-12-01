package handler

import (
	"tiny-tg/internal/dtos"
	"tiny-tg/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	data, err := shouldBind[models.User](c)
	if hasErr(c, err) {
		return
	}

	res, err := h.service.Auth.Register(data)

	finish(c, res, err)
}

func (h *Handler) Login(c *gin.Context) {
	data, err := shouldBind[dtos.Login](c)
	if hasErr(c, err) {
		return
	}

	res, err := h.service.Auth.Login(data)

	finish(c, res, err)
}
