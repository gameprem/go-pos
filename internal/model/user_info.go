package model

type UserInfo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"not null;unique" json:"user_id"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
