package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/octalope/stat-service/modules/stats"
	"github.com/octalope/stat-service/modules/util"
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
	log := util.Log()
	log.Info().
		Str("verb", r.Method).
		Str("path", "/lsf").
		Msg("request")

	if r.Method == http.MethodPost {

		var data SampleDataBody
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process the 'data' struct
		fmt.Printf("Received data: data=%v, x_col=%d, y_col=%d\n", data.Data, data.XCol, data.YCol)

		var m, dm, b, db, rSquared float64 = stats.LeastSquaresFit(data.Data, data.XCol, data.YCol)

		var result RegressionResult = RegressionResult{
			M:        m,
			Dm:       dm,
			B:        b,
			Db:       db,
			RSquared: rSquared,
		}

		// Encode the RegressionResult struct into JSON and write it to the response body.
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			// If there's an error encoding the JSON, log it and send an internal server error.
			// log.Printf("Error encoding JSON response: %v", err)
			log.Info().
				Str("verb", r.Method).
				Str("path", "/health").
				Msg("request")

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
