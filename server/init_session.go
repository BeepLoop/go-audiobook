package server

import (
	"github.com/BeepLoop/go-audiobook/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func InitSession() cookie.Store {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		Domain:   config.Env.IP,
		MaxAge:   60 * 60 * 24,
	})

	return store
}
