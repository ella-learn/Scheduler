package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// func PrintAgenda() {
// 	logger.Info("Berhasil print agenda")
// }

func main() {
	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	defer scheduler.Stop()

	logger, _ := zap.NewDevelopment()

	defer logger.Sync()

	scheduler.AddFunc("*/1 * * * *", func() { logger.Info("This is from zap!") })

	go scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
