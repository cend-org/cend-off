package message

import (
	"duval/internal/authentication"
	"duval/internal/utils"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Discussion struct {
	Id                  uint            `json:"id"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
	DeletedAt           *time.Time      `json:"deleted_at"`
	Name                string          `json:"name"`
	LastMessageSentDate time.Time       `json:"last_message_sent_date"`
	Actor               DiscussionActor `json:"actor" q:"_"`
	Message             Message         `json:"message" q:"_"`
}

type DiscussionActor struct {
	Id           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	DiscussionId uint       `json:"discussion_id"`
	UserId       uint       `json:"user_id"`
}

type Message struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	ActorId   uint       `json:"actor_id"`
	Content   string     `json:"content"`
	Status    uint       `json:"status"`
}

func CreateDiscussion(ctx *gin.Context) {
	var (
		err        error
		discussion Discussion
		actors     []DiscussionActor
		tok        *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	_ = tok

	err = ctx.ShouldBindJSON(actors)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	// If this is not a discussion group but a normal chat between two person
	if len(actors) == 1 {
		// check first if there is already a chat between actors[0] and the tok.user_id

	}

	ctx.JSON(http.StatusOK, discussion)
	return
}

func GetUserDiscussion(ctx *gin.Context) {
	var (
		err        error
		tok        *authentication.Token
		discussion []Discussion
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	query := `SELECT d.* 
				FROM discussion_actor da
					JOIN discussion d ON da.discussion_id = d.id
			  WHERE da.user_id = ? ORDER BY d.last_message_sent_date desc
    		`

	err = database.Client.Select(&discussion, query, tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, discussion)
	return
}
