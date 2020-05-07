package service

import (
	"github.com/mvgv/forum-service/model"
	"github.com/mvgv/forum-service/repository"
	"github.com/mvgv/forum-service/validation"
)

var (
	/*PostServiceImpl Variavel que expoe o post service para outros pacotes*/
	PostServiceImpl PostService
)

type postServiceImpl struct{}

func init() {
	PostServiceImpl = &postServiceImpl{}
}

/*GetPosts servico de negocio para listar posts a partir de um topico*/
func (postServiceImpl *postServiceImpl) GetPosts(topicID uint64) (*[]model.Post, *validation.ApplicationError) {
	return repository.PostRepositoryImpl.GetPostsByTopicID(topicID)
}
