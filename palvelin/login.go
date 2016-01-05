package main

import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"github.com/stretchr/signature"
	"html/template"
	"net/http"
)

var cookieEncrypter = securecookie.New(securecookie.GenerateRandomKey(64), nil)

const cookieName = "user"
const userKey = 0

func init() {

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New("712450255947-o5uppfp82sc3sb2c1dum7es8k93ras4q.apps.googleusercontent.com", "TrFX1I5QJZ_unJ-P5UgYLF4N", PalvelimenOsoite+"/authcallback/google"),
	)

	gob.Register(User{})
}

func authentication(w http.ResponseWriter, r *http.Request) {

	user, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := User{
		Email:     user.Email(),
		Name:      user.Name(),
		Nickname:  user.Nickname(),
		AvatarURL: user.AvatarURL(),
	}

	validatableUser, err := cookieEncrypter.Encode(cookieName, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: validatableUser,
		Path:  "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getUser(r *http.Request) (user common.User, err error) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		panic(err)
	}

	omap, err := objx.FromURLQuery(r.URL.RawQuery)
	if err != nil {
		return
	}

	creds, err := provider.CompleteAuth(omap)
	if err != nil {
		return
	}

	user, err = provider.GetUser(creds)
	return
}

type User struct {
	Email, Name, Nickname, AvatarURL string
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func loginRequired(handler func(http.ResponseWriter, *http.Request, User)) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if user, err := userFromCookie(r); err == nil {
			handler(w, r, user)
		} else {
			login(w, r)
		}
	}
}

func userFromCookie(r *http.Request) (user User, err error) {
	var cookie *http.Cookie
	if cookie, err = r.Cookie(cookieName); err == nil {
		user = User{}
		if err = cookieEncrypter.Decode(cookieName, cookie.Value, &user); err == nil {
			return
		}
	}
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	template, err := template.New("loginpage").Parse(`
			<!DOCTYPE html>
				<body>
				<h2>Log in with...</h2>
				<ul>
					<li>
					<a href="{{.}}">Google</a>
					</li>
				</ul>
				</body>
		`)

	provider, err := gomniauth.Provider("google")
	if err != nil {
		panic(err)
	}

	state := gomniauth.NewState("after", "success")
	authUrl, err := provider.GetBeginAuthURL(state, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template.Execute(w, authUrl)
}