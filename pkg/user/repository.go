package user

import (
	"github.com/jinzhu/gorm"
	"github.com/rithikjain/quickscan-backend/pkg"
	"github.com/rithikjain/quickscan-backend/pkg/entities"
)

type Repository interface {
	FindByID(id float64) (*entities.User, error)

	FindByEmail(email string) (*entities.User, error)

	FindByUUID(uuid string) (*entities.User, error)

	Register(user *entities.User) (*entities.User, error)

	DoesEmailExist(email string) (bool, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindByID(id float64) (*entities.User, error) {
	user := &entities.User{}
	r.DB.Where("id = ?", id).First(user)
	if user.Email == "" {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) Register(user *entities.User) (*entities.User, error) {
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) DoesEmailExist(email string) (bool, error) {
	user := &entities.User{}
	if r.DB.Where("email = ?", email).First(user).RecordNotFound() {
		return false, nil
	}
	return true, nil
}

func (r *repo) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	result := r.DB.Where("email = ?", email).First(user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) FindByUUID(uuid string) (*entities.User, error) {
	user := &entities.User{}
	result := r.DB.Where("uuid = ?", uuid).First(user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}
