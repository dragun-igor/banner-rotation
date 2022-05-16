package resources

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dragun-igor/banner-rotation/internal/config"

	_ "github.com/lib/pq"
)

type Resources struct {
	DB     *sql.DB
	Config *config.Config
}

func GetResources(ctx context.Context) *Resources {
	config := config.New()
	return &Resources{
		DB:     ConnectDB(*config, ctx),
		Config: config,
	}
}

func ConnectDB(config config.Config, ctx context.Context) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully connected to database!")
	go func() {
		<-ctx.Done()
		_ = db.Close()
		fmt.Println("Connection to database has closed!")
	}()
	return db
}
