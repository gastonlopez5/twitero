package routers

import (
	"encoding/json"
	"net/http"

	"githut.com/gastonlopez5/twitero/bd"
	"githut.com/gastonlopez5/twitero/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se actuaizó el registro correctamente", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}
