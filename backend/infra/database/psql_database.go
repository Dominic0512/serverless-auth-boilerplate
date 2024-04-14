package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/config"
)

type PSQLDatabase struct {
	Client *ent.Client
}

func NewPSQLDatabase(config *config.Config) (*PSQLDatabase, error) {
	driver := config.DBDriver
	username := config.DBUsername
	password := config.DBPassword
	host := config.DBHost
	port := config.DBPort
	dbname := config.DBName

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, username, dbname, password)

	client, err := ent.Open(driver, dataSource)
	if err != nil {
		log.Fatalf("The ent open database connection error : %s", err)
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("The ent background auto-migration error : %s", err)
		return nil, err
	}

	return &PSQLDatabase{
		Client: client,
	}, nil
}
