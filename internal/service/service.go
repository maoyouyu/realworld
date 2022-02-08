package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "helloworld/api/realworld/v1"
	"helloworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)


// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealworldServer

	uc  *biz.UserUsercase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uc *biz.UserUsercase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}

