// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
    ABI: `[{"inputs":[{"internalType":"address[2]","name":"owners","type":"address[2]"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[],"name":"UpdateConfig","type":"event"},{"inputs":[],"name":"CONFIG_VERSION","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"addChainID","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"addTokenID","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getAllChainConfig","outputs":[{"components":[{"internalType":"uint256","name":"ChainID","type":"uint256"},{"internalType":"string","name":"BlockChain","type":"string"},{"internalType":"string","name":"RouterContract","type":"string"},{"internalType":"uint64","name":"Confirmations","type":"uint64"},{"internalType":"uint64","name":"InitialHeight","type":"uint64"},{"internalType":"string","name":"Extra","type":"string"}],"internalType":"struct RouterConfig.ChainConfig2[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getAllChainIDLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getAllChainIDs","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"getAllFeeConfigs","outputs":[{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwapFee","type":"uint256"},{"internalType":"uint256","name":"MinimumSwapFee","type":"uint256"},{"internalType":"uint256","name":"SwapFeeRatePerMillion","type":"uint256"}],"internalType":"struct RouterConfig.FeeConfig2[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"getAllMultichainTokenConfig","outputs":[{"components":[{"internalType":"uint256","name":"ChainID","type":"uint256"},{"internalType":"uint8","name":"Decimals","type":"uint8"},{"internalType":"string","name":"ContractAddress","type":"string"},{"internalType":"uint256","name":"ContractVersion","type":"uint256"},{"internalType":"string","name":"RouterContract","type":"string"},{"internalType":"string","name":"Extra","type":"string"}],"internalType":"struct RouterConfig.TokenConfig2[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"getAllMultichainTokens","outputs":[{"components":[{"internalType":"uint256","name":"ChainID","type":"uint256"},{"internalType":"string","name":"TokenAddress","type":"string"}],"internalType":"struct RouterConfig.MultichainToken[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"getAllSwapConfigs","outputs":[{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwap","type":"uint256"},{"internalType":"uint256","name":"MinimumSwap","type":"uint256"},{"internalType":"uint256","name":"BigValueThreshold","type":"uint256"}],"internalType":"struct RouterConfig.SwapConfig2[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getAllTokenIDLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getAllTokenIDs","outputs":[{"internalType":"string[]","name":"result","type":"string[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"getChainConfig","outputs":[{"components":[{"internalType":"string","name":"BlockChain","type":"string"},{"internalType":"string","name":"RouterContract","type":"string"},{"internalType":"uint64","name":"Confirmations","type":"uint64"},{"internalType":"uint64","name":"InitialHeight","type":"uint64"},{"internalType":"string","name":"Extra","type":"string"}],"internalType":"struct RouterConfig.ChainConfig","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getChainIDByIndex","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"key","type":"string"}],"name":"getCustomConfig","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"getExtraConfig","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"srcChainID","type":"uint256"},{"internalType":"uint256","name":"dstChainID","type":"uint256"}],"name":"getFeeConfig","outputs":[{"components":[{"internalType":"uint256","name":"MaximumSwapFee","type":"uint256"},{"internalType":"uint256","name":"MinimumSwapFee","type":"uint256"},{"internalType":"uint256","name":"SwapFeeRatePerMillion","type":"uint256"}],"internalType":"struct RouterConfig.FeeConfig","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getFeeConfigAtIndex","outputs":[{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwapFee","type":"uint256"},{"internalType":"uint256","name":"MinimumSwapFee","type":"uint256"},{"internalType":"uint256","name":"SwapFeeRatePerMillion","type":"uint256"}],"internalType":"struct RouterConfig.FeeConfig2","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"startIndex","type":"uint256"},{"internalType":"uint256","name":"endIndex","type":"uint256"}],"name":"getFeeConfigAtIndexRange","outputs":[{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwapFee","type":"uint256"},{"internalType":"uint256","name":"MinimumSwapFee","type":"uint256"},{"internalType":"uint256","name":"SwapFeeRatePerMillion","type":"uint256"}],"internalType":"struct RouterConfig.FeeConfig2[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"getFeeConfigsCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"mpcAddress","type":"string"}],"name":"getMPCPubkey","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"getMultichainToken","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"getOriginalTokenConfig","outputs":[{"components":[{"internalType":"uint8","name":"Decimals","type":"uint8"},{"internalType":"string","name":"ContractAddress","type":"string"},{"internalType":"uint256","name":"ContractVersion","type":"uint256"},{"internalType":"string","name":"RouterContract","type":"string"},{"internalType":"string","name":"Extra","type":"string"}],"internalType":"struct RouterConfig.TokenConfig","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"srcChainID","type":"uint256"},{"internalType":"uint256","name":"dstChainID","type":"uint256"}],"name":"getSwapConfig","outputs":[{"components":[{"internalType":"uint256","name":"MaximumSwap","type":"uint256"},{"internalType":"uint256","name":"MinimumSwap","type":"uint256"},{"internalType":"uint256","name":"BigValueThreshold","type":"uint256"}],"internalType":"struct RouterConfig.SwapConfig","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getSwapConfigAtIndex","outputs":[{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwap","type":"uint256"},{"internalType":"uint256","name":"MinimumSwap","type":"uint256"},{"internalType":"uint256","name":"BigValueThreshold","type":"uint256"}],"internalType":"struct RouterConfig.SwapConfig2","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"startIndex","type":"uint256"},{"internalType":"uint256","name":"endIndex","type":"uint256"}],"name":"getSwapConfigAtIndexRange","outputs":[{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwap","type":"uint256"},{"internalType":"uint256","name":"MinimumSwap","type":"uint256"},{"internalType":"uint256","name":"BigValueThreshold","type":"uint256"}],"internalType":"struct RouterConfig.SwapConfig2[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"getSwapConfigsCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"getTokenConfig","outputs":[{"components":[{"internalType":"uint8","name":"Decimals","type":"uint8"},{"internalType":"string","name":"ContractAddress","type":"string"},{"internalType":"uint256","name":"ContractVersion","type":"uint256"},{"internalType":"string","name":"RouterContract","type":"string"},{"internalType":"string","name":"Extra","type":"string"}],"internalType":"struct RouterConfig.TokenConfig","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"tokenAddress","type":"string"}],"name":"getTokenID","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getTokenIDByIndex","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"isChainIDExist","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"isTokenIDExist","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owners","outputs":[{"internalType":"address[2]","name":"","type":"address[2]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"removeAllChainConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"dstChainIDs","type":"uint256[]"}],"name":"removeAllDestFeeConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"dstChainIDs","type":"uint256[]"}],"name":"removeAllDestSwapConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"removeAllFeeConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"removeAllMultichainTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"srcChainIDs","type":"uint256[]"}],"name":"removeAllSrcFeeConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"srcChainIDs","type":"uint256[]"}],"name":"removeAllSrcSwapConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"removeAllSwapConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"removeAllTokenConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256[]","name":"chainIDs","type":"uint256[]"}],"name":"removeChainConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"removeChainID","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"srcChainIDs","type":"uint256[]"},{"internalType":"uint256[]","name":"dstChainIDs","type":"uint256[]"}],"name":"removeFeeConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"addr","type":"string"}],"name":"removeMPCPubkey","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"}],"name":"removeMultichainToken","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"srcChainID","type":"uint256"},{"internalType":"uint256","name":"dstChainID","type":"uint256"}],"name":"removeSwapAndFeeConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"srcChainIDs","type":"uint256[]"},{"internalType":"uint256[]","name":"dstChainIDs","type":"uint256[]"}],"name":"removeSwapConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256[]","name":"chainIDs","type":"uint256[]"}],"name":"removeTokenConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"}],"name":"removeTokenID","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"blockChain","type":"string"},{"internalType":"string","name":"routerContract","type":"string"},{"internalType":"uint64","name":"confirmations","type":"uint64"},{"internalType":"uint64","name":"initialHeight","type":"uint64"},{"internalType":"string","name":"extra","type":"string"}],"name":"setChainConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"extra","type":"string"}],"name":"setChainExtraConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"key","type":"string"},{"internalType":"string","name":"data","type":"string"}],"name":"setCustomConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"},{"internalType":"string","name":"data","type":"string"}],"name":"setExtraConfig","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"srcChainID","type":"uint256"},{"internalType":"uint256","name":"dstChainID","type":"uint256"},{"internalType":"uint256","name":"maxFee","type":"uint256"},{"internalType":"uint256","name":"minFee","type":"uint256"},{"internalType":"uint256","name":"feeRate","type":"uint256"}],"name":"setFeeConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwapFee","type":"uint256"},{"internalType":"uint256","name":"MinimumSwapFee","type":"uint256"},{"internalType":"uint256","name":"SwapFeeRatePerMillion","type":"uint256"}],"internalType":"struct RouterConfig.FeeConfig2[]","name":"configs","type":"tuple[]"}],"name":"setFeeConfigs","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"addr","type":"string"},{"internalType":"string","name":"pubkey","type":"string"}],"name":"setMPCPubkey","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"token","type":"string"}],"name":"setMultichainToken","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"srcChainID","type":"uint256"},{"internalType":"uint256","name":"dstChainID","type":"uint256"},{"internalType":"uint256","name":"maxSwap","type":"uint256"},{"internalType":"uint256","name":"minSwap","type":"uint256"},{"internalType":"uint256","name":"bigSwap","type":"uint256"},{"internalType":"uint256","name":"maxFee","type":"uint256"},{"internalType":"uint256","name":"minFee","type":"uint256"},{"internalType":"uint256","name":"feeRate","type":"uint256"}],"name":"setSwapAndFeeConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"srcChainID","type":"uint256"},{"internalType":"uint256","name":"dstChainID","type":"uint256"},{"internalType":"uint256","name":"maxSwap","type":"uint256"},{"internalType":"uint256","name":"minSwap","type":"uint256"},{"internalType":"uint256","name":"bigSwap","type":"uint256"}],"name":"setSwapConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"components":[{"internalType":"uint256","name":"FromChainID","type":"uint256"},{"internalType":"uint256","name":"ToChainID","type":"uint256"},{"internalType":"uint256","name":"MaximumSwap","type":"uint256"},{"internalType":"uint256","name":"MinimumSwap","type":"uint256"},{"internalType":"uint256","name":"BigValueThreshold","type":"uint256"}],"internalType":"struct RouterConfig.SwapConfig2[]","name":"configs","type":"tuple[]"}],"name":"setSwapConfigs","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"tokenAddr","type":"string"},{"internalType":"uint8","name":"decimals","type":"uint8"},{"internalType":"uint256","name":"version","type":"uint256"},{"internalType":"string","name":"routerContract","type":"string"},{"internalType":"string","name":"extra","type":"string"}],"name":"setTokenConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"extra","type":"string"}],"name":"setTokenExtraConfig","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenID","type":"string"},{"internalType":"uint256","name":"chainID","type":"uint256"},{"internalType":"string","name":"routerContract","type":"string"}],"name":"setTokenRouterContract","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"updateConfig","outputs":[],"stateMutability":"nonpayable","type":"function"}]`,
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// CONFIGVERSION is a free data retrieval call binding the contract method 0x34e30097.
//
// Solidity: function CONFIG_VERSION() view returns(uint256)
func (_Main *MainCaller) CONFIGVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "CONFIG_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGVERSION is a free data retrieval call binding the contract method 0x34e30097.
//
// Solidity: function CONFIG_VERSION() view returns(uint256)
func (_Main *MainSession) CONFIGVERSION() (*big.Int, error) {
	return _Main.Contract.CONFIGVERSION(&_Main.CallOpts)
}

