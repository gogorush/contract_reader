package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAllChainConfig(contract, postURL, targetFile string) {
	// Connect to an Ethereum node
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_goerli")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	// Set the smart contract address
	contractAddress := common.HexToAddress(contract)

	// Create a new instance of the RouterConfig contract
	router, err := NewMainCaller(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create a new instance of the RouterConfig contract: %v", err)
	}

	// Call the getAllChainConfig function
	allChainConfigs, err := router.GetAllChainConfig(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to call getAllChainConfig: %v", err)
	}

	var scanList []ScanInfo
	// Process and display the result
	for _, chainConfig := range allChainConfigs {
		var info ScanInfo
		info.ChainID = chainConfig.ChainID.String()
		info.ChainName = chainConfig.BlockChain
		info.Interval = 1
		info.StableBlockHeight = 100
		info.BlockStep = 1000
		info.PostURL = fmt.Sprintf(postURL, chainConfig.ChainID)
		info.TopicList = []string{
			"0x17dac14bf31c4070ebb2dc182fc25ae5df58f14162a7f24a65b103e22385af0d",
			"0x9ca1de98ebed0a9c38ace93d3ca529edacbbe199cf1b6f0f416ae9b724d4a81c",
			"0xcaac11c45e5fdb5c513e20ac229a3f9f99143580b5eb08d0fecbdd5ae8c81ef5",
			"0xa17aef042e1a5dd2b8e68f0d0d92f9a6a0b35dc25be1d12c0cb3135bfd8951c9",
			"0x36850177870d3e3dca07a29dcdc3994356392b81c60f537c1696468b1a01e61d",
		}
		info.RouterList = []string{chainConfig.RouterContract}
		scanList = append(scanList, info)
	}
	var file ymlfile
	file.TokenList = scanList
	writeScanInfoToYAMLFile(file, targetFile)
}

type ymlfile struct {
	TokenList []ScanInfo `yaml:"TokenList"`
}

func writeScanInfoToYAMLFile(i interface{}, filename string) error {

	// Marshal the list of ScanInfo structs into YAML format
	yamlData, err := yaml.Marshal(&i)
	if err != nil {
		return fmt.Errorf("failed to marshal ScanInfo list to YAML: %v", err)
	}

	// Write the YAML data to a file
	err = ioutil.WriteFile(filename, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML data to file: %v", err)
	}

	return nil
}
