package handlers

import (
	"encoding/json"
	"github.com/octalope/stat-service/modules/stats"
	"github.com/rs/zerolog"
	"net/http"
)

type SampleDataBody struct {
	Data [][]float64 `json:"data"`
	XCol int         `json:"x_col"`
	YCol int         `json:"y_col"`
}

type RegressionResult struct {
	M        float64 `json:"m"`
	Dm       float64 `json:"dm"`
	B        float64 `json:"b"`
	Db       float64 `json:"db"`
	RSquared float64 `json:"rSquared"`
}

func LsfHandler(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	if r.Method == http.MethodPost {

		var data SampleDataBody
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Error().Err(err).Msg("Error decoding JSON request body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Interface("request_body", SampleDataBody(data))
		})

		var m, dm, b, db, rSquared float64 = stats.LeastSquaresFit(data.Data, data.XCol, data.YCol)

		var result RegressionResult = RegressionResult{
			M:        m,
			Dm:       dm,
			B:        b,
			Db:       db,
			RSquared: rSquared,
		}

		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Interface("response_body", RegressionResult(result))
		})

		// Encode the RegressionResult struct into JSON and write it to the response body.
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			// If there's an error encoding the JSON, log it and send an internal server error.
			log.Error().Err(err).Msg("Error encoding JSON response")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json.
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
