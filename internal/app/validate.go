package app

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"

	pb "github.com/rubberduckkk/credit-card/api/pb/credit-card"
	"github.com/rubberduckkk/credit-card/internal/infra/luhn"
	util2 "github.com/rubberduckkk/credit-card/pkg/util"
)

type creditCardServerImpl struct {
	pb.UnimplementedCreditCardServer
}

var creditCardServerInstance *creditCardServerImpl

func init() {
	creditCardServerInstance = &creditCardServerImpl{}
}

func GetCreditCardServer() pb.CreditCardServer {
	return creditCardServerInstance
}

func (c *creditCardServerImpl) Validate(ctx context.Context, card *pb.Card) (*pb.Status, error) {
	if card == nil {
		return util2.PBStatus(codes.InvalidArgument, "err: credit card is empty"), nil
	}

	digits, err := util2.ConvertStringToDigitArray(card.Number)
	if err != nil {
		return util2.PBStatus(codes.InvalidArgument, "err: invalid credit card number"), nil
	}

	valid, err := luhn.IsValidLuhn(digits)
	if err != nil {
		return util2.PBStatus(codes.Internal, fmt.Sprintf("validate credit card failed: %v", err)), nil
	}

	if !valid {
		return util2.PBStatus(codes.InvalidArgument, "err: credit card number is not valid"), nil
	}

	return util2.PBStatus(codes.OK, "success: credit card number is valid"), nil
}
