package model

import (
	"github.com/cend-org/duval/internal/database"
	"time"
)

type UserAcademicCoursePreferenceCollector struct{}

func (c *UserAcademicCoursePreferenceCollector) UserAcademicCourseId(a *int) (r UserAcademicCoursePreference, err error) {
	err = database.Get(&r, `SELECT * FROM user_academic_course_preference WHERE user_academic_course_id = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserAcademicCoursePreferenceCollector) IsOnline(a *bool) (r UserAcademicCoursePreference, err error) {
	err = database.Get(&r, `SELECT * FROM user_academic_course_preference WHERE is_online = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserAcademicCoursePreferenceCollector) Availability(a *time.Time) (r UserAcademicCoursePreference, err error) {
	err = database.Get(&r, `SELECT * FROM user_academic_course_preference WHERE availability = ? ORDER BY created_at DESC LIMIT 1`, a)
	return r, err
}

func (c *UserAcademicCoursePreferenceCollector) _b() {
	_ = time.Now()
}

func MapUserAcademicCoursePreferenceInputToUserAcademicCoursePreference(a UserAcademicCoursePreferenceInput, e UserAcademicCoursePreference) UserAcademicCoursePreference {
	if a.UserAcademicCourseId != nil {
		e.UserAcademicCourseId = *a.UserAcademicCourseId
	}
	if a.IsOnline != nil {
		e.IsOnline = *a.IsOnline
	}
	if a.Availability != nil {
		e.Availability = *a.Availability
	}
	return e
}
