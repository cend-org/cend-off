package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type LanguageResourceCollector struct {} 
func (c *LanguageResourceCollector) ResourceRef(a *string) (r LanguageResource, err error) { 
		err = database.Get(&r, `SELECT * FROM language_resource WHERE resource_ref = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *LanguageResourceCollector) ResourceLanguage(a *int) (r LanguageResource, err error) { 
		err = database.Get(&r, `SELECT * FROM language_resource WHERE resource_language = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *LanguageResourceCollector) ResourceMessage(a *string) (r LanguageResource, err error) { 
		err = database.Get(&r, `SELECT * FROM language_resource WHERE resource_message = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *LanguageResourceCollector) _b(){ 
		_ = time.Now()
}

func MapLanguageResourceInputToLanguageResource(a LanguageResourceInput, e LanguageResource) LanguageResource { 
 if a.ResourceRef != nil { 
	e.ResourceRef = *a.ResourceRef 
 }
 if a.ResourceLanguage != nil { 
	e.ResourceLanguage = *a.ResourceLanguage 
 }
 if a.ResourceMessage != nil { 
	e.ResourceMessage = *a.ResourceMessage 
 }
  return e
}
