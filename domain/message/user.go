package message

type User struct {
    ID             uint   `json:"id"`
    LocationId     uint   `json:"location_id"`
    BusinessUnitID uint   `json:"business_unit_id"`
    Active         bool   `json:"active"`
    Internal       bool   `json:"internal"`
    FirstLogin     bool   `json:"first_login"`
    Email          string `json:"email"`
    Password       string `json:"password"`
    CasbinUser     string `json:"casbin_user"`
    AuthSource     string `json:"auth_source"`
    FullName       string `json:"full_name"`
    FirstName      string `json:"first_name"`
    LastName       string `json:"last_name"`
    Note           string `json:"note"`
    CareerMission  string `json:"career_mission"`
    FreedomDate    string `json:"freedom_date"`
    LastLoginTime  string `json:"last_login_time"`
    AccountType    string `json:"account_type"`
    BillingStatus  string `json:"billing_status"`
    CreatedAt      string `json:"created_at"`
    UpdatedAt      string `json:"updated_at"`
}

type ListUserRequest struct {
    Page    uint `json:"page"`
    PerPage uint `json:"per_page"`
}

type ListUserResponse struct {
    Items      []*User `json:"items"`
    TotalCount uint    `json:"total_count"`
    MaxPage    uint    `json:"max_page"`
    Page       uint    `json:"page"`
    PerPage    uint    `json:"per_page"`
}

type LoginUserRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginUserResponse struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    Token string `json:"token"`
}
