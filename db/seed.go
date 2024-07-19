package db

import (
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/id"
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
			ID:       id.New(),
			ActionID: "read-team",
			RoleID:   "admin",
		},
		{
			ID:       id.New(),
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
