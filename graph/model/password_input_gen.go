package model

import (
	"github.com/cend-org/duval/internal/database"
	"time"
)

type PasswordCollector struct{}

func (c *PasswordCollector) Hash(a *string) (r Password, err error) {
	err = database.Get(&r, `SELECT * FROM password WHERE hash = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *PasswordCollector) _b() {
	_ = time.Now()
}

func MapPasswordInputToPassword(a PasswordInput, e Password) Password {
	if a.Hash != nil {
		e.Hash = *a.Hash
	}
	return e
}
