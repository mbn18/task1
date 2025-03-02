package main

import (
	"context"
	"fmt"
	"github.com/mbn18/dream/datagen"
	"github.com/mbn18/dream/internal/mapper/upsert"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func main() {
	driver, err := neo4j.NewDriverWithContext("neo4j://localhost", neo4j.NoAuth())
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal("Failed to connect to Neo4j:", err)
	}
	fmt.Println("Connected to Neo4j successfully")

	sCtx := context.Background()
	session := driver.NewSession(sCtx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(sCtx)

	for i := 0; i < 10; i++ {
		host := datagen.Generate(1)
		err = upsert.Do(sCtx, session, host)
		if err != nil {
			log.Fatal("Failed to upsert data:", err)
		}
	}
}
