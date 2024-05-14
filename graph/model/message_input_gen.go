package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type MessageCollector struct {} 
func (c *MessageCollector) ResourceType(a *int) (r Message, err error) { 
		err = database.Get(&r, `SELECT * FROM message WHERE resource_type = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MessageCollector) ResourceValue(a *int) (r Message, err error) { 
		err = database.Get(&r, `SELECT * FROM message WHERE resource_value = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MessageCollector) ResourceNumber(a *int) (r Message, err error) { 
		err = database.Get(&r, `SELECT * FROM message WHERE resource_number = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MessageCollector) ResourceLabel(a *string) (r Message, err error) { 
		err = database.Get(&r, `SELECT * FROM message WHERE resource_label = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MessageCollector) ResourceLanguage(a *int) (r Message, err error) { 
		err = database.Get(&r, `SELECT * FROM message WHERE resource_language = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *MessageCollector) _b(){ 
		_ = time.Now()
}

func MapMessageInputToMessage(a MessageInput, e Message) Message { 
 if a.ResourceType != nil { 
	e.ResourceType = *a.ResourceType 
 }
 if a.ResourceValue != nil { 
	e.ResourceValue = *a.ResourceValue 
 }
 if a.ResourceNumber != nil { 
	e.ResourceNumber = *a.ResourceNumber 
 }
 if a.ResourceLabel != nil { 
	e.ResourceLabel = *a.ResourceLabel 
 }
 if a.ResourceLanguage != nil { 
	e.ResourceLanguage = *a.ResourceLanguage 
 }
  return e
}
