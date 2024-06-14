package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type UserAcademicCoursePreferenceCollector struct {} 
func (c *UserAcademicCoursePreferenceCollector) IsOnline(a *bool) (r UserAcademicCoursePreference, err error) { 
		err = database.Get(&r, `SELECT * FROM user_academic_course_preference WHERE is_online = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *UserAcademicCoursePreferenceCollector) _b(){ 
		_ = time.Now()
}

func MapUserAcademicCoursePreferenceInputToUserAcademicCoursePreference(a UserAcademicCoursePreferenceInput, e UserAcademicCoursePreference) UserAcademicCoursePreference { 
 if a.IsOnline != nil { 
	e.IsOnline = *a.IsOnline 
 }
  return e
}
