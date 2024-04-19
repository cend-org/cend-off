package translator

import (
	"context"
	"duval/internal/graph/model"
	"duval/internal/pkg/translator/resource"
	"duval/internal/utils/errx"
)

func GetMessages(ctx *context.Context) ([]*model.Message, error) {
	var (
		messages    []model.Message
		gqlMessages []*model.Message
		err         error
	)

	messages, err = resource.GetMessages()
	if err != nil {
		return gqlMessages, errx.Lambda(err)
	}

	for _, msg := range messages {
		gqlMessages = append(gqlMessages, &model.Message{
			Id:               msg.Id,
			CreatedAt:        msg.CreatedAt,
			UpdatedAt:        msg.UpdatedAt,
			DeletedAt:        msg.DeletedAt,
			ResourceType:     msg.ResourceType,
			ResourceNumber:   msg.ResourceNumber,
			ResourceValue:    msg.ResourceValue,
			ResourceLabel:    msg.ResourceLabel,
			ResourceLanguage: msg.ResourceLanguage,
		})
	}
	return gqlMessages, nil
}

func GetMessagesInLanguage(ctx *context.Context, language int) ([]*model.Message, error) {
	var (
		messages    []model.Message
		gqlMessages []*model.Message
		err         error
	)

	messages, err = resource.GetMessagesInLanguage(language)
	if err != nil {

		return gqlMessages, errx.Lambda(err)
	}

	for _, msg := range messages {
		gqlMessages = append(gqlMessages, &model.Message{
			Id:               msg.Id,
			CreatedAt:        msg.CreatedAt,
			UpdatedAt:        msg.UpdatedAt,
			DeletedAt:        msg.DeletedAt,
			ResourceType:     msg.ResourceType,
			ResourceNumber:   msg.ResourceNumber,
			ResourceValue:    msg.ResourceValue,
			ResourceLabel:    msg.ResourceLabel,
			ResourceLanguage: msg.ResourceLanguage,
		})
	}
	return gqlMessages, nil
}

func GetMessage(ctx *context.Context, language int, resourceNumber int) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message, err = resource.GetMessage(resourceNumber, language)
	if err != nil {
		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func NewMessage(ctx *context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message.ResourceValue = input.ResourceValue
	message.ResourceType = input.ResourceType
	message.ResourceLanguage = input.ResourceLanguage
	message.ResourceNumber = input.ResourceNumber
	//
	//err = ctx.ShouldBindJSON(&message)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
	//		Message: errx.Lambda(err),
	//	})
	//	return
	//}

	message, err = resource.NewMessage(message.ResourceLabel, message.ResourceLanguage)
	if err != nil {

		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func DelMessage(ctx *context.Context, language int, messageNumber int) (*string, error) {
	var (
		err    error
		status string
	)

	err = resource.DeleteMessage(messageNumber, language)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	status = "ok"

	return &status, nil
}

func UpdMessage(ctx *context.Context, input *model.MessageUpdateInput) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message.Id = input.Id
	message.ResourceValue = input.ResourceValue
	message.ResourceType = input.ResourceType
	message.ResourceLanguage = input.ResourceLanguage
	message.ResourceNumber = input.ResourceNumber
	//err = ctx.ShouldBind(&message)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
	//		Message: errx.Lambda(err),
	//	})
	//	return
	//}

	message, err = resource.UpdateMessage(message)
	if err != nil {
		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func GetMenuList(ctx *context.Context) ([]*model.Message, error) {
	var (
		messages    []model.Message
		gqlMessages []*model.Message
		err         error
	)

	messages, err = resource.GetMenuList()
	if err != nil {
		return gqlMessages, errx.Lambda(err)
	}
	for _, msg := range messages {
		gqlMessages = append(gqlMessages, &model.Message{
			Id:               msg.Id,
			CreatedAt:        msg.CreatedAt,
			UpdatedAt:        msg.UpdatedAt,
			DeletedAt:        msg.DeletedAt,
			ResourceType:     msg.ResourceType,
			ResourceNumber:   msg.ResourceNumber,
			ResourceValue:    msg.ResourceValue,
			ResourceLabel:    msg.ResourceLabel,
			ResourceLanguage: msg.ResourceLanguage,
		})
	}

	return gqlMessages, nil
}

func GetMenuItems(ctx *context.Context, language int, menuNumber int) ([]*model.Message, error) {
	var (
		messages    []model.Message
		gqlMessages []*model.Message
		err         error
	)

	messages, err = resource.GetMenuItems(menuNumber, language)
	if err != nil {
		return gqlMessages, errx.Lambda(err)
	}

	for _, msg := range messages {
		gqlMessages = append(gqlMessages, &model.Message{
			Id:               msg.Id,
			CreatedAt:        msg.CreatedAt,
			UpdatedAt:        msg.UpdatedAt,
			DeletedAt:        msg.DeletedAt,
			ResourceType:     msg.ResourceType,
			ResourceNumber:   msg.ResourceNumber,
			ResourceValue:    msg.ResourceValue,
			ResourceLabel:    msg.ResourceLabel,
			ResourceLanguage: msg.ResourceLanguage,
		})
	}

	return gqlMessages, nil
}

func NewMenu(ctx *context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message.ResourceValue = input.ResourceValue
	message.ResourceType = input.ResourceType
	message.ResourceLanguage = input.ResourceLanguage
	message.ResourceNumber = input.ResourceNumber

	//err = ctx.ShouldBindJSON(&message)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
	//		Message: errx.Lambda(err),
	//	})
	//	return
	//}

	message, err = resource.NewMenu(message.ResourceLabel, message.ResourceLanguage)
	if err != nil {
		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func DelMenu(ctx *context.Context, menuNumber int) (*string, error) {
	var (
		err    error
		status string
	)

	err = resource.DeleteMenu(menuNumber)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	status = "ok"

	return &status, nil
}

func NewMenuItem(ctx *context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		menu model.Message
		err  error
	)

	menu.ResourceValue = input.ResourceValue
	menu.ResourceType = input.ResourceType
	menu.ResourceLanguage = input.ResourceLanguage
	menu.ResourceNumber = input.ResourceNumber
	//err = ctx.ShouldBind(&menu)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
	//		Message: errx.Lambda(err),
	//	})
	//	return
	//}

	menu, err = resource.NewMenuItem(menu.ResourceLabel, menu.ResourceNumber, menu.ResourceLanguage)
	if err != nil {
		return &menu, errx.Lambda(err)
	}

	return &menu, nil
}

func DelMenuItem(ctx *context.Context, input *model.MessageUpdateInput) (*string, error) {
	var (
		menu   model.Message
		err    error
		status string
	)

	menu.Id = input.Id
	menu.ResourceValue = input.ResourceValue
	menu.ResourceType = input.ResourceType
	menu.ResourceLanguage = input.ResourceLanguage
	menu.ResourceNumber = input.ResourceNumber
	//err = ctx.ShouldBind(&menu)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
	//		Message: errx.Lambda(err),
	//	})
	//	return
	//}

	err = resource.DeleteMenuItem(menu.ResourceNumber, menu.ResourceValue, menu.ResourceLanguage)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	status = "ok"
	return &status, nil
}
