package sdk

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

// just a temporary test

const (
	name   = "alice"
	passWd = "12345678"
	// sender's mnemonic
	mnemonic = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	addr     = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	// target address
	mnemonic1 = "pepper basket run install fury scheme journey worry tumble toddler swap change"
	addr1     = "okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7"
	// validator address
	valAddr1 = "okchainvaloper10q0rk5qnyag7wfvvt7rtphlw589m7frs863s3m"
	valAddr2 = "okchainvaloper1g7znsf24w4jc3xfca88pq9kmlyjdare6mph5rx"
	valAddr3 = "okchainvaloper1alq9na49n9yycysh889rl90g9nhe58lcs50wu5"
	valAddr4 = "okchainvaloper1svzxp4ts5le2s4zugx34ajt6shz2hg42a3gl7g"
	// validator mnemonic
	valMnemonic = "race imitate stay curtain puppy suggest spend toe old bridge sunset pride"
	valName     = "validator"
)

func TestQueryValidators(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	vals, err := client.Staking().QueryValidators()
	require.NoError(t, err)
	for _, v := range vals {
		fmt.Println(v)
	}
}

func TestDelegate(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	resp, err := client.Staking().Delegate(fromInfo, passWd, "1024.1024okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestUnbond(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	resp, err := client.Staking().Unbond(fromInfo, passWd, "0.1024okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestVote(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	valsToVoted := []string{"okchainvaloper1dcsxvxgj374dv3wt9szflf9nz6342juz7grk2y", "okchainvaloper1fntm5xy7umzwmj6uxkateygmuhqf23e3uur68s"}
	resp, err := client.Staking().Vote(fromInfo, passWd, valsToVoted, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestDestroyValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(
		"relief food aim snap pumpkin black ginger badge flock citizen agree stone", name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount("okchain1yjsqpggz0tglf4mhtd40rwnwk3c08xrz908yge")
	require.NoError(t, err)

	resp, err := client.Staking().DestroyValidator(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestCreateValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	pubkeyStr := "okchainvalconspub1zcjduepqghrtvkngejwese62wg49ewskz4r93vkyj3md5mg5rf7twcc6jduqpqw66q"
	resp, err := client.Staking().CreateValidator(fromInfo, passWd, pubkeyStr, "my moniker", "my identity",
		"my website", "my details", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestEditValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(
		"sand hour across excess rocket usage cotton used install orient piano where", name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount("okchain12gzxhpah2nulpeqx7kpf7fmfpgmp3hwhczf8pu")
	require.NoError(t, err)

	resp, err := client.Staking().EditValidator(fromInfo, passWd, "my moniker", "my identity", "my website",
		"my details", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)

}

func TestRegisterProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	sequence := accInfo.GetSequence()
	res, err := client.Staking().Delegate(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	sequence++
	res, err = client.Staking().RegisterProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)
	fmt.Println(res)
}

func TestUnregisterProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	res, err := client.Staking().UnregisterProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestBindProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	valMnemo := "pride mystery melody cycle frog tag march style away dash they gold"
	// validator becomes a proxy
	valAcc, _, err := utils.CreateAccountWithMnemo(valMnemo, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(valAcc.GetAddress().String())
	require.NoError(t, err)

	sequence := accInfo.GetSequence()
	res, err := client.Staking().Delegate(valAcc, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	sequence++
	res, err = client.Staking().RegisterProxy(valAcc, passWd, "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	// delegator tries to bind proxy
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err = client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	sequence = accInfo.GetSequence()
	res, err = client.Staking().Delegate(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	sequence++
	res, err = client.Staking().BindProxy(fromInfo, passWd, valAcc.GetAddress().String(), "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)
	fmt.Println(res)

}
