package upsert

import (
	"context"
	"github.com/mbn18/dream/internal/entity"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"time"
)

func Do(ctx context.Context, session neo4j.SessionWithContext, host *entity.Host) error {

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		params := extractQueryParams(host)
		params["createdAt"] = host.Processes.CreatedAt.UTC().Format(time.RFC3339)

		q, err := queryBuilder(params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), q, params)
		return nil, err

	})
	return err
}
