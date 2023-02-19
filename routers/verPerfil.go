package routers

import (
	"encoding/json"
	"net/http"

	"githut.com/gastonlopez5/twitero/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)
}