// CONFIGVERSION is a free data retrieval call binding the contract method 0x34e30097.
//
// Solidity: function CONFIG_VERSION() view returns(uint256)
func (_Main *MainCallerSession) CONFIGVERSION() (*big.Int, error) {
	return _Main.Contract.CONFIGVERSION(&_Main.CallOpts)
}

// GetAllChainConfig is a free data retrieval call binding the contract method 0x71a4a947.
//
// Solidity: function getAllChainConfig() view returns((uint256,string,string,uint64,uint64,string)[])
func (_Main *MainCaller) GetAllChainConfig(opts *bind.CallOpts) ([]RouterConfigChainConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllChainConfig")

	if err != nil {
		return *new([]RouterConfigChainConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigChainConfig2)).(*[]RouterConfigChainConfig2)

	return out0, err

}

// GetAllChainConfig is a free data retrieval call binding the contract method 0x71a4a947.
//
// Solidity: function getAllChainConfig() view returns((uint256,string,string,uint64,uint64,string)[])
func (_Main *MainSession) GetAllChainConfig() ([]RouterConfigChainConfig2, error) {
	return _Main.Contract.GetAllChainConfig(&_Main.CallOpts)
}

// GetAllChainConfig is a free data retrieval call binding the contract method 0x71a4a947.
//
// Solidity: function getAllChainConfig() view returns((uint256,string,string,uint64,uint64,string)[])
func (_Main *MainCallerSession) GetAllChainConfig() ([]RouterConfigChainConfig2, error) {
	return _Main.Contract.GetAllChainConfig(&_Main.CallOpts)
}

// GetAllChainIDLength is a free data retrieval call binding the contract method 0x72c93288.
//
// Solidity: function getAllChainIDLength() view returns(uint256)
func (_Main *MainCaller) GetAllChainIDLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllChainIDLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllChainIDLength is a free data retrieval call binding the contract method 0x72c93288.
//
// Solidity: function getAllChainIDLength() view returns(uint256)
func (_Main *MainSession) GetAllChainIDLength() (*big.Int, error) {
	return _Main.Contract.GetAllChainIDLength(&_Main.CallOpts)
}

// GetAllChainIDLength is a free data retrieval call binding the contract method 0x72c93288.
//
// Solidity: function getAllChainIDLength() view returns(uint256)
func (_Main *MainCallerSession) GetAllChainIDLength() (*big.Int, error) {
	return _Main.Contract.GetAllChainIDLength(&_Main.CallOpts)
}

// GetAllChainIDs is a free data retrieval call binding the contract method 0xe27112d5.
//
// Solidity: function getAllChainIDs() view returns(uint256[])
func (_Main *MainCaller) GetAllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllChainIDs is a free data retrieval call binding the contract method 0xe27112d5.
//
// Solidity: function getAllChainIDs() view returns(uint256[])
func (_Main *MainSession) GetAllChainIDs() ([]*big.Int, error) {
	return _Main.Contract.GetAllChainIDs(&_Main.CallOpts)
}

// GetAllChainIDs is a free data retrieval call binding the contract method 0xe27112d5.
//
// Solidity: function getAllChainIDs() view returns(uint256[])
func (_Main *MainCallerSession) GetAllChainIDs() ([]*big.Int, error) {
	return _Main.Contract.GetAllChainIDs(&_Main.CallOpts)
}

// GetAllFeeConfigs is a free data retrieval call binding the contract method 0x6a3ea04f.
//
// Solidity: function getAllFeeConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCaller) GetAllFeeConfigs(opts *bind.CallOpts, tokenID string) ([]RouterConfigFeeConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllFeeConfigs", tokenID)

	if err != nil {
		return *new([]RouterConfigFeeConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigFeeConfig2)).(*[]RouterConfigFeeConfig2)

	return out0, err

}

// GetAllFeeConfigs is a free data retrieval call binding the contract method 0x6a3ea04f.
//
// Solidity: function getAllFeeConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainSession) GetAllFeeConfigs(tokenID string) ([]RouterConfigFeeConfig2, error) {
	return _Main.Contract.GetAllFeeConfigs(&_Main.CallOpts, tokenID)
}

// GetAllFeeConfigs is a free data retrieval call binding the contract method 0x6a3ea04f.
//
// Solidity: function getAllFeeConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCallerSession) GetAllFeeConfigs(tokenID string) ([]RouterConfigFeeConfig2, error) {
	return _Main.Contract.GetAllFeeConfigs(&_Main.CallOpts, tokenID)
}

// GetAllMultichainTokenConfig is a free data retrieval call binding the contract method 0x160dcc6f.
//
// Solidity: function getAllMultichainTokenConfig(string tokenID) view returns((uint256,uint8,string,uint256,string,string)[])
func (_Main *MainCaller) GetAllMultichainTokenConfig(opts *bind.CallOpts, tokenID string) ([]RouterConfigTokenConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllMultichainTokenConfig", tokenID)

	if err != nil {
		return *new([]RouterConfigTokenConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigTokenConfig2)).(*[]RouterConfigTokenConfig2)

	return out0, err

}

// GetAllMultichainTokenConfig is a free data retrieval call binding the contract method 0x160dcc6f.
//
// Solidity: function getAllMultichainTokenConfig(string tokenID) view returns((uint256,uint8,string,uint256,string,string)[])
func (_Main *MainSession) GetAllMultichainTokenConfig(tokenID string) ([]RouterConfigTokenConfig2, error) {
	return _Main.Contract.GetAllMultichainTokenConfig(&_Main.CallOpts, tokenID)
}

// GetAllMultichainTokenConfig is a free data retrieval call binding the contract method 0x160dcc6f.
//
// Solidity: function getAllMultichainTokenConfig(string tokenID) view returns((uint256,uint8,string,uint256,string,string)[])
func (_Main *MainCallerSession) GetAllMultichainTokenConfig(tokenID string) ([]RouterConfigTokenConfig2, error) {
	return _Main.Contract.GetAllMultichainTokenConfig(&_Main.CallOpts, tokenID)
}

// GetAllMultichainTokens is a free data retrieval call binding the contract method 0x8fcb62a3.
//
// Solidity: function getAllMultichainTokens(string tokenID) view returns((uint256,string)[])
func (_Main *MainCaller) GetAllMultichainTokens(opts *bind.CallOpts, tokenID string) ([]RouterConfigMultichainToken, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllMultichainTokens", tokenID)

	if err != nil {
		return *new([]RouterConfigMultichainToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigMultichainToken)).(*[]RouterConfigMultichainToken)

	return out0, err

}

// GetAllMultichainTokens is a free data retrieval call binding the contract method 0x8fcb62a3.
//
// Solidity: function getAllMultichainTokens(string tokenID) view returns((uint256,string)[])
func (_Main *MainSession) GetAllMultichainTokens(tokenID string) ([]RouterConfigMultichainToken, error) {
	return _Main.Contract.GetAllMultichainTokens(&_Main.CallOpts, tokenID)
}

// GetAllMultichainTokens is a free data retrieval call binding the contract method 0x8fcb62a3.
//
// Solidity: function getAllMultichainTokens(string tokenID) view returns((uint256,string)[])
func (_Main *MainCallerSession) GetAllMultichainTokens(tokenID string) ([]RouterConfigMultichainToken, error) {
	return _Main.Contract.GetAllMultichainTokens(&_Main.CallOpts, tokenID)
}

// GetAllSwapConfigs is a free data retrieval call binding the contract method 0x3c6b1a8f.
//
// Solidity: function getAllSwapConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCaller) GetAllSwapConfigs(opts *bind.CallOpts, tokenID string) ([]RouterConfigSwapConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllSwapConfigs", tokenID)

	if err != nil {
		return *new([]RouterConfigSwapConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigSwapConfig2)).(*[]RouterConfigSwapConfig2)

	return out0, err

}

