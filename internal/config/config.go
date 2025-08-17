package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string        `mapstructure:"telegram_token"`
	AppEnv        string        `mapstructure:"app_env"` // dev|prod
	LogPath       string        `mapstructure:"log_path"`
	DefaultTZ     string        `mapstructure:"default_tz"`
	SummaryFreq   string        `mapstructure:"summary_freq"` // off|daily|weekly|monthly
	ReadTimeout   time.Duration `mapstructure:"read_timeout"`
	WriteTimeout  time.Duration `mapstructure:"write_timeout"`
}

func setDefaults() {
	viper.SetDefault("telegram_token", "")
	viper.SetDefault("app_env", "dev")
	viper.SetDefault("log_path", "logs/app.log")
	viper.SetDefault("default_tz", "UTC")
	viper.SetDefault("summary_freq", "off")
	viper.SetDefault("read_timeout", "5s")
	viper.SetDefault("write_timeout", "5s")
}

func initEnv() {
	viper.SetEnvPrefix("DSB")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func BindFlags(fs *pflag.FlagSet) {
	fs.String("config", "", "Path to config file (yaml)")
	fs.String("token", "", "Telegram bot token")
	fs.String("env", "dev", "Environment (dev|prod)")
	fs.String("log", "logs/app.log", "Log file path")
	fs.String("tz", "UTC", "Default time zone")
	fs.String("summary", "off", "Daily summary: off|daily|weekly")
	fs.Duration("read-timeout", 5*time.Second, "Read timeout")
	fs.Duration("write-timeout", 5*time.Second, "Write timeout")

	// Bind к ключам viper
	_ = viper.BindPFlag("config", fs.Lookup("config"))
	_ = viper.BindPFlag("telegram_token", fs.Lookup("token"))
	_ = viper.BindPFlag("app_env", fs.Lookup("env"))
	_ = viper.BindPFlag("log_path", fs.Lookup("log"))
	_ = viper.BindPFlag("default_tz", fs.Lookup("tz"))
	_ = viper.BindPFlag("summary_freq", fs.Lookup("summary"))
	_ = viper.BindPFlag("read_timeout", fs.Lookup("read-timeout"))
	_ = viper.BindPFlag("write_timeout", fs.Lookup("write-timeout"))
}

func Load(configPath string) (Config, error) {
	setDefaults()
	initEnv()

	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	_ = viper.ReadInConfig()
	fmt.Println("cfg file:", viper.ConfigFileUsed())

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("unable to decode config into struct, %w", err)
	}
	if err := validate(cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func validate(c Config) error {
	if c.AppEnv != "dev" && c.AppEnv != "prod" {
		return fmt.Errorf("app_env must be either 'dev' or 'prod'")
	}
	switch c.SummaryFreq {
	case "off", "daily", "weekly":
	default:
		return fmt.Errorf("summary_freq must be off|daily|weekly, got %q", c.SummaryFreq)
	}
	if c.DefaultTZ == "" {
		return fmt.Errorf("default_tz must be set")
	}
	return nil
}
