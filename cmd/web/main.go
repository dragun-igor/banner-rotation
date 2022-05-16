package main

import (
	"context"

	"github.com/dragun-igor/banner-rotation/internal/resources"
	"github.com/dragun-igor/banner-rotation/internal/rotator"
	"github.com/dragun-igor/banner-rotation/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Context(context.Background())
	res := resources.GetResources(ctx)

	rr := createRotator(res)
	initServer(ctx, rr, res)
}

func initServer(ctx context.Context, r *rotator.Rotator, res *resources.Resources) {
	s := server.NewServer(
		r,
		res,
	)
	log.Fatal().Err(s.Run(ctx))
}

func createRotator(res *resources.Resources) *rotator.Rotator {
	return rotator.NewRotator(res)
}
