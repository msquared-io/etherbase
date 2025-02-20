package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// Config holds the runtime configuration
type Config struct {
	ReaderPort string
	WriterPort string
	RpcURL      string
	Environment          string
	Port                 int
	EtherbaseAddress       common.Address
	MulticallAddress     common.Address
	LocalPrivateKey      string // For local usage
	PrivateKey          string
	PollingIntervalMS    int
}

func LoadConfig() (*Config, error) {
	readerPort := os.Getenv("READER_PORT")
	if readerPort == "" {
		readerPort = "8082"
	}

	writerPort := os.Getenv("WRITER_PORT")
	if writerPort == "" {
		writerPort = "8081"
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "wss://dream-rpc.somnia.network/ws"
	}

	intervalStr := os.Getenv("POLL_INTERVAL_MS")
	if intervalStr == "" {
		intervalStr = "1000"
	}
	interval, _ := strconv.Atoi(intervalStr)

	return &Config{
		ReaderPort:         readerPort,
		WriterPort:         writerPort,
		RpcURL:             rpcURL,
		Environment:       os.Getenv("ENV"),
		LocalPrivateKey:   os.Getenv("LOCAL_PRIVATE_KEY"),
		PrivateKey:        os.Getenv("PRIVATE_KEY"),
		PollingIntervalMS: interval,
		EtherbaseAddress:    common.HexToAddress("0x4B5c05B16Ef39a3a3E1AA8A5162acE704Fd6DC61"),
		MulticallAddress:    common.HexToAddress("0x6ac4600Cc9F3c94ba5Ef63cAFB0b51f43083F31D"),
	}, nil
}

func (c *Config) String() string {
	return fmt.Sprintf("Env=%s Port=%d Etherbase=%s Multicall=%s RPC=%s", 
		c.Environment, c.Port, c.EtherbaseAddress.Hex(), c.MulticallAddress.Hex(), c.RpcURL)
}
