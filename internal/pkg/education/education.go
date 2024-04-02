package education

import (
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Education struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
}

type Subject struct {
	Id               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	EducationLevelId uint       `json:"education_level_id"`
	Name             string     `json:"name"`
}

func GetSubjects(ctx *gin.Context) {
	var (
		err      error
		subjects []Subject
		eduId    int
	)

	eduId, err = strconv.Atoi(ctx.Param("edu"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to retrieve params",
		})
	}

	err = database.Select(&subjects, `SELECT * FROM subject WHERE education_level_id = ?`, eduId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, subjects)
	return
}

func GetEducation(ctx *gin.Context) {
	var (
		err  error
		edus []Education
	)

	err = database.Select(&edus, `SELECT * FROM education WHERE id > 0 ORDER BY  created_at`)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, edus)
	return
}
