package model

// ProfileScope represents scope of user's profile
type ProfileScope int

const (
	// MembersOnly represents that only members can view the profile
	MembersOnly ProfileScope = iota
	// Public represents that the profile can be seen all users
	Public
	// Private represents that no one can see it(like email)
	Private
)
