package main

import(
	"backend/models"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/pascaldekloe/jwt"
	"fmt"
	"time"
	"errors"
)

var validUser = models.User{
	ID : 10,
	Email : "me@here.com",
	Password : "$2a$12$Afdyg2cgnztKWzrU4RpRK.AerRCy4Kt4Sdh18ds8zF//zh2ElDiAG",
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) SignIn(w http.ResponseWriter, r *http.Request){
	var creds Credentials
	
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}
	
	hashedPassword := validUser.Password
	
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}
	
	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}
	
	jwtBytes, errSign := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if errSign != nil {
		app.errorJSON(w, errors.New("Error Sign"))
		return
	}
	
	app.writeJSON(w, http.StatusOK,"response", jwtBytes)
	

}