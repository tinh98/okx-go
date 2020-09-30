package governance

import (
	"github.com/okex/okexchain-go-sdk/module/governance/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
)

// QueryProposals gets all proposals
// Note:
//	optional:
//		status option - DepositPeriod|VotingPeriod|Passed|Rejected. Defaults to all proposals by ""
//		depositorAddrStr - filter by proposals deposited on by depositor. Defaults to all proposals by ""
//		voterAddrStr - filter by proposals voted on by voted. Defaults to all proposals by ""
// 		numLimit - limit to latest [number] proposals. Defaults to all proposals by 0
func (gc govClient) QueryProposals(depositorAddrStr, voterAddrStr, status string, numLimit uint64) (
	proposals []types.Proposal, err error) {
	var depositorAddr, voterAddr sdk.AccAddress
	var proposalStatus types.ProposalStatus
	proposalParams := params.NewQueryProposalsParams(proposalStatus, numLimit, depositorAddr, voterAddr)

	if len(depositorAddrStr) != 0 {
		depositorAddr, err = sdk.AccAddressFromBech32(depositorAddrStr)
		if err != nil {
			return
		}
		proposalParams.Depositor = depositorAddr
	}
	if len(voterAddrStr) != 0 {
		voterAddr, err = sdk.AccAddressFromBech32(voterAddrStr)
		if err != nil {
			return
		}
		proposalParams.Voter = voterAddr
	}
	if len(status) != 0 {
		proposalStatus, err = types.ProposalStatusFromString(status)
		if err != nil {
			return
		}
		proposalParams.ProposalStatus = proposalStatus
	}

	jsonBytes, err := gc.GetCodec().MarshalJSON(proposalParams)
	if err != nil {
		return proposals, utils.ErrMarshalJSON(err.Error())
	}

	res, err := gc.Query(types.ProposalsPath, jsonBytes)
	if err != nil {
		return proposals, utils.ErrClientQuery(err.Error())
	}

	if err = gc.GetCodec().UnmarshalJSON(res, &proposals); err != nil {
		return proposals, utils.ErrUnmarshalJSON(err.Error())
	}
	return

}