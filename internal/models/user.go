package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Avatar   string `json:"avatar"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts"`
}
