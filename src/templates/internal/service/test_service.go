package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"threeton-starter/internal/model"
	"threeton-starter/internal/repository"
)

type TestService struct {
	db             *gorm.DB
	validate       *validator.Validate
	log            *zap.SugaredLogger
	testRepository *repository.TestRepository
}

func NewTestService(db *gorm.DB, validate *validator.Validate, logger *zap.SugaredLogger, testRepository *repository.TestRepository) *TestService {
	return &TestService{
		db:             db,
		validate:       validate,
		log:            logger,
		testRepository: testRepository,
	}
}

func (s *TestService) SayHello(ctx context.Context, request *model.SayHelloRequest) (string, error) {
	if err := s.validate.StructCtx(ctx, request); err != nil {
		s.log.Errorf("failed to validate request body: %v", err)
		return "", nil
	}

	return s.testRepository.SayHello(request.Name), nil
}
