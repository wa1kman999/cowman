package global

import (
	"github.com/wa1kman999/cowman/config"
	"gorm.io/gorm"
)

var (
	CMConfig config.Config
	CMMysql  *gorm.DB
)
