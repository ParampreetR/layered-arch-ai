package common_di

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	common_config "github.com/parampreetr/layered_arch_ai/internal/common/config"
	common_models "github.com/parampreetr/layered_arch_ai/internal/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Container *ProjectDependencies = nil

func InitDB(cfg *common_config.Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db = db.Debug()

	// Add OpenTelemetry tracing
	// if err := db.Use(tracing.NewPlugin()); err != nil {
	// 	return nil, fmt.Errorf("failed to setup database tracing: %w", err)
	// }

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func initializeDeps(cfg *common_config.Config, router fiber.Router) error {

	envVars := common_config.Load()

	fmt.Println(envVars)
	db, err := InitDB(cfg)
	if err != nil {
		return err
	}

	common_models.RunCommonMigrations(db)

	// Provide DB interface
	Container.DatabaseInstance = db

	return nil
}

func Init(cfg *common_config.Config, router fiber.Router) {
	if err := initializeDeps(cfg, router); err != nil {
		log.Fatalf("error while initializing app dependencies: %v\n", err)
	}

	// common_router.SetupRouter(router, cfg)
}
