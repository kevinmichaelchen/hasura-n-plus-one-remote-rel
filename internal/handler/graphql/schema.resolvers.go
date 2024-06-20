package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/kevinmichaelchen/hasura-n-plus-one-remote-rel/internal/handler/graphql/generated"
	"github.com/kevinmichaelchen/hasura-n-plus-one-remote-rel/internal/handler/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"strconv"
)

// CreateOwner is the resolver for the createOwner field.
func (r *mutationResolver) CreateOwner(ctx context.Context, input model.CreateOwnerInput) (int, error) {
	panic(fmt.Errorf("not implemented: CreateOwner - createOwner"))
}

// OwnerNickname is the resolver for the ownerNickname field.
func (r *queryResolver) OwnerNickname(ctx context.Context, ownerID int) (string, error) {
	ctx, span := trace.SpanFromContext(ctx).
		TracerProvider().
		Tracer("").
		Start(ctx, "OwnerNickname",
			trace.WithAttributes(
				attribute.String("owner_id", strconv.Itoa(ownerID)),
			),
		)
	defer span.End()

	log.Info("OwnerNickname", "ownerID", ownerID)

	return "fancyOwnerNickname", nil
}

// PetNickname is the resolver for the petNickname field.
func (r *queryResolver) PetNickname(ctx context.Context, petID int) (string, error) {
	ctx, span := trace.SpanFromContext(ctx).
		TracerProvider().
		Tracer("").
		Start(ctx, "PetNickname",
			trace.WithAttributes(
				attribute.String("pet_id", strconv.Itoa(petID)),
			),
		)
	defer span.End()

	log.Info("PetNickname", "petID", petID)

	return "fancyPetNickname", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
