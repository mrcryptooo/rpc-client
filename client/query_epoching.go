package client

import (
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	"github.com/strangelove-ventures/lens/client/query"
)

// QueryEpochingParams queries epoching module's parameters via ChainClient
// code is adapted from https://github.com/strangelove-ventures/lens/blob/v0.5.1/client/query/staking.go#L7-L18
func (c *Client) QueryEpochingParams() (*epochingtypes.Params, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return &epochingtypes.Params{}, err
	}
	return &resp.Params, nil
}

// QueryCurrentEpoch queries the current epoch number via ChainClient
func (c *Client) QueryCurrentEpoch() (uint64, error) {
	var (
		epochNum uint64
		err      error
	)

	q := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryCurrentEpochRequest{}
	resp, err := queryClient.CurrentEpoch(ctx, req)
	if err != nil {
		return epochNum, err
	}

	return resp.CurrentEpoch, nil
}
