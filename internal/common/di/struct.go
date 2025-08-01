package common_di

import "gorm.io/gorm"

type ProjectDependencies struct {
	DatabaseInstance *gorm.DB
}
