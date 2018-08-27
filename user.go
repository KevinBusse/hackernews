package hackernews

import (
	"fmt"
)

// User represents the user data
type User struct {
	ID        string `json:"id"`        // The user's unique username. Case-sensitive. Required.
	Delay     int    `json:"delay"`     // Delay in minutes between a comment's creation and its visibility to other users.
	Created   int    `json:"created"`   // Creation date of the user, in Unix Time.
	Karma     int    `json:"karma"`     // The user's karma.
	About     string `json:"about"`     // The user's optional self-description. HTML.
	Submitted []int  `json:"submitted"` // List of the user's stories, polls and comments.
}

func (u User) String() string {
	return fmt.Sprintf("@%s has %d karma, created %d item(s)", u.ID, u.Karma, len(u.Submitted))
}
