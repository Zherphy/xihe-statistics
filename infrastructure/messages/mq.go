package messages

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
)

func Init(cfg mq.MQConfig, log *logrus.Entry, topic config.Topics) error {
	topics = topic

	err := kafka.Init(
		mq.Addresses(cfg.Addresses...),
		mq.Log(log),
	)
	if err != nil {
		return err
	}

	return kafka.Connect()
}

func Exit(log *logrus.Entry) {
	if err := kafka.Disconnect(); err != nil {
		log.Errorf("exit kafka, err:%v", err)
	}
}

func NewHandler(cfg *config.SrvConfig, log *logrus.Entry) *Handler {

	bigModelRecord := repositories.NewBigModelRecordRepository(
		// infrastructure.mongodb -> infrastructure.repositories (mapper)
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

	registerRecord := repositories.NewRegisterRecordRepository(
		pgsql.NewRegisterRecordMapper(pgsql.RegisterRecord{}),
	)

	bs := app.NewBigModelRecordMessageService(bigModelRecord)
	rs := app.NewRepoRecordMessageService(repoRecord)
	rr := app.NewRegisterRecordMessageService(registerRecord)

	return &Handler{
		Log:      log,
		MaxRetry: config.Conf.MaxRetry,

		BigModel: bs,
		Repo:     rs,
		Register: rr,
	}
}

func Run(h *Handler, log *logrus.Entry) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	defer wg.Wait()

	called := false
	ctx, done := context.WithCancel(context.Background())

	defer func() {
		if !called {
			called = true
			done()
		}
	}()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		select {
		case <-ctx.Done():
			log.Info("receive done. exit normally")
			return

		case <-sig:
			log.Info("receive exit signal")
			done()
			called = true
			return
		}
	}(ctx)

	if err := Subscribe(ctx, h, log); err != nil {
		log.Errorf("subscribe failed, err:%v", err)
	}
}