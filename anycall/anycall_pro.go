// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package anycall

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RouterConfigChainConfig is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigChainConfig struct {
	BlockChain     string
	RouterContract string
	Confirmations  uint64
	InitialHeight  uint64
	Extra          string
}

// RouterConfigChainConfig2 is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigChainConfig2 struct {
	ChainID        *big.Int
	BlockChain     string
	RouterContract string
	Confirmations  uint64
	InitialHeight  uint64
	Extra          string
}

// RouterConfigFeeConfig is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigFeeConfig struct {
	MaximumSwapFee        *big.Int
	MinimumSwapFee        *big.Int
	SwapFeeRatePerMillion *big.Int
}

// RouterConfigFeeConfig2 is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigFeeConfig2 struct {
	FromChainID           *big.Int
	ToChainID             *big.Int
	MaximumSwapFee        *big.Int
	MinimumSwapFee        *big.Int
	SwapFeeRatePerMillion *big.Int
}

// RouterConfigMultichainToken is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigMultichainToken struct {
	ChainID      *big.Int
	TokenAddress string
}

// RouterConfigSwapConfig is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigSwapConfig struct {
	MaximumSwap       *big.Int
	MinimumSwap       *big.Int
	BigValueThreshold *big.Int
}

// RouterConfigSwapConfig2 is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigSwapConfig2 struct {
	FromChainID       *big.Int
	ToChainID         *big.Int
	MaximumSwap       *big.Int
	MinimumSwap       *big.Int
	BigValueThreshold *big.Int
}

// RouterConfigTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigTokenConfig struct {
	Decimals        uint8
	ContractAddress string
	ContractVersion *big.Int
	RouterContract  string
	Extra           string
}

// RouterConfigTokenConfig2 is an auto generated low-level Go binding around an user-defined struct.
type RouterConfigTokenConfig2 struct {
	ChainID         *big.Int
	Decimals        uint8
	ContractAddress string
	ContractVersion *big.Int
	RouterContract  string
	Extra           string
}

// AnycallMetaData contains all meta data concerning the Anycall contract.
var AnycallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"owners\",\"type\":\"address[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdateConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIG_VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"addChainID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"addTokenID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ChainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"BlockChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RouterContract\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"Confirmations\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"InitialHeight\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"Extra\",\"type\":\"string\"}],\"internalType\":\"structRouterConfig.ChainConfig2[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllChainIDLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getAllFeeConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SwapFeeRatePerMillion\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.FeeConfig2[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getAllMultichainTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"Decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"ContractAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ContractVersion\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"RouterContract\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Extra\",\"type\":\"string\"}],\"internalType\":\"structRouterConfig.TokenConfig2[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getAllMultichainTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ChainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"TokenAddress\",\"type\":\"string\"}],\"internalType\":\"structRouterConfig.MultichainToken[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getAllSwapConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BigValueThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.SwapConfig2[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTokenIDLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTokenIDs\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"result\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"BlockChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RouterContract\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"Confirmations\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"InitialHeight\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"Extra\",\"type\":\"string\"}],\"internalType\":\"structRouterConfig.ChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChainIDByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getCustomConfig\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getExtraConfig\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainID\",\"type\":\"uint256\"}],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"MaximumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SwapFeeRatePerMillion\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getFeeConfigAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SwapFeeRatePerMillion\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.FeeConfig2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getFeeConfigAtIndexRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SwapFeeRatePerMillion\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.FeeConfig2[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getFeeConfigsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mpcAddress\",\"type\":\"string\"}],\"name\":\"getMPCPubkey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMultichainToken\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getOriginalTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"Decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"ContractAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ContractVersion\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"RouterContract\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Extra\",\"type\":\"string\"}],\"internalType\":\"structRouterConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainID\",\"type\":\"uint256\"}],\"name\":\"getSwapConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"MaximumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BigValueThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.SwapConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSwapConfigAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BigValueThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.SwapConfig2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getSwapConfigAtIndexRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BigValueThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.SwapConfig2[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getSwapConfigsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"Decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"ContractAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ContractVersion\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"RouterContract\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Extra\",\"type\":\"string\"}],\"internalType\":\"structRouterConfig.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"}],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getTokenIDByIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"isChainIDExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"isTokenIDExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address[2]\",\"name\":\"\",\"type\":\"address[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllChainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"dstChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeAllDestFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"dstChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeAllDestSwapConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"removeAllFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"removeAllMultichainTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"srcChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeAllSrcFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"srcChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeAllSrcSwapConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"removeAllSwapConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"removeAllTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeChainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeChainID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"srcChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"dstChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"removeMPCPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeMultichainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainID\",\"type\":\"uint256\"}],\"name\":\"removeSwapAndFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"srcChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"dstChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeSwapConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"removeTokenID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"blockChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"routerContract\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"confirmations\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initialHeight\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"setChainConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"setChainExtraConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"setCustomConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"setExtraConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SwapFeeRatePerMillion\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.FeeConfig2[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"}],\"name\":\"setMPCPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"}],\"name\":\"setMultichainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bigSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"setSwapAndFeeConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bigSwap\",\"type\":\"uint256\"}],\"name\":\"setSwapConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"FromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ToChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MaximumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MinimumSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BigValueThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structRouterConfig.SwapConfig2[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"setSwapConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddr\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"routerContract\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"setTokenConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"setTokenExtraConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"routerContract\",\"type\":\"string\"}],\"name\":\"setTokenRouterContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AnycallABI is the input ABI used to generate the binding from.
// Deprecated: Use AnycallMetaData.ABI instead.
var AnycallABI = AnycallMetaData.ABI

// Anycall is an auto generated Go binding around an Ethereum contract.
type Anycall struct {
	AnycallCaller     // Read-only binding to the contract
	AnycallTransactor // Write-only binding to the contract
	AnycallFilterer   // Log filterer for contract events
}

// AnycallCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnycallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnycallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnycallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnycallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnycallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnycallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnycallSession struct {
	Contract     *Anycall          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnycallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnycallCallerSession struct {
	Contract *AnycallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AnycallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnycallTransactorSession struct {
	Contract     *AnycallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AnycallRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnycallRaw struct {
	Contract *Anycall // Generic contract binding to access the raw methods on
}

// AnycallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnycallCallerRaw struct {
	Contract *AnycallCaller // Generic read-only contract binding to access the raw methods on
}

// AnycallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnycallTransactorRaw struct {
	Contract *AnycallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnycall creates a new instance of Anycall, bound to a specific deployed contract.
func NewAnycall(address common.Address, backend bind.ContractBackend) (*Anycall, error) {
	contract, err := bindAnycall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Anycall{AnycallCaller: AnycallCaller{contract: contract}, AnycallTransactor: AnycallTransactor{contract: contract}, AnycallFilterer: AnycallFilterer{contract: contract}}, nil
}

// NewAnycallCaller creates a new read-only instance of Anycall, bound to a specific deployed contract.
func NewAnycallCaller(address common.Address, caller bind.ContractCaller) (*AnycallCaller, error) {
	contract, err := bindAnycall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnycallCaller{contract: contract}, nil
}

// NewAnycallTransactor creates a new write-only instance of Anycall, bound to a specific deployed contract.
func NewAnycallTransactor(address common.Address, transactor bind.ContractTransactor) (*AnycallTransactor, error) {
	contract, err := bindAnycall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnycallTransactor{contract: contract}, nil
}

// NewAnycallFilterer creates a new log filterer instance of Anycall, bound to a specific deployed contract.
func NewAnycallFilterer(address common.Address, filterer bind.ContractFilterer) (*AnycallFilterer, error) {
	contract, err := bindAnycall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnycallFilterer{contract: contract}, nil
}

// bindAnycall binds a generic wrapper to an already deployed contract.
func bindAnycall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnycallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anycall *AnycallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anycall.Contract.AnycallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anycall *AnycallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anycall.Contract.AnycallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anycall *AnycallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anycall.Contract.AnycallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anycall *AnycallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anycall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anycall *AnycallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anycall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anycall *AnycallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anycall.Contract.contract.Transact(opts, method, params...)
}

// CONFIGVERSION is a free data retrieval call binding the contract method 0x34e30097.
//
// Solidity: function CONFIG_VERSION() view returns(uint256)
func (_Anycall *AnycallCaller) CONFIGVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "CONFIG_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGVERSION is a free data retrieval call binding the contract method 0x34e30097.
//
// Solidity: function CONFIG_VERSION() view returns(uint256)
func (_Anycall *AnycallSession) CONFIGVERSION() (*big.Int, error) {
	return _Anycall.Contract.CONFIGVERSION(&_Anycall.CallOpts)
}

// CONFIGVERSION is a free data retrieval call binding the contract method 0x34e30097.
//
// Solidity: function CONFIG_VERSION() view returns(uint256)
func (_Anycall *AnycallCallerSession) CONFIGVERSION() (*big.Int, error) {
	return _Anycall.Contract.CONFIGVERSION(&_Anycall.CallOpts)
}

// GetAllChainConfig is a free data retrieval call binding the contract method 0x71a4a947.
//
// Solidity: function getAllChainConfig() view returns((uint256,string,string,uint64,uint64,string)[])
func (_Anycall *AnycallCaller) GetAllChainConfig(opts *bind.CallOpts) ([]RouterConfigChainConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllChainConfig")

	if err != nil {
		return *new([]RouterConfigChainConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigChainConfig2)).(*[]RouterConfigChainConfig2)

	return out0, err

}

// GetAllChainConfig is a free data retrieval call binding the contract method 0x71a4a947.
//
// Solidity: function getAllChainConfig() view returns((uint256,string,string,uint64,uint64,string)[])
func (_Anycall *AnycallSession) GetAllChainConfig() ([]RouterConfigChainConfig2, error) {
	return _Anycall.Contract.GetAllChainConfig(&_Anycall.CallOpts)
}

// GetAllChainConfig is a free data retrieval call binding the contract method 0x71a4a947.
//
// Solidity: function getAllChainConfig() view returns((uint256,string,string,uint64,uint64,string)[])
func (_Anycall *AnycallCallerSession) GetAllChainConfig() ([]RouterConfigChainConfig2, error) {
	return _Anycall.Contract.GetAllChainConfig(&_Anycall.CallOpts)
}

// GetAllChainIDLength is a free data retrieval call binding the contract method 0x72c93288.
//
// Solidity: function getAllChainIDLength() view returns(uint256)
func (_Anycall *AnycallCaller) GetAllChainIDLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllChainIDLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllChainIDLength is a free data retrieval call binding the contract method 0x72c93288.
//
// Solidity: function getAllChainIDLength() view returns(uint256)
func (_Anycall *AnycallSession) GetAllChainIDLength() (*big.Int, error) {
	return _Anycall.Contract.GetAllChainIDLength(&_Anycall.CallOpts)
}

// GetAllChainIDLength is a free data retrieval call binding the contract method 0x72c93288.
//
// Solidity: function getAllChainIDLength() view returns(uint256)
func (_Anycall *AnycallCallerSession) GetAllChainIDLength() (*big.Int, error) {
	return _Anycall.Contract.GetAllChainIDLength(&_Anycall.CallOpts)
}

// GetAllChainIDs is a free data retrieval call binding the contract method 0xe27112d5.
//
// Solidity: function getAllChainIDs() view returns(uint256[])
func (_Anycall *AnycallCaller) GetAllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllChainIDs is a free data retrieval call binding the contract method 0xe27112d5.
//
// Solidity: function getAllChainIDs() view returns(uint256[])
func (_Anycall *AnycallSession) GetAllChainIDs() ([]*big.Int, error) {
	return _Anycall.Contract.GetAllChainIDs(&_Anycall.CallOpts)
}

// GetAllChainIDs is a free data retrieval call binding the contract method 0xe27112d5.
//
// Solidity: function getAllChainIDs() view returns(uint256[])
func (_Anycall *AnycallCallerSession) GetAllChainIDs() ([]*big.Int, error) {
	return _Anycall.Contract.GetAllChainIDs(&_Anycall.CallOpts)
}

// GetAllFeeConfigs is a free data retrieval call binding the contract method 0x6a3ea04f.
//
// Solidity: function getAllFeeConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCaller) GetAllFeeConfigs(opts *bind.CallOpts, tokenID string) ([]RouterConfigFeeConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllFeeConfigs", tokenID)

	if err != nil {
		return *new([]RouterConfigFeeConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigFeeConfig2)).(*[]RouterConfigFeeConfig2)

	return out0, err

}

