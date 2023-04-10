package domain

import "gorm.io/gorm"

// user schema for user table to get listed all users
type User struct {
	// gorm.Model

	IdLogin      int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserName     string `json:"username" gorm:"not null;unique" binding:"required,email"`
	Password     string `json:"password"  binding:"required,min=5"`
	UserType     string `json:"usertype" postgres:"type:ENUM('admin', 'worker', 'user')" gorm:"not null"`
	Verification bool   `json:"-" gorm:"default:false"`
	Status       string `json:"-" gorm:"default:newuser"`
}
type Profile struct {
	IdUser        int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	LoginId       int    `json:"-" gorm:"unique"`
	Login         *User  `json:"-" gorm:"foreignKey:LoginId;references:IdLogin"`
	Name          string `json:"name" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
	DateOfBirth   string `json:"dateofbirth" binding:"required"`
	HouseName     string `json:"housename" binding:"required"`
	Place         string `json:"place" binding:"required"`
	Post          string `json:"post" binding:"required"`
	Pin           string `json:"pin" binding:"required"`
	ContactNumber string `gorm:"unique" json:"contactnumber" binding:"required,min=10"`
	EmailID       string `gorm:"unique" json:"emailid" binding:"required"`
	Photo         string `json:"photo" binding:"required"`
}

//to store mail verification details

type Verification struct {
	gorm.Model
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Category struct {
	IdCategory int    `json:"id_category" gorm:"primaryKey;autoIncrement:true;unique"`
	Category   string `gorm:"unique" json:"category" binding:"required"`
}

type Job struct {
	IdJob       int       `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	CategoryId  int       `json:"categoryid" gorm:"not null"`
	Category    *Category `json:"-" gorm:"foreignKey:CategoryId;references:IdCategory"`
	IdWorker    int       `json:"-" gorm:"not null"`
	Login       *User     `json:"-,omitempty" bson:",omitempty" gorm:"foreignKey:IdWorker;references:IdLogin"`
	Wage        int       `json:"wage" gorm:"not null" binding:"required"`
	Description string    `json:"desctription" binding:"required"`
}

type Favorite struct {
	IdFavorite int   `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId     int   `json:"-"`
	User       *User `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	JobId      int   `json:"jobid" binding:"required"`
	job        *Job  `json:"-" gorm:"foreignKey:JobId;references:IdJob;unique"`
}

type Address struct {
	IdAddress int `gorm:"primaryKey;autoIncrement:true;unique"`
	UserId    int
	User      *User  `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	HouseName string `binding:"required"`
	Place     string `binding:"required"`
	City      string `binding:"required"`
	Post      string `binding:"required"`
	Pin       string `binding:"required"`
	Phone     string `binding:"required,min=10"`
}

type Request struct {
	IdRequset int `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId    int
	User      *User `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	JobId     int
	Job       *Job     `json:"-" gorm:"foreignKey:JobId;references:IdJob"`
	AddressId int      `binding:"required"`
	Address   *Address `json:"-" gorm:"foreignKey:AddressId;references:IdAddress"`
	Status    string   `json:"-" gorm:"default:pending"`
	Date      string   `jsom:"-"`
}

type JobPayment struct {
	IdPayment     int `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	RequestId     int
	Request       *Request `json:"-" gorm:"foreignKey:RequestId;references:IdRequset"`
	OrderId       string
	RazorPaymetId string
	UserId        int
	User          *User `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	Amount        int
	Date          string
	PaymentStatus string `json:"-" gorm:"default:orderd"`
}
