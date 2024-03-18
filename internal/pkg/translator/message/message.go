package message

import (
	"database/sql"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"github.com/joinverse/xid"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Identifier string     `json:"identifier"`
	Number     int        `json:"-"`
	Xid        string     `json:"-"`
	Label      string     `json:"label"`
	Language   int        `json:"language"`
}

type Menu struct {
	Id            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Identifier    string     `json:"identifier"`
	Number        int        `json:"-"`
	MessageNumber uint       `json:"-"`
	Message       Message    `json:"message" q:"_"`
	Items         []MenuItem `json:"items" q:"_"`
}

type MenuItem struct {
	Id                     uint       `json:"id"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at"`
	Identifier             string     `json:"identifier"`
	Number                 int        `json:"-"`
	MenuTitleMessageNumber int        `json:"-"`
	MessageNumber          uint       `json:"-"`
	Message                Message    `json:"message" q:"_"`
}

func GetMessages() (messages []Message, err error) {
	err = database.Select(&messages, `SELECT * FROM  message WHERE id > 0 ORDER BY number desc `)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func GetMessagesInLanguage(language int) (messages []Message, err error) {
	err = database.Select(&messages, `SELECT * FROM message WHERE language = ? ORDER BY number desc `, language)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func GetMessage(identifier string) (messages []Message, err error) {
	err = database.Select(&messages, `SELECT * FROM message WHERE identifier = ? ORDER BY language`, identifier)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func GetMessageInLanguage(identifier string, language int) (message Message, err error) {
	query := `SELECT 
				COALESCE(m1.id, m.id) as 'id',
				COALESCE(m1.created_at, m.created_at) as 'created_at',
				COALESCE(m1.identifier, m.identifier) as 'identifier',
				COALESCE(m1.language, m.language) as 'language',
				COALESCE(m1.number, m.number) as 'number',
				COALESCE(m1.label, m.language) as 'label',
				COALESCE(m1.xid, m.xid) as 'xid'
				FROM message m
					LEFT JOIN message m1 ON m.identifier = m1.identifier AND m1.language = ?
				WHERE m.identifier = ? AND m.language = 0
			`

	err = database.Select(&message, query, language, identifier)
	if err != nil {
		return message, err
	}

	return message, err
}

func DeleteMessage(message Message) (err error) {
	if message.Identifier == state.EMPTY {
		return errors.New("message should have identifier")
	}

	if message.Language == state.ZERO {
		err = database.Exec(`DELETE FROM message WHERE identifier = ?`, message.Identifier)
		if err != nil {
			return err
		}
		return err
	}

	err = database.Exec(`DELETE FROM message WHERE identifier = ? AND language = ?`, message.Identifier, message.Language)
	if err != nil {
		return err
	}

	return err
}

func UpdateMessage(message Message) (msg Message, err error) {
	if message.Id == state.ZERO {
		return message, errors.New("id is not set")
	}

	msg, err = GetMessageInLanguage(message.Identifier, message.Language)
	if err != nil {
		return message, errors.New("cannot get message")
	}

	message.Number = msg.Number
	message.Xid = msg.Xid

	err = database.Update(message)
	if err != nil {
		return message, err
	}

	return message, err
}

func NewMessage(message Message) (msg Message, err error) {
	if strings.TrimSpace(message.Label) == state.EMPTY {
		return msg, errors.New("message label cannot be empty")
	}

	message.Number, err = getNewMessageNumber()
	if err != nil {
		return message, err
	}

	message.Xid = xid.New().String()
	if strings.TrimSpace(message.Identifier) == state.EMPTY {
		message.Identifier = strconv.Itoa(message.Number)
	}

	if message.Language > state.ZERO {
		var orgMessage Message = message
		orgMessage.Language = 0
		_, err = NewMessage(orgMessage)
		if err != nil {
			return message, err
		}
	}

	msg.Id, err = database.InsertOne(message)
	if err != nil {
		return message, err
	}

	return message, err
}

func getNewMessageNumber() (nb int, err error) {
	err = database.Get(&nb, `SELECT number from message ORDER BY number DESC LIMIT 1`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nb, nil
		}
		return nb, err
	}
	return nb + 1, err
}
