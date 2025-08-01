package common_models

import (
	"fmt"
	"strings"

	master_models "github.com/parampreetr/layered_arch_ai/internal/common/models/master"
	"gorm.io/gorm"
)

func IgnoreExistError(err error) error {
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		// Only panic if it's NOT an "already exists" error
		return err
	}

	fmt.Printf("Ignored Exist error: %v\n", err)
	return nil
}

func RunCommonMigrations(db *gorm.DB) error {

	db.Exec(`CREATE SCHEMA IF NOT EXISTS stl`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS master`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS pao`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS bank`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS spc`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS session`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS sub`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS sec`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS acc`)
	db.Exec(`CREATE SCHEMA IF NOT EXISTS icra`)

	var err []error
	if errMigrate := db.AutoMigrate(&master_models.PraoMst{}); errMigrate != nil {
		//append table name and error
		err = append(err, fmt.Errorf("table: %s, error: %v", (master_models.PraoMst{}).TableName(), errMigrate))
	}

	db.Exec(`
  DO $$
  DECLARE
    r RECORD;
  BEGIN
    FOR r IN
      SELECT schemaname, tablename
      FROM pg_tables
      WHERE schemaname NOT IN ('pg_catalog', 'information_schema')
    LOOP
      EXECUTE format('ALTER TABLE %I.%I REPLICA IDENTITY FULL;', r.schemaname, r.tablename);
    END LOOP;
  END
  $$;
`)

	// err = append(err, MigrationFunctionForAnyModule(db))

	for _, e := range err {
		if e != nil {
			fmt.Printf("ERROR: %v", e)
		}
	}

	for _, e := range err {
		if e != nil {
			return e
		}
	}

	return nil
}
