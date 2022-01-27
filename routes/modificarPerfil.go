package routes

import (
	"encoding/json"
	"net/http"

	"github.com/mxaxaxbx/twitter-clone/bd"
	"github.com/mxaxaxbx/twitter-clone/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el registro. Intente de nuevo "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado midificar el registro del usuario ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
