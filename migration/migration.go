package migration

import (
    "github.com/dinhtp/project-recess/migration/versions"
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    db = db.Debug()

    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID:      "20221101000000",
            Migrate: versions.Version20221101000000,
        },
    })

    return m.Migrate()
}
