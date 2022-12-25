package app

import (
	"Bookstore/internal/app/config"
	"Bookstore/internal/app/dsn"
	"Bookstore/internal/app/redis"
	"Bookstore/internal/app/repository"
	"context"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	config *config.Config
	repo   *repository.Repository
	redis  *redis.Client
}

func New(ctx context.Context) (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	repo, err := repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.New(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}

	return &Application{
		config: cfg,
		repo:   repo,
		redis:  redisClient,
	}, nil
}

func (a *Application) Run() error {
	log.Println("application start running")
	a.StartServer()
	log.Println("application shut down")

	return nil
}
