package versions

import (
    "time"

    "gorm.io/gorm"
)

func Version20221101000000(tx *gorm.DB) error {
    type User struct {
        ID        uint `gorm:"TYPE:BIGINT(20) UNSIGNED AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
        CreatedAt time.Time
        UpdatedAt time.Time
        DeletedAt gorm.DeletedAt `gorm:"index"`

        Email          string `gorm:"TYPE:VARCHAR(255)"`
        Password       string `gorm:"TYPE:VARCHAR(255)"`
        CasbinUser     string `gorm:"TYPE:VARCHAR(255)"`
        AuthSource     string `gorm:"TYPE:VARCHAR(255)"`
        FullName       string `gorm:"TYPE:VARCHAR(255)"`
        FirstName      string `gorm:"TYPE:VARCHAR(255)"`
        LastName       string `gorm:"TYPE:VARCHAR(255)"`
        Note           string `gorm:"TYPE:VARCHAR(255)"`
        Active         bool
        Internal       bool
        LocationId     uint   `gorm:"TYPE:BIGINT(20)"`
        CareerMission  string `gorm:"TYPE:VARCHAR(255)"`
        FreeDomDate    time.Time
        BusinessUnitID uint `gorm:"TYPE:BIGINT(20)"`
        LastLoginTime  time.Time
        FirstLogin     bool
        AccountType    string `gorm:"TYPE:VARCHAR(255)"`
        BillingStatus  string `gorm:"TYPE:VARCHAR(255)"`
    }

    return tx.AutoMigrate(&User{})
}
