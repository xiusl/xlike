package biz

import "github.com/go-kratos/kratos/v2/log"

// UseCase is .
type UseCase struct {
    repo UserRepo
    log *log.Helper
}

// NewUseCase is .
func NewUseCase(repo UserRepo, logger log.Logger) *UseCase {
    return &UseCase{
        repo: repo,
        log: log.NewHelper(log.With(logger, "moudle", "user/usecase")),
    }
}