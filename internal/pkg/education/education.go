package education

import (
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Education struct {
	Id           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	LevelOfStudy string     `json:"level_of_study"`
}

var levels = []string{
	"Primaire 1",
	"Primaire 2",
	"Primaire 3",
	"Primaire 4",
	"Primaire 5",
	"Primaire 6",
	"Secondaire 1",
	"Secondaire 2",
	"Secondaire 3",
	"Secondaire 4",
	"Secondaire 5",
	"Cégep",
	"Universités",
}

func InsertLevelOfStudy(ctx *gin.Context) {
	var (
		education []Education
		err       error
	)
	for _, level := range levels {
		education = append(education, Education{
			LevelOfStudy: level,
		})
	}
	err = InsertEducations(education)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}

/*
	UTILS
*/

func InsertEducations(educations []Education) (err error) {
	for _, education := range educations {
		_, err := database.InsertOne(education)
		if err != nil {
			return err
		}
	}
	return err
}
