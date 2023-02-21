package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"githut.com/gastonlopez5/twitero/bd"
)

func LeoTwwets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", 400)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", 400)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", 400)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoTwwets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los tweets", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
