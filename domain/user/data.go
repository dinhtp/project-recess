package user

import (
    "time"

    "github.com/dinhtp/project-recess/database/models"
    "github.com/dinhtp/project-recess/domain/message"
)

func prepareDataToResponse(o *models.User) *message.User {
    data := &message.User{
        ID:             o.ID,
        LocationId:     o.LocationId,
        BusinessUnitID: o.BusinessUnitID,
        Active:         o.Active,
        Internal:       o.Internal,
        FirstLogin:     o.FirstLogin,
        Email:          o.Email,
        Password:       o.Password,
        CasbinUser:     o.CasbinUser,
        AuthSource:     o.AuthSource,
        FullName:       o.FullName,
        FirstName:      o.FirstName,
        LastName:       o.LastName,
        Note:           o.Note,
        CareerMission:  o.CareerMission,
        AccountType:    o.AccountType,
        BillingStatus:  o.BillingStatus,
        CreatedAt:      o.CreatedAt.Format(time.RFC3339),
        UpdatedAt:      o.UpdatedAt.Format(time.RFC3339),
    }

    if o.FreeDomDate != nil {
        data.FreedomDate = o.FreeDomDate.Format(time.RFC3339)
    }

    if o.LastLoginTime != nil {
        data.LastLoginTime = o.LastLoginTime.Format(time.RFC3339)
    }

    return data
}
