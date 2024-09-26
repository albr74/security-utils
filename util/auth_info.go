package util

import (
	"net/http"
)

type key int

const SESSION_KEY key = iota

// TODO: colocar aqui a relação de roles que devem ser buscada do Redis
type AuthenticationInfo struct {
	UserUuid   string
	TenantUuid string
	SessionId  string
}

func GetAuthInfo(r *http.Request) *AuthenticationInfo {
	intf := r.Context().Value(SESSION_KEY)
	if intf != nil {
		auth := intf.(AuthenticationInfo)
		return &auth
	} else {
		return nil
	}
}
