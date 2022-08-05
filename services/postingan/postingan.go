package postingan

import (
	"opinia/entities"
	postRepository "opinia/repository/postingan"
)

type PostService struct {
	postRepo postRepository.PostinganRepositoryInterface
}

func NewPostService(repository postRepository.PostinganRepositoryInterface) *PostService {
	return &PostService{
		postRepo: repository,
	}
}

func (ps *PostService) CreatePost(user_id uint, request entities.Postingan) (entities.Postingan, error) {

	respond, err := ps.postRepo.InsertPost(request)
	if err != nil {
		return entities.Postingan{}, err
	}

	return respond, nil
}

func (ps *PostService) GetbyID(id uint) (entities.Postingan, error) {

	get, err := ps.postRepo.GetPostID(id)
	if err != nil {
		return entities.Postingan{}, err
	}

	return get, err
}

func (ps *PostService) UpdatePost(user_id uint, post_id uint, request entities.Postingan) (entities.Postingan, error) {

	// Get user by ID via repository
	update, err := ps.postRepo.UpdatePost(post_id, request)
	if err != nil {
		return entities.Postingan{}, err
	}

	return update, err
}

func (ps *PostService) DeletePost(user_id, post_id uint) error {

	// Delete via repository
	err := ps.postRepo.DeletePost(post_id)
	return err
}
