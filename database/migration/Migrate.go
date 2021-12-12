package migration

import (
	"github.com/teoferizovic/senator/database"
	"github.com/teoferizovic/senator/model"
)

func Migrate()  {
	database.DBCon.AutoMigrate(&model.User{}, &model.Article{})
}