// GetAllSwapConfigs is a free data retrieval call binding the contract method 0x3c6b1a8f.
//
// Solidity: function getAllSwapConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainSession) GetAllSwapConfigs(tokenID string) ([]RouterConfigSwapConfig2, error) {
	return _Main.Contract.GetAllSwapConfigs(&_Main.CallOpts, tokenID)
}

// GetAllSwapConfigs is a free data retrieval call binding the contract method 0x3c6b1a8f.
//
// Solidity: function getAllSwapConfigs(string tokenID) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCallerSession) GetAllSwapConfigs(tokenID string) ([]RouterConfigSwapConfig2, error) {
	return _Main.Contract.GetAllSwapConfigs(&_Main.CallOpts, tokenID)
}

// GetAllTokenIDLength is a free data retrieval call binding the contract method 0xa43f37d7.
//
// Solidity: function getAllTokenIDLength() view returns(uint256)
func (_Main *MainCaller) GetAllTokenIDLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllTokenIDLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllTokenIDLength is a free data retrieval call binding the contract method 0xa43f37d7.
//
// Solidity: function getAllTokenIDLength() view returns(uint256)
func (_Main *MainSession) GetAllTokenIDLength() (*big.Int, error) {
	return _Main.Contract.GetAllTokenIDLength(&_Main.CallOpts)
}

// GetAllTokenIDLength is a free data retrieval call binding the contract method 0xa43f37d7.
//
// Solidity: function getAllTokenIDLength() view returns(uint256)
func (_Main *MainCallerSession) GetAllTokenIDLength() (*big.Int, error) {
	return _Main.Contract.GetAllTokenIDLength(&_Main.CallOpts)
}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_Main *MainCaller) GetAllTokenIDs(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllTokenIDs")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_Main *MainSession) GetAllTokenIDs() ([]string, error) {
	return _Main.Contract.GetAllTokenIDs(&_Main.CallOpts)
}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_Main *MainCallerSession) GetAllTokenIDs() ([]string, error) {
	return _Main.Contract.GetAllTokenIDs(&_Main.CallOpts)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x19ed16dc.
//
// Solidity: function getChainConfig(uint256 chainID) view returns((string,string,uint64,uint64,string))
func (_Main *MainCaller) GetChainConfig(opts *bind.CallOpts, chainID *big.Int) (RouterConfigChainConfig, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getChainConfig", chainID)

	if err != nil {
		return *new(RouterConfigChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigChainConfig)).(*RouterConfigChainConfig)

	return out0, err

}

// GetChainConfig is a free data retrieval call binding the contract method 0x19ed16dc.
//
// Solidity: function getChainConfig(uint256 chainID) view returns((string,string,uint64,uint64,string))
func (_Main *MainSession) GetChainConfig(chainID *big.Int) (RouterConfigChainConfig, error) {
	return _Main.Contract.GetChainConfig(&_Main.CallOpts, chainID)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x19ed16dc.
//
// Solidity: function getChainConfig(uint256 chainID) view returns((string,string,uint64,uint64,string))
func (_Main *MainCallerSession) GetChainConfig(chainID *big.Int) (RouterConfigChainConfig, error) {
	return _Main.Contract.GetChainConfig(&_Main.CallOpts, chainID)
}

// GetChainIDByIndex is a free data retrieval call binding the contract method 0x21576975.
//
// Solidity: function getChainIDByIndex(uint256 index) view returns(uint256)
func (_Main *MainCaller) GetChainIDByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getChainIDByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainIDByIndex is a free data retrieval call binding the contract method 0x21576975.
//
// Solidity: function getChainIDByIndex(uint256 index) view returns(uint256)
func (_Main *MainSession) GetChainIDByIndex(index *big.Int) (*big.Int, error) {
	return _Main.Contract.GetChainIDByIndex(&_Main.CallOpts, index)
}

// GetChainIDByIndex is a free data retrieval call binding the contract method 0x21576975.
//
// Solidity: function getChainIDByIndex(uint256 index) view returns(uint256)
func (_Main *MainCallerSession) GetChainIDByIndex(index *big.Int) (*big.Int, error) {
	return _Main.Contract.GetChainIDByIndex(&_Main.CallOpts, index)
}

// GetCustomConfig is a free data retrieval call binding the contract method 0x61387d61.
//
// Solidity: function getCustomConfig(uint256 chainID, string key) view returns(string)
func (_Main *MainCaller) GetCustomConfig(opts *bind.CallOpts, chainID *big.Int, key string) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getCustomConfig", chainID, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCustomConfig is a free data retrieval call binding the contract method 0x61387d61.
//
// Solidity: function getCustomConfig(uint256 chainID, string key) view returns(string)
func (_Main *MainSession) GetCustomConfig(chainID *big.Int, key string) (string, error) {
	return _Main.Contract.GetCustomConfig(&_Main.CallOpts, chainID, key)
}

// GetCustomConfig is a free data retrieval call binding the contract method 0x61387d61.
//
// Solidity: function getCustomConfig(uint256 chainID, string key) view returns(string)
func (_Main *MainCallerSession) GetCustomConfig(chainID *big.Int, key string) (string, error) {
	return _Main.Contract.GetCustomConfig(&_Main.CallOpts, chainID, key)
}

// GetExtraConfig is a free data retrieval call binding the contract method 0x340a5f2d.
//
// Solidity: function getExtraConfig(string key) view returns(string)
func (_Main *MainCaller) GetExtraConfig(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getExtraConfig", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetExtraConfig is a free data retrieval call binding the contract method 0x340a5f2d.
//
// Solidity: function getExtraConfig(string key) view returns(string)
func (_Main *MainSession) GetExtraConfig(key string) (string, error) {
	return _Main.Contract.GetExtraConfig(&_Main.CallOpts, key)
}

// GetExtraConfig is a free data retrieval call binding the contract method 0x340a5f2d.
//
// Solidity: function getExtraConfig(string key) view returns(string)
func (_Main *MainCallerSession) GetExtraConfig(key string) (string, error) {
	return _Main.Contract.GetExtraConfig(&_Main.CallOpts, key)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x1aed1c97.
//
// Solidity: function getFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Main *MainCaller) GetFeeConfig(opts *bind.CallOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigFeeConfig, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getFeeConfig", tokenID, srcChainID, dstChainID)

	if err != nil {
		return *new(RouterConfigFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigFeeConfig)).(*RouterConfigFeeConfig)

	return out0, err

}

// GetFeeConfig is a free data retrieval call binding the contract method 0x1aed1c97.
//
// Solidity: function getFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Main *MainSession) GetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigFeeConfig, error) {
	return _Main.Contract.GetFeeConfig(&_Main.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetFeeConfig is a free data retrieval call binding the contract method 0x1aed1c97.
//
// Solidity: function getFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Main *MainCallerSession) GetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigFeeConfig, error) {
	return _Main.Contract.GetFeeConfig(&_Main.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetFeeConfigAtIndex is a free data retrieval call binding the contract method 0x8b665a51.
//
// Solidity: function getFeeConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Main *MainCaller) GetFeeConfigAtIndex(opts *bind.CallOpts, tokenID string, index *big.Int) (RouterConfigFeeConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getFeeConfigAtIndex", tokenID, index)

	if err != nil {
		return *new(RouterConfigFeeConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigFeeConfig2)).(*RouterConfigFeeConfig2)

	return out0, err

}

// GetFeeConfigAtIndex is a free data retrieval call binding the contract method 0x8b665a51.
//
// Solidity: function getFeeConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Main *MainSession) GetFeeConfigAtIndex(tokenID string, index *big.Int) (RouterConfigFeeConfig2, error) {
	return _Main.Contract.GetFeeConfigAtIndex(&_Main.CallOpts, tokenID, index)
}

// GetFeeConfigAtIndex is a free data retrieval call binding the contract method 0x8b665a51.
//
// Solidity: function getFeeConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Main *MainCallerSession) GetFeeConfigAtIndex(tokenID string, index *big.Int) (RouterConfigFeeConfig2, error) {
	return _Main.Contract.GetFeeConfigAtIndex(&_Main.CallOpts, tokenID, index)
}

// GetFeeConfigAtIndexRange is a free data retrieval call binding the contract method 0xbdd5ef4a.
//
// Solidity: function getFeeConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCaller) GetFeeConfigAtIndexRange(opts *bind.CallOpts, tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigFeeConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getFeeConfigAtIndexRange", tokenID, startIndex, endIndex)

	if err != nil {
		return *new([]RouterConfigFeeConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigFeeConfig2)).(*[]RouterConfigFeeConfig2)

	return out0, err

}

// GetFeeConfigAtIndexRange is a free data retrieval call binding the contract method 0xbdd5ef4a.
//
// Solidity: function getFeeConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainSession) GetFeeConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigFeeConfig2, error) {
	return _Main.Contract.GetFeeConfigAtIndexRange(&_Main.CallOpts, tokenID, startIndex, endIndex)
}

// GetFeeConfigAtIndexRange is a free data retrieval call binding the contract method 0xbdd5ef4a.
//
// Solidity: function getFeeConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCallerSession) GetFeeConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigFeeConfig2, error) {
	return _Main.Contract.GetFeeConfigAtIndexRange(&_Main.CallOpts, tokenID, startIndex, endIndex)
}

