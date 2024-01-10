package migration

import (
	"log"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/storage"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

func RunMigrateGorm() {
	dbGorm := storage.NewDBGormPostgreSQL()
	err := dbGorm.Connect()
	if err != nil {
		log.Printf("%sError to connect db to migrate. Error: %s at %s%s\n", colorconst.ANSIColorRed, err, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}

	err = dbGorm.DB.AutoMigrate(
		entities.User{},
	)
	if err != nil {
		log.Printf("%sError run migrate. Error: %s at %s%s\n", colorconst.ANSIColorRed, err, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
		return
	}

	log.Printf("%sRun migrate successfully at %s%s\n", colorconst.ANSIColorPurple, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
}
