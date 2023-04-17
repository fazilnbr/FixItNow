package domain

type Signup struct {
	CountryCode string `json:"countrycode"`
	PhoneNumber string `json:"phonenumber"`
	Otp         string `json:"otp"`
}

type UserData struct {
	UserId       int
	Email        string
	FirstName    string
	LastName     string
	Gender       string
	Dob          string
	ProfilePhoto string `json:"profilephoto"  binding:"required"`
}
