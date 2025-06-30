package model

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Name      Name      `gorm:"embedded"`
	Address   Address   `gorm:"embedded"`
	IsActive  bool      `gorm:"default:false"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      UserRole  `gorm:"type:varchar(20);default:student;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Name struct {
	FirstName  string
	MiddleName string
	LastName   string
}

type Address struct {
	PermanentAddress string
	CurrentAddress   string
	PostalCode       int32
}

type UserRole string

const (
	RoleStudent UserRole = "student"
	RoleFaculty UserRole = "faculty"
	RoleMLS     UserRole = "mls"
	RoleGurdian UserRole = "gurdian"
)

func (r UserRole) IsUserRoleValid() bool {
	switch r {
	case RoleStudent, RoleFaculty, RoleGurdian, RoleMLS:
		return true
	}
	return false
}
