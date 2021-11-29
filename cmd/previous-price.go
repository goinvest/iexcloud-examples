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
	rootCmd.AddCommand(previousPriceCmd)
}

var previousPriceCmd = &cobra.Command{
	Use:   "previous-price [stock]",
	Short: "Retrieve the previous price for stock symbol(s)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		prices := make(map[string]iex.PreviousDay)
		if len(args) == 1 {
			stock := args[0]
			price, err := client.PreviousDay(context.Background(), stock)
			if err != nil {
				log.Fatalf("Error getting stock price: %s", err)
			}
			prices[stock] = price
		} else {
			prices, err = client.BatchPrevious(context.Background(), args)
			if err != nil {
				log.Fatalf("Error getting stock prices: %s", err)
			}
		}

		b, err := json.MarshalIndent(prices, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Previous Day ##")
		fmt.Println(string(b))
	},
}
