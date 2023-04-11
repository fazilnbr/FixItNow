package domain

type Signup struct {
	CountryCode string `json:"countrycode"`
	PhoneNumber string `json:"phonenumber"`
	Otp         string `json:"otp"`
}
