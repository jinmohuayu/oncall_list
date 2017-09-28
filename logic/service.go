package logic

import (
	"database/sql"
	"log"

	"git.elenet.me/DA/oncall_list/data/mysql"
)

// ServiceConfig 权限服务配置
type ServiceConfig struct {
	DB *sql.DB
}

// Service 权限服务
type Service struct {
	db             *sql.DB
	config         ServiceConfig
	logger         *log.Logger
	userRepository *mysql.UserRepository
}

// NewService 新建权限服务
func NewService(config ServiceConfig, logger *log.Logger) *Service {
	return &Service{
		db:             config.DB,
		config:         config,
		logger:         logger,
		userRepository: mysql.NewUserRepository(config.DB),
	}
}
