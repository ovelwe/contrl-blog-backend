package models

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	UserID    uint      `json:"user_id" binding:"required"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Likes     int       `json:"likes"`
	CreatedAt string    `json:"created_at"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}
