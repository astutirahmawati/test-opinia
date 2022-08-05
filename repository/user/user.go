package user

import (
	"errors"
	"opinia/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewrUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (ur *UserRepo) Login(email string, password string) (entities.User, error) {
	users := []entities.User{}

	if err := ur.Db.Where("email = ? AND password = ?", email, password).First(&users).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot select data")
	}

	return users[0], nil
}

func (ur *UserRepo) InsertUser(newUser entities.User) (entities.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot insert data")
	}

	log.Info()
	return newUser, nil
}

func (ur *UserRepo) GetUserID(ID uint) (entities.User, error) {
	arrUser := []entities.User{}

	if err := ur.Db.Where("id = ?", ID).Find(&arrUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot select data")
	}

	if len(arrUser) == 0 {
		log.Warn("data not found")
		return entities.User{}, errors.New("not found")
	}

	log.Info()
	return arrUser[0], nil
}

func (ur *UserRepo) FindByUser(value string) (entities.User, error) {
	user := entities.User{}
	tx := ur.Db.Where("email = ?", value).First(&user)
	if tx.Error != nil {
		return entities.User{}, tx.Error
	}
	return user, nil
}

func (ur *UserRepo) UpdateUser(ID uint, update entities.User) (entities.User, error) {
	var res entities.User
	if err := ur.Db.Where("id = ?", ID).Updates(&update).Find(&res).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot update")
	}

	log.Info()
	return res, nil
}

func (ur *UserRepo) DeleteUser(id uint) error {

	// Delete from database
	tx := ur.Db.Delete(&entities.User{}, id)
	if tx.Error != nil {

		// return kode 500 jika error
		return tx.Error
	}
	return nil
}
