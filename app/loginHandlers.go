package app

import (
	"encoding/json"
	"github.com/Darkhackit/banking-auth/dto"
	"github.com/Darkhackit/banking-auth/service"
	"io"
	"net/http"
	"time"
)

type loginHandlers struct {
	service service.LoginService
}

func (handler *loginHandlers) loginHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		user, appError := handler.service.Login(request)
		if appError != nil {
			writeResponse(w, http.StatusUnauthorized, appError.Message)
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    *user,
				Expires:  time.Now().Add(2 * time.Hour),
				HttpOnly: true,
				Secure:   true,
				Path:     "/",
			})
			writeResponse(w, http.StatusOK, user)
		}
	}
}

func (handler *loginHandlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		user, appError := handler.service.Register(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusOK, user)
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)
}

func writeResponse(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
