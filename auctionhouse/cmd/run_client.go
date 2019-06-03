package cmd

import (
	"context"
	"fmt"

	"github.com/calvinfeng/go-academy/auctionhouse/protobuf"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var RunClientCmd = &cobra.Command{
	Use:   "runclient",
	Short: "run HTTP/gRPC client to dial auction server",
	RunE:  runClient,
}

func runClient(cmd *cobra.Command, args []string) error {
	cc, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		return err
	}

	cli := protobuf.NewAuctionClient(cc)
	resp, err := cli.Bid(context.Background(), &protobuf.BidRequest{})
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
