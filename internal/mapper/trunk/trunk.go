package trunk

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const query = "MATCH (n) DETACH DELETE (n)"

func Do(ctx context.Context, session neo4j.SessionWithContext) error {
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		_, err := tx.Run(ctx, query, map[string]interface{}{})
		return nil, err

	})
	return err

}