// GetFeeConfigsCount is a free data retrieval call binding the contract method 0x15393217.
//
// Solidity: function getFeeConfigsCount(string tokenID) view returns(uint256)
func (_Main *MainCaller) GetFeeConfigsCount(opts *bind.CallOpts, tokenID string) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getFeeConfigsCount", tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeConfigsCount is a free data retrieval call binding the contract method 0x15393217.
//
// Solidity: function getFeeConfigsCount(string tokenID) view returns(uint256)
func (_Main *MainSession) GetFeeConfigsCount(tokenID string) (*big.Int, error) {
	return _Main.Contract.GetFeeConfigsCount(&_Main.CallOpts, tokenID)
}

// GetFeeConfigsCount is a free data retrieval call binding the contract method 0x15393217.
//
// Solidity: function getFeeConfigsCount(string tokenID) view returns(uint256)
func (_Main *MainCallerSession) GetFeeConfigsCount(tokenID string) (*big.Int, error) {
	return _Main.Contract.GetFeeConfigsCount(&_Main.CallOpts, tokenID)
}

// GetMPCPubkey is a free data retrieval call binding the contract method 0x9f1cdedd.
//
// Solidity: function getMPCPubkey(string mpcAddress) view returns(string)
func (_Main *MainCaller) GetMPCPubkey(opts *bind.CallOpts, mpcAddress string) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMPCPubkey", mpcAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMPCPubkey is a free data retrieval call binding the contract method 0x9f1cdedd.
//
// Solidity: function getMPCPubkey(string mpcAddress) view returns(string)
func (_Main *MainSession) GetMPCPubkey(mpcAddress string) (string, error) {
	return _Main.Contract.GetMPCPubkey(&_Main.CallOpts, mpcAddress)
}

// GetMPCPubkey is a free data retrieval call binding the contract method 0x9f1cdedd.
//
// Solidity: function getMPCPubkey(string mpcAddress) view returns(string)
func (_Main *MainCallerSession) GetMPCPubkey(mpcAddress string) (string, error) {
	return _Main.Contract.GetMPCPubkey(&_Main.CallOpts, mpcAddress)
}

// GetMultichainToken is a free data retrieval call binding the contract method 0xb735ab5a.
//
// Solidity: function getMultichainToken(string tokenID, uint256 chainID) view returns(string)
func (_Main *MainCaller) GetMultichainToken(opts *bind.CallOpts, tokenID string, chainID *big.Int) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMultichainToken", tokenID, chainID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMultichainToken is a free data retrieval call binding the contract method 0xb735ab5a.
//
// Solidity: function getMultichainToken(string tokenID, uint256 chainID) view returns(string)
func (_Main *MainSession) GetMultichainToken(tokenID string, chainID *big.Int) (string, error) {
	return _Main.Contract.GetMultichainToken(&_Main.CallOpts, tokenID, chainID)
}

// GetMultichainToken is a free data retrieval call binding the contract method 0xb735ab5a.
//
// Solidity: function getMultichainToken(string tokenID, uint256 chainID) view returns(string)
func (_Main *MainCallerSession) GetMultichainToken(tokenID string, chainID *big.Int) (string, error) {
	return _Main.Contract.GetMultichainToken(&_Main.CallOpts, tokenID, chainID)
}

// GetOriginalTokenConfig is a free data retrieval call binding the contract method 0x339d7ed8.
//
// Solidity: function getOriginalTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Main *MainCaller) GetOriginalTokenConfig(opts *bind.CallOpts, tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getOriginalTokenConfig", tokenID, chainID)

	if err != nil {
		return *new(RouterConfigTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigTokenConfig)).(*RouterConfigTokenConfig)

	return out0, err

}

// GetOriginalTokenConfig is a free data retrieval call binding the contract method 0x339d7ed8.
//
// Solidity: function getOriginalTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Main *MainSession) GetOriginalTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Main.Contract.GetOriginalTokenConfig(&_Main.CallOpts, tokenID, chainID)
}

// GetOriginalTokenConfig is a free data retrieval call binding the contract method 0x339d7ed8.
//
// Solidity: function getOriginalTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Main *MainCallerSession) GetOriginalTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Main.Contract.GetOriginalTokenConfig(&_Main.CallOpts, tokenID, chainID)
}

// GetSwapConfig is a free data retrieval call binding the contract method 0x4da7163c.
//
// Solidity: function getSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Main *MainCaller) GetSwapConfig(opts *bind.CallOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigSwapConfig, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSwapConfig", tokenID, srcChainID, dstChainID)

	if err != nil {
		return *new(RouterConfigSwapConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigSwapConfig)).(*RouterConfigSwapConfig)

	return out0, err

}

// GetSwapConfig is a free data retrieval call binding the contract method 0x4da7163c.
//
// Solidity: function getSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Main *MainSession) GetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigSwapConfig, error) {
	return _Main.Contract.GetSwapConfig(&_Main.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetSwapConfig is a free data retrieval call binding the contract method 0x4da7163c.
//
// Solidity: function getSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) view returns((uint256,uint256,uint256))
func (_Main *MainCallerSession) GetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (RouterConfigSwapConfig, error) {
	return _Main.Contract.GetSwapConfig(&_Main.CallOpts, tokenID, srcChainID, dstChainID)
}

// GetSwapConfigAtIndex is a free data retrieval call binding the contract method 0xd825b3ef.
//
// Solidity: function getSwapConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Main *MainCaller) GetSwapConfigAtIndex(opts *bind.CallOpts, tokenID string, index *big.Int) (RouterConfigSwapConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSwapConfigAtIndex", tokenID, index)

	if err != nil {
		return *new(RouterConfigSwapConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigSwapConfig2)).(*RouterConfigSwapConfig2)

	return out0, err

}

// GetSwapConfigAtIndex is a free data retrieval call binding the contract method 0xd825b3ef.
//
// Solidity: function getSwapConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Main *MainSession) GetSwapConfigAtIndex(tokenID string, index *big.Int) (RouterConfigSwapConfig2, error) {
	return _Main.Contract.GetSwapConfigAtIndex(&_Main.CallOpts, tokenID, index)
}

// GetSwapConfigAtIndex is a free data retrieval call binding the contract method 0xd825b3ef.
//
// Solidity: function getSwapConfigAtIndex(string tokenID, uint256 index) view returns((uint256,uint256,uint256,uint256,uint256))
func (_Main *MainCallerSession) GetSwapConfigAtIndex(tokenID string, index *big.Int) (RouterConfigSwapConfig2, error) {
	return _Main.Contract.GetSwapConfigAtIndex(&_Main.CallOpts, tokenID, index)
}

// GetSwapConfigAtIndexRange is a free data retrieval call binding the contract method 0xc03411ba.
//
// Solidity: function getSwapConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCaller) GetSwapConfigAtIndexRange(opts *bind.CallOpts, tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigSwapConfig2, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSwapConfigAtIndexRange", tokenID, startIndex, endIndex)

	if err != nil {
		return *new([]RouterConfigSwapConfig2), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterConfigSwapConfig2)).(*[]RouterConfigSwapConfig2)

	return out0, err

}

// GetSwapConfigAtIndexRange is a free data retrieval call binding the contract method 0xc03411ba.
//
// Solidity: function getSwapConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainSession) GetSwapConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigSwapConfig2, error) {
	return _Main.Contract.GetSwapConfigAtIndexRange(&_Main.CallOpts, tokenID, startIndex, endIndex)
}

// GetSwapConfigAtIndexRange is a free data retrieval call binding the contract method 0xc03411ba.
//
// Solidity: function getSwapConfigAtIndexRange(string tokenID, uint256 startIndex, uint256 endIndex) view returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Main *MainCallerSession) GetSwapConfigAtIndexRange(tokenID string, startIndex *big.Int, endIndex *big.Int) ([]RouterConfigSwapConfig2, error) {
	return _Main.Contract.GetSwapConfigAtIndexRange(&_Main.CallOpts, tokenID, startIndex, endIndex)
}

// GetSwapConfigsCount is a free data retrieval call binding the contract method 0x06e9a46a.
//
// Solidity: function getSwapConfigsCount(string tokenID) view returns(uint256)
func (_Main *MainCaller) GetSwapConfigsCount(opts *bind.CallOpts, tokenID string) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getSwapConfigsCount", tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapConfigsCount is a free data retrieval call binding the contract method 0x06e9a46a.
//
// Solidity: function getSwapConfigsCount(string tokenID) view returns(uint256)
func (_Main *MainSession) GetSwapConfigsCount(tokenID string) (*big.Int, error) {
	return _Main.Contract.GetSwapConfigsCount(&_Main.CallOpts, tokenID)
}

// GetSwapConfigsCount is a free data retrieval call binding the contract method 0x06e9a46a.
//
// Solidity: function getSwapConfigsCount(string tokenID) view returns(uint256)
func (_Main *MainCallerSession) GetSwapConfigsCount(tokenID string) (*big.Int, error) {
	return _Main.Contract.GetSwapConfigsCount(&_Main.CallOpts, tokenID)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x459511d1.
//
// Solidity: function getTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Main *MainCaller) GetTokenConfig(opts *bind.CallOpts, tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getTokenConfig", tokenID, chainID)

	if err != nil {
		return *new(RouterConfigTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterConfigTokenConfig)).(*RouterConfigTokenConfig)

	return out0, err

}

// GetTokenConfig is a free data retrieval call binding the contract method 0x459511d1.
//
// Solidity: function getTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Main *MainSession) GetTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Main.Contract.GetTokenConfig(&_Main.CallOpts, tokenID, chainID)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x459511d1.
//
// Solidity: function getTokenConfig(string tokenID, uint256 chainID) view returns((uint8,string,uint256,string,string))
func (_Main *MainCallerSession) GetTokenConfig(tokenID string, chainID *big.Int) (RouterConfigTokenConfig, error) {
	return _Main.Contract.GetTokenConfig(&_Main.CallOpts, tokenID, chainID)
}

