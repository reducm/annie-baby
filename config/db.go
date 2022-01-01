package config

import (
	"annie-baby/ent"
	"context"

	_ "github.com/lib/pq"

	"github.com/kr/pretty"
)

var DBClient *ent.Client

func init() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=scott password=tiger dbname=annie_baby sslmode=disable")
	if err != nil {
		pretty.Printf("db err: %v", err)
		panic("DB Connect error")
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		pretty.Errorf("failed creating schema resources: %v", err)
	}

	pretty.Logf("client db ok: %v", *client)

	DBClient = client
}

func GetDB() *ent.Client {
	return DBClient
}
