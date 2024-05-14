package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type PostTagCollector struct {} 
func (c *PostTagCollector) PostId(a *int) (r PostTag, err error) { 
		err = database.Get(&r, `SELECT * FROM post_tag WHERE post_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PostTagCollector) TagContent(a *string) (r PostTag, err error) { 
		err = database.Get(&r, `SELECT * FROM post_tag WHERE tag_content = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PostTagCollector) _b(){ 
		_ = time.Now()
}

func MapPostTagInputToPostTag(a PostTagInput, e PostTag) PostTag { 
 if a.PostId != nil { 
	e.PostId = *a.PostId 
 }
 if a.TagContent != nil { 
	e.TagContent = *a.TagContent 
 }
  return e
}
