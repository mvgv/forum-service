package repository

import (
	"net/http"
	"time"

	"github.com/mvgv/forum-service/model"
	"github.com/mvgv/forum-service/validation"
)

var (
	firstPost = model.Post{123, model.User{}, time.Now(), time.Now(), "lorem ipsum"}

	secondPost = model.Post{125, model.User{}, time.Now(), time.Now(), "lorem amet"}

	selectedTopic = map[uint64]*[]model.Post{
		123: &[]model.Post{firstPost, secondPost},
	}

	/*PostRepositoryImpl expoe a implementacao do repository para outros pacotes*/
	PostRepositoryImpl PostRepository
)

func init() {
	PostRepositoryImpl = &postRepositoryImpl{}
}

type postRepositoryImpl struct{}

/*GetPostsByTopicID implementa busca da lista de posts associados a um topico pelo seu id*/
func (postRepositoryImpl *postRepositoryImpl) GetPostsByTopicID(topicID uint64) (*[]model.Post, *validation.ApplicationError) {
	selectedPostList := selectedTopic[topicID]
	if selectedPostList == nil {
		return nil, &validation.ApplicationError{
			Message:    "Topic not found",
			StatusCode: http.StatusNotFound,
		}
	}
	return selectedPostList, nil
}
