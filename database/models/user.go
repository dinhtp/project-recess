package models

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email          string
    Password       string
    CasbinUser     string
    AuthSource     string
    FullName       string
    FirstName      string
    LastName       string
    Note           string
    Active         bool
    Internal       bool
    LocationId     uint
    CareerMission  string
    FreeDomDate    *time.Time
    BusinessUnitID uint
    LastLoginTime  *time.Time
    FirstLogin     bool
    AccountType    string
    BillingStatus  string
}
