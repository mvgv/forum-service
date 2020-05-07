package service

import (
	"github.com/mvgv/forum-service/model"
	"github.com/mvgv/forum-service/validation"
)

/*PostService é a interface responsavel pela abstracao das operacoes sobre um post do forum*/
type PostService interface {
	GetPosts(topicID uint64) (*[]model.Post, *validation.ApplicationError)
}
