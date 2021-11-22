package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authorization []Auth

func (authorization Authorization) authorizationTable() map[string][]byte {
	m := make(map[string][]byte)
	for _, item := range authorization {
		passwordHash := sha256.Sum256([]byte(item.Password))
		m[item.Username] = passwordHash[:]
	}
	return m
}

func (authorization Authorization) authorizationHandler(w http.ResponseWriter, r *http.Request) bool {
	m := authorization.authorizationTable()
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"private\"")
		w.WriteHeader(401)
		_, _ = w.Write([]byte("401 Unauthorized\n"))
		return false
	}
	passwordHash := sha256.Sum256([]byte(password))
	if subtle.ConstantTimeCompare(m[username], passwordHash[:]) == 0 {
		w.WriteHeader(401)
		_, _ = w.Write([]byte("401 Unauthorized\n"))
		return false
	}
	return true
}
