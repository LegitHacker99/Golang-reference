package users

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserRole string

const (
	Admin  UserRole = "Admin"
	Member UserRole = "Member"
	Guest  UserRole = "Guest"
)

type UserActivities struct {
	Id                  uint      `json:"id"`
	Role_Updated_At     time.Time `json:"role_updated_at"`
	Role_Update_Detail  string    `json:"role_update_detail"`
	Other_Updated_At    time.Time `json:"other_updated_at"`
	Other_Update_Detail string    `json:"other_update_detail"`
	User_Uuid           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type User struct {
	User_Uuid  uuid.UUID `json:"userUuid" gorm:"type:uuid; default:uuid_generate_v4(); not null"`
	User_Name  string    `json:"username"`
	Email      string    `json:"email" validate:"email"`
	Password   string    `json:"password"`
	Role       UserRole  `json:"role"`
	Created_At time.Time `json:"createdAt"`
}
