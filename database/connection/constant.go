package connection

import (
    "time"
)

const (
    DbTypeMySql  = "mysql"
    DbTypeSqLite = "sqlite"

    DefaultIdleConnection = 10
    DefaultMaxConnection  = 20
    DefaultLifeTime       = 300 * time.Minute
)
