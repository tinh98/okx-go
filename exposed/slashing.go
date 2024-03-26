package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
)

// Slashing shows the expected behavior for inner slashing client
type Slashing interface {
	gosdktypes.Module
	SlashingTx
}

// SlashingTx shows the expected tx behavior for inner slashing client
type SlashingTx interface {
	Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}
