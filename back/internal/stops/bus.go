package stops

import (
	"fmt"
	"math/rand"

	"github.com/Bruoche/messagerie-tr/pkg/request"
	"github.com/gin-gonic/gin"
)

const MinsPerHours = 60

type minutes int

func At(hour int, mins minutes) minutes {
	return minutes(hour*MinsPerHours) + mins
}

func (mins minutes) ToString() string {
	hours := mins / MinsPerHours
	format := "%dh%02d"
	days := hours / 24
	if days > 0 {
		format = fmt.Sprintf("%dj", days) + "%02dh%02d"
	}
	return fmt.Sprintf(format, hours%24, mins%MinsPerHours)
}

type Stop struct {
	ID     string
	Name   string
	Passes []minutes
}

func makeStop(id int, name string, passes []minutes) Stop {
	return Stop{
		ID:     fmt.Sprintf("S31%03d", id),
		Name:   name,
		Passes: passes,
	}
}

func Stops(names ...string) []Stop {
	stops := make([]Stop, len(names))
	for id, name := range names {
		start := At(7, minutes(rand.Intn(150)))
		nbPasses := 10 + rand.Intn(10)
		frequency := 5 + rand.Intn(40)
		passes := make([]minutes, nbPasses)
		i := 0
		for i < nbPasses {
			passes[i] = start + minutes(frequency*i)
			i++
		}
		stops[id] = makeStop(id+1, name, passes)
	}
	return stops
}

var stops = Stops(
	"Plantanes",
	"Glamancourt",
	"Morneaux",
	"Pormes Carroussel",
	"Boiteaux Blanc",
	"Glannes",
	"Gémistres",
	"Gosansseaux",
	"Poissac",
	"Clamoreaux",
	"Clomansourt",
	"Pluituis",
	"Flip",
	"Glaps",
	"Clommes",
)

var _ request.Loadable = &Stop{}

func (q *Stop) Load(c *gin.Context) error {
	id, err := request.RequireParam(c, "id")
	if err != nil {
		return err
	}
	for _, stop := range stops {
		if stop.ID == id {
			*q = stop
			return nil
		}
	}
	return &request.ErrorResponse{
		Code: 404,
		Err:  fmt.Sprintf("L'arrêt \"%s\" n'a pas été trouvé.", id),
	}
}

type StopResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func All(*request.NoParam, *request.NoParam) ([]StopResponse, error) {
	response := make([]StopResponse, len(stops))
	for i, stop := range stops {
		response[i] = StopResponse{
			ID:   stop.ID,
			Name: stop.Name,
		}
	}
	return response, nil
}

func ListHoursAt(_ *request.NoParam, stop *Stop) ([]string, error) {
	passes := make([]string, len(stop.Passes))
	for i, pass := range stop.Passes {
		passes[i] = pass.ToString()
	}
	return passes, nil
}
