package server

import "time"

// type Server struct {
// 	Resources     *resources.Resources
// 	rotatorServer *rotator.Server
// }

// func NewServer(r *rotator.Rotator, res *resources.Resources) *Server {
// 	s := &Server{
// 		Resources: res,
// 	}
// 	s.rotatorServer = rotator.NewRotatorServer(r)

// 	return s
// }

// func (s *Server) Run(ctx context.Context) error {
// 	eg, ctx := errgroup.WithContext(ctx)
// 	eg.Go(func() error {
// 		return s.grpcInit(ctx)
// 	})

// 	return eg.Wait()
// }

// func (s *Server) grpcInit(ctx context.Context) error {
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterRotatorServer(grpcServer, s.rotatorServer)

// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Resources.Config.GRPCPort))
// 	if err != nil {
// 		log.Fatal("unable to create grpc listener:", err)
// 	}

// 	go func() {
// 		<-ctx.Done()
// 		go func() {
// 			<-ctx.Done()
// 			log.Println("The grpc-server is shutting down...")
// 			grpcServer.GracefulStop()
// 			log.Println("The grpc-server was successfully stopped")
// 		}()
// 		<-ctx.Done()
// 		time.Sleep(time.Second * 10)
// 		grpcServer.Stop()
// 		log.Println("The grpc-server was successfully stopped")
// 	}()

// 	return grpcServer.Serve(listener)
// }

const (
	gracefulTimeout        = 2 * time.Second
	logDBLongQueryDuration = 1 * time.Second
)