// GetAllFeeConfigs is a free data retrieval call binding the contract method 0x6a3ea04f.
//
// Solidity: function getAllFeeConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallSession) GetAllFeeConfigs(tokenID string) ([]RouterConfigFeeConfig2, error) {
	return _Anycall.Contract.GetAllFeeConfigs(&_Anycall.CallOpts, tokenID)
}

// GetAllFeeConfigs is a free data retrieval call binding the contract method 0x6a3ea04f.
//
// Solidity: function getAllFeeConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCallerSession) GetAllFeeConfigs(tokenID string) ([]RouterConfigFeeConfig2, error) {
	return _Anycall.Contract.GetAllFeeConfigs(&_Anycall.CallOpts, tokenID)
}

// GetAllMultichainTokenConfig is a free data retrieval call binding the contract method 0x160dcc6f.
//
// Solidity: function getAllMultichainTokenConfig(string tokenID) view returns((uint256,uint8,string,uint256,string,string)[])
func (_Anycall *AnycallCaller) GetAllMultichainTokenConfig(opts *bind.CallOpts, tokenID string) ([]RouterConfigTokenConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllMultichainTokenConfig", tokenID)

	if err != nil {
		return *new([]RouterConfigTokenConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigTokenConfig2)).(*[]RouterConfigTokenConfig2)

	return out0, err

}

// GetAllMultichainTokenConfig is a free data retrieval call binding the contract method 0x160dcc6f.
//
// Solidity: function getAllMultichainTokenConfig(string tokenID) view returns((uint256,uint8,string,uint256,string,string)[])
func (_Anycall *AnycallSession) GetAllMultichainTokenConfig(tokenID string) ([]RouterConfigTokenConfig2, error) {
	return _Anycall.Contract.GetAllMultichainTokenConfig(&_Anycall.CallOpts, tokenID)
}

// GetAllMultichainTokenConfig is a free data retrieval call binding the contract method 0x160dcc6f.
//
// Solidity: function getAllMultichainTokenConfig(string tokenID) view returns((uint256,uint8,string,uint256,string,string)[])
func (_Anycall *AnycallCallerSession) GetAllMultichainTokenConfig(tokenID string) ([]RouterConfigTokenConfig2, error) {
	return _Anycall.Contract.GetAllMultichainTokenConfig(&_Anycall.CallOpts, tokenID)
}

// GetAllMultichainTokens is a free data retrieval call binding the contract method 0x8fcb62a3.
//
// Solidity: function getAllMultichainTokens(string tokenID) view returns((uint256,string)[])
func (_Anycall *AnycallCaller) GetAllMultichainTokens(opts *bind.CallOpts, tokenID string) ([]RouterConfigMultichainToken, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllMultichainTokens", tokenID)

	if err != nil {
		return *new([]RouterConfigMultichainToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigMultichainToken)).(*[]RouterConfigMultichainToken)

	return out0, err

}

// GetAllMultichainTokens is a free data retrieval call binding the contract method 0x8fcb62a3.
//
// Solidity: function getAllMultichainTokens(string tokenID) view returns((uint256,string)[])
func (_Anycall *AnycallSession) GetAllMultichainTokens(tokenID string) ([]RouterConfigMultichainToken, error) {
	return _Anycall.Contract.GetAllMultichainTokens(&_Anycall.CallOpts, tokenID)
}

// GetAllMultichainTokens is a free data retrieval call binding the contract method 0x8fcb62a3.
//
// Solidity: function getAllMultichainTokens(string tokenID) view returns((uint256,string)[])
func (_Anycall *AnycallCallerSession) GetAllMultichainTokens(tokenID string) ([]RouterConfigMultichainToken, error) {
	return _Anycall.Contract.GetAllMultichainTokens(&_Anycall.CallOpts, tokenID)
}

// GetAllSwapConfigs is a free data retrieval call binding the contract method 0x3c6b1a8f.
//
// Solidity: function getAllSwapConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCaller) GetAllSwapConfigs(opts *bind.CallOpts, tokenID string) ([]RouterConfigSwapConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllSwapConfigs", tokenID)

	if err != nil {
		return *new([]RouterConfigSwapConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigSwapConfig2)).(*[]RouterConfigSwapConfig2)

	return out0, err

}

// GetAllSwapConfigs is a free data retrieval call binding the contract method 0x3c6b1a8f.
//
// Solidity: function getAllSwapConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallSession) GetAllSwapConfigs(tokenID string) ([]RouterConfigSwapConfig2, error) {
	return _Anycall.Contract.GetAllSwapConfigs(&_Anycall.CallOpts, tokenID)
}

// GetAllSwapConfigs is a free data retrieval call binding the contract method 0x3c6b1a8f.
//
// Solidity: function getAllSwapConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCallerSession) GetAllSwapConfigs(tokenID string) ([]RouterConfigSwapConfig2, error) {
	return _Anycall.Contract.GetAllSwapConfigs(&_Anycall.CallOpts, tokenID)
}

// GetAllTokenIDLength is a free data retrieval call binding the contract method 0xa43f37d7.
//
// Solidity: function getAllTokenIDLength() view returns(uint256)
func (_Anycall *AnycallCaller) GetAllTokenIDLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllTokenIDLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllTokenIDLength is a free data retrieval call binding the contract method 0xa43f37d7.
//
// Solidity: function getAllTokenIDLength() view returns(uint256)
func (_Anycall *AnycallSession) GetAllTokenIDLength() (*big.Int, error) {
	return _Anycall.Contract.GetAllTokenIDLength(&_Anycall.CallOpts)
}

