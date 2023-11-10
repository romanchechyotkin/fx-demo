package service

import "go.uber.org/zap"

type Service struct {
	log    *zap.Logger
	worker chan struct{}
}
