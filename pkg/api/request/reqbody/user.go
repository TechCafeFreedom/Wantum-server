package reqbody

type UserCreate struct {
	AuthID    string
	UserName  string
	Mail      string
	Name      string
	Thumbnail string
	Bio       string
	Gender    int
	Phone     string
	Place     string
	Birth     string
}
