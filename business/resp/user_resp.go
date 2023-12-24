package resp

type UserResp struct {
	UserId      int64  `json:"userId"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Gender      int8   `json:"gender"`
	Country     string `json:"country"`
	Province    string `json:"province"`
	City        string `json:"city"`
}
