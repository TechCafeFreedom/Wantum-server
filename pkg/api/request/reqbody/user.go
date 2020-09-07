package reqbody

type UserCreate struct {
	AuthID    string `json:"auth_id"`
	UserName  string `json:"user_name"`
	Mail      string `json:"mail"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Bio       string `json:"bio"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	Place     string `json:"place"`
	Birth     int    `json:"birth"`
}
