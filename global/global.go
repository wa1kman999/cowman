package global

import (
	"github.com/wa1kman999/cowman/config"
	"gorm.io/gorm"
)

var (
	GBConfig config.Config
	GBMysql  *gorm.DB
)