// GetTokenID is a free data retrieval call binding the contract method 0xabad7b5c.
//
// Solidity: function getTokenID(uint256 chainID, string tokenAddress) view returns(string)
func (_Main *MainCaller) GetTokenID(opts *bind.CallOpts, chainID *big.Int, tokenAddress string) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getTokenID", chainID, tokenAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0xabad7b5c.
//
// Solidity: function getTokenID(uint256 chainID, string tokenAddress) view returns(string)
func (_Main *MainSession) GetTokenID(chainID *big.Int, tokenAddress string) (string, error) {
	return _Main.Contract.GetTokenID(&_Main.CallOpts, chainID, tokenAddress)
}

// GetTokenID is a free data retrieval call binding the contract method 0xabad7b5c.
//
// Solidity: function getTokenID(uint256 chainID, string tokenAddress) view returns(string)
func (_Main *MainCallerSession) GetTokenID(chainID *big.Int, tokenAddress string) (string, error) {
	return _Main.Contract.GetTokenID(&_Main.CallOpts, chainID, tokenAddress)
}

// GetTokenIDByIndex is a free data retrieval call binding the contract method 0xe828d00f.
//
// Solidity: function getTokenIDByIndex(uint256 index) view returns(string)
func (_Main *MainCaller) GetTokenIDByIndex(opts *bind.CallOpts, index *big.Int) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getTokenIDByIndex", index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenIDByIndex is a free data retrieval call binding the contract method 0xe828d00f.
//
// Solidity: function getTokenIDByIndex(uint256 index) view returns(string)
func (_Main *MainSession) GetTokenIDByIndex(index *big.Int) (string, error) {
	return _Main.Contract.GetTokenIDByIndex(&_Main.CallOpts, index)
}

// GetTokenIDByIndex is a free data retrieval call binding the contract method 0xe828d00f.
//
// Solidity: function getTokenIDByIndex(uint256 index) view returns(string)
func (_Main *MainCallerSession) GetTokenIDByIndex(index *big.Int) (string, error) {
	return _Main.Contract.GetTokenIDByIndex(&_Main.CallOpts, index)
}

// IsChainIDExist is a free data retrieval call binding the contract method 0xfd15ea70.
//
// Solidity: function isChainIDExist(uint256 chainID) view returns(bool)
func (_Main *MainCaller) IsChainIDExist(opts *bind.CallOpts, chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isChainIDExist", chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChainIDExist is a free data retrieval call binding the contract method 0xfd15ea70.
//
// Solidity: function isChainIDExist(uint256 chainID) view returns(bool)
func (_Main *MainSession) IsChainIDExist(chainID *big.Int) (bool, error) {
	return _Main.Contract.IsChainIDExist(&_Main.CallOpts, chainID)
}

// IsChainIDExist is a free data retrieval call binding the contract method 0xfd15ea70.
//
// Solidity: function isChainIDExist(uint256 chainID) view returns(bool)
func (_Main *MainCallerSession) IsChainIDExist(chainID *big.Int) (bool, error) {
	return _Main.Contract.IsChainIDExist(&_Main.CallOpts, chainID)
}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_Main *MainCaller) IsTokenIDExist(opts *bind.CallOpts, tokenID string) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isTokenIDExist", tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_Main *MainSession) IsTokenIDExist(tokenID string) (bool, error) {
	return _Main.Contract.IsTokenIDExist(&_Main.CallOpts, tokenID)
}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_Main *MainCallerSession) IsTokenIDExist(tokenID string) (bool, error) {
	return _Main.Contract.IsTokenIDExist(&_Main.CallOpts, tokenID)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[2])
func (_Main *MainCaller) Owners(opts *bind.CallOpts) ([2]common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[2])
func (_Main *MainSession) Owners() ([2]common.Address, error) {
	return _Main.Contract.Owners(&_Main.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[2])
func (_Main *MainCallerSession) Owners() ([2]common.Address, error) {
	return _Main.Contract.Owners(&_Main.CallOpts)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns(bool)
func (_Main *MainTransactor) AddChainID(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addChainID", chainID)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns(bool)
func (_Main *MainSession) AddChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddChainID(&_Main.TransactOpts, chainID)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns(bool)
func (_Main *MainTransactorSession) AddChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddChainID(&_Main.TransactOpts, chainID)
}

// AddTokenID is a paid mutator transaction binding the contract method 0xe3e01664.
//
// Solidity: function addTokenID(string tokenID) returns(bool)
func (_Main *MainTransactor) AddTokenID(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addTokenID", tokenID)
}

// AddTokenID is a paid mutator transaction binding the contract method 0xe3e01664.
//
// Solidity: function addTokenID(string tokenID) returns(bool)
func (_Main *MainSession) AddTokenID(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.AddTokenID(&_Main.TransactOpts, tokenID)
}

// AddTokenID is a paid mutator transaction binding the contract method 0xe3e01664.
//
// Solidity: function addTokenID(string tokenID) returns(bool)
func (_Main *MainTransactorSession) AddTokenID(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.AddTokenID(&_Main.TransactOpts, tokenID)
}

// RemoveAllChainConfig is a paid mutator transaction binding the contract method 0x909b4225.
//
// Solidity: function removeAllChainConfig() returns()
func (_Main *MainTransactor) RemoveAllChainConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllChainConfig")
}

// RemoveAllChainConfig is a paid mutator transaction binding the contract method 0x909b4225.
//
// Solidity: function removeAllChainConfig() returns()
func (_Main *MainSession) RemoveAllChainConfig() (*types.Transaction, error) {
	return _Main.Contract.RemoveAllChainConfig(&_Main.TransactOpts)
}

// RemoveAllChainConfig is a paid mutator transaction binding the contract method 0x909b4225.
//
// Solidity: function removeAllChainConfig() returns()
func (_Main *MainTransactorSession) RemoveAllChainConfig() (*types.Transaction, error) {
	return _Main.Contract.RemoveAllChainConfig(&_Main.TransactOpts)
}

// RemoveAllDestFeeConfig is a paid mutator transaction binding the contract method 0xb5f59a96.
//
// Solidity: function removeAllDestFeeConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Main *MainTransactor) RemoveAllDestFeeConfig(opts *bind.TransactOpts, tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllDestFeeConfig", tokenID, dstChainIDs)
}

// RemoveAllDestFeeConfig is a paid mutator transaction binding the contract method 0xb5f59a96.
//
// Solidity: function removeAllDestFeeConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Main *MainSession) RemoveAllDestFeeConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllDestFeeConfig(&_Main.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllDestFeeConfig is a paid mutator transaction binding the contract method 0xb5f59a96.
//
// Solidity: function removeAllDestFeeConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Main *MainTransactorSession) RemoveAllDestFeeConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllDestFeeConfig(&_Main.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllDestSwapConfig is a paid mutator transaction binding the contract method 0x267d8ea6.
//
// Solidity: function removeAllDestSwapConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Main *MainTransactor) RemoveAllDestSwapConfig(opts *bind.TransactOpts, tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllDestSwapConfig", tokenID, dstChainIDs)
}

// RemoveAllDestSwapConfig is a paid mutator transaction binding the contract method 0x267d8ea6.
//
// Solidity: function removeAllDestSwapConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Main *MainSession) RemoveAllDestSwapConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllDestSwapConfig(&_Main.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllDestSwapConfig is a paid mutator transaction binding the contract method 0x267d8ea6.
//
// Solidity: function removeAllDestSwapConfig(string tokenID, uint256[] dstChainIDs) returns()
func (_Main *MainTransactorSession) RemoveAllDestSwapConfig(tokenID string, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllDestSwapConfig(&_Main.TransactOpts, tokenID, dstChainIDs)
}

// RemoveAllFeeConfig is a paid mutator transaction binding the contract method 0x0a2b6b13.
//
// Solidity: function removeAllFeeConfig(string tokenID) returns()
func (_Main *MainTransactor) RemoveAllFeeConfig(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllFeeConfig", tokenID)
}

// RemoveAllFeeConfig is a paid mutator transaction binding the contract method 0x0a2b6b13.
//
// Solidity: function removeAllFeeConfig(string tokenID) returns()
func (_Main *MainSession) RemoveAllFeeConfig(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllFeeConfig(&_Main.TransactOpts, tokenID)
}

// RemoveAllFeeConfig is a paid mutator transaction binding the contract method 0x0a2b6b13.
//
// Solidity: function removeAllFeeConfig(string tokenID) returns()
func (_Main *MainTransactorSession) RemoveAllFeeConfig(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllFeeConfig(&_Main.TransactOpts, tokenID)
}

// RemoveAllMultichainTokens is a paid mutator transaction binding the contract method 0x83f5d5d4.
//
// Solidity: function removeAllMultichainTokens(string tokenID) returns()
func (_Main *MainTransactor) RemoveAllMultichainTokens(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllMultichainTokens", tokenID)
}

// RemoveAllMultichainTokens is a paid mutator transaction binding the contract method 0x83f5d5d4.
//
// Solidity: function removeAllMultichainTokens(string tokenID) returns()
func (_Main *MainSession) RemoveAllMultichainTokens(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllMultichainTokens(&_Main.TransactOpts, tokenID)
}

// RemoveAllMultichainTokens is a paid mutator transaction binding the contract method 0x83f5d5d4.
//
// Solidity: function removeAllMultichainTokens(string tokenID) returns()
func (_Main *MainTransactorSession) RemoveAllMultichainTokens(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllMultichainTokens(&_Main.TransactOpts, tokenID)
}

