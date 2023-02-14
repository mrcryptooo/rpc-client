package query

import (
	"context"

	zctypes "github.com/babylonchain/babylon/x/zoneconcierge/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// QueryZoneConcierge queries the ZoneConcierge module of the Babylon node
// according to the given function
func (c *QueryClient) QueryZoneConcierge(f func(ctx context.Context, queryClient zctypes.QueryClient) error) error {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.RPCClient}
	queryClient := zctypes.NewQueryClient(clientCtx)

	return f(ctx, queryClient)
}

// FinalizedConnectedChainInfo queries the zoneconcierge module to get the finalization information for a connected chain
func (c *QueryClient) FinalizedConnectedChainInfo(chainID string) (*zctypes.QueryFinalizedChainInfoResponse, error) {
	var resp *zctypes.QueryFinalizedChainInfoResponse
	err := c.QueryZoneConcierge(func(ctx context.Context, queryClient zctypes.QueryClient) error {
		var err error
		req := &zctypes.QueryFinalizedChainInfoRequest{
			ChainId: chainID,
		}
		resp, err = queryClient.FinalizedChainInfo(ctx, req)
		return err
	})

	return resp, err
}

// ConnectedChainInfo queries the zoneconcierge module to get information for a connected chain
func (c *QueryClient) ConnectedChainInfo(chainID string) (*zctypes.QueryChainInfoResponse, error) {
	var resp *zctypes.QueryChainInfoResponse
	err := c.QueryZoneConcierge(func(ctx context.Context, queryClient zctypes.QueryClient) error {
		var err error
		req := &zctypes.QueryChainInfoRequest{
			ChainId: chainID,
		}
		resp, err = queryClient.ChainInfo(ctx, req)
		return err
	})

	return resp, err
}

// ConnectedChainList queries the zoneconierge module for the chain IDs of the connected chains
func (c *QueryClient) ConnectedChainList() (*zctypes.QueryChainListResponse, error) {
	var resp *zctypes.QueryChainListResponse
	err := c.QueryZoneConcierge(func(ctx context.Context, queryClient zctypes.QueryClient) error {
		var err error
		req := &zctypes.QueryChainListRequest{}
		resp, err = queryClient.ChainList(ctx, req)
		return err
	})

	return resp, err
}

// ConnectedChainHeaders queries the zoneconcierge module for the headers of a connected chain
func (c *QueryClient) ConnectedChainHeaders(chainID string, pagination *sdkquerytypes.PageRequest) (*zctypes.QueryListHeadersResponse, error) {
	var resp *zctypes.QueryListHeadersResponse
	err := c.QueryZoneConcierge(func(ctx context.Context, queryClient zctypes.QueryClient) error {
		var err error
		req := &zctypes.QueryListHeadersRequest{
			ChainId:    chainID,
			Pagination: pagination,
		}
		resp, err = queryClient.ListHeaders(ctx, req)
		return err
	})

	return resp, err
}

// ConnectedChainEpochInfo queries the zoneconcierge module for the chain information of a connected chain at a particular epoch
func (c *QueryClient) ConnectedChainEpochInfo(chainID string, epochNum uint64) (*zctypes.QueryEpochChainInfoResponse, error) {
	var resp *zctypes.QueryEpochChainInfoResponse
	err := c.QueryZoneConcierge(func(ctx context.Context, queryClient zctypes.QueryClient) error {
		var err error
		req := &zctypes.QueryEpochChainInfoRequest{
			ChainId:  chainID,
			EpochNum: epochNum,
		}
		resp, err = queryClient.EpochChainInfo(ctx, req)
		return err
	})

	return resp, err
}