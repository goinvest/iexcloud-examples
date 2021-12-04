module github.com/goinvest/iexcloud-examples

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/goinvest/iexcloud/v2 v2.11.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/goinvest/iexcloud/v2 => ../iexcloud/

go 1.11
