package message

import (
	"context"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/socketio"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"time"
)

type MessageQuery struct{}
type MessageMutation struct{}
type MessageSubscription struct{}

// publish message

func (m *MessageMutation) NewMessage(ctx context.Context, new model.MessageInput) (*model.Message, error) {
	var (
		tok         *token.Token
		message     model.Message
		userMessage model.UserMessage
		err         error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	message = model.MapMessageInputToMessage(new, message)
	message.SenderId = tok.UserId

	time.Sleep(100)
	err = socketio.SendMessage(message)
	if err != nil {
		return &message, errx.SupportError
	}

	messageId, err := database.InsertOne(message)
	if err != nil {
		return &message, errx.SupportError
	}

	userMessage.UserId = tok.UserId
	userMessage.MessageId = messageId

	_, err = database.InsertOne(userMessage)
	if err != nil {
		return &message, errx.SupportError
	}

	return &message, nil
}

//subscribe message

func (s *MessageSubscription) MessageReceived(ctx context.Context) (<-chan *model.Message, error) {
	msgChan := make(chan *model.Message, 1)

	go func() {
		defer close(msgChan)
		c, err := gosocketio.Dial(
			gosocketio.GetUrl("localhost", 3811, false),
			transport.GetDefaultWebsocketTransport())
		if err != nil {
			return
		}

		c.On("/message", func(h *gosocketio.Channel, args model.Message) {
			now := time.Now()
			msgChan <- &model.Message{
				Id:        args.Id,
				CreatedAt: now,
				UpdatedAt: now,
				DeletedAt: nil,
				Channel:   args.Channel,
				Text:      args.Text,
			}
		})
	}()

	return msgChan, nil
}

//get all message

func (q *MessageQuery) Messages(ctx context.Context) ([]model.Message, error) {
	var (
		tok      *token.Token
		messages []model.Message
		err      error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	messages, err = GetMessages(tok.UserId)
	if err != nil {
		return nil, errx.SupportError
	}
	return messages, nil
}

// delete selected Message

func (m *MessageMutation) RemoveMessage(ctx context.Context, messageID int) (*bool, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if tok.UserId == 0 {
		return nil, errx.UnAuthorizedError
	}
	return DeleteMessage(messageID)
}

func (m *MessageMutation) NewDefaultGroup(ctx context.Context) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageMutation) NewGroup(ctx context.Context) (*bool, error) {
	//TODO implement me
	panic("implement me")
}
