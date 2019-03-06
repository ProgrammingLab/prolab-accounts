package model

// UserID represents an user id
type UserID int64

// Authority represents user's authority
type Authority int32

const (
	// Member represents member's authority
	Member Authority = iota
	// Admin represents admin's authority
	Admin
)