// GetAllTokenIDLength is a free data retrieval call binding the contract method 0xa43f37d7.
//
// Solidity: function getAllTokenIDLength() view returns(uint256)
func (_Anycall *AnycallCallerSession) GetAllTokenIDLength() (*big.Int, error) {
	return _Anycall.Contract.GetAllTokenIDLength(&_Anycall.CallOpts)
}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_Anycall *AnycallCaller) GetAllTokenIDs(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getAllTokenIDs")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_Anycall *AnycallSession) GetAllTokenIDs() ([]string, error) {
	return _Anycall.Contract.GetAllTokenIDs(&_Anycall.CallOpts)
}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_Anycall *AnycallCallerSession) GetAllTokenIDs() ([]string, error) {
	return _Anycall.Contract.GetAllTokenIDs(&_Anycall.CallOpts)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x19ed16dc.
//
// Solidity: function getChainConfig(uint256 chainID) view returns((string,string,uint64,uint64,string))
func (_Anycall *AnycallCaller) GetChainConfig(opts *bind.CallOpts, chainID *big.Int) (RouterConfigChainConfig, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getChainConfig", chainID)

	if err != nil {
		return *new(RouterConfigChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigChainConfig)).(*RouterConfigChainConfig)

	return out0, err

}

// GetChainConfig is a free data retrieval call binding the contract method 0x19ed16dc.
//
// Solidity: function getChainConfig(uint256 chainID) view returns((string,string,uint64,uint64,string))
func (_Anycall *AnycallSession) GetChainConfig(chainID *big.Int) (RouterConfigChainConfig, error) {
	return _Anycall.Contract.GetChainConfig(&_Anycall.CallOpts, chainID)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x19ed16dc.
//
// Solidity: function getChainConfig(uint256 chainID) view returns((string,string,uint64,uint64,string))
func (_Anycall *AnycallCallerSession) GetChainConfig(chainID *big.Int) (RouterConfigChainConfig, error) {
	return _Anycall.Contract.GetChainConfig(&_Anycall.CallOpts, chainID)
}

// GetChainIDByIndex is a free data retrieval call binding the contract method 0x21576975.
//
// Solidity: function getChainIDByIndex(uint256 index) view returns(uint256)
func (_Anycall *AnycallCaller) GetChainIDByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getChainIDByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainIDByIndex is a free data retrieval call binding the contract method 0x21576975.
//
// Solidity: function getChainIDByIndex(uint256 index) view returns(uint256)
func (_Anycall *AnycallSession) GetChainIDByIndex(index *big.Int) (*big.Int, error) {
	return _Anycall.Contract.GetChainIDByIndex(&_Anycall.CallOpts, index)
}

// GetChainIDByIndex is a free data retrieval call binding the contract method 0x21576975.
//
// Solidity: function getChainIDByIndex(uint256 index) view returns(uint256)
func (_Anycall *AnycallCallerSession) GetChainIDByIndex(index *big.Int) (*big.Int, error) {
	return _Anycall.Contract.GetChainIDByIndex(&_Anycall.CallOpts, index)
}

// GetCustomConfig is a free data retrieval call binding the contract method 0x61387d61.
//
// Solidity: function getCustomConfig(uint256 chainID, string key) view returns(string)
func (_Anycall *AnycallCaller) GetCustomConfig(opts *bind.CallOpts, chainID *big.Int, key string) (string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getCustomConfig", chainID, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCustomConfig is a free data retrieval call binding the contract method 0x61387d61.
//
// Solidity: function getCustomConfig(uint256 chainID, string key) view returns(string)
func (_Anycall *AnycallSession) GetCustomConfig(chainID *big.Int, key string) (string, error) {
	return _Anycall.Contract.GetCustomConfig(&_Anycall.CallOpts, chainID, key)
}

// GetCustomConfig is a free data retrieval call binding the contract method 0x61387d61.
//
// Solidity: function getCustomConfig(uint256 chainID, string key) view returns(string)
func (_Anycall *AnycallCallerSession) GetCustomConfig(chainID *big.Int, key string) (string, error) {
	return _Anycall.Contract.GetCustomConfig(&_Anycall.CallOpts, chainID, key)
}

// GetExtraConfig is a free data retrieval call binding the contract method 0x340a5f2d.
//
// Solidity: function getExtraConfig(string key) view returns(string)
func (_Anycall *AnycallCaller) GetExtraConfig(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getExtraConfig", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetExtraConfig is a free data retrieval call binding the contract method 0x340a5f2d.
//
// Solidity: function getExtraConfig(string key) view returns(string)
func (_Anycall *AnycallSession) GetExtraConfig(key string) (string, error) {
	return _Anycall.Contract.GetExtraConfig(&_Anycall.CallOpts, key)
}

// GetExtraConfig is a free data retrieval call binding the contract method 0x340a5f2d.
//
// Solidity: function getExtraConfig(string key) view returns(string)
func (_Anycall *AnycallCallerSession) GetExtraConfig(key string) (string, error) {
	return _Anycall.Contract.GetExtraConfig(&_Anycall.CallOpts, key)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x1aed1c97.
//
// Solidity: function getFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Anycall *AnycallCaller) GetFeeConfig(opts *bind.CallOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigFeeConfig, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getFeeConfig", tokenID, srcChainID, dstChainID)

	if err != nil {
		return *new(RouterConfigFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigFeeConfig)).(*RouterConfigFeeConfig)

	return out0, err

}

// GetFeeConfig is a free data retrieval call binding the contract method 0x1aed1c97.
//
// Solidity: function getFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Anycall *AnycallSession) GetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigFeeConfig, error) {
	return _Anycall.Contract.GetFeeConfig(&_Anycall.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x1aed1c97.
//
// Solidity: function getFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Anycall *AnycallCallerSession) GetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigFeeConfig, error) {
	return _Anycall.Contract.GetFeeConfig(&_Anycall.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetFeeConfigAtIndex is a free data retrieval call binding the contract method 0x8b665a51.
//
// Solidity: function getFeeConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Anycall *AnycallCaller) GetFeeConfigAtIndex(opts *bind.CallOpts, tokenID string, index *big.Int) (RouterConfigFeeConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getFeeConfigAtIndex", tokenID, index)

	if err != nil {
		return *new(RouterConfigFeeConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigFeeConfig2)).(*RouterConfigFeeConfig2)

	return out0, err

}

// GetFeeConfigAtIndex is a free data retrieval call binding the contract method 0x8b665a51.
//
// Solidity: function getFeeConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Anycall *AnycallSession) GetFeeConfigAtIndex(tokenID string, index *big.Int) (RouterConfigFeeConfig2, error) {
	return _Anycall.Contract.GetFeeConfigAtIndex(&_Anycall.CallOpts, tokenID, index)
}

// GetFeeConfigAtIndex is a free data retrieval call binding the contract method 0x8b665a51.
//
// Solidity: function getFeeConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Anycall *AnycallCallerSession) GetFeeConfigAtIndex(tokenID string, index *big.Int) (RouterConfigFeeConfig2, error) {
	return _Anycall.Contract.GetFeeConfigAtIndex(&_Anycall.CallOpts, tokenID, index)
}

// GetFeeConfigAtIndexRange is a free data retrieval call binding the contract method 0xbdd5ef4a.
//
// Solidity: function getFeeConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCaller) GetFeeConfigAtIndexRange(opts *bind.CallOpts, tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigFeeConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getFeeConfigAtIndexRange", tokenID, startIndex, endIndex)

	if err != nil {
		return *new([]RouterConfigFeeConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigFeeConfig2)).(*[]RouterConfigFeeConfig2)

	return out0, err

}

// GetFeeConfigAtIndexRange is a free data retrieval call binding the contract method 0xbdd5ef4a.
//
// Solidity: function getFeeConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallSession) GetFeeConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigFeeConfig2, error) {
	return _Anycall.Contract.GetFeeConfigAtIndexRange(&_Anycall.CallOpts, tokenID, startIndex, endIndex)
}

// GetFeeConfigAtIndexRange is a free data retrieval call binding the contract method 0xbdd5ef4a.
//
// Solidity: function getFeeConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCallerSession) GetFeeConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigFeeConfig2, error) {
	return _Anycall.Contract.GetFeeConfigAtIndexRange(&_Anycall.CallOpts, tokenID, startIndex, endIndex)
}

