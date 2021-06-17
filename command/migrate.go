package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.ebupt.com/lets/app"
	"gotutorial/database/Seeds"
	"gotutorial/model"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "运行数据迁移服务",
	Run: func(cmd *cobra.Command, args []string) {
		startMigrate()
	},
}

/*
	数据迁移和初始数据Seed
*/
func startMigrate() {
	fmt.Println("数据迁移 && Seed")
	app.LDB.AutoMigrate(&model.Admin{}, &model.Role{}, &model.Permission{})
	Seeds.InitSeed()
	fmt.Println("Successful migration!")
}
