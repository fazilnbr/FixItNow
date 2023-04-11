package domain

import "gorm.io/gorm"

// user schema for user table to get listed all users
type User struct {
	IdUser       int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserName     string `json:"username" gorm:"not null;unique" binding:"required"`
	Phone        string `json:"phonenumber" gorm:"not null;unique" binding:"required"`
	Email        string `json:"email" gorm:"not null;unique" binding:"required,email"`
	Password     string `json:"password"  binding:"required,min=5"`
	UserType     string `json:"usertype" postgres:"type:ENUM('admin', 'worker', 'user')" gorm:"not null"`
	Verification bool   `json:"-" gorm:"default:false"`
	Status       string `json:"-" gorm:"default:newuser"`
	Profilephoto string `json:"profilephoto"  binding:"required"`
}

type Address struct {
	IdAddress       int `gorm:"primaryKey;autoIncrement:true;unique"`
	UserId          int
	User            *User  `json:"-" gorm:"foreignKey:UserId;references:IdUser"`
	AddressCategory string `json:"category"`
	Mapcoordinates  string `json:"mapcoordinates"`
	Housenumber     string `json:"housenumber"`
	Floor           string `json:"floor"`
	BlockorTower    string `json:"blockortower"`
	Landmark        string `json:"landmark"`
}

type Verification struct {
	gorm.Model
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Category struct {
	IdCategory   int    `json:"id_category" gorm:"primaryKey;autoIncrement:true;unique"`
	Category     string `gorm:"unique" json:"category" binding:"required"`
	CategoryIcon string `gorm:"unique" json:"categoryicon" binding:"required"`
}

type Job struct {
	IdJob       int       `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	IdWorker    int       `json:"-" gorm:"not null"`
	User        *User     `json:"-,omitempty" bson:",omitempty" gorm:"foreignKey:IdWorker;references:IdUser"`
	CategoryId  int       `json:"categoryid" gorm:"not null"`
	Category    *Category `json:"-" gorm:"foreignKey:CategoryId;references:IdCategory"`
	Expirience  string    `json:"expirience" binding:"required"`
	Description string    `json:"desctription" binding:"required"`
	FullDayWage int       `json:"fuldaywage" gorm:"not null" binding:"required"`
	HalfDayWage int       `json:"halfdaywage" gorm:"not null" binding:"required"`
	Openwork    bool      `json:"openwork" gorm:"default:true"`
	Priority    bool      `json:"priority"`
}

type Favorite struct {
	IdFavorite int   `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId     int   `json:"-"`
	User       *User `json:"-" gorm:"foreignKey:UserId;references:IdUser"`
	JobId      int   `json:"jobid" binding:"required"`
	job        *Job  `json:"-" gorm:"foreignKey:JobId;references:IdJob;unique"`
}

type Request struct {
	IdRequset int `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId    int
	User      *User `json:"-" gorm:"foreignKey:UserId;references:IdUser"`
	JobId     int
	Job       *Job     `json:"-" gorm:"foreignKey:JobId;references:IdJob"`
	AddressId int      `binding:"required"`
	Address   *Address `json:"-" gorm:"foreignKey:AddressId;references:IdAddress"`
	Status    string   `json:"-" gorm:"default:pending"`
	Date      string   `jsom:"-"`
	BidAmount string   `jsom:"bidamount"`
}

type Ratings struct {
	IdRatings int `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	WorkerId  int
	User      *User `json:"-" gorm:"foreignKey:WorkerId;references:IdUser"`
	UserId    int
	Users     *User `json:"-" gorm:"foreignKey:WorkerId;references:IdUser"`
	Rating    int   `json:"rating" gorm:"not null"`
}

type Banner struct {
	IdBanner int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	Image    string `json:"image" gorm:"not null"`
	Deeplink string `json:"deeplink" gorm:"not null"`
}
