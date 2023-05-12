package main

// Config : global config struct
type FileConfig struct {
	TokenList []ScanInfo `yaml:"TokenList"`
	RedisConf RedisConf  `yaml:"RedisConf"`
}

type ScanInfo struct {
	ChainID           string   `yaml:"ChainID"`
	ChainName         string   `yaml:"ChainName,omitempty"`
	TxType            string   `yaml:"TxType,omitempty"`
	PostURL           string   `yaml:"PostURL,omitempty"`
	StartBlock        int      `yaml:"StartBlock,omitempty"`
	BlockStep         int      `yaml:"BlockStep,omitempty"`
	StableBlockHeight int      `yaml:"StableBlockHeight"` // if not 0, scan backward StableBlockHeight
	TopicList         []string `yaml:"TopicList,omitempty"`
	RouterList        []string `yaml:"RouterList,omitempty"`
	WhiteList         []string `yaml:"WhiteList,omitempty"`
	Interval          int      `yaml:"Interval,omitempty"`
	GetReceipt        bool     `yaml:"GetReceipt,omitempty"`
}

type RedisConf struct {
	Ip               string `yaml:"Ip,omitempty"`
	Port             string `yaml:"Port,omitempty"`
	DbId             int    `yaml:"DbId,omitempty"`
	Key              string `yaml:"Key,omitempty"`
	ExpireTimeoutSec int    `yaml:"ExpireTimeoutSec,omitempty"`
	PoolSize         int    `yaml:"PoolSize,omitempty"`
	DialTimeoutSec   int    `yaml:"DialTimeoutSec,omitempty"`
}
