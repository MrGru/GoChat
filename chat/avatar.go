package main

import "errors"

// ErrNoAvatar is the error that is returned when the Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

// Avatar represents types capable of representing user profile pictures.
type Avatar interface {
	//GetAvatarURL gets the avatar URL for the specified client, or returns an error if something goes wrong.
	// a URL for the specific clients.
	GetAvatarURL(c *client) (string, error)
}
