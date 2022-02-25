package repository

import (
	"service/internal/config"
	"service/internal/infrastructure/database"
	"service/internal/model"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Config *config.Configuration
	Db     *database.DB
}

func NewUserRepository(
	cf *config.Configuration,
	db *database.DB,
) *UserRepository {
	return &UserRepository{
		Config: cf,
		Db: db,
	}
}

func (s *UserRepository) CreateUser(users *model.Users) (bool, error) {
	result := s.Db.Connection.Create(users) // pass pointer of data to Create
	if result != nil {
		return true, nil
	}
	return false, nil
}

func (s *UserRepository) CreateOrUpdateUser(user *model.Users) error {
	if err := s.Db.Connection.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"username": user.Username,
			"mobile_no": user.MobileNo,
		}),
	}).Create(&user).Error; err != nil {
		return err
	}
	
	return nil
}