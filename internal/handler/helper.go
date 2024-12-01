package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	errInvalidUserId   = errors.New("invalid user id")
	errInvalidUserRole = errors.New("invalid user role")
)

type UserInfo struct {
	Id int
}

func getUserInfo(ctx *gin.Context) (UserInfo, error) {
	var userInfo UserInfo

	userId, ok := ctx.Get(UserIdCtx)
	if !ok {
		return UserInfo{}, errInvalidUserId
	}

	userInfo.Id, ok = userId.(int)
	if !ok {
		return UserInfo{}, errInvalidUserId
	}

	return userInfo, nil
}

func getPageQuery(c *gin.Context) (int, error) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 0, errors.New("invalid page parameter")
	}
	return page, nil
}

func getLimitQuery(c *gin.Context) (int, error) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return 0, errors.New("invalid limit parameter")
	}
	return limit, nil
}

func calculatePagination(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return offset, limit
}

func shouldBind[D any](c *gin.Context) (*D, error) {
	data := new(D)
	err := c.ShouldBind(data)

	return data, err
}
