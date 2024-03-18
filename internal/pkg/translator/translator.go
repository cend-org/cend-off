package translator

import (
	"duval/internal/pkg/translator/message"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Messages(ctx *gin.Context) {
	var (
		messages []message.Message
		err      error
	)

	messages, err = message.GetMessages()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}

func Message(ctx *gin.Context) {
	var (
		msg message.Message
		err error
	)

	err = ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	if strings.TrimSpace(msg.Identifier) != state.EMPTY {
		msg, err = message.GetMessageInLanguage(msg.Identifier, msg.Language)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.Lambda(err),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, msg)
	return
}

func NewMessage(ctx *gin.Context) {
	var (
		msg message.Message
		err error
	)

	err = ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	msg, err = message.NewMessage(msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, msg)
	return
}

func UpdMessage(ctx *gin.Context) {
	var (
		err error
		msg message.Message
	)

	err = ctx.ShouldBind(&msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	msg, err = message.UpdateMessage(msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, msg)
	return
}

func DeleteMessage(ctx *gin.Context) {
	var (
		err error
		msg message.Message
	)

	err = ctx.ShouldBind(&msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	err = message.DeleteMessage(msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.Status(http.StatusOK)
	return
}
