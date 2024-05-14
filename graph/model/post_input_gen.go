package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type PostCollector struct {} 
func (c *PostCollector) PublisherId(a *int) (r Post, err error) { 
		err = database.Get(&r, `SELECT * FROM post WHERE publisher_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PostCollector) Content(a *string) (r Post, err error) { 
		err = database.Get(&r, `SELECT * FROM post WHERE content = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PostCollector) ExpirationDate(a *time.Time) (r Post, err error) { 
		err = database.Get(&r, `SELECT * FROM post WHERE expiration_date = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PostCollector) _b(){ 
		_ = time.Now()
}

func MapPostInputToPost(a PostInput, e Post) Post { 
 if a.PublisherId != nil { 
	e.PublisherId = *a.PublisherId 
 }
 if a.Content != nil { 
	e.Content = *a.Content 
 }
 if a.ExpirationDate != nil { 
	e.ExpirationDate = *a.ExpirationDate 
 }
  return e
}
