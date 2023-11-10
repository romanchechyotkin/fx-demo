package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	log    *zap.Logger
	worker chan struct{}
	store  map[int]struct{}
}

func NewService(log *zap.Logger) *Service {
	return &Service{
		log:    log,
		worker: make(chan struct{}),
		store: map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
			5: {},
		},
	}
}

func (s *Service) Run() {
	workerCtx := context.Background()
	s.log.Info("worker started")

	for {
		select {
		case <-s.worker:
			workerCtx.Done()
			return
		case <-time.After(3 * time.Second):
			s.log.Info("service job")

			for k, _ := range s.store {
				fmt.Println(k)
			}
		}
	}

}
