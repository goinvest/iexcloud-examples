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
	"time"

	"github.com/goinvest/iexcloud-examples/domain"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(quoteStreamCmd)
}

var quoteStreamCmd = &cobra.Command{
	Use:   "quote-stream [stocks]",
	Short: "Open a stream to report stock quotes for the given symbols.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithSSEBaseURL(cfg.SSEBaseURL))

		fmt.Println("## Opening SSE Stream, Press Ctrl-C to Exit ##\n")
		if err := client.QuoteStream(context.Background(), args, false, func(quotes []iex.Quote) {
			b, err := json.MarshalIndent(quotes, "", "  ")
			if err != nil {
				log.Fatalf("Error marshaling into JSON: %s", err)
			}
			fmt.Printf("## Quote(s) Received at %v ##\n", time.Now())
			fmt.Println(string(b))
			fmt.Println()
		}); err != nil {
			log.Fatalf("Error streaming quotes: %s", err)
		}
	},
}
