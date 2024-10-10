package app

import (
	"context"
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	creditcard "github.com/rubberduckkk/credit-card/api/pb/credit-card"
	pb "github.com/rubberduckkk/credit-card/api/pb/credit-card"
	"github.com/rubberduckkk/credit-card/internal/infra/luhn"
	"github.com/rubberduckkk/credit-card/pkg/util"
)

func TestGetCreditCardServer(t *testing.T) {
	tests := []struct {
		name string
		want creditcard.CreditCardServer
	}{
		{
			name: "case normal",
			want: creditCardServerInstance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetCreditCardServer())
		})
	}
}

func Test_creditCardServerImpl_Validate(t *testing.T) {
	type args struct {
		ctx  context.Context
		card *pb.Card
	}
	tests := []struct {
		name     string
		args     args
		mocks    func() *gomonkey.Patches
		wantCode int32
		wantErr  bool
	}{
		{
			name: "case nil card",
			args: args{
				ctx:  context.Background(),
				card: nil,
			},
			wantCode: int32(codes.InvalidArgument),
			wantErr:  false,
		},
		{
			name: "case failed to convert to digits array",
			args: args{
				ctx:  context.Background(),
				card: &pb.Card{Number: "abc"},
			},
			mocks: func() *gomonkey.Patches {
				return gomonkey.ApplyFunc(util.ConvertStringToDigitArray, func(s string) ([]int, error) {
					return nil, errors.New("mock convert string to digit array error")
				})
			},
			wantCode: int32(codes.InvalidArgument),
			wantErr:  false,
		},
		{
			name: "case validluhn error",
			args: args{
				ctx:  context.Background(),
				card: &pb.Card{Number: "12345"},
			},
			mocks: func() *gomonkey.Patches {
				return gomonkey.ApplyFunc(luhn.IsValidLuhn, func(digits []int) (bool, error) {
					return false, errors.New("mock isvalidluhn error")
				})
			},
			wantCode: int32(codes.Internal),
			wantErr:  false,
		},
		{
			name: "case invalid credit card number",
			args: args{
				ctx:  context.Background(),
				card: &pb.Card{Number: "123"},
			},
			mocks: func() *gomonkey.Patches {
				return gomonkey.ApplyFunc(luhn.IsValidLuhn, func(digits []int) (bool, error) {
					return false, nil
				})
			},
			wantCode: int32(codes.InvalidArgument),
			wantErr:  false,
		},
		{
			name: "case success",
			args: args{
				ctx:  context.Background(),
				card: &pb.Card{Number: "123"},
			},
			mocks: func() *gomonkey.Patches {
				return gomonkey.ApplyFunc(luhn.IsValidLuhn, func(digits []int) (bool, error) {
					return true, nil
				})
			},
			wantCode: int32(codes.OK),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &creditCardServerImpl{}
			if tt.mocks != nil {
				if patches := tt.mocks(); patches != nil {
					defer patches.Reset()
				}
			}
			got, err := c.Validate(tt.args.ctx, tt.args.card)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantCode, got.GetCode())
		})
	}
}
