package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type MarkCollector struct {} 
func (c *MarkCollector) UserId(a *int) (r Mark, err error) { 
		err = database.Get(&r, `SELECT * FROM mark WHERE user_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MarkCollector) AuthorId(a *int) (r Mark, err error) { 
		err = database.Get(&r, `SELECT * FROM mark WHERE author_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MarkCollector) AuthorComment(a *string) (r Mark, err error) { 
		err = database.Get(&r, `SELECT * FROM mark WHERE author_comment = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MarkCollector) AuthorMark(a *int) (r Mark, err error) { 
		err = database.Get(&r, `SELECT * FROM mark WHERE author_mark = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MarkCollector) _b(){ 
		_ = time.Now()
}

func MapMarkInputToMark(a MarkInput, e Mark) Mark { 
 if a.UserId != nil { 
	e.UserId = *a.UserId 
 }
 if a.AuthorId != nil { 
	e.AuthorId = *a.AuthorId 
 }
 if a.AuthorComment != nil { 
	e.AuthorComment = *a.AuthorComment 
 }
 if a.AuthorMark != nil { 
	e.AuthorMark = *a.AuthorMark 
 }
  return e
}
