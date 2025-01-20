package conn

import (
	"context"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgSQL struct {
	Err error
	Ctx context.Context
	Raw *gorm.DB
}

func NewPgsql(url string) (*PgSQL, error) {

	t := &PgSQL{}

	t.Raw, t.Err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if t.Err != nil {
		return nil, t.Err
	}
	return t, nil
}

func (t *PgSQL) Close() {
	sqlDB, err := t.Raw.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}
	sqlDB.Close()
}

// 连接数据库
func (t *PgSQL) connectDB(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// ping
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	t.Raw = db

}