// RemoveAllSrcFeeConfig is a paid mutator transaction binding the contract method 0xdcd59579.
//
// Solidity: function removeAllSrcFeeConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Main *MainTransactor) RemoveAllSrcFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllSrcFeeConfig", tokenID, srcChainIDs)
}

// RemoveAllSrcFeeConfig is a paid mutator transaction binding the contract method 0xdcd59579.
//
// Solidity: function removeAllSrcFeeConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Main *MainSession) RemoveAllSrcFeeConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllSrcFeeConfig(&_Main.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSrcFeeConfig is a paid mutator transaction binding the contract method 0xdcd59579.
//
// Solidity: function removeAllSrcFeeConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Main *MainTransactorSession) RemoveAllSrcFeeConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllSrcFeeConfig(&_Main.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSrcSwapConfig is a paid mutator transaction binding the contract method 0x103a24a3.
//
// Solidity: function removeAllSrcSwapConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Main *MainTransactor) RemoveAllSrcSwapConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllSrcSwapConfig", tokenID, srcChainIDs)
}

// RemoveAllSrcSwapConfig is a paid mutator transaction binding the contract method 0x103a24a3.
//
// Solidity: function removeAllSrcSwapConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Main *MainSession) RemoveAllSrcSwapConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllSrcSwapConfig(&_Main.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSrcSwapConfig is a paid mutator transaction binding the contract method 0x103a24a3.
//
// Solidity: function removeAllSrcSwapConfig(string tokenID, uint256[] srcChainIDs) returns()
func (_Main *MainTransactorSession) RemoveAllSrcSwapConfig(tokenID string, srcChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllSrcSwapConfig(&_Main.TransactOpts, tokenID, srcChainIDs)
}

// RemoveAllSwapConfig is a paid mutator transaction binding the contract method 0xa7bdc603.
//
// Solidity: function removeAllSwapConfig(string tokenID) returns()
func (_Main *MainTransactor) RemoveAllSwapConfig(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllSwapConfig", tokenID)
}

// RemoveAllSwapConfig is a paid mutator transaction binding the contract method 0xa7bdc603.
//
// Solidity: function removeAllSwapConfig(string tokenID) returns()
func (_Main *MainSession) RemoveAllSwapConfig(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllSwapConfig(&_Main.TransactOpts, tokenID)
}

// RemoveAllSwapConfig is a paid mutator transaction binding the contract method 0xa7bdc603.
//
// Solidity: function removeAllSwapConfig(string tokenID) returns()
func (_Main *MainTransactorSession) RemoveAllSwapConfig(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllSwapConfig(&_Main.TransactOpts, tokenID)
}

// RemoveAllTokenConfig is a paid mutator transaction binding the contract method 0x0c4fe066.
//
// Solidity: function removeAllTokenConfig(string tokenID) returns()
func (_Main *MainTransactor) RemoveAllTokenConfig(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAllTokenConfig", tokenID)
}

// RemoveAllTokenConfig is a paid mutator transaction binding the contract method 0x0c4fe066.
//
// Solidity: function removeAllTokenConfig(string tokenID) returns()
func (_Main *MainSession) RemoveAllTokenConfig(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllTokenConfig(&_Main.TransactOpts, tokenID)
}

// RemoveAllTokenConfig is a paid mutator transaction binding the contract method 0x0c4fe066.
//
// Solidity: function removeAllTokenConfig(string tokenID) returns()
func (_Main *MainTransactorSession) RemoveAllTokenConfig(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveAllTokenConfig(&_Main.TransactOpts, tokenID)
}

// RemoveChainConfig is a paid mutator transaction binding the contract method 0xb00f5cc8.
//
// Solidity: function removeChainConfig(uint256[] chainIDs) returns()
func (_Main *MainTransactor) RemoveChainConfig(opts *bind.TransactOpts, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeChainConfig", chainIDs)
}

// RemoveChainConfig is a paid mutator transaction binding the contract method 0xb00f5cc8.
//
// Solidity: function removeChainConfig(uint256[] chainIDs) returns()
func (_Main *MainSession) RemoveChainConfig(chainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveChainConfig(&_Main.TransactOpts, chainIDs)
}

// RemoveChainConfig is a paid mutator transaction binding the contract method 0xb00f5cc8.
//
// Solidity: function removeChainConfig(uint256[] chainIDs) returns()
func (_Main *MainTransactorSession) RemoveChainConfig(chainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveChainConfig(&_Main.TransactOpts, chainIDs)
}

// RemoveChainID is a paid mutator transaction binding the contract method 0xe90768c6.
//
// Solidity: function removeChainID(uint256 chainID) returns()
func (_Main *MainTransactor) RemoveChainID(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeChainID", chainID)
}

// RemoveChainID is a paid mutator transaction binding the contract method 0xe90768c6.
//
// Solidity: function removeChainID(uint256 chainID) returns()
func (_Main *MainSession) RemoveChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveChainID(&_Main.TransactOpts, chainID)
}

// RemoveChainID is a paid mutator transaction binding the contract method 0xe90768c6.
//
// Solidity: function removeChainID(uint256 chainID) returns()
func (_Main *MainTransactorSession) RemoveChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveChainID(&_Main.TransactOpts, chainID)
}

// RemoveFeeConfig is a paid mutator transaction binding the contract method 0xb4a51ac4.
//
// Solidity: function removeFeeConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Main *MainTransactor) RemoveFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeFeeConfig", tokenID, srcChainIDs, dstChainIDs)
}

// RemoveFeeConfig is a paid mutator transaction binding the contract method 0xb4a51ac4.
//
// Solidity: function removeFeeConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Main *MainSession) RemoveFeeConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFeeConfig(&_Main.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveFeeConfig is a paid mutator transaction binding the contract method 0xb4a51ac4.
//
// Solidity: function removeFeeConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Main *MainTransactorSession) RemoveFeeConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFeeConfig(&_Main.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveMPCPubkey is a paid mutator transaction binding the contract method 0x04b75939.
//
// Solidity: function removeMPCPubkey(string addr) returns()
func (_Main *MainTransactor) RemoveMPCPubkey(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeMPCPubkey", addr)
}

// RemoveMPCPubkey is a paid mutator transaction binding the contract method 0x04b75939.
//
// Solidity: function removeMPCPubkey(string addr) returns()
func (_Main *MainSession) RemoveMPCPubkey(addr string) (*types.Transaction, error) {
	return _Main.Contract.RemoveMPCPubkey(&_Main.TransactOpts, addr)
}

// RemoveMPCPubkey is a paid mutator transaction binding the contract method 0x04b75939.
//
// Solidity: function removeMPCPubkey(string addr) returns()
func (_Main *MainTransactorSession) RemoveMPCPubkey(addr string) (*types.Transaction, error) {
	return _Main.Contract.RemoveMPCPubkey(&_Main.TransactOpts, addr)
}

// RemoveMultichainToken is a paid mutator transaction binding the contract method 0x9026b205.
//
// Solidity: function removeMultichainToken(string tokenID, uint256 chainID) returns()
func (_Main *MainTransactor) RemoveMultichainToken(opts *bind.TransactOpts, tokenID string, chainID *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeMultichainToken", tokenID, chainID)
}

// RemoveMultichainToken is a paid mutator transaction binding the contract method 0x9026b205.
//
// Solidity: function removeMultichainToken(string tokenID, uint256 chainID) returns()
func (_Main *MainSession) RemoveMultichainToken(tokenID string, chainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveMultichainToken(&_Main.TransactOpts, tokenID, chainID)
}

// RemoveMultichainToken is a paid mutator transaction binding the contract method 0x9026b205.
//
// Solidity: function removeMultichainToken(string tokenID, uint256 chainID) returns()
func (_Main *MainTransactorSession) RemoveMultichainToken(tokenID string, chainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveMultichainToken(&_Main.TransactOpts, tokenID, chainID)
}

// RemoveSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x8f2cfccb.
//
// Solidity: function removeSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) returns()
func (_Main *MainTransactor) RemoveSwapAndFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeSwapAndFeeConfig", tokenID, srcChainID, dstChainID)
}

// RemoveSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x8f2cfccb.
//
// Solidity: function removeSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) returns()
func (_Main *MainSession) RemoveSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveSwapAndFeeConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID)
}

// RemoveSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x8f2cfccb.
//
// Solidity: function removeSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID) returns()
func (_Main *MainTransactorSession) RemoveSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveSwapAndFeeConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID)
}

// RemoveSwapConfig is a paid mutator transaction binding the contract method 0x9fdc851a.
//
// Solidity: function removeSwapConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Main *MainTransactor) RemoveSwapConfig(opts *bind.TransactOpts, tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeSwapConfig", tokenID, srcChainIDs, dstChainIDs)
}

