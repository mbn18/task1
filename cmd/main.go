package main

import (
	"context"
	"github.com/mbn18/dream/datagen"
	"github.com/mbn18/dream/internal/mapper"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func main() {
	ctx := context.Background()
	driver := initNeo4J(ctx)
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	for i := 0; i < 10; i++ {
		host := datagen.Generate(1)
		err := mapper.Upsert(ctx, session, host)
		if err != nil {
			log.Fatal("Failed to upsert data:", err)
		}
	}
}

func initNeo4J(ctx context.Context) neo4j.DriverWithContext {
	driver, err := neo4j.NewDriverWithContext("neo4j://localhost", neo4j.NoAuth())
	if err != nil {
		log.Fatal(err)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal("Failed to connect to Neo4j:", err)
	}
	log.Println("Connected to Neo4j successfully")
	return driver
}
