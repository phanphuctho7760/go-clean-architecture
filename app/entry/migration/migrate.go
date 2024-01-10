package main

import (
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/migration"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/configs"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

func main() {
	// config the same .env at root, but need to separate for avoid dev mistake environment run
	configs.NewConfigEnvGlobalInstance(helpers.GetProjectRootPath())
	//Change what migrate want to run
	migration.RunMigrateGorm()
}
