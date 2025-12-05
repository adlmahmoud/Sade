package helpers

import (
	"net/http"
	"net/url"
	"strconv"
)

// RedirectToError redirige vers /error en passant code + message en query string.
// À utiliser partout où tu veux sortir proprement vers une page d'erreur "user-friendly".
func RedirectToError(w http.ResponseWriter, r *http.Request, code int, message string) {
	// Construction sûre de la query string (Encode gère l'échappement)
	params := url.Values{}
	if code > 0 {
		params.Set("code", strconv.Itoa(code))
	}
	if message != "" {
		params.Set("message", message)
	}

	pathTarget := "/error"
	// Cible finale : /error ou /error?code=...&message=...
	if encodeParams := params.Encode(); encodeParams != "" {
		pathTarget += "?" + encodeParams
	}
	http.Redirect(w, r, pathTarget, http.StatusSeeOther)
	// Pense à "return" juste après l'appel côté handler appelant !
}