// GetFeeConfigsCount is a free data retrieval call binding the contract method 0x15393217.
//
// Solidity: function getFeeConfigsCount(string tokenID) view returns(uint256)
func (_Anycall *AnycallCaller) GetFeeConfigsCount(opts *bind.CallOpts, tokenID string) (*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getFeeConfigsCount", tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeConfigsCount is a free data retrieval call binding the contract method 0x15393217.
//
// Solidity: function getFeeConfigsCount(string tokenID) view returns(uint256)
func (_Anycall *AnycallSession) GetFeeConfigsCount(tokenID string) (*big.Int, error) {
	return _Anycall.Contract.GetFeeConfigsCount(&_Anycall.CallOpts, tokenID)
}

// GetFeeConfigsCount is a free data retrieval call binding the contract method 0x15393217.
//
// Solidity: function getFeeConfigsCount(string tokenID) view returns(uint256)
func (_Anycall *AnycallCallerSession) GetFeeConfigsCount(tokenID string) (*big.Int, error) {
	return _Anycall.Contract.GetFeeConfigsCount(&_Anycall.CallOpts, tokenID)
}

// GetMPCPubkey is a free data retrieval call binding the contract method 0x9f1cdedd.
//
// Solidity: function getMPCPubkey(string mpcAddress) view returns(string)
func (_Anycall *AnycallCaller) GetMPCPubkey(opts *bind.CallOpts, mpcAddress string) (string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getMPCPubkey", mpcAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMPCPubkey is a free data retrieval call binding the contract method 0x9f1cdedd.
//
// Solidity: function getMPCPubkey(string mpcAddress) view returns(string)
func (_Anycall *AnycallSession) GetMPCPubkey(mpcAddress string) (string, error) {
	return _Anycall.Contract.GetMPCPubkey(&_Anycall.CallOpts, mpcAddress)
}

// GetMPCPubkey is a free data retrieval call binding the contract method 0x9f1cdedd.
//
// Solidity: function getMPCPubkey(string mpcAddress) view returns(string)
func (_Anycall *AnycallCallerSession) GetMPCPubkey(mpcAddress string) (string, error) {
	return _Anycall.Contract.GetMPCPubkey(&_Anycall.CallOpts, mpcAddress)
}

// GetMultichainToken is a free data retrieval call binding the contract method 0xb735ab5a.
//
// Solidity: function getMultichainToken(string tokenID, uint256 chainID) view returns(string)
func (_Anycall *AnycallCaller) GetMultichainToken(opts *bind.CallOpts, tokenID string, chainID *big.Int) (string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getMultichainToken", tokenID, chainID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMultichainToken is a free data retrieval call binding the contract method 0xb735ab5a.
//
// Solidity: function getMultichainToken(string tokenID, uint256 chainID) view returns(string)
func (_Anycall *AnycallSession) GetMultichainToken(tokenID string, chainID *big.Int) (string, error) {
	return _Anycall.Contract.GetMultichainToken(&_Anycall.CallOpts, tokenID, chainID)
}

// GetMultichainToken is a free data retrieval call binding the contract method 0xb735ab5a.
//
// Solidity: function getMultichainToken(string tokenID, uint256 chainID) view returns(string)
func (_Anycall *AnycallCallerSession) GetMultichainToken(tokenID string, chainID *big.Int) (string, error) {
	return _Anycall.Contract.GetMultichainToken(&_Anycall.CallOpts, tokenID, chainID)
}

// GetOriginalTokenConfig is a free data retrieval call binding the contract method 0x339d7ed8.
//
// Solidity: function getOriginalTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Anycall *AnycallCaller) GetOriginalTokenConfig(opts *bind.CallOpts, tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getOriginalTokenConfig", tokenID, chainID)

	if err != nil {
		return *new(RouterConfigTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigTokenConfig)).(*RouterConfigTokenConfig)

	return out0, err

}

// GetOriginalTokenConfig is a free data retrieval call binding the contract method 0x339d7ed8.
//
// Solidity: function getOriginalTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Anycall *AnycallSession) GetOriginalTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Anycall.Contract.GetOriginalTokenConfig(&_Anycall.CallOpts, tokenID, chainID)
}

// GetOriginalTokenConfig is a free data retrieval call binding the contract method 0x339d7ed8.
//
// Solidity: function getOriginalTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Anycall *AnycallCallerSession) GetOriginalTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Anycall.Contract.GetOriginalTokenConfig(&_Anycall.CallOpts, tokenID, chainID)
}

// GetSwapConfig is a free data retrieval call binding the contract method 0x4da7163c.
//
// Solidity: function getSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Anycall *AnycallCaller) GetSwapConfig(opts *bind.CallOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigSwapConfig, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getSwapConfig", tokenID, srcChainID, dstChainID)

	if err != nil {
		return *new(RouterConfigSwapConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigSwapConfig)).(*RouterConfigSwapConfig)

	return out0, err

}

// GetSwapConfig is a free data retrieval call binding the contract method 0x4da7163c.
//
// Solidity: function getSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Anycall *AnycallSession) GetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigSwapConfig, error) {
	return _Anycall.Contract.GetSwapConfig(&_Anycall.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetSwapConfig is a free data retrieval call binding the contract method 0x4da7163c.
//
// Solidity: function getSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Anycall *AnycallCallerSession) GetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigSwapConfig, error) {
	return _Anycall.Contract.GetSwapConfig(&_Anycall.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetSwapConfigAtIndex is a free data retrieval call binding the contract method 0xd825b3ef.
//
// Solidity: function getSwapConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Anycall *AnycallCaller) GetSwapConfigAtIndex(opts *bind.CallOpts, tokenID string, index *big.Int) (RouterConfigSwapConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getSwapConfigAtIndex", tokenID, index)

	if err != nil {
		return *new(RouterConfigSwapConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigSwapConfig2)).(*RouterConfigSwapConfig2)

	return out0, err

}

// GetSwapConfigAtIndex is a free data retrieval call binding the contract method 0xd825b3ef.
//
// Solidity: function getSwapConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Anycall *AnycallSession) GetSwapConfigAtIndex(tokenID string, index *big.Int) (RouterConfigSwapConfig2, error) {
	return _Anycall.Contract.GetSwapConfigAtIndex(&_Anycall.CallOpts, tokenID, index)
}

// GetSwapConfigAtIndex is a free data retrieval call binding the contract method 0xd825b3ef.
//
// Solidity: function getSwapConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Anycall *AnycallCallerSession) GetSwapConfigAtIndex(tokenID string, index *big.Int) (RouterConfigSwapConfig2, error) {
	return _Anycall.Contract.GetSwapConfigAtIndex(&_Anycall.CallOpts, tokenID, index)
}

// GetSwapConfigAtIndexRange is a free data retrieval call binding the contract method 0xc03411ba.
//
// Solidity: function getSwapConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCaller) GetSwapConfigAtIndexRange(opts *bind.CallOpts, tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigSwapConfig2, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getSwapConfigAtIndexRange", tokenID, startIndex, endIndex)

	if err != nil {
		return *new([]RouterConfigSwapConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigSwapConfig2)).(*[]RouterConfigSwapConfig2)

	return out0, err

}

// GetSwapConfigAtIndexRange is a free data retrieval call binding the contract method 0xc03411ba.
//
// Solidity: function getSwapConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallSession) GetSwapConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigSwapConfig2, error) {
	return _Anycall.Contract.GetSwapConfigAtIndexRange(&_Anycall.CallOpts, tokenID, startIndex, endIndex)
}

// GetSwapConfigAtIndexRange is a free data retrieval call binding the contract method 0xc03411ba.
//
// Solidity: function getSwapConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Anycall *AnycallCallerSession) GetSwapConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigSwapConfig2, error) {
	return _Anycall.Contract.GetSwapConfigAtIndexRange(&_Anycall.CallOpts, tokenID, startIndex, endIndex)
}

// GetSwapConfigsCount is a free data retrieval call binding the contract method 0x06e9a46a.
//
// Solidity: function getSwapConfigsCount(string tokenID) view returns(uint256)
func (_Anycall *AnycallCaller) GetSwapConfigsCount(opts *bind.CallOpts, tokenID string) (*big.Int, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getSwapConfigsCount", tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapConfigsCount is a free data retrieval call binding the contract method 0x06e9a46a.
//
// Solidity: function getSwapConfigsCount(string tokenID) view returns(uint256)
func (_Anycall *AnycallSession) GetSwapConfigsCount(tokenID string) (*big.Int, error) {
	return _Anycall.Contract.GetSwapConfigsCount(&_Anycall.CallOpts, tokenID)
}

// GetSwapConfigsCount is a free data retrieval call binding the contract method 0x06e9a46a.
//
// Solidity: function getSwapConfigsCount(string tokenID) view returns(uint256)
func (_Anycall *AnycallCallerSession) GetSwapConfigsCount(tokenID string) (*big.Int, error) {
	return _Anycall.Contract.GetSwapConfigsCount(&_Anycall.CallOpts, tokenID)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x459511d1.
//
// Solidity: function getTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Anycall *AnycallCaller) GetTokenConfig(opts *bind.CallOpts, tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getTokenConfig", tokenID, chainID)

	if err != nil {
		return *new(RouterConfigTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigTokenConfig)).(*RouterConfigTokenConfig)

	return out0, err

}

// GetTokenConfig is a free data retrieval call binding the contract method 0x459511d1.
//
// Solidity: function getTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Anycall *AnycallSession) GetTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Anycall.Contract.GetTokenConfig(&_Anycall.CallOpts, tokenID, chainID)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x459511d1.
//
// Solidity: function getTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Anycall *AnycallCallerSession) GetTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Anycall.Contract.GetTokenConfig(&_Anycall.CallOpts, tokenID, chainID)
}

