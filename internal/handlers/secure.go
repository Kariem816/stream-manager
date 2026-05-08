package handlers

import (
	"net/http"

	"github.com/Kariem816/webrtc-control/config"
)

func Secure(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != config.Cfg.SharedSecretKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
