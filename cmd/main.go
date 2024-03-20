package main

import (
	diskspace "diskspacereminder"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Disk Space reminder started")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}

	var cfg diskspace.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	c := time.NewTicker(13 * time.Hour)

	telegram := diskspace.NewTelegram(cfg.Telegram.BaseUrl, cfg.Telegram.Token, cfg.Telegram.GroupId)

	diskInstance := &diskspace.DiskSpace{
		Telegram: telegram,
		Cfg:      cfg.Disk,
	}

	go func() {
		for {
			select {
			case <-c.C:
				diskInstance.WarnDiskSpace()
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("Stop")
	c.Stop()
	fmt.Println("good bye")
}
