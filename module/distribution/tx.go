package distribution

import (
	"fmt"

	"github.com/okex/exchain-go-sdk/types/params"
	"github.com/okx/okbchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	distrtypes "github.com/okx/okbchain/x/distribution/types"
)

// SetWithdrawAddr changes the withdraw address of validator to receive rewards
func (dc distrClient) SetWithdrawAddr(fromInfo keys.Info, passWd, withdrawAddrStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckSendParams(fromInfo, passWd, withdrawAddrStr); err != nil {
		return
	}

	withdrawAddr, err := sdk.AccAddressFromBech32(withdrawAddrStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse Address [%s] error: %s", withdrawAddrStr, err)
	}

	msg := distrtypes.NewMsgSetWithdrawAddress(fromInfo.GetAddress(), withdrawAddr)
	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// WithdrawRewards withdraws the rewards of validator by himself
func (dc distrClient) WithdrawRewards(fromInfo keys.Info, passWd, valAddrStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
	if err != nil {
		return resp, fmt.Errorf("failed. invalid validator address: %s", valAddrStr)
	}

	msg := distrtypes.NewMsgWithdrawValidatorCommission(valAddr)
	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
