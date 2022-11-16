package user

import (
    "context"
    "math"

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

func (s *Service) List(ctx context.Context, r *message.ListUserRequest) (*message.ListUserResponse, error) {
    results, total, err := s.repo.List(r)
    if err != nil {
        return nil, err
    }

    var list []*message.User

    for _, result := range results {
        list = append(list, prepareDataToResponse(result))
    }

    return &message.ListUserResponse{
        Items:      list,
        TotalCount: uint(total),
        MaxPage:    uint(math.Ceil(float64(total) / float64(r.PerPage))),
        Page:       r.Page,
        PerPage:    r.PerPage,
    }, nil
}

func (s *Service) Create(ctx context.Context, r *message.User) (*message.User, error) {
    result, err := s.repo.Insert(prepareDataToCreate(r))
    if err != nil {
        return nil, err
    }

    return prepareDataToResponse(result), nil
}

func (s *Service) Update(ctx context.Context, r *message.User) (*message.User, error) {
    _, err := s.repo.Read(r.ID)
    if err != nil {
        return nil, err
    }

    result, err := s.repo.Update(prepareDataToUpdate(r))
    if err != nil {
        return nil, err
    }

    return s.Get(ctx, result.ID)
}

func (s *Service) Delete(ctx context.Context, ID uint) error {
    result, err := s.repo.Read(ID)
    if err != nil {
        return err
    }

    return s.repo.Delete(result)
}
