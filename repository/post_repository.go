package repository

import (
	"github.com/mvgv/forum-service/model"
	"github.com/mvgv/forum-service/validation"
)

/*PostRepository interface  que abstrai a comunicacao com a base de posts*/
type PostRepository interface {
	GetPostsByTopicID(topicID uint64) (*[]model.Post, *validation.ApplicationError)
}
