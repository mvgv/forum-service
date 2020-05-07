package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mvgv/forum-service/validation"

	"github.com/mvgv/forum-service/service"
)

/*GetPosts é reponsavel por recuperar todos os posts contidos em um topico*/
func GetPosts(context *gin.Context) {
	topicID, err := strconv.ParseUint(context.Param("topic_id"), 10, 64)

	if err != nil {
		topicErr := &validation.ApplicationError{
			Message:    "topic_id must be a number",
			StatusCode: http.StatusBadRequest,
		}
		context.JSON(topicErr.StatusCode, topicErr)
		return
	}

	topic, topicErr := service.PostServiceImpl.GetPosts(topicID)
	if topicErr != nil {
		context.JSON(topicErr.StatusCode, topicErr)
		return
	}

	context.JSON(http.StatusOK, topic)
}
