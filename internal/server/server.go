package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/dragun-igor/banner-rotation/internal/resources"
	"github.com/dragun-igor/banner-rotation/internal/rotator"
	"github.com/dragun-igor/banner-rotation/pkg/pb"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	Resources     *resources.Resources
	rotatorServer *rotator.Server
}

func NewServer(r *rotator.Rotator, res *resources.Resources) *Server {
	s := &Server{
		Resources: res,
	}
	s.rotatorServer = rotator.NewRotatorServer(r)

	return s
}

func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return s.grpcInit(ctx)
	})

	return eg.Wait()
}

func (s *Server) grpcInit(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Resources.Config.GRPCPort))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRotatorServer(grpcServer, s.rotatorServer)

	go func() {
		go func() {
			<-ctx.Done()
			log.Info().Msg("The grpc-server is shutting down...")
			grpcServer.GracefulStop()
			log.Info().Msg("The grpc-server was successfully stopped")
		}()
		<-ctx.Done()
		time.Sleep(time.Second * 10)
		grpcServer.Stop()
		log.Info().Msg("The grpc-server was successfully stopped")
	}()

	return grpcServer.Serve(lis)
}
