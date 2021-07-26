package resp


type User struct {
	UserId int `json:"userId" gorm:"column:user_id"`
	UserName string `json:"userName" gorm:"column:user_name"`
	UserAge int `json:"userAge" gorm:"column:user_age"`
}

