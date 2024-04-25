package translator

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
)

func NewMessage(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: NewMessage - newMessage"))
}

func UpdMessage(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: UpdMessage - updMessage"))
}

func DelMessage(ctx context.Context, language int, messageNumber int) (*string, error) {
	panic(fmt.Errorf("not implemented: DelMessage - delMessage"))
}

func NewMenu(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: NewMenu - newMenu"))
}

func DelMenu(ctx context.Context, menuNumber int) (*string, error) {
	panic(fmt.Errorf("not implemented: DelMenu - delMenu"))
}

func NewMenuItem(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: NewMenuItem - newMenuItem"))
}

func DelMenuItem(ctx context.Context, input *model.MessageInput) (*string, error) {
	panic(fmt.Errorf("not implemented: DelMenuItem - delMenuItem"))
}

func GetMessages(ctx context.Context) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: GetMessages - getMessages"))
}

func GetMessagesInLanguage(ctx context.Context, language int) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: GetMessagesInLanguage - getMessagesInLanguage"))
}

func GetMessage(ctx context.Context, language int, resourceNumber int) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: GetMessage - getMessage"))
}

func GetMenuList(ctx context.Context) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: GetMenuList - getMenuList"))
}

func GetMenuItems(ctx context.Context, language int, menuNumber int) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: GetMenuItems - getMenuItems"))
}
