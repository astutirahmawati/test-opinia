package post_type

import (
	"errors"
	"opinia/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type TypeRepo struct {
	Db *gorm.DB
}

func NewTypeRepo(db *gorm.DB) *TypeRepo {
	return &TypeRepo{
		Db: db,
	}
}

func (ur *TypeRepo) InsertType(newType entities.PostType) (entities.PostType, error) {
	if err := ur.Db.Create(&newType).Error; err != nil {
		log.Warn(err)
		return entities.PostType{}, errors.New("cannot insert data")
	}

	log.Info()
	return newType, nil
}

func (ur *TypeRepo) GetTypeID(ID uint) (entities.PostType, error) {
	arrType := []entities.PostType{}

	if err := ur.Db.Where("id = ?", ID).Find(&arrType).Error; err != nil {
		log.Warn(err)
		return entities.PostType{}, errors.New("cannot select data")
	}

	if len(arrType) == 0 {
		log.Warn("data not found")
		return entities.PostType{}, errors.New("not found")
	}

	log.Info()
	return arrType[0], nil
}

func (ur *TypeRepo) UpdateType(ID uint, update entities.PostType) (entities.PostType, error) {
	var res entities.PostType
	if err := ur.Db.Where("id = ?", ID).Updates(&update).Find(&res).Error; err != nil {
		log.Warn(err)
		return entities.PostType{}, errors.New("cannot update")
	}

	log.Info()
	return res, nil
}

func (ur *TypeRepo) DeleteType(id uint) error {

	// Delete from database
	tx := ur.Db.Delete(&entities.PostType{}, id)
	if tx.Error != nil {

		// return kode 500 jika error
		return tx.Error
	}
	return nil
}
