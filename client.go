package sdk

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module"
	"github.com/okex/okchain-go-sdk/module/auth"
	"github.com/okex/okchain-go-sdk/module/staking"
	"github.com/okex/okchain-go-sdk/types"
)

// Client defines the main client of okchain gosdk
type Client struct {
	cdc     types.SDKCodec
	modules map[string]types.Module
}

// NewClient creates a new instance of Client
func NewClient(config types.ClientConfig) Client {
	cdc := types.NewCodec()
	pBaseClient := module.NewBaseClient(cdc, config)

	pClient := &Client{
		cdc:     cdc,
		modules: make(map[string]types.Module),
	}

	pClient.registerModule(
		auth.NewAuthClient(pBaseClient),
		staking.NewStakingClient(pBaseClient),
	)

	return *pClient
}

func (cli *Client) registerModule(modules ...types.Module) {
	for _, module := range modules {
		moduleName := module.Name()
		if _, ok := cli.modules[module.Name()]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		module.RegisterCodec(cli.cdc)
		cli.modules[moduleName] = module
	}
	types.RegisterBasicCodec(cli.cdc)
}

// nolint
func (cli *Client) Auth() exposed.Auth       { return cli.modules[auth.ModuleName].(exposed.Auth) }
func (cli *Client) Staking() exposed.Staking { return cli.modules[staking.ModuleName].(exposed.Staking) }