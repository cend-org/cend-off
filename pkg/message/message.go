package message

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
)

func DeleteMessage(messageId int) (*bool, error) {
	var (
		message model.Message
		err     error
		status  bool
	)

	message, err = GetMessage(messageId)
	if err != nil {
		return &status, errx.SupportError
	}

	err = RmMessage(message)
	if err != nil {
		return &status, errx.SupportError
	}
	status = true
	return &status, nil
}

/*

	UTILS

*/

func GetMessage(messageId int) (message model.Message, err error) {
	err = database.Get(&message, `SELECT * FROM message WHERE  id = ?`, messageId)
	if err != nil {
		return message, err
	}

	return message, nil
}

func GetMessages(userId int) (message []model.Message, err error) {
	err = database.Select(&message, `SELECT * FROM message JOIN user_message um on message.id = um.user_id = ?`, userId)
	if err != nil {
		return message, err
	}
	return message, nil
}

func RmMessage(message model.Message) error {
	err := database.Delete(message)
	if err != nil {
		return err
	}
	return nil
}
