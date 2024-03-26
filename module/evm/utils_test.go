package evm

import (
	"math/big"
	"testing"

	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/okex/exchain-go-sdk/mocks"
	"github.com/okex/exchain-go-sdk/module/auth"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	privKeyHex     = "89c81c304704e9890025a5a91898802294658d6e4034a11c6116f4b129ea12d3"
	expectedTxHash = "0xe03ce7281cee90953b1bee78509b80be34e116821d8a9018a26a0feda09ac41d"
)

func TestEvmClient_GetTxHash(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	privateKeyECDSA, err := crypto.HexToECDSA(privKeyHex)
	require.NoError(t, err)

	// details in okt transfer
	chainID := big.NewInt(65)
	toAddress := ethcmn.HexToAddress(recAddrEth)
	value, err := sdk.NewDecFromStr("1")
	require.NoError(t, err)
	gasLimit := uint64(21000)
	gasPrice, err := sdk.NewDecFromStr("0.000000001")
	require.NoError(t, err)
	unsignedTx := ethcore.NewTransaction(591, toAddress, value.BigInt(), gasLimit, gasPrice.BigInt(), nil)
	signedTx, err := ethcore.SignTx(unsignedTx, ethcore.NewEIP155Signer(chainID), privateKeyECDSA)
	require.NoError(t, err)

	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewEvmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)

	txHash, err := mockCli.Evm().GetTxHash(signedTx)
	require.NotNil(t, txHash)
	require.Equal(t, expectedTxHash, txHash.String())
}
