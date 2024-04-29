package translator

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/pkg/translator/resource"
	"github.com/cend-org/duval/internal/utils/errx"
)

func NewMessage(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message = model.MapMessageInputToMessage(*input, message)

	message, err = resource.NewMessage(message.ResourceLabel, message.ResourceLanguage)
	if err != nil {

		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func UpdMessage(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message = model.MapMessageInputToMessage(*input, message)

	message, err = resource.UpdateMessage(message)
	if err != nil {
		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func DelMessage(ctx context.Context, language int, messageNumber int) (*string, error) {
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

func NewMenu(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		message model.Message
		err     error
	)

	message = model.MapMessageInputToMessage(*input, message)

	message, err = resource.NewMenu(message.ResourceLabel, message.ResourceLanguage)
	if err != nil {
		return &message, errx.Lambda(err)
	}

	return &message, nil
}

func DelMenu(ctx context.Context, menuNumber int) (*string, error) {
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

func NewMenuItem(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	var (
		menu model.Message
		err  error
	)

	menu = model.MapMessageInputToMessage(*input, menu)

	menu, err = resource.NewMenuItem(menu.ResourceLabel, menu.ResourceNumber, menu.ResourceLanguage)
	if err != nil {
		return &menu, errx.Lambda(err)
	}

	return &menu, nil
}

func DelMenuItem(ctx context.Context, input *model.MessageInput) (*string, error) {
	var (
		menu   model.Message
		err    error
		status string
	)
	menu = model.MapMessageInputToMessage(*input, menu)

	err = resource.DeleteMenuItem(menu.ResourceNumber, menu.ResourceValue, menu.ResourceLanguage)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	status = "ok"
	return &status, nil
}

func GetMessages(ctx context.Context) ([]model.Message, error) {
	var (
		messages []model.Message
		err      error
	)

	messages, err = resource.GetMessages()
	if err != nil {
		return messages, errx.Lambda(err)
	}

	return messages, nil
}

func GetMessagesInLanguage(ctx context.Context, language int) ([]model.Message, error) {
	var (
		messages []model.Message
		err      error
	)

	messages, err = resource.GetMessagesInLanguage(language)
	if err != nil {

		return messages, errx.Lambda(err)
	}

	return messages, nil
}

func GetMessage(ctx context.Context, language int, resourceNumber int) (*model.Message, error) {
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

func GetMenuList(ctx context.Context) ([]model.Message, error) {
	var (
		messages []model.Message
		err      error
	)

	messages, err = resource.GetMenuList()
	if err != nil {
		return messages, errx.Lambda(err)
	}

	return messages, nil
}

func GetMenuItems(ctx context.Context, language int, menuNumber int) ([]model.Message, error) {
	var (
		messages []model.Message
		err      error
	)

	messages, err = resource.GetMenuItems(menuNumber, language)
	if err != nil {
		return messages, errx.Lambda(err)
	}

	return messages, nil
}
