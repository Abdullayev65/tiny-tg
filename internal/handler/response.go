package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tiny-tg/internal/pkg/app_errors"
)

func success(c *gin.Context, res any, code ...int) {
	//c.JSON(getCode(200, code), map[string]any{
	//	"res":    res,
	//	"status": true,
	//})
	c.JSON(getCode(200, code), res)
}

func fail(c *gin.Context, msg string, code ...int) {
	c.JSON(getCode(400, code), map[string]any{
		"status":  false,
		"message": msg,
	})
}

func failErr(c *gin.Context, err error, code ...int) {

	var appErr *app_errors.AppError
	if errors.As(err, &appErr) {
		fail(c, appErr.Message, appErr.Status)
		return
	}

	fail(c, err.Error(), code...)
}

// others

func hasErr(c *gin.Context, err error, code ...int) bool {

	if err != nil {
		failErr(c, err, code...)
		return true
	}

	return false
}

func finish(c *gin.Context, res any, err error, code ...int) {

	if err != nil {
		failErr(c, err, code...)
		return
	}

	success(c, res, code...)
}

func getCode(c int, code []int) int {
	if len(code) == 1 {
		return code[0]
	}

	return c
}
