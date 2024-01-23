package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/config"
	"github.com/google/wire"
)

type Database struct {
	Client *ent.Client
}

func NewDatabase(config *config.Config) (*Database, error) {
	driver := config.DBDriver
	username := config.DBUsername
	password := config.DBPassword
	host := config.DBHost
	port := config.DBPort
	dbname := config.DBName

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, username, dbname, password)

	client, err := ent.Open(driver, dataSource)
	if err != nil {
		log.Fatal("The ent open database connection error.")
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("The ent background auto-migration error.")
		return nil, err
	}

	return &Database{
		Client: client,
	}, nil
}

var ProviderSet = wire.NewSet(NewDatabase)
