package translator

import (
	"duval/internal/pkg/translator/resource"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMessages(ctx *gin.Context) {
	var (
		messages []resource.Message
		err      error
	)

	messages, err = resource.GetMessages()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}

func GetMessagesInLanguage(ctx *gin.Context) {
	var (
		messages []resource.Message
		language int
		err      error
	)

	language, err = strconv.Atoi(ctx.Param("language"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	messages, err = resource.GetMessagesInLanguage(language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}

func GetMessage(ctx *gin.Context) {
	var (
		message        resource.Message
		resourceNumber int
		language       int
		err            error
	)

	language, err = strconv.Atoi(ctx.Param("language"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	resourceNumber, err = strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	message, err = resource.GetMessage(resourceNumber, language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, message)
	return
}

func NewMessage(ctx *gin.Context) {
	var (
		message resource.Message
		err     error
	)

	err = ctx.ShouldBindJSON(&message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	message, err = resource.NewMessage(message.ResourceLabel, message.ResourceLanguage)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, message)
	return
}

func DelMessage(ctx *gin.Context) {
	var (
		messageNumber int
		language      int
		err           error
	)

	language, err = strconv.Atoi(ctx.Param("language"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	messageNumber, err = strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
	}

	err = resource.DeleteMessage(messageNumber, language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func UpdMessage(ctx *gin.Context) {
	var (
		message resource.Message
		err     error
	)

	err = ctx.ShouldBind(&message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	message, err = resource.UpdateMessage(message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, message)
	return
}

func GetMenuList(ctx *gin.Context) {
	var (
		messages []resource.Message
		err      error
	)

	messages, err = resource.GetMenuList()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}

func GetMenuItems(ctx *gin.Context) {
	var (
		messages   []resource.Message
		language   int
		menuNumber int
		err        error
	)

	language, err = strconv.Atoi(ctx.Param("language"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	menuNumber, err = strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	messages, err = resource.GetMenuItems(menuNumber, language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}

func NewMenu(ctx *gin.Context) {
	var (
		message resource.Message
		err     error
	)

	err = ctx.ShouldBindJSON(&message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	message, err = resource.NewMenu(message.ResourceLabel, message.ResourceLanguage)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, message)
	return
}

func DelMenu(ctx *gin.Context) {
	var (
		menuNumber int
		err        error
	)

	menuNumber, err = strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	err = resource.DeleteMenu(menuNumber)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func NewMenuItem(ctx *gin.Context) {
	var (
		menu resource.Message
		err  error
	)

	err = ctx.ShouldBind(&menu)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	menu, err = resource.NewMenuItem(menu.ResourceLabel, menu.ResourceNumber, menu.ResourceLanguage)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, menu)
	return
}

func DelMenuItem(ctx *gin.Context) {
	var (
		menu resource.Message
		err  error
	)

	err = ctx.ShouldBind(&menu)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	err = resource.DeleteMenuItem(menu.ResourceNumber, menu.ResourceValue, menu.ResourceLanguage)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
