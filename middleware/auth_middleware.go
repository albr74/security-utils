package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/albr74/commons/redis"
	"github.com/albr74/commons/util"
	sec "github.com/albr74/security-utils/util"
	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processRequest(next, w, r)
	})
}

func processRequest(next http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	value := r.Header.Get("Authorization")
	var tokenString string
	if value == "" {
		token, err := r.Cookie("jwt-session")
		if err != nil {
			util.MakeMessage(w, http.StatusForbidden, sec.ForbiddenResponse())
			return
		}
		log.Printf("Autentication by Cookie")
		tokenString = token.Value
	} else {
		split := strings.Split(value, " ")
		log.Printf("Autentication by Header Authorization: Bearer")
		tokenString = strings.TrimSpace(split[1])
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret, err := sec.GetJwtSecret()
		if err != nil {
			return nil, fmt.Errorf("error getting JWT secret")
		}
		return secret, nil
	})
	if err != nil {
		log.Println("AuthValidator: ", err)
		util.MakeMessage(w, http.StatusForbidden, sec.ForbiddenResponse())
		return
	}
	data := sha256.Sum256([]byte(tokenString))
	encodedToken := base64.StdEncoding.EncodeToString(data[:])
	/* Este tipo de teste acarretará impacto significativo no desempenho
	revoked, err := tokenWasRevoked(encodedToken)
	if err != nil {
		util.MakeMessage(w, http.StatusInternalServerError, util.UnexpectedError())
		return
	}
	if revoked {
		log.Printf("Revoked token: %s", encodedToken)
		util.MakeMessage(w, http.StatusUnauthorized, sec.TokenRevoked())
		return
	}
	*/
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		unix := claims["exp"].(float64)
		expires := time.Unix(int64(unix), 0)
		if expires.Before(time.Now()) {
			util.MakeMessage(w, http.StatusForbidden, sec.ForbiddenResponse())
			return
		}
		value := claims["sub"]
		userUuid := fmt.Sprintf("%v", value)

		background := context.Background()
		tenantUuid := redis.RedisConnection.Get(background, "users::"+userUuid+"::session::"+encodedToken)
		if tenantUuid.Val() == "" {
			log.Println("Not logged")
			util.MakeMessage(w, http.StatusForbidden, sec.ForbiddenResponse())
			return
		}

		authInfo := sec.AuthenticationInfo{
			UserUuid:   userUuid,
			TenantUuid: tenantUuid.Val(),
			SessionId:  encodedToken,
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, sec.SESSION_KEY, authInfo)
		ro := r.WithContext(ctx)
		next.ServeHTTP(w, ro)
		return
	}
	next.ServeHTTP(w, r)
}
