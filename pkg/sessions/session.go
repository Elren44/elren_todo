package sessions

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

func NewSessions() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	return session
}