// GetTokenID is a free data retrieval call binding the contract method 0xabad7b5c.
//
// Solidity: function getTokenID(uint256 chainID, string tokenAddress) view returns(string)
func (_Anycall *AnycallCaller) GetTokenID(opts *bind.CallOpts, chainID *big.Int, tokenAddress string) (string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getTokenID", chainID, tokenAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0xabad7b5c.
//
// Solidity: function getTokenID(uint256 chainID, string tokenAddress) view returns(string)
func (_Anycall *AnycallSession) GetTokenID(chainID *big.Int, tokenAddress string) (string, error) {
	return _Anycall.Contract.GetTokenID(&_Anycall.CallOpts, chainID, tokenAddress)
}

// GetTokenID is a free data retrieval call binding the contract method 0xabad7b5c.
//
// Solidity: function getTokenID(uint256 chainID, string tokenAddress) view returns(string)
func (_Anycall *AnycallCallerSession) GetTokenID(chainID *big.Int, tokenAddress string) (string, error) {
	return _Anycall.Contract.GetTokenID(&_Anycall.CallOpts, chainID, tokenAddress)
}

// GetTokenIDByIndex is a free data retrieval call binding the contract method 0xe828d00f.
//
// Solidity: function getTokenIDByIndex(uint256 index) view returns(string)
func (_Anycall *AnycallCaller) GetTokenIDByIndex(opts *bind.CallOpts, index *big.Int) (string, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "getTokenIDByIndex", index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenIDByIndex is a free data retrieval call binding the contract method 0xe828d00f.
//
// Solidity: function getTokenIDByIndex(uint256 index) view returns(string)
func (_Anycall *AnycallSession) GetTokenIDByIndex(index *big.Int) (string, error) {
	return _Anycall.Contract.GetTokenIDByIndex(&_Anycall.CallOpts, index)
}

// GetTokenIDByIndex is a free data retrieval call binding the contract method 0xe828d00f.
//
// Solidity: function getTokenIDByIndex(uint256 index) view returns(string)
func (_Anycall *AnycallCallerSession) GetTokenIDByIndex(index *big.Int) (string, error) {
	return _Anycall.Contract.GetTokenIDByIndex(&_Anycall.CallOpts, index)
}

// IsChainIDExist is a free data retrieval call binding the contract method 0xfd15ea70.
//
// Solidity: function isChainIDExist(uint256 chainID) view returns(bool)
func (_Anycall *AnycallCaller) IsChainIDExist(opts *bind.CallOpts, chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "isChainIDExist", chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChainIDExist is a free data retrieval call binding the contract method 0xfd15ea70.
//
// Solidity: function isChainIDExist(uint256 chainID) view returns(bool)
func (_Anycall *AnycallSession) IsChainIDExist(chainID *big.Int) (bool, error) {
	return _Anycall.Contract.IsChainIDExist(&_Anycall.CallOpts, chainID)
}

// IsChainIDExist is a free data retrieval call binding the contract method 0xfd15ea70.
//
// Solidity: function isChainIDExist(uint256 chainID) view returns(bool)
func (_Anycall *AnycallCallerSession) IsChainIDExist(chainID *big.Int) (bool, error) {
	return _Anycall.Contract.IsChainIDExist(&_Anycall.CallOpts, chainID)
}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_Anycall *AnycallCaller) IsTokenIDExist(opts *bind.CallOpts, tokenID string) (bool, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "isTokenIDExist", tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_Anycall *AnycallSession) IsTokenIDExist(tokenID string) (bool, error) {
	return _Anycall.Contract.IsTokenIDExist(&_Anycall.CallOpts, tokenID)
}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_Anycall *AnycallCallerSession) IsTokenIDExist(tokenID string) (bool, error) {
	return _Anycall.Contract.IsTokenIDExist(&_Anycall.CallOpts, tokenID)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[2])
func (_Anycall *AnycallCaller) Owners(opts *bind.CallOpts) ([2]common.Address, error) {
	var out []interface{}
	err := _Anycall.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[2])
func (_Anycall *AnycallSession) Owners() ([2]common.Address, error) {
	return _Anycall.Contract.Owners(&_Anycall.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[2])
func (_Anycall *AnycallCallerSession) Owners() ([2]common.Address, error) {
	return _Anycall.Contract.Owners(&_Anycall.CallOpts)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns(bool)
func (_Anycall *AnycallTransactor) AddChainID(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "addChainID", chainID)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns(bool)
func (_Anycall *AnycallSession) AddChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.AddChainID(&_Anycall.TransactOpts, chainID)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns(bool)
func (_Anycall *AnycallTransactorSession) AddChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.AddChainID(&_Anycall.TransactOpts, chainID)
}

// AddTokenID is a paid mutator transaction binding the contract method 0xe3e01664.
//
// Solidity: function addTokenID(string tokenID) returns(bool)
func (_Anycall *AnycallTransactor) AddTokenID(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "addTokenID", tokenID)
}

// AddTokenID is a paid mutator transaction binding the contract method 0xe3e01664.
//
// Solidity: function addTokenID(string tokenID) returns(bool)
func (_Anycall *AnycallSession) AddTokenID(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.AddTokenID(&_Anycall.TransactOpts, tokenID)
}

// AddTokenID is a paid mutator transaction binding the contract method 0xe3e01664.
//
// Solidity: function addTokenID(string tokenID) returns(bool)
func (_Anycall *AnycallTransactorSession) AddTokenID(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.AddTokenID(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllChainConfig is a paid mutator transaction binding the contract method 0x909b4225.
//
// Solidity: function removeAllChainConfig() returns()
func (_Anycall *AnycallTransactor) RemoveAllChainConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllChainConfig")
}

// RemoveAllChainConfig is a paid mutator transaction binding the contract method 0x909b4225.
//
// Solidity: function removeAllChainConfig() returns()
func (_Anycall *AnycallSession) RemoveAllChainConfig() (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllChainConfig(&_Anycall.TransactOpts)
}

// RemoveAllChainConfig is a paid mutator transaction binding the contract method 0x909b4225.
//
// Solidity: function removeAllChainConfig() returns()
func (_Anycall *AnycallTransactorSession) RemoveAllChainConfig() (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllChainConfig(&_Anycall.TransactOpts)
}

// RemoveAllDestFeeConfig is a paid mutator transaction binding the contract method 0xb5f59a96.
//
// Solidity: function removeAllDestFeeConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveAllDestFeeConfig(opts *bind.TransactOpts, tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllDestFeeConfig", tokenID, dstChainIDs)
}

// RemoveAllDestFeeConfig is a paid mutator transaction binding the contract method 0xb5f59a96.
//
// Solidity: function removeAllDestFeeConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallSession) RemoveAllDestFeeConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllDestFeeConfig(&_Anycall.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllDestFeeConfig is a paid mutator transaction binding the contract method 0xb5f59a96.
//
// Solidity: function removeAllDestFeeConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllDestFeeConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllDestFeeConfig(&_Anycall.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllDestSwapConfig is a paid mutator transaction binding the contract method 0x267d8ea6.
//
// Solidity: function removeAllDestSwapConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveAllDestSwapConfig(opts *bind.TransactOpts, tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllDestSwapConfig", tokenID, dstChainIDs)
}

// RemoveAllDestSwapConfig is a paid mutator transaction binding the contract method 0x267d8ea6.
//
// Solidity: function removeAllDestSwapConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallSession) RemoveAllDestSwapConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllDestSwapConfig(&_Anycall.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllDestSwapConfig is a paid mutator transaction binding the contract method 0x267d8ea6.
//
// Solidity: function removeAllDestSwapConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllDestSwapConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllDestSwapConfig(&_Anycall.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllFeeConfig is a paid mutator transaction binding the contract method 0x0a2b6b13.
//
// Solidity: function removeAllFeeConfig(string tokenID) returns()
func (_Anycall *AnycallTransactor) RemoveAllFeeConfig(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllFeeConfig", tokenID)
}

