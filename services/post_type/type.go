package post_type

import (
	"opinia/entities"
	typeRepository "opinia/repository/post_type"
)

type TypeService struct {
	typeRepo typeRepository.TypeRepositoryInterface
}

func NewTypeService(repository typeRepository.TypeRepositoryInterface) *TypeService {
	return &TypeService{
		typeRepo: repository,
	}
}

func (ts *TypeService) CreateType(user_id uint, request entities.PostType) (entities.PostType, error) {

	// Konversi user request menjadi domain untuk diteruskan ke repository
	// types := entities.PostType{}
	// copier.Copy(&user, &internalRequest)

	// Insert ke sistem melewati repository
	respond, err := ts.typeRepo.InsertType(request)
	if err != nil {
		return entities.PostType{}, err
	}

	return respond, nil
}

func (ts *TypeService) GetbyID(id uint) (entities.PostType, error) {

	typee, err := ts.typeRepo.GetTypeID(id)
	if err != nil {
		return entities.PostType{}, err
	}

	return typee, err
}

func (ts *TypeService) UpdateType(type_id uint, request entities.PostType) (entities.PostType, error) {

	// Get user by ID via repository
	update, err := ts.typeRepo.UpdateType(type_id, request)
	if err != nil {
		return entities.PostType{}, err
	}

	return update, err
}

func (ts *TypeService) DeleteType(user_id, type_id uint) error {

	// Delete via repository
	err := ts.typeRepo.DeleteType(type_id)
	return err
}