// RemoveSwapConfig is a paid mutator transaction binding the contract method 0x9fdc851a.
//
// Solidity: function removeSwapConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Main *MainSession) RemoveSwapConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveSwapConfig(&_Main.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveSwapConfig is a paid mutator transaction binding the contract method 0x9fdc851a.
//
// Solidity: function removeSwapConfig(string tokenID, uint256[] srcChainIDs, uint256[] dstChainIDs) returns()
func (_Main *MainTransactorSession) RemoveSwapConfig(tokenID string, srcChainIDs []*big.Int, dstChainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveSwapConfig(&_Main.TransactOpts, tokenID, srcChainIDs, dstChainIDs)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xb9b85d34.
//
// Solidity: function removeTokenConfig(string tokenID, uint256[] chainIDs) returns()
func (_Main *MainTransactor) RemoveTokenConfig(opts *bind.TransactOpts, tokenID string, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeTokenConfig", tokenID, chainIDs)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xb9b85d34.
//
// Solidity: function removeTokenConfig(string tokenID, uint256[] chainIDs) returns()
func (_Main *MainSession) RemoveTokenConfig(tokenID string, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveTokenConfig(&_Main.TransactOpts, tokenID, chainIDs)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xb9b85d34.
//
// Solidity: function removeTokenConfig(string tokenID, uint256[] chainIDs) returns()
func (_Main *MainTransactorSession) RemoveTokenConfig(tokenID string, chainIDs []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveTokenConfig(&_Main.TransactOpts, tokenID, chainIDs)
}

// RemoveTokenID is a paid mutator transaction binding the contract method 0x5d58d0a1.
//
// Solidity: function removeTokenID(string tokenID) returns()
func (_Main *MainTransactor) RemoveTokenID(opts *bind.TransactOpts, tokenID string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeTokenID", tokenID)
}

// RemoveTokenID is a paid mutator transaction binding the contract method 0x5d58d0a1.
//
// Solidity: function removeTokenID(string tokenID) returns()
func (_Main *MainSession) RemoveTokenID(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveTokenID(&_Main.TransactOpts, tokenID)
}

// RemoveTokenID is a paid mutator transaction binding the contract method 0x5d58d0a1.
//
// Solidity: function removeTokenID(string tokenID) returns()
func (_Main *MainTransactorSession) RemoveTokenID(tokenID string) (*types.Transaction, error) {
	return _Main.Contract.RemoveTokenID(&_Main.TransactOpts, tokenID)
}

// SetChainConfig is a paid mutator transaction binding the contract method 0x2383006c.
//
// Solidity: function setChainConfig(uint256 chainID, string blockChain, string routerContract, uint64 confirmations, uint64 initialHeight, string extra) returns(bool)
func (_Main *MainTransactor) SetChainConfig(opts *bind.TransactOpts, chainID *big.Int, blockChain string, routerContract string, confirmations uint64, initialHeight uint64, extra string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setChainConfig", chainID, blockChain, routerContract, confirmations, initialHeight, extra)
}

// SetChainConfig is a paid mutator transaction binding the contract method 0x2383006c.
//
// Solidity: function setChainConfig(uint256 chainID, string blockChain, string routerContract, uint64 confirmations, uint64 initialHeight, string extra) returns(bool)
func (_Main *MainSession) SetChainConfig(chainID *big.Int, blockChain string, routerContract string, confirmations uint64, initialHeight uint64, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetChainConfig(&_Main.TransactOpts, chainID, blockChain, routerContract, confirmations, initialHeight, extra)
}

// SetChainConfig is a paid mutator transaction binding the contract method 0x2383006c.
//
// Solidity: function setChainConfig(uint256 chainID, string blockChain, string routerContract, uint64 confirmations, uint64 initialHeight, string extra) returns(bool)
func (_Main *MainTransactorSession) SetChainConfig(chainID *big.Int, blockChain string, routerContract string, confirmations uint64, initialHeight uint64, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetChainConfig(&_Main.TransactOpts, chainID, blockChain, routerContract, confirmations, initialHeight, extra)
}

// SetChainExtraConfig is a paid mutator transaction binding the contract method 0xd236a09b.
//
// Solidity: function setChainExtraConfig(uint256 chainID, string extra) returns(bool)
func (_Main *MainTransactor) SetChainExtraConfig(opts *bind.TransactOpts, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setChainExtraConfig", chainID, extra)
}

// SetChainExtraConfig is a paid mutator transaction binding the contract method 0xd236a09b.
//
// Solidity: function setChainExtraConfig(uint256 chainID, string extra) returns(bool)
func (_Main *MainSession) SetChainExtraConfig(chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetChainExtraConfig(&_Main.TransactOpts, chainID, extra)
}

// SetChainExtraConfig is a paid mutator transaction binding the contract method 0xd236a09b.
//
// Solidity: function setChainExtraConfig(uint256 chainID, string extra) returns(bool)
func (_Main *MainTransactorSession) SetChainExtraConfig(chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetChainExtraConfig(&_Main.TransactOpts, chainID, extra)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0xe671242a.
//
// Solidity: function setCustomConfig(uint256 chainID, string key, string data) returns()
func (_Main *MainTransactor) SetCustomConfig(opts *bind.TransactOpts, chainID *big.Int, key string, data string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setCustomConfig", chainID, key, data)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0xe671242a.
//
// Solidity: function setCustomConfig(uint256 chainID, string key, string data) returns()
func (_Main *MainSession) SetCustomConfig(chainID *big.Int, key string, data string) (*types.Transaction, error) {
	return _Main.Contract.SetCustomConfig(&_Main.TransactOpts, chainID, key, data)
}

// SetCustomConfig is a paid mutator transaction binding the contract method 0xe671242a.
//
// Solidity: function setCustomConfig(uint256 chainID, string key, string data) returns()
func (_Main *MainTransactorSession) SetCustomConfig(chainID *big.Int, key string, data string) (*types.Transaction, error) {
	return _Main.Contract.SetCustomConfig(&_Main.TransactOpts, chainID, key, data)
}

// SetExtraConfig is a paid mutator transaction binding the contract method 0x04cbc75b.
//
// Solidity: function setExtraConfig(string key, string data) returns()
func (_Main *MainTransactor) SetExtraConfig(opts *bind.TransactOpts, key string, data string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setExtraConfig", key, data)
}

// SetExtraConfig is a paid mutator transaction binding the contract method 0x04cbc75b.
//
// Solidity: function setExtraConfig(string key, string data) returns()
func (_Main *MainSession) SetExtraConfig(key string, data string) (*types.Transaction, error) {
	return _Main.Contract.SetExtraConfig(&_Main.TransactOpts, key, data)
}

// SetExtraConfig is a paid mutator transaction binding the contract method 0x04cbc75b.
//
// Solidity: function setExtraConfig(string key, string data) returns()
func (_Main *MainTransactorSession) SetExtraConfig(key string, data string) (*types.Transaction, error) {
	return _Main.Contract.SetExtraConfig(&_Main.TransactOpts, key, data)
}

// SetFeeConfig is a paid mutator transaction binding the contract method 0x2f50cf14.
//
// Solidity: function setFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Main *MainTransactor) SetFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFeeConfig", tokenID, srcChainID, dstChainID, maxFee, minFee, feeRate)
}

// SetFeeConfig is a paid mutator transaction binding the contract method 0x2f50cf14.
//
// Solidity: function setFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Main *MainSession) SetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetFeeConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID, maxFee, minFee, feeRate)
}

// SetFeeConfig is a paid mutator transaction binding the contract method 0x2f50cf14.
//
// Solidity: function setFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Main *MainTransactorSession) SetFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetFeeConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID, maxFee, minFee, feeRate)
}

// SetFeeConfigs is a paid mutator transaction binding the contract method 0x9a29eaf2.
//
// Solidity: function setFeeConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Main *MainTransactor) SetFeeConfigs(opts *bind.TransactOpts, tokenID string, configs []RouterConfigFeeConfig2) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFeeConfigs", tokenID, configs)
}

// SetFeeConfigs is a paid mutator transaction binding the contract method 0x9a29eaf2.
//
// Solidity: function setFeeConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Main *MainSession) SetFeeConfigs(tokenID string, configs []RouterConfigFeeConfig2) (*types.Transaction, error) {
	return _Main.Contract.SetFeeConfigs(&_Main.TransactOpts, tokenID, configs)
}

// SetFeeConfigs is a paid mutator transaction binding the contract method 0x9a29eaf2.
//
// Solidity: function setFeeConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Main *MainTransactorSession) SetFeeConfigs(tokenID string, configs []RouterConfigFeeConfig2) (*types.Transaction, error) {
	return _Main.Contract.SetFeeConfigs(&_Main.TransactOpts, tokenID, configs)
}

// SetMPCPubkey is a paid mutator transaction binding the contract method 0xb4d8d47c.
//
// Solidity: function setMPCPubkey(string addr, string pubkey) returns()
func (_Main *MainTransactor) SetMPCPubkey(opts *bind.TransactOpts, addr string, pubkey string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setMPCPubkey", addr, pubkey)
}

// SetMPCPubkey is a paid mutator transaction binding the contract method 0xb4d8d47c.
//
// Solidity: function setMPCPubkey(string addr, string pubkey) returns()
func (_Main *MainSession) SetMPCPubkey(addr string, pubkey string) (*types.Transaction, error) {
	return _Main.Contract.SetMPCPubkey(&_Main.TransactOpts, addr, pubkey)
}

// SetMPCPubkey is a paid mutator transaction binding the contract method 0xb4d8d47c.
//
// Solidity: function setMPCPubkey(string addr, string pubkey) returns()
func (_Main *MainTransactorSession) SetMPCPubkey(addr string, pubkey string) (*types.Transaction, error) {
	return _Main.Contract.SetMPCPubkey(&_Main.TransactOpts, addr, pubkey)
}

// SetMultichainToken is a paid mutator transaction binding the contract method 0xf4730a7e.
//
// Solidity: function setMultichainToken(string tokenID, uint256 chainID, string token) returns()
func (_Main *MainTransactor) SetMultichainToken(opts *bind.TransactOpts, tokenID string, chainID *big.Int, token string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setMultichainToken", tokenID, chainID, token)
}

// SetMultichainToken is a paid mutator transaction binding the contract method 0xf4730a7e.
//
// Solidity: function setMultichainToken(string tokenID, uint256 chainID, string token) returns()
func (_Main *MainSession) SetMultichainToken(tokenID string, chainID *big.Int, token string) (*types.Transaction, error) {
	return _Main.Contract.SetMultichainToken(&_Main.TransactOpts, tokenID, chainID, token)
}

