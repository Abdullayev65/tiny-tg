package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiny-tg/internal/dtos"
)

func (h *Handler) GetChat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("chat_id"))
	if hasErr(c, err) {
		return
	}

	res, err := h.service.Chats.GetGroupChat(id)

	finish(c, res, err)
}

func (h *Handler) GetPersonalChat(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if hasErr(c, err) {
		return
	}

	info, err := getUserInfo(c)
	if hasErr(c, err) {
		return
	}

	res, err := h.service.Chats.MustGetPersonalChat([2]int{info.Id, userId})

	finish(c, res, err)
}

func (h *Handler) SearchChat(c *gin.Context) {
	opt := new(dtos.ListOpts)
	err := c.BindQuery(opt)
	if hasErr(c, err) {
		return
	}

	res, err := h.service.Chats.SearchChat(opt)

	finish(c, res, err)

}
