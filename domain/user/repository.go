package user

import (
    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/models"
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Read(ID uint, email string) (*models.User, error) {
    var result *models.User

    query := r.db.Model(&models.User{})

    if ID > 0 {
        query = query.Where("id = ?", ID)
    }

    if email != "" {
        query = query.Where("email = ?", email)
    }

    if err := query.First(&result).Error; nil != err {
        return nil, err
    }

    return result, nil
}

func (r *Repository) List(req *message.ListUserRequest) ([]*models.User, int64, error) {
    var totalCount int64
    var results []*models.User

    limit := int(req.PerPage)
    offset := limit * (int(req.Page) - 1)
    query := r.db.Model(&models.User{}).Order("updated_at DESC")

    if err := query.Select("id").Count(&totalCount).Error; err != nil {
        return nil, 0, err
    }

    if err := query.Select("*").Limit(limit).Offset(offset).Find(&results).Error; err != nil {
        return nil, 0, err
    }

    return results, totalCount, nil
}

func (r *Repository) Insert(o *models.User) (*models.User, error) {
    err := r.db.Create(o).Error
    if err != nil {
        return nil, err
    }

    return o, nil
}

func (r *Repository) Update(o *models.User) (*models.User, error) {
    query := r.db.Select("*").Omit("id,created_at").Where("id = ?", o.ID)

    err := query.Updates(o).Error
    if err != nil {
        return nil, err
    }

    return o, nil
}

func (r *Repository) Delete(o *models.User) error {
    return r.db.Delete(o).Error
}
