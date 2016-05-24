package main

import (
	"encoding/gob"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/github.com/gorilla/securecookie"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/github.com/stretchr/gomniauth"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/github.com/stretchr/gomniauth/common"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/github.com/stretchr/gomniauth/providers/google"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/github.com/stretchr/objx"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/github.com/stretchr/signature"
	"html/template"
	"log"
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

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func init() {
	http.HandleFunc("/authcallback/google", authentication)
}

func authentication(w http.ResponseWriter, r *http.Request) {

	user, redirectTarget, err := getUser(r)
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

	http.Redirect(w, r, redirectTarget, http.StatusSeeOther)
}

func getUser(r *http.Request) (user common.User, target string, err error) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatal(err)
	}

	omap, err := objx.FromURLQuery(r.URL.RawQuery)
	if err != nil {
		return
	}

	creds, err := provider.CompleteAuth(omap)
	if err != nil {
		return
	}

	state, err := gomniauth.StateFromParam(omap.Get("state").String())
	if err != nil {
		return
	}

	target = state.Get("redirect").String()

	user, err = provider.GetUser(creds)
	return
}

type User struct {
	Email, Name, Nickname, AvatarURL string
}

func loginRequired(path string, handler func(http.ResponseWriter, *http.Request, User)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if user, err := userFromCookie(r); err == nil {
			handler(w, r, user)
		} else {
			login(w, r, path)
		}
	})
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

var loginPage *template.Template

func init() {
	var err error
	loginPage, err = template.ParseFiles("data/login.html")
	if err != nil {
		log.Fatal(err)
	}
}

type Provider struct {
	URL  template.URL
	Name string
}

func login(w http.ResponseWriter, r *http.Request, redirect string) {

	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatal(err)
	}

	state := gomniauth.NewState("redirect", redirect)
	authUrl, err := provider.GetBeginAuthURL(state, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	loginPage.Execute(w, []Provider{{URL: template.URL(authUrl), Name: "Google"}})
}
