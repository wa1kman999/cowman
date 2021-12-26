package initialize

import (
	"os"

	"github.com/sirupsen/logrus"
	userModel "github.com/wa1kman999/cowman/pkg/user/model"
	"gorm.io/gorm"
)

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&userModel.User{},
	)

	if err != nil {
		logrus.Errorf("register table failed: %s", err.Error())
		os.Exit(0)
	}
	logrus.Info("register table success")
}
