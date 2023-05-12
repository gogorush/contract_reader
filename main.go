package main

import (
	"flag"
)

var (
	fileConfig           string // config file path
	chainID              string
	rpcFile              string
	smartContractAddress string
    targetFile string
)

// init program from config
func init() {
    flag.StringVar(&fileConfig, "f", "/Users/hou/Documents/workspace/golang/src/erc20_scanner/default.yml", "-f: config file path")
	flag.StringVar(&rpcFile, "r", "/Users/hou/Documents/workspace/golang/src/scan/rpc.json", "-r: rpc json file")
	flag.StringVar(&smartContractAddress, "d", "0x64d5baf5ac030e2b7c435add967f787ae94d0205", "-s: smart contract address which needed to be read from")
	flag.StringVar(&targetFile, "t", "anycall_test.yml", "-t: targe yml file")
}

var anycall_test_postURL = "http://8.210.58.66:31599/swap/register/%s/%%s"

func main() {
    flag.Parse()
    GetAllChainConfig(smartContractAddress, anycall_test_postURL, targetFile)
}
