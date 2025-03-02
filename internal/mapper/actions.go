package mapper

import (
	"context"
	"github.com/mbn18/dream/internal/entity"
	"github.com/mbn18/dream/internal/mapper/trunk"
	"github.com/mbn18/dream/internal/mapper/upsert"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Upsert(ctx context.Context, session neo4j.SessionWithContext, host *entity.Host) error {
	return upsert.Do(ctx, session, host)
}

func Trunk(ctx context.Context, session neo4j.SessionWithContext) error {
	return trunk.Do(ctx, session)
}
