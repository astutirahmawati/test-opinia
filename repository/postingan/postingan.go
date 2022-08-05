package postingan

import (
	"errors"
	"opinia/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type PostRepo struct {
	Db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		Db: db,
	}
}

func (pr *PostRepo) InsertPost(newPost entities.Postingan) (entities.Postingan, error) {
	if err := pr.Db.Create(&newPost).Error; err != nil {
		log.Warn(err)
		return entities.Postingan{}, errors.New("cannot insert data")
	}

	log.Info()
	return newPost, nil
}

func (pr *PostRepo) GetPostID(ID uint) (entities.Postingan, error) {
	arrPost := []entities.Postingan{}

	if err := pr.Db.Where("id = ?", ID).Find(&arrPost).Error; err != nil {
		log.Warn(err)
		return entities.Postingan{}, errors.New("cannot select data")
	}

	if len(arrPost) == 0 {
		log.Warn("data not found")
		return entities.Postingan{}, errors.New("not found")
	}

	log.Info()
	return arrPost[0], nil
}

func (pr *PostRepo) UpdatePost(ID uint, update entities.Postingan) (entities.Postingan, error) {
	var res entities.Postingan
	if err := pr.Db.Where("id = ?", ID).Updates(&update).Find(&res).Error; err != nil {
		log.Warn(err)
		return entities.Postingan{}, errors.New("cannot update")
	}

	log.Info()
	return res, nil
}

func (pr *PostRepo) DeletePost(id uint) error {

	// Delete from database
	tx := pr.Db.Delete(&entities.Postingan{}, id)
	if tx.Error != nil {

		// return kode 500 jika error
		return tx.Error
	}
	return nil
}
