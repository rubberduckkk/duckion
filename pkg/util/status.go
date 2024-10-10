package util

import (
	pbany "github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc/codes"

	pb "github.com/rubberduckkk/credit-card/api/pb/credit-card"
)

func PBStatus(code codes.Code, msg string, detailed ...*pbany.Any) *pb.Status {
	return &pb.Status{
		Code:     int32(code),
		Message:  msg,
		Detailed: detailed,
	}
}
