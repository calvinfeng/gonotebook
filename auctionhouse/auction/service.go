package auction

import (
	"context"

	"github.com/calvinfeng/go-academy/auctionhouse/protobuf"
)

type Service interface {
	protobuf.AuctionServer
	Customers() ([]*Customer, error)
	Listings() ([]*Listing, error)
	BiddingsByCustomer(uint) ([]*Bidding, error)
	BiddingsByListing(uint) ([]*Bidding, error)
	BiddingsCreate(*Customer, *Listing) error
}

var _ Service = (*ServiceServer)(nil)

type ServiceServer struct {
}

func (srv *ServiceServer) Bid(ctx context.Context, req *protobuf.BidRequest) (*protobuf.BidResponse, error) {
	return &protobuf.BidResponse{
		Status: &protobuf.BidStatus{
			Code:        protobuf.BidStatus_SUCCESS,
			Description: "successfully bidded on something",
		},
	}, nil
}

func (srv ServiceServer) Customers() ([]*Customer, error) {
	return nil, nil
}

func (srv *ServiceServer) Listings() ([]*Listing, error) {
	return nil, nil
}

func (srv *ServiceServer) BiddingsByCustomer(customerID uint) ([]*Bidding, error) {
	return nil, nil
}

func (srv *ServiceServer) BiddingsByListing(listingID uint) ([]*Bidding, error) {
	return nil, nil
}

func (srv *ServiceServer) BiddingsCreate(*Customer, *Listing) error {
	return nil
}
