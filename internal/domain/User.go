package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified"`
	UserType  string    `json:"user_type"`
}
