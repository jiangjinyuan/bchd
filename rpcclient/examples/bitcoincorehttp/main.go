// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/gcash/bchd/rpcclient"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:19888",
		User:         "bitcoinrpc1",
		Pass:         "mytestpass123",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	// Get the mempool info.
	mempoolInfo, err := client.GetMempoolInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("mempoolInfo: %v", mempoolInfo)
	// Get the estimate feeRate.
	feeRate, err := client.EstimateFee()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("feeRate: %f", feeRate)
}
