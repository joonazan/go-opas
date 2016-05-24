package spotify

import (
	"strconv"

	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

const (
	spotifyKeyID         string = "id"
	spotifyKeyName       string = "display_name"
	spotifyKeyNickname   string = "username"
	spotifyKeyEmail      string = "email"
	spotifyKeyPictureUrl string = "images[0].url"
)

type User struct {
	data objx.Map
}

// NewUser builds a new User object for Github.
func NewUser(data objx.Map, creds *common.Credentials, provider common.Provider) *User {
	user := &User{data}

	creds.Set(common.CredentialsKeyID, data[spotifyKeyID])
	// set provider credentials
	user.data[common.UserKeyProviderCredentials] = map[string]*common.Credentials{
		provider.Name(): creds,
	}

	return user
}

// Email gets the users email address.
func (u *User) Email() string {
	return u.Data().Get(spotifyKeyEmail).Str()
}

// Name gets the users full name.
func (u *User) Name() string {
	return u.Data().Get(spotifyKeyName).Str()

}

// Nickname gets the users nickname or username.
func (u *User) Nickname() string {
	return u.Data().Get(spotifyKeyNickname).Str()

}

// AvatarURL gets the URL of an image representing the user.
func (u *User) AvatarURL() string {
	return u.Data().Get(spotifyKeyPictureUrl).Str()
}

// ProviderCredentials gets a map of Credentials (by provider name).
func (u *User) ProviderCredentials() map[string]*common.Credentials {
	return u.Data().Get(common.UserKeyProviderCredentials).Data().(map[string]*common.Credentials)
}

// IDForProvider gets the ID value for the specified provider name for
// this user from the ProviderCredentials data.
func (u *User) IDForProvider(provider string) string {
	id := u.ProviderCredentials()[provider].Get(common.CredentialsKeyID).Data()
	switch id.(type) {
	case string:
		return id.(string)
	case float64:
		return strconv.FormatFloat(id.(float64), 'f', 0, 64)
	}
	return ""
}

// AuthCode gets this user's globally unique ID (generated by the host program)
func (u *User) AuthCode() string {
	return u.Data().Get(common.UserKeyAuthCode).Str()
}

// GetValue gets any User field by name.
func (u *User) Data() objx.Map {
	return u.data
}

func (u *User) PublicData(options map[string]interface{}) (publicData interface{}, err error) {
	return u.data, nil
}
