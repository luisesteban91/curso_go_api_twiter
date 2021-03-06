package banner

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/luisesteban91/curso_go_api_twiter/bd/usuario"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

/*UploadBanner ruta para subir banner*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + procesarToken.IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.Usuario
	var status bool
	user.Banner = procesarToken.IDUsuario + "." + extension
	status, err = usuario.EditarUsuario(user, procesarToken.IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la BD !"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
