package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type UserAcademicCourseCollector struct {} 
func (c *UserAcademicCourseCollector) CourseId(a *int) (r UserAcademicCourse, err error) { 
		err = database.Get(&r, `SELECT * FROM user_academic_course WHERE course_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *UserAcademicCourseCollector) _b(){ 
		_ = time.Now()
}

func MapUserAcademicCourseInputToUserAcademicCourse(a UserAcademicCourseInput, e UserAcademicCourse) UserAcademicCourse { 
 if a.CourseId != nil { 
	e.CourseId = *a.CourseId 
 }
  return e
}
