package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nitkumar91296/banking-auth/domain"
	"github.com/nitkumar91296/banking-auth/dto"
	"github.com/nitkumar91296/banking-auth/service"
)

type AuthHandler struct {
	repo    domain.AuthRepository
	service service.AuthService
}

func (h AuthHandler) NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler not implemented!!!")
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error while decoding login request: ", err.Error())
		w.WriteHeader(http.StatusBadGateway)
	} else {
		token, err := h.service.Login(loginRequest)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprint(w, *token)
		}
	}

}
