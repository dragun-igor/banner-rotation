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
}

func NewRotatorServer(rotator *Rotator) *Server {
	return &Server{rotator: rotator}
}

func (r Server) SelectBanner(ctx context.Context, request *pb.SelectBannerRequest) (*pb.SelectBannerResponse, error) {
	res, err := r.rotator.SelectBanner(int(request.SlotId), int(request.UserGroupId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.SelectBannerResponse{BannerId: uint32(res)}, nil
}

func (r Server) AddBannerToSlot(ctx context.Context, request *pb.AddBannerToSlotRequest) (*pb.AddBannerToSlotResponse, error) {
	err := r.rotator.AddBannerToSlot(int(request.BannerId), int(request.SlotId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.AddBannerToSlotResponse{}, nil
}

func (r Server) RemoveBannerFromSlot(
	ctx context.Context,
	request *pb.RemoveBannerFromSlotRequest,
) (*pb.RemoveBannerFromSlotResponse, error) {
	err := r.rotator.RemoveBannerFromSlot(int(request.BannerId), int(request.SlotId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.RemoveBannerFromSlotResponse{}, nil
}

func (r Server) handleErr(err error) error {
	var errNotFound *NotFoundError

	if err == nil {
		return nil
	}

	if errors.As(err, &errNotFound) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