// SetMultichainToken is a paid mutator transaction binding the contract method 0xf4730a7e.
//
// Solidity: function setMultichainToken(string tokenID, uint256 chainID, string token) returns()
func (_Main *MainTransactorSession) SetMultichainToken(tokenID string, chainID *big.Int, token string) (*types.Transaction, error) {
	return _Main.Contract.SetMultichainToken(&_Main.TransactOpts, tokenID, chainID, token)
}

// SetSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x7926422c.
//
// Solidity: function setSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Main *MainTransactor) SetSwapAndFeeConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setSwapAndFeeConfig", tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap, maxFee, minFee, feeRate)
}

// SetSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x7926422c.
//
// Solidity: function setSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Main *MainSession) SetSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetSwapAndFeeConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap, maxFee, minFee, feeRate)
}

// SetSwapAndFeeConfig is a paid mutator transaction binding the contract method 0x7926422c.
//
// Solidity: function setSwapAndFeeConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap, uint256 maxFee, uint256 minFee, uint256 feeRate) returns(bool)
func (_Main *MainTransactorSession) SetSwapAndFeeConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int, maxFee *big.Int, minFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetSwapAndFeeConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap, maxFee, minFee, feeRate)
}

// SetSwapConfig is a paid mutator transaction binding the contract method 0x1dc463b4.
//
// Solidity: function setSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap) returns(bool)
func (_Main *MainTransactor) SetSwapConfig(opts *bind.TransactOpts, tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setSwapConfig", tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap)
}

// SetSwapConfig is a paid mutator transaction binding the contract method 0x1dc463b4.
//
// Solidity: function setSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap) returns(bool)
func (_Main *MainSession) SetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetSwapConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap)
}

// SetSwapConfig is a paid mutator transaction binding the contract method 0x1dc463b4.
//
// Solidity: function setSwapConfig(string tokenID, uint256 srcChainID, uint256 dstChainID, uint256 maxSwap, uint256 minSwap, uint256 bigSwap) returns(bool)
func (_Main *MainTransactorSession) SetSwapConfig(tokenID string, srcChainID *big.Int, dstChainID *big.Int, maxSwap *big.Int, minSwap *big.Int, bigSwap *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetSwapConfig(&_Main.TransactOpts, tokenID, srcChainID, dstChainID, maxSwap, minSwap, bigSwap)
}

// SetSwapConfigs is a paid mutator transaction binding the contract method 0x6da91269.
//
// Solidity: function setSwapConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Main *MainTransactor) SetSwapConfigs(opts *bind.TransactOpts, tokenID string, configs []RouterConfigSwapConfig2) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setSwapConfigs", tokenID, configs)
}

// SetSwapConfigs is a paid mutator transaction binding the contract method 0x6da91269.
//
// Solidity: function setSwapConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Main *MainSession) SetSwapConfigs(tokenID string, configs []RouterConfigSwapConfig2) (*types.Transaction, error) {
	return _Main.Contract.SetSwapConfigs(&_Main.TransactOpts, tokenID, configs)
}

// SetSwapConfigs is a paid mutator transaction binding the contract method 0x6da91269.
//
// Solidity: function setSwapConfigs(string tokenID, (uint256,uint256,uint256,uint256,uint256)[] configs) returns()
func (_Main *MainTransactorSession) SetSwapConfigs(tokenID string, configs []RouterConfigSwapConfig2) (*types.Transaction, error) {
	return _Main.Contract.SetSwapConfigs(&_Main.TransactOpts, tokenID, configs)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x34fd0135.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddr, uint8 decimals, uint256 version, string routerContract, string extra) returns(bool)
func (_Main *MainTransactor) SetTokenConfig(opts *bind.TransactOpts, tokenID string, chainID *big.Int, tokenAddr string, decimals uint8, version *big.Int, routerContract string, extra string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setTokenConfig", tokenID, chainID, tokenAddr, decimals, version, routerContract, extra)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x34fd0135.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddr, uint8 decimals, uint256 version, string routerContract, string extra) returns(bool)
func (_Main *MainSession) SetTokenConfig(tokenID string, chainID *big.Int, tokenAddr string, decimals uint8, version *big.Int, routerContract string, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetTokenConfig(&_Main.TransactOpts, tokenID, chainID, tokenAddr, decimals, version, routerContract, extra)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x34fd0135.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddr, uint8 decimals, uint256 version, string routerContract, string extra) returns(bool)
func (_Main *MainTransactorSession) SetTokenConfig(tokenID string, chainID *big.Int, tokenAddr string, decimals uint8, version *big.Int, routerContract string, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetTokenConfig(&_Main.TransactOpts, tokenID, chainID, tokenAddr, decimals, version, routerContract, extra)
}

// SetTokenExtraConfig is a paid mutator transaction binding the contract method 0x84cd5b5a.
//
// Solidity: function setTokenExtraConfig(string tokenID, uint256 chainID, string extra) returns(bool)
func (_Main *MainTransactor) SetTokenExtraConfig(opts *bind.TransactOpts, tokenID string, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setTokenExtraConfig", tokenID, chainID, extra)
}

// SetTokenExtraConfig is a paid mutator transaction binding the contract method 0x84cd5b5a.
//
// Solidity: function setTokenExtraConfig(string tokenID, uint256 chainID, string extra) returns(bool)
func (_Main *MainSession) SetTokenExtraConfig(tokenID string, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetTokenExtraConfig(&_Main.TransactOpts, tokenID, chainID, extra)
}

// SetTokenExtraConfig is a paid mutator transaction binding the contract method 0x84cd5b5a.
//
// Solidity: function setTokenExtraConfig(string tokenID, uint256 chainID, string extra) returns(bool)
func (_Main *MainTransactorSession) SetTokenExtraConfig(tokenID string, chainID *big.Int, extra string) (*types.Transaction, error) {
	return _Main.Contract.SetTokenExtraConfig(&_Main.TransactOpts, tokenID, chainID, extra)
}

// SetTokenRouterContract is a paid mutator transaction binding the contract method 0x1910dbdd.
//
// Solidity: function setTokenRouterContract(string tokenID, uint256 chainID, string routerContract) returns(bool)
func (_Main *MainTransactor) SetTokenRouterContract(opts *bind.TransactOpts, tokenID string, chainID *big.Int, routerContract string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setTokenRouterContract", tokenID, chainID, routerContract)
}

// SetTokenRouterContract is a paid mutator transaction binding the contract method 0x1910dbdd.
//
// Solidity: function setTokenRouterContract(string tokenID, uint256 chainID, string routerContract) returns(bool)
func (_Main *MainSession) SetTokenRouterContract(tokenID string, chainID *big.Int, routerContract string) (*types.Transaction, error) {
	return _Main.Contract.SetTokenRouterContract(&_Main.TransactOpts, tokenID, chainID, routerContract)
}

// SetTokenRouterContract is a paid mutator transaction binding the contract method 0x1910dbdd.
//
// Solidity: function setTokenRouterContract(string tokenID, uint256 chainID, string routerContract) returns(bool)
func (_Main *MainTransactorSession) SetTokenRouterContract(tokenID string, chainID *big.Int, routerContract string) (*types.Transaction, error) {
	return _Main.Contract.SetTokenRouterContract(&_Main.TransactOpts, tokenID, chainID, routerContract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x747fd44c.
//
// Solidity: function updateConfig() returns()
func (_Main *MainTransactor) UpdateConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "updateConfig")
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x747fd44c.
//
// Solidity: function updateConfig() returns()
func (_Main *MainSession) UpdateConfig() (*types.Transaction, error) {
	return _Main.Contract.UpdateConfig(&_Main.TransactOpts)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x747fd44c.
//
// Solidity: function updateConfig() returns()
func (_Main *MainTransactorSession) UpdateConfig() (*types.Transaction, error) {
	return _Main.Contract.UpdateConfig(&_Main.TransactOpts)
}

// MainUpdateConfigIterator is returned from FilterUpdateConfig and is used to iterate over the raw logs and unpacked data for UpdateConfig events raised by the Main contract.
type MainUpdateConfigIterator struct {
	Event *MainUpdateConfig // Event containing the contract specifics and raw log

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
func (it *MainUpdateConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUpdateConfig)
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
		it.Event = new(MainUpdateConfig)
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
func (it *MainUpdateConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUpdateConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUpdateConfig represents a UpdateConfig event raised by the Main contract.
type MainUpdateConfig struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateConfig is a free log retrieval operation binding the contract event 0x22590461e7ba17e1fe7580cb0ea47f283d3b2248f04873dfbe926d08fe4c5ab9.
//
// Solidity: event UpdateConfig()
func (_Main *MainFilterer) FilterUpdateConfig(opts *bind.FilterOpts) (*MainUpdateConfigIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "UpdateConfig")
	if err != nil {
		return nil, err
	}
	return &MainUpdateConfigIterator{contract: _Main.contract, event: "UpdateConfig", logs: logs, sub: sub}, nil
}

// WatchUpdateConfig is a free log subscription operation binding the contract event 0x22590461e7ba17e1fe7580cb0ea47f283d3b2248f04873dfbe926d08fe4c5ab9.
//
// Solidity: event UpdateConfig()
func (_Main *MainFilterer) WatchUpdateConfig(opts *bind.WatchOpts, sink chan<- *MainUpdateConfig) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "UpdateConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUpdateConfig)
				if err := _Main.contract.UnpackLog(event, "UpdateConfig", log); err != nil {
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
func (_Main *MainFilterer) ParseUpdateConfig(log types.Log) (*MainUpdateConfig, error) {
	event := new(MainUpdateConfig)
	if err := _Main.contract.UnpackLog(event, "UpdateConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
