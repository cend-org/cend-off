package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type SubjectCollector struct {} 
func (c *SubjectCollector) EducationLevelId(a *int) (r Subject, err error) { 
		err = database.Get(&r, `SELECT * FROM subject WHERE education_level_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *SubjectCollector) Name(a *string) (r Subject, err error) { 
		err = database.Get(&r, `SELECT * FROM subject WHERE name = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *SubjectCollector) _b(){ 
		_ = time.Now()
}

func MapSubjectInputToSubject(a SubjectInput, e Subject) Subject { 
 if a.EducationLevelId != nil { 
	e.EducationLevelId = *a.EducationLevelId 
 }
 if a.Name != nil { 
	e.Name = *a.Name 
 }
  return e
}
