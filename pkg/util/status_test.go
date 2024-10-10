package util

import (
	"testing"

	pbany "github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	pb "github.com/rubberduckkk/credit-card/api/pb/credit-card"
)

func TestPBStatus(t *testing.T) {
	type args struct {
		code     codes.Code
		msg      string
		detailed []*pbany.Any
	}
	tests := []struct {
		name string
		args args
		want *pb.Status
	}{
		{
			name: "case normal",
			args: args{
				code:     0,
				msg:      "success",
				detailed: nil,
			},
			want: &pb.Status{
				Code:     0,
				Message:  "success",
				Detailed: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PBStatus(tt.args.code, tt.args.msg, tt.args.detailed...)
			assert.Equal(t, tt.want, got)
		})
	}
}
