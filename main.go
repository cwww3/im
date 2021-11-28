package main

import (
	"context"
	"fmt"
	"github.com/cwww3/go-tools/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var env *string

func init() {
	env = pflag.String("env", "prod", "开发环境")
	pflag.Parse()
	loadConfig()
	logger.InitLogger(viper.GetString("log.path"),viper.GetString("log.name"))
}

func main() {
	server := http.Server{
		Addr: fmt.Sprintf("%v:%v", viper.GetString("server.host"), viper.GetInt("server.port")),
	}
	go func() {
		logger.Infof("server run listening url=%v env=%v", fmt.Sprintf("http://%v:%v", viper.GetString("server.host"), viper.GetString("server.port")), *env)
		err := server.ListenAndServe()
		if err != nil {
			logger.Errorf("server run err = %v", err)
		}
	}()

	// 优雅关闭
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Errorf("server shutdown err = %v", err)
		} else {
			cancel()
			logger.Infof("server shutdown gracefully")
		}
	}
}

func loadConfig() {
	viper.AddConfigPath("configs/")
	viper.SetConfigName(fmt.Sprintf("settings-%v", *env))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
