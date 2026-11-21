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
const EnvPort = "TISSEO_BACK_PORT"

func main() {
	serv, ok := getEnv(EnvServer)
	if !ok {
		log.Info("Aucun nom de serveur défini.")
	} else {
		log.Info("Running on %s.", serv)
	}
	port, ok := getEnv(EnvPort)
	if !ok {
		log.Debug("No port given at environnement variable %s", EnvPort)
		port = "8080"
	}
	request.Serve("localhost:"+port, false,
		request.MakePath(ext.GET, "/", request.String(request.Index)),
		request.MakePath(ext.GET, "/server", request.String(GetServer)),
		request.MakePath(ext.GET, "/stops/", request.Json(stops.All)),
		request.MakePath(ext.GET, "/stop/:id", request.Json(stops.ListHoursAt)),
	)
}

func GetServer(*request.NoParam, *request.NoParam) (string, error) {
	server, ok := getEnv(EnvServer)
	if !ok {
		return "", &request.ErrorResponse{
			Code: 404,
			Err:  fmt.Sprintf("Nom du serveur non défini. Veuillez définir la variable d'environnement \"%s\".", EnvServer),
		}
	}
	return server, nil
}

func getEnv(key string) (string, bool) {
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if pair[0] == key {
			return pair[1], true
		}
	}
	return "", false
}
