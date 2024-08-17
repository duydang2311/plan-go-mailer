package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"plan/internal/handlers"
	"plan/internal/mailing"
	"plan/internal/runtime"
	"plan/internal/utils"

	"github.com/nats-io/nats.go"
)

func main() {
	defer log.Println("exiting...")

	RESEND_API_KEY := utils.RequireEnv("RESEND_API_KEY")
	NATS_URL := utils.RequireEnv("NATS_URL")
	NATS_USER := utils.RequireEnv("NATS_USER")
	NATS_PASSWORD := utils.RequireEnv("NATS_PASSWORD")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	var conn *nats.Conn
	var err error
outer:
	for {
		select {
		case <-quit:
			return
		default:
			conn, err = nats.Connect(NATS_URL, nats.UserInfo(NATS_USER, NATS_PASSWORD))

			if err != nil {
				log.Println(err)
				log.Println("retrying...")
				time.Sleep(time.Second)
			} else {
				break outer
			}
		}
	}

	log.Println("connected to nats")
	defer conn.Close()

	runtime := runtime.Runtime{
		Context: &context,
		Nats:    conn,
		Mailer: mailing.NewResend(&mailing.ResendOptions{
			ApiKey: RESEND_API_KEY,
		}),
	}

	sub, err := handlers.SendMailHandler(&runtime)
	if err != nil {
		defer sub.Unsubscribe()
	}

	<-quit
}
