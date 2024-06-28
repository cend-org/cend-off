package model

import (
	"github.com/cend-org/duval/internal/database"
	"time"
)

type MessageCollector struct{}

func (c *MessageCollector) Text(a *string) (r Message, err error) {
	err = database.Get(&r, `SELECT * FROM message WHERE text = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *MessageCollector) _b() {
	_ = time.Now()
}

func MapMessageInputToMessage(a MessageInput, e Message) Message {
	if a.Text != nil {
		e.Text = *a.Text
	}
	return e
}
