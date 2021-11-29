// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/goinvest/iexcloud-examples/domain"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(batchQuoteCmd)
}

var batchQuoteCmd = &cobra.Command{
	Use:   "batch-quote [stocks]",
	Short: "Retrieve the quote data for the list of stock symbols",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		quote, err := client.BatchQuote(context.Background(), args)
		if err != nil {
			log.Fatalf("Error getting quote: %s", err)
		}
		b, err := json.MarshalIndent(quote, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Quote ##")
		fmt.Println(string(b))
	},
}
