package resolver

import (
	"context"

	"github.com/private-project-pp/graphql-api-service/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) (out []model.Todo, err error) {
	out = append(out, model.Todo{
		ID:   "1",
		Text: "Ini teks",
	})
	return out, nil
}
