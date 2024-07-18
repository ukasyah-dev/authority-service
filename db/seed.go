package db

import (
	"github.com/ukasyah-dev/authority-service/model"
	"gorm.io/gorm/clause"
)

func seed() error {
	err := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&[]model.Role{
		{
			ID:   "admin",
			Name: "Admin",
		},
	}).Error
	if err != nil {
		return err
	}

	return nil
}
