package db

import (
	"github.com/ukasyah-dev/authority-service/model"
	"gorm.io/gorm/clause"
)

func seed() {
	actions := []model.Action{
		{
			ID:   "read-team",
			Name: "Read team",
		},
		{
			ID:   "write-team",
			Name: "Write team",
		},
	}

	roles := []model.Role{
		{
			ID:   "admin",
			Name: "Admin",
		},
	}

	permissions := []model.Permission{
		{
			ActionID: "read-team",
			RoleID:   "admin",
		},
		{
			ActionID: "write-team",
			RoleID:   "admin",
		},
	}

	create(actions)
	create(roles)
	create(permissions)
}

func create(data any) {
	err := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(data).Error
	if err != nil {
		panic(err)
	}
}
