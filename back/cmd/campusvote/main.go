package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Bruoche/messagerie-tr/pkg/ext"
	"github.com/Bruoche/messagerie-tr/pkg/log"
	"github.com/Bruoche/messagerie-tr/pkg/request"
	"github.com/Bruoche/mvp-tisseo/back/internal/stops"
)

const EnvServer = "TISSEO_BACK_SERV_NAME"

func main() {
	serv, err := GetServer(&request.NoParam{}, &request.NoParam{})
	if err != nil {
		log.Info("Aucun nom de serveur défini.")
	} else {
		log.Info("Running on %s.", serv)
	}
	request.Serve("localhost:8080", false,
		request.MakePath(ext.GET, "/", request.String(request.Index)),
		request.MakePath(ext.GET, "/server", request.String(GetServer)),
		request.MakePath(ext.GET, "/stops/", request.Json(stops.All)),
		request.MakePath(ext.GET, "/stop/:id", request.Json(stops.ListHoursAt)),
	)
}

func GetServer(*request.NoParam, *request.NoParam) (string, error) {
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if pair[0] == EnvServer {
			return pair[1], nil
		}
	}
	return "", &request.ErrorResponse{
		Code: 404,
		Err:  fmt.Sprintf("Nom du serveur non défini. Veuillez définir la variable d'environnement \"%s\".", EnvServer),
	}
}