// RemoveAllFeeConfig is a paid mutator transaction binding the contract method 0x0a2b6b13.
//
// Solidity: function removeAllFeeConfig(string tokenID) returns()
func (_Anycall *AnycallSession) RemoveAllFeeConfig(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllFeeConfig(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllFeeConfig is a paid mutator transaction binding the contract method 0x0a2b6b13.
//
// Solidity: function removeAllFeeConfig(string tokenID) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllFeeConfig(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllFeeConfig(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllMultichainTokens is a paid mutator transaction binding the contract method 0x83f5d5d4.
//
// Solidity: function removeAllMultichainTokens(string tokenID) returns()
func (_Anycall *AnycallTransactor) RemoveAllMultichainTokens(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllMultichainTokens", tokenID)
}

// RemoveAllMultichainTokens is a paid mutator transaction binding the contract method 0x83f5d5d4.
//
// Solidity: function removeAllMultichainTokens(string tokenID) returns()
func (_Anycall *AnycallSession) RemoveAllMultichainTokens(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllMultichainTokens(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllMultichainTokens is a paid mutator transaction binding the contract method 0x83f5d5d4.
//
// Solidity: function removeAllMultichainTokens(string tokenID) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllMultichainTokens(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllMultichainTokens(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllSrcFeeConfig is a paid mutator transaction binding the contract method 0xdcd59579.
//
// Solidity: function removeAllSrcFeeConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveAllSrcFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllSrcFeeConfig", tokenID, srcChainIDs)
}

// RemoveAllSrcFeeConfig is a paid mutator transaction binding the contract method 0xdcd59579.
//
// Solidity: function removeAllSrcFeeConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Anycall *AnycallSession) RemoveAllSrcFeeConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllSrcFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSrcFeeConfig is a paid mutator transaction binding the contract method 0xdcd59579.
//
// Solidity: function removeAllSrcFeeConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllSrcFeeConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllSrcFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSrcSwapConfig is a paid mutator transaction binding the contract method 0x103a24a3.
//
// Solidity: function removeAllSrcSwapConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveAllSrcSwapConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllSrcSwapConfig", tokenID, srcChainIDs)
}

// RemoveAllSrcSwapConfig is a paid mutator transaction binding the contract method 0x103a24a3.
//
// Solidity: function removeAllSrcSwapConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Anycall *AnycallSession) RemoveAllSrcSwapConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllSrcSwapConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSrcSwapConfig is a paid mutator transaction binding the contract method 0x103a24a3.
//
// Solidity: function removeAllSrcSwapConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllSrcSwapConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllSrcSwapConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSwapConfig is a paid mutator transaction binding the contract method 0xa7bdc603.
//
// Solidity: function removeAllSwapConfig(string tokenID) returns()
func (_Anycall *AnycallTransactor) RemoveAllSwapConfig(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllSwapConfig", tokenID)
}

// RemoveAllSwapConfig is a paid mutator transaction binding the contract method 0xa7bdc603.
//
// Solidity: function removeAllSwapConfig(string tokenID) returns()
func (_Anycall *AnycallSession) RemoveAllSwapConfig(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllSwapConfig(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllSwapConfig is a paid mutator transaction binding the contract method 0xa7bdc603.
//
// Solidity: function removeAllSwapConfig(string tokenID) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllSwapConfig(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllSwapConfig(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllTokenConfig is a paid mutator transaction binding the contract method 0x0c4fe066.
//
// Solidity: function removeAllTokenConfig(string tokenID) returns()
func (_Anycall *AnycallTransactor) RemoveAllTokenConfig(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeAllTokenConfig", tokenID)
}

// RemoveAllTokenConfig is a paid mutator transaction binding the contract method 0x0c4fe066.
//
// Solidity: function removeAllTokenConfig(string tokenID) returns()
func (_Anycall *AnycallSession) RemoveAllTokenConfig(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllTokenConfig(&_Anycall.TransactOpts, tokenID)
}

// RemoveAllTokenConfig is a paid mutator transaction binding the contract method 0x0c4fe066.
//
// Solidity: function removeAllTokenConfig(string tokenID) returns()
func (_Anycall *AnycallTransactorSession) RemoveAllTokenConfig(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveAllTokenConfig(&_Anycall.TransactOpts, tokenID)
}

// RemoveChainConfig is a paid mutator transaction binding the contract method 0xb00f5cc8.
//
// Solidity: function removeChainConfig(uint256[] chainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveChainConfig(opts *bind.TransactOpts, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeChainConfig", chainIDs)
}

// RemoveChainConfig is a paid mutator transaction binding the contract method 0xb00f5cc8.
//
// Solidity: function removeChainConfig(uint256[] chainIDs) returns()
func (_Anycall *AnycallSession) RemoveChainConfig(chainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveChainConfig(&_Anycall.TransactOpts, chainIDs)
}

// RemoveChainConfig is a paid mutator transaction binding the contract method 0xb00f5cc8.
//
// Solidity: function removeChainConfig(uint256[] chainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveChainConfig(chainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveChainConfig(&_Anycall.TransactOpts, chainIDs)
}

// RemoveChainID is a paid mutator transaction binding the contract method 0xe90768c6.
//
// Solidity: function removeChainID(uint256 chainID) returns()
func (_Anycall *AnycallTransactor) RemoveChainID(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeChainID", chainID)
}

// RemoveChainID is a paid mutator transaction binding the contract method 0xe90768c6.
//
// Solidity: function removeChainID(uint256 chainID) returns()
func (_Anycall *AnycallSession) RemoveChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveChainID(&_Anycall.TransactOpts, chainID)
}

// RemoveChainID is a paid mutator transaction binding the contract method 0xe90768c6.
//
// Solidity: function removeChainID(uint256 chainID) returns()
func (_Anycall *AnycallTransactorSession) RemoveChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveChainID(&_Anycall.TransactOpts, chainID)
}

// RemoveFeeConfig is a paid mutator transaction binding the contract method 0xb4a51ac4.
//
// Solidity: function removeFeeConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeFeeConfig", tokenID, srcChainIDs, dstChainIDs)
}

// RemoveFeeConfig is a paid mutator transaction binding the contract method 0xb4a51ac4.
//
// Solidity: function removeFeeConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallSession) RemoveFeeConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveFeeConfig is a paid mutator transaction binding the contract method 0xb4a51ac4.
//
// Solidity: function removeFeeConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveFeeConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveMPCPubkey is a paid mutator transaction binding the contract method 0x04b75939.
//
// Solidity: function removeMPCPubkey(string addr) returns()
func (_Anycall *AnycallTransactor) RemoveMPCPubkey(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeMPCPubkey", addr)
}

// RemoveMPCPubkey is a paid mutator transaction binding the contract method 0x04b75939.
//
// Solidity: function removeMPCPubkey(string addr) returns()
func (_Anycall *AnycallSession) RemoveMPCPubkey(addr string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveMPCPubkey(&_Anycall.TransactOpts, addr)
}

// RemoveMPCPubkey is a paid mutator transaction binding the contract method 0x04b75939.
//
// Solidity: function removeMPCPubkey(string addr) returns()
func (_Anycall *AnycallTransactorSession) RemoveMPCPubkey(addr string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveMPCPubkey(&_Anycall.TransactOpts, addr)
}

// RemoveMultichainToken is a paid mutator transaction binding the contract method 0x9026b205.
//
// Solidity: function removeMultichainToken(string tokenID, uint256 chainID) returns()
func (_Anycall *AnycallTransactor) RemoveMultichainToken(opts *bind.TransactOpts, tokenID string, chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeMultichainToken", tokenID, chainID)
}

// RemoveMultichainToken is a paid mutator transaction binding the contract method 0x9026b205.
//
// Solidity: function removeMultichainToken(string tokenID, uint256 chainID) returns()
func (_Anycall *AnycallSession) RemoveMultichainToken(tokenID string, chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveMultichainToken(&_Anycall.TransactOpts, tokenID, chainID)
}

// RemoveMultichainToken is a paid mutator transaction binding the contract method 0x9026b205.
//
// Solidity: function removeMultichainToken(string tokenID, uint256 chainID) returns()
func (_Anycall *AnycallTransactorSession) RemoveMultichainToken(tokenID string, chainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveMultichainToken(&_Anycall.TransactOpts, tokenID, chainID)
}

// RemoveSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x8f2cfccb.
//
// Solidity: function removeSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) returns()
func (_Anycall *AnycallTransactor) RemoveSwapAndFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeSwapAndFeeConfig", tokenID, srcChainID, dstChainID)
}

// RemoveSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x8f2cfccb.
//
// Solidity: function removeSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) returns()
func (_Anycall *AnycallSession) RemoveSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveSwapAndFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID)
}

// RemoveSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x8f2cfccb.
//
// Solidity: function removeSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) returns()
func (_Anycall *AnycallTransactorSession) RemoveSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveSwapAndFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID)
}

// RemoveSwapConfig is a paid mutator transaction binding the contract method 0x9fdc851a.
//
// Solidity: function removeSwapConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveSwapConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeSwapConfig", tokenID, srcChainIDs, dstChainIDs)
}

