package bizmodel

type User struct {
    ID       int    `json:"user_id"`
    Username string `json:"nickname"`
    Password string `json:"password"`
}

