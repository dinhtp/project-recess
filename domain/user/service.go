package user

import (
    "context"

    "github.com/dinhtp/project-recess/domain/message"
    "gorm.io/gorm"
)

type Service struct {
    repo *Repository
}

func NewService(db *gorm.DB) *Service {
    return &Service{repo: NewRepository(db)}
}

func (s *Service) Get(ctx context.Context, ID uint) (*message.User, error) {
    result, err := s.repo.Read(ID)
    if err != nil {
        return nil, err
    }

    return prepareDataToResponse(result), nil
}

func (s *Service) List(ctx context.Context) {

}

func (s *Service) Create(ctx context.Context) {

}

func (s *Service) Update(ctx context.Context) {

}

func (s *Service) Delete(ctx context.Context) {

}
