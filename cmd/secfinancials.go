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
	rootCmd.AddCommand(secfinancialsCmd)
}

var secfinancialsCmd = &cobra.Command{
	Use:   "secfinancials [stock]",
	Short: "Retrieve the financials as reported for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		financials, err := client.QuarterlyFinancialsAsReported(context.Background(), stock, 4)
		if err != nil {
			log.Fatalf("Error getting quarterly SEC financials: %s", err)
		}
		b, err := json.MarshalIndent(financials, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Quarterly SEC 10-Q financials ##")
		fmt.Println(string(b))
	},
}
