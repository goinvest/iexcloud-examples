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
	rootCmd.AddCommand(iposTodayCmd)
}

var iposTodayCmd = &cobra.Command{
	Use:   "ipos-today",
	Short: "Retrieve today's IPOs",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		ipos, err := client.IPOsToday(context.Background())
		if err != nil {
			log.Fatalf("Error getting IPOs: %s", err)
		}
		b, err := json.MarshalIndent(ipos, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## IPOs ##")
		fmt.Println(string(b))
	},
}
