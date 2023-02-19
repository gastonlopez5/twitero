package bd

import (
	"githut.com/gastonlopez5/twitero/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passBytes := []byte(password)
	passDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passDB, passBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
