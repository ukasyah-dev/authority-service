package db

import (
	"os"

	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/db/pool"
	identityModel "github.com/ukasyah-dev/identity-service/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() {
	var err error

	DB, err = pool.Open(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(
		&identityModel.User{},
		&model.Team{},
	)
	if err != nil {
		panic(err)
	}
}

func Close() error {
	sql, _ := DB.DB()
	return sql.Close()
}
