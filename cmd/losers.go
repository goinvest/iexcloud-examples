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
	"strconv"

	"github.com/goinvest/iexcloud-examples/domain"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(losersCmd)
}

var losersCmd = &cobra.Command{
	Use:   "losers [limit]",
	Short: "Retrieve the list of losers.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var limit int
		if len(args) == 1 {
			var err error
			if limit, err = strconv.Atoi(args[0]); err != nil {
				log.Fatalf("Invalid limit, must be a number: %s", args[0])
			}
		}
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		quote, err := client.Losers(context.Background(), limit)
		if err != nil {
			log.Fatalf("Error getting list: %s", err)
		}
		b, err := json.MarshalIndent(quote, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Gainers ##")
		fmt.Println(string(b))
	},
}
