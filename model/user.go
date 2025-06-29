package model

import uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"

type User struct {
	ID       uuid.UUID
	Name     Name
	Address  Address
	Email    string
	Password string
}

type Name struct {
	FirstName  string
	MiddleName string
	LastName   string
}

type Address struct {
	permanentAddress string
	currentAddress   string
	postalCode       int32
}
