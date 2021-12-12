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
	rootCmd.AddCommand(newsCmd)
}

var newsCmd = &cobra.Command{
	Use:   "news [symbol] [limit]",
	Short: "Retrieve news articles about the given symbol. Pass the special symbol 'market' to get market news about a variety of stocks.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		num, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Invalid limit: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		bs, err := client.News(context.Background(), args[0], num)
		if err != nil {
			log.Fatalf("Error getting news: %s", err)
		}
		b, err := json.MarshalIndent(bs, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## News ##")
		fmt.Println(string(b))
	},
}
