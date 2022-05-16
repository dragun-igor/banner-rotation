package rotator

import (
	"context"
	"errors"

	"github.com/dragun-igor/banner-rotation/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	rotator *Rotator
	pb.UnimplementedRotatorServer
}

func NewRotatorServer(rotator *Rotator) *Server {
	return &Server{rotator: rotator}
}

// func (s Server) AddBanner(ctx context.Context, request *pb.AddBanner) error {
// 	err := s.rotator.AddBanner(request.Description)
// 	return handleErr(err)
// }

// func (s Server) AddSlot(ctx context.Context, request *pb.AddSlot) error {
// 	err := s.rotator.AddSlot(request.Description)
// 	return handleErr(err)
// }

// func (s Server) AddUserGroup(ctx context.Context, request *pb.AddUserGroup) error {
// 	err := s.rotator.AddUserGroup(request.Description)
// 	return handleErr(err)
// }

func (s Server) SelectBanner(ctx context.Context, request *pb.SelectBannerRequest) (*pb.SelectBannerResponse, error) {
	res, err := s.rotator.SelectBanner(int(request.SlotId), int(request.UserGroupId))
	err = handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.SelectBannerResponse{BannerId: uint32(res)}, nil
}

func (s Server) AddBannerToSlot(ctx context.Context, request *pb.AddBannerToSlotRequest) (*pb.AddBannerToSlotResponse, error) {
	err := s.rotator.AddBannerToSlot(int(request.BannerId), int(request.SlotId))
	err = handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.AddBannerToSlotResponse{}, nil
}

func (s Server) RemoveBannerFromSlot(ctx context.Context, request *pb.RemoveBannerFromSlotRequest) (*pb.RemoveBannerFromSlotResponse, error) {
	err := s.rotator.RemoveBannerFromSlot(int(request.BannerId), int(request.SlotId))
	err = handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.RemoveBannerFromSlotResponse{}, nil
}

func (s Server) RotatorServer() {

}

func handleErr(err error) error {
	var errNotFound *NotFoundError

	if err == nil {
		return nil
	}

	if errors.As(err, &errNotFound) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
