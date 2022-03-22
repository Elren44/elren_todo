package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

var s *scs.SessionManager

func SendSessions(manager *scs.SessionManager) {
	s = manager
}

//NoSurf adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {

	scrfHandler := nosurf.New(next)
	scrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return scrfHandler
}

//SessionLoad loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	return s.LoadAndSave(next)
}
