package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/FollowLille/daily-secretary-bot/internal/config"
)

func main() {
	_ = godotenv.Load(".env")
	config.BindFlags(pflag.CommandLine)
	pflag.Parse()

	cfgPath := viper.GetString("config")
	cfg, err := config.Load(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Config loaded: env=%s tz=%s summary=%s log=%s\n",
		cfg.AppEnv, cfg.DefaultTZ, cfg.SummaryFreq, cfg.LogPath)

	if cfg.TelegramToken == "" {
		fmt.Println("TELEGRAM token is empty (ok for now).")
	}
}