// RemoveSwapConfig is a paid mutator transaction binding the contract method 0x9fdc851a.
//
// Solidity: function removeSwapConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallSession) RemoveSwapConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveSwapConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveSwapConfig is a paid mutator transaction binding the contract method 0x9fdc851a.
//
// Solidity: function removeSwapConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveSwapConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveSwapConfig(&_Anycall.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xb9b85d34.
//
// Solidity: function removeTokenConfig(string tokenID, uint256[] chainIDs) returns()
func (_Anycall *AnycallTransactor) RemoveTokenConfig(opts *bind.TransactOpts, tokenID string, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeTokenConfig", tokenID, chainIDs)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xb9b85d34.
//
// Solidity: function removeTokenConfig(string tokenID, uint256[] chainIDs) returns()
func (_Anycall *AnycallSession) RemoveTokenConfig(tokenID string, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveTokenConfig(&_Anycall.TransactOpts, tokenID, chainIDs)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xb9b85d34.
//
// Solidity: function removeTokenConfig(string tokenID, uint256[] chainIDs) returns()
func (_Anycall *AnycallTransactorSession) RemoveTokenConfig(tokenID string, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveTokenConfig(&_Anycall.TransactOpts, tokenID, chainIDs)
}

// RemoveTokenID is a paid mutator transaction binding the contract method 0x5d58d0a1.
//
// Solidity: function removeTokenID(string tokenID) returns()
func (_Anycall *AnycallTransactor) RemoveTokenID(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "removeTokenID", tokenID)
}

// RemoveTokenID is a paid mutator transaction binding the contract method 0x5d58d0a1.
//
// Solidity: function removeTokenID(string tokenID) returns()
func (_Anycall *AnycallSession) RemoveTokenID(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveTokenID(&_Anycall.TransactOpts, tokenID)
}

// RemoveTokenID is a paid mutator transaction binding the contract method 0x5d58d0a1.
//
// Solidity: function removeTokenID(string tokenID) returns()
func (_Anycall *AnycallTransactorSession) RemoveTokenID(tokenID string) (*types.Transaction, error) {
	return _Anycall.Contract.RemoveTokenID(&_Anycall.TransactOpts, tokenID)
}

// SetChainConfig is a paid mutator transaction binding the contract method 0x2383006c.
//
// Solidity: function setChainConfig(uint256 chainID, string blockChain, string routerContract, uint64 confirmations, uint64 initialHeight, string extra) returns(bool)
func (_Anycall *AnycallTransactor) SetChainConfig(opts *bind.TransactOpts, chainID *big.Int, blockChain string, routerContract string, confirmations uint64, initialHeight uint64, extra string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setChainConfig", chainID, blockChain, routerContract, confirmations, initialHeight, extra)
}

// SetChainConfig is a paid mutator transaction binding the contract method 0x2383006c.
//
// Solidity: function setChainConfig(uint256 chainID, string blockChain, string routerContract, uint64 confirmations, uint64 initialHeight, string extra) returns(bool)
func (_Anycall *AnycallSession) SetChainConfig(chainID *big.Int, blockChain string, routerContract string, confirmations uint64, initialHeight uint64, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetChainConfig(&_Anycall.TransactOpts, chainID, blockChain, routerContract, confirmations, initialHeight, extra)
}

// SetChainConfig is a paid mutator transaction binding the contract method 0x2383006c.
//
// Solidity: function setChainConfig(uint256 chainID, string blockChain, string routerContract, uint64 confirmations, uint64 initialHeight, string extra) returns(bool)
func (_Anycall *AnycallTransactorSession) SetChainConfig(chainID *big.Int, blockChain string, routerContract string, confirmations uint64, initialHeight uint64, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetChainConfig(&_Anycall.TransactOpts, chainID, blockChain, routerContract, confirmations, initialHeight, extra)
}

// SetChainExtraConfig is a paid mutator transaction binding the contract method 0xd236a09b.
//
// Solidity: function setChainExtraConfig(uint256 chainID, string extra) returns(bool)
func (_Anycall *AnycallTransactor) SetChainExtraConfig(opts *bind.TransactOpts, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setChainExtraConfig", chainID, extra)
}

// SetChainExtraConfig is a paid mutator transaction binding the contract method 0xd236a09b.
//
// Solidity: function setChainExtraConfig(uint256 chainID, string extra) returns(bool)
func (_Anycall *AnycallSession) SetChainExtraConfig(chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetChainExtraConfig(&_Anycall.TransactOpts, chainID, extra)
}

// SetChainExtraConfig is a paid mutator transaction binding the contract method 0xd236a09b.
//
// Solidity: function setChainExtraConfig(uint256 chainID, string extra) returns(bool)
func (_Anycall *AnycallTransactorSession) SetChainExtraConfig(chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetChainExtraConfig(&_Anycall.TransactOpts, chainID, extra)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0xe671242a.
//
// Solidity: function setCustomConfig(uint256 chainID, string key, string data) returns()
func (_Anycall *AnycallTransactor) SetCustomConfig(opts *bind.TransactOpts, chainID *big.Int, key string, data string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setCustomConfig", chainID, key, data)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0xe671242a.
//
// Solidity: function setCustomConfig(uint256 chainID, string key, string data) returns()
func (_Anycall *AnycallSession) SetCustomConfig(chainID *big.Int, key string, data string) (*types.Transaction, error) {
	return _Anycall.Contract.SetCustomConfig(&_Anycall.TransactOpts, chainID, key, data)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0xe671242a.
//
// Solidity: function setCustomConfig(uint256 chainID, string key, string data) returns()
func (_Anycall *AnycallTransactorSession) SetCustomConfig(chainID *big.Int, key string, data string) (*types.Transaction, error) {
	return _Anycall.Contract.SetCustomConfig(&_Anycall.TransactOpts, chainID, key, data)
}

// SetExtraConfig is a paid mutator transaction binding the contract method 0x04cbc75b.
//
// Solidity: function setExtraConfig(string key, string data) returns()
func (_Anycall *AnycallTransactor) SetExtraConfig(opts *bind.TransactOpts, key string, data string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setExtraConfig", key, data)
}

// SetExtraConfig is a paid mutator transaction binding the contract method 0x04cbc75b.
//
// Solidity: function setExtraConfig(string key, string data) returns()
func (_Anycall *AnycallSession) SetExtraConfig(key string, data string) (*types.Transaction, error) {
	return _Anycall.Contract.SetExtraConfig(&_Anycall.TransactOpts, key, data)
}

// SetExtraConfig is a paid mutator transaction binding the contract method 0x04cbc75b.
//
// Solidity: function setExtraConfig(string key, string data) returns()
func (_Anycall *AnycallTransactorSession) SetExtraConfig(key string, data string) (*types.Transaction, error) {
	return _Anycall.Contract.SetExtraConfig(&_Anycall.TransactOpts, key, data)
}

// SetFeeConfig is a paid mutator transaction binding the contract method 0x2f50cf14.
//
// Solidity: function setFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Anycall *AnycallTransactor) SetFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setFeeConfig", tokenID, srcChainID, dstChainID, maxFee, minFee, feeRate)
}

// SetFeeConfig is a paid mutator transaction binding the contract method 0x2f50cf14.
//
// Solidity: function setFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Anycall *AnycallSession) SetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.SetFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID, maxFee, minFee, feeRate)
}

// SetFeeConfig is a paid mutator transaction binding the contract method 0x2f50cf14.
//
// Solidity: function setFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Anycall *AnycallTransactorSession) SetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.SetFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID, maxFee, minFee, feeRate)
}

// SetFeeConfigs is a paid mutator transaction binding the contract method 0x9a29eaf2.
//
// Solidity: function setFeeConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Anycall *AnycallTransactor) SetFeeConfigs(opts *bind.TransactOpts, tokenID string, configs []RouterConfigFeeConfig2) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setFeeConfigs", tokenID, configs)
}

// SetFeeConfigs is a paid mutator transaction binding the contract method 0x9a29eaf2.
//
// Solidity: function setFeeConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Anycall *AnycallSession) SetFeeConfigs(tokenID string, configs []RouterConfigFeeConfig2) (*types.Transaction, error) {
	return _Anycall.Contract.SetFeeConfigs(&_Anycall.TransactOpts, tokenID, configs)
}

// SetFeeConfigs is a paid mutator transaction binding the contract method 0x9a29eaf2.
//
// Solidity: function setFeeConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Anycall *AnycallTransactorSession) SetFeeConfigs(tokenID string, configs []RouterConfigFeeConfig2) (*types.Transaction, error) {
	return _Anycall.Contract.SetFeeConfigs(&_Anycall.TransactOpts, tokenID, configs)
}

// SetMPCPubkey is a paid mutator transaction binding the contract method 0xb4d8d47c.
//
// Solidity: function setMPCPubkey(string addr, string pubkey) returns()
func (_Anycall *AnycallTransactor) SetMPCPubkey(opts *bind.TransactOpts, addr string, pubkey string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setMPCPubkey", addr, pubkey)
}

// SetMPCPubkey is a paid mutator transaction binding the contract method 0xb4d8d47c.
//
// Solidity: function setMPCPubkey(string addr, string pubkey) returns()
func (_Anycall *AnycallSession) SetMPCPubkey(addr string, pubkey string) (*types.Transaction, error) {
	return _Anycall.Contract.SetMPCPubkey(&_Anycall.TransactOpts, addr, pubkey)
}

// SetMPCPubkey is a paid mutator transaction binding the contract method 0xb4d8d47c.
//
// Solidity: function setMPCPubkey(string addr, string pubkey) returns()
func (_Anycall *AnycallTransactorSession) SetMPCPubkey(addr string, pubkey string) (*types.Transaction, error) {
	return _Anycall.Contract.SetMPCPubkey(&_Anycall.TransactOpts, addr, pubkey)
}

// SetMultichainToken is a paid mutator transaction binding the contract method 0xf4730a7e.
//
// Solidity: function setMultichainToken(string tokenID, uint256 chainID, string token) returns()
func (_Anycall *AnycallTransactor) SetMultichainToken(opts *bind.TransactOpts, tokenID string, chainID *big.Int, token string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setMultichainToken", tokenID, chainID, token)
}

// SetMultichainToken is a paid mutator transaction binding the contract method 0xf4730a7e.
//
// Solidity: function setMultichainToken(string tokenID, uint256 chainID, string token) returns()
func (_Anycall *AnycallSession) SetMultichainToken(tokenID string, chainID *big.Int, token string) (*types.Transaction, error) {
	return _Anycall.Contract.SetMultichainToken(&_Anycall.TransactOpts, tokenID, chainID, token)
}

// SetMultichainToken is a paid mutator transaction binding the contract method 0xf4730a7e.
//
// Solidity: function setMultichainToken(string tokenID, uint256 chainID, string token) returns()
func (_Anycall *AnycallTransactorSession) SetMultichainToken(tokenID string, chainID *big.Int, token string) (*types.Transaction, error) {
	return _Anycall.Contract.SetMultichainToken(&_Anycall.TransactOpts, tokenID, chainID, token)
}

// SetSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x7926422c.
//
// Solidity: function setSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Anycall *AnycallTransactor) SetSwapAndFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setSwapAndFeeConfig", tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap, maxFee, minFee, feeRate)
}

// SetSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x7926422c.
//
// Solidity: function setSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Anycall *AnycallSession) SetSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.SetSwapAndFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap, maxFee, minFee, feeRate)
}

// SetSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x7926422c.
//
// Solidity: function setSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Anycall *AnycallTransactorSession) SetSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.SetSwapAndFeeConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap, maxFee, minFee, feeRate)
}

// SetSwapConfig is a paid mutator transaction binding the contract method 0x1dc463b4.
//
// Solidity: function setSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap) returns(bool)
func (_Anycall *AnycallTransactor) SetSwapConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setSwapConfig", tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap)
}

// SetSwapConfig is a paid mutator transaction binding the contract method 0x1dc463b4.
//
// Solidity: function setSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap) returns(bool)
func (_Anycall *AnycallSession) SetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.SetSwapConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap)
}

// SetSwapConfig is a paid mutator transaction binding the contract method 0x1dc463b4.
//
// Solidity: function setSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap) returns(bool)
func (_Anycall *AnycallTransactorSession) SetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int) (*types.Transaction, error) {
	return _Anycall.Contract.SetSwapConfig(&_Anycall.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap)
}

// SetSwapConfigs is a paid mutator transaction binding the contract method 0x6da91269.
//
// Solidity: function setSwapConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Anycall *AnycallTransactor) SetSwapConfigs(opts *bind.TransactOpts, tokenID string, configs []RouterConfigSwapConfig2) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setSwapConfigs", tokenID, configs)
}

// SetSwapConfigs is a paid mutator transaction binding the contract method 0x6da91269.
//
// Solidity: function setSwapConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Anycall *AnycallSession) SetSwapConfigs(tokenID string, configs []RouterConfigSwapConfig2) (*types.Transaction, error) {
	return _Anycall.Contract.SetSwapConfigs(&_Anycall.TransactOpts, tokenID, configs)
}

// SetSwapConfigs is a paid mutator transaction binding the contract method 0x6da91269.
//
// Solidity: function setSwapConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Anycall *AnycallTransactorSession) SetSwapConfigs(tokenID string, configs []RouterConfigSwapConfig2) (*types.Transaction, error) {
	return _Anycall.Contract.SetSwapConfigs(&_Anycall.TransactOpts, tokenID, configs)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x34fd0135.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddr, uint8 decimals, uint256 version, string routerContract, string extra) returns(bool)
func (_Anycall *AnycallTransactor) SetTokenConfig(opts *bind.TransactOpts, tokenID string, chainID *big.Int, tokenAddr string, decimals uint8, version *big.Int, routerContract string, extra string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setTokenConfig", tokenID, chainID, tokenAddr, decimals, version, routerContract, extra)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x34fd0135.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddr, uint8 decimals, uint256 version, string routerContract, string extra) returns(bool)
func (_Anycall *AnycallSession) SetTokenConfig(tokenID string, chainID *big.Int, tokenAddr string, decimals uint8, version *big.Int, routerContract string, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetTokenConfig(&_Anycall.TransactOpts, tokenID, chainID, tokenAddr, decimals, version, routerContract, extra)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x34fd0135.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddr, uint8 decimals, uint256 version, string routerContract, string extra) returns(bool)
func (_Anycall *AnycallTransactorSession) SetTokenConfig(tokenID string, chainID *big.Int, tokenAddr string, decimals uint8, version *big.Int, routerContract string, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetTokenConfig(&_Anycall.TransactOpts, tokenID, chainID, tokenAddr, decimals, version, routerContract, extra)
}

// SetTokenExtraConfig is a paid mutator transaction binding the contract method 0x84cd5b5a.
//
// Solidity: function setTokenExtraConfig(string tokenID, uint256 chainID, string extra) returns(bool)
func (_Anycall *AnycallTransactor) SetTokenExtraConfig(opts *bind.TransactOpts, tokenID string, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setTokenExtraConfig", tokenID, chainID, extra)
}

// SetTokenExtraConfig is a paid mutator transaction binding the contract method 0x84cd5b5a.
//
// Solidity: function setTokenExtraConfig(string tokenID, uint256 chainID, string extra) returns(bool)
func (_Anycall *AnycallSession) SetTokenExtraConfig(tokenID string, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetTokenExtraConfig(&_Anycall.TransactOpts, tokenID, chainID, extra)
}

// SetTokenExtraConfig is a paid mutator transaction binding the contract method 0x84cd5b5a.
//
// Solidity: function setTokenExtraConfig(string tokenID, uint256 chainID, string extra) returns(bool)
func (_Anycall *AnycallTransactorSession) SetTokenExtraConfig(tokenID string, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Anycall.Contract.SetTokenExtraConfig(&_Anycall.TransactOpts, tokenID, chainID, extra)
}

// SetTokenRouterContract is a paid mutator transaction binding the contract method 0x1910dbdd.
//
// Solidity: function setTokenRouterContract(string tokenID, uint256 chainID, string routerContract) returns(bool)
func (_Anycall *AnycallTransactor) SetTokenRouterContract(opts *bind.TransactOpts, tokenID string, chainID *big.Int, routerContract string) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "setTokenRouterContract", tokenID, chainID, routerContract)
}

// SetTokenRouterContract is a paid mutator transaction binding the contract method 0x1910dbdd.
//
// Solidity: function setTokenRouterContract(string tokenID, uint256 chainID, string routerContract) returns(bool)
func (_Anycall *AnycallSession) SetTokenRouterContract(tokenID string, chainID *big.Int, routerContract string) (*types.Transaction, error) {
	return _Anycall.Contract.SetTokenRouterContract(&_Anycall.TransactOpts, tokenID, chainID, routerContract)
}

// SetTokenRouterContract is a paid mutator transaction binding the contract method 0x1910dbdd.
//
// Solidity: function setTokenRouterContract(string tokenID, uint256 chainID, string routerContract) returns(bool)
func (_Anycall *AnycallTransactorSession) SetTokenRouterContract(tokenID string, chainID *big.Int, routerContract string) (*types.Transaction, error) {
	return _Anycall.Contract.SetTokenRouterContract(&_Anycall.TransactOpts, tokenID, chainID, routerContract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Anycall *AnycallTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Anycall *AnycallSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Anycall.Contract.TransferOwnership(&_Anycall.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Anycall *AnycallTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Anycall.Contract.TransferOwnership(&_Anycall.TransactOpts, newOwner)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x747fd44c.
//
// Solidity: function updateConfig() returns()
func (_Anycall *AnycallTransactor) UpdateConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anycall.contract.Transact(opts, "updateConfig")
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x747fd44c.
//
// Solidity: function updateConfig() returns()
func (_Anycall *AnycallSession) UpdateConfig() (*types.Transaction, error) {
	return _Anycall.Contract.UpdateConfig(&_Anycall.TransactOpts)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x747fd44c.
//
// Solidity: function updateConfig() returns()
func (_Anycall *AnycallTransactorSession) UpdateConfig() (*types.Transaction, error) {
	return _Anycall.Contract.UpdateConfig(&_Anycall.TransactOpts)
}

// AnycallUpdateConfigIterator is returned from FilterUpdateConfig and is used to iterate over the raw logs and unpacked data for UpdateConfig events raised by the Anycall contract.
type AnycallUpdateConfigIterator struct {
	Event *AnycallUpdateConfig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AnycallUpdateConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnycallUpdateConfig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AnycallUpdateConfig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AnycallUpdateConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnycallUpdateConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnycallUpdateConfig represents a UpdateConfig event raised by the Anycall contract.
type AnycallUpdateConfig struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateConfig is a free log retrieval operation binding the contract event 0x22590461e7ba17e1fe7580cb0ea47f283d3b2248f04873dfbe926d08fe4c5ab9.
//
// Solidity: event UpdateConfig()
func (_Anycall *AnycallFilterer) FilterUpdateConfig(opts *bind.FilterOpts) (*AnycallUpdateConfigIterator, error) {

	logs, sub, err := _Anycall.contract.FilterLogs(opts, "UpdateConfig")
	if err != nil {
		return nil, err
	}
	return &AnycallUpdateConfigIterator{contract: _Anycall.contract, event: "UpdateConfig", logs: logs, sub: sub}, nil
}

// WatchUpdateConfig is a free log subscription operation binding the contract event 0x22590461e7ba17e1fe7580cb0ea47f283d3b2248f04873dfbe926d08fe4c5ab9.
//
// Solidity: event UpdateConfig()
func (_Anycall *AnycallFilterer) WatchUpdateConfig(opts *bind.WatchOpts, sink chan<- *AnycallUpdateConfig) (event.Subscription, error) {

	logs, sub, err := _Anycall.contract.WatchLogs(opts, "UpdateConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnycallUpdateConfig)
				if err := _Anycall.contract.UnpackLog(event, "UpdateConfig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateConfig is a log parse operation binding the contract event 0x22590461e7ba17e1fe7580cb0ea47f283d3b2248f04873dfbe926d08fe4c5ab9.
//
// Solidity: event UpdateConfig()
func (_Anycall *AnycallFilterer) ParseUpdateConfig(log types.Log) (*AnycallUpdateConfig, error) {
	event := new(AnycallUpdateConfig)
	if err := _Anycall.contract.UnpackLog(event, "UpdateConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
