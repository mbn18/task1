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
	// The driver should be kept somewhere else, possibly as a struct with repo methods. Preferable not as global variable.
	driver := initNeo4J(ctx)
	defer driver.Close(ctx)

	// In real scenario it is recommended to create a new session per received message. For the given poc we will use 1.
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// This mimic how message will be handled after pulled from the topic
	for i := 10; i > 0; i-- {
		host := datagen.Generate(i)
		err := mapper.Upsert(ctx, session, host)
		if err != nil {
			log.Fatal("Failed to upsert data:", err)
			// dispatch the message to a dead-letter topic & return
		}
		//	commit the kafka message
	}
}

// In a proper micro-service the configuration should be supplied using the ENV keys. And then constructed in a Config struct upon initialization
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
