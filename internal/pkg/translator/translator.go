package translator

import (
	"duval/internal/utils"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Translator struct {
	Id               uint         `json:"id"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        *time.Time   `json:"deleted_at"`
	Msg              string       `json:"msg"`
	Number           int          `json:"number"`
	Language         uint         `json:"language"`
	MenuParentNumber int          `json:"menu_parent_number"`
	Items            []Translator `json:"items" q:"_"`
}

func GetAllTranslation(ctx *gin.Context) {
	var (
		err         error
		translators []Translator
	)

	err = database.Select(&translators, `SELECT * FROM translator ORDER BY number desc `)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, translators)
	return
}

func NewTranslation(ctx *gin.Context) {
	var (
		err         error
		translation Translator
	)

	err = ctx.ShouldBindJSON(&translation)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if translation.Number == state.ZERO {
		translation.Number = generateAUniqueTranslatorNumber()
	}

	_, err = database.InsertOne(translation)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, translation)
	return
}

func GetTranslation(ctx *gin.Context) {
	var (
		err        error
		translator Translator
		language   int
		nb         int
	)

	language, err = strconv.Atoi(ctx.Param("language"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	nb, err = strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	translator, err = getTranslator(nb, language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, translator)
	return
}

func DeleteTranslation(ctx *gin.Context) {
	var (
		err        error
		translator Translator
		language   int
		nb         int
	)

	language, err = strconv.Atoi(ctx.Param("language"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	nb, err = strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Get(&translator, `SELECT * FROM translator WHERE number = ? AND language = ?`, nb, language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	err = database.Delete(translator)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func UpdateTranslation(ctx *gin.Context) {
	var (
		err        error
		translator Translator
	)

	err = ctx.ShouldBindJSON(translator)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if translator.Id == state.ZERO || translator.Number == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errors.New("id or number must be provided"),
		})
		return
	}

	err = database.Update(translator)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, translator)
	return
}

func generateAUniqueTranslatorNumber() (number int) {
	var lastNumber int
	var err error

	err = database.Get(&lastNumber, `SELECT number from translator ORDER BY  number desc  limit 1`)
	if err != nil {
		return state.ZERO
	}

	return lastNumber + 1
}

func getTranslator(number, language int) (translator Translator, err error) {
	Q := `SELECT COALESCE(t1.id, t.id) as 'id',
				 COALESCE(t1.msg, t.msg) as 'msg',
				 COALESCE(t1.number, t.number) as 'number',
				 COALESCE(t1.menu_parent_number, t.menu_parent_number) as 'menu_parent_number',
				 COALESCE(t1.language, t.language) as 'language'
			FROM translator t 
    			LEFT JOIN translator t1 ON t.number = t1.number 
    		WHERE t.number = ? AND t.language = 0 AND t1.language = ?`

	err = database.Get(&translator, Q, number, language)
	if err != nil {
		return translator, err
	}

	//query all message item
	translator.Items, err = getTranslatorItems(translator.Number, language)
	if err != nil {
		return translator, err
	}

	return translator, err
}

func getTranslatorItems(parentNumber, language int) (items []Translator, err error) {
	Q := `SELECT COALESCE(t1.id, t.id) as 'id',
				 COALESCE(t1.msg, t.msg) as 'msg',
				 COALESCE(t1.number, t.number) as 'number',
				 COALESCE(t1.menu_parent_number, t.menu_parent_number) as 'menu_parent_number',
				 COALESCE(t1.language, t.language) as 'language'
			FROM translator t 
    			LEFT JOIN translator t1 ON t.number = t1.number 
    		WHERE t.menu_parent_number = ? AND t.language = 0 AND t1.language = ?`

	err = database.Select(&items, Q, parentNumber, language)
	if err != nil {
		return nil, err
	}

	return items, err
}
