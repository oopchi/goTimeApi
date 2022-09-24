package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type TimeTz struct {
	CurrentTime string `json:"current_time" xml:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tz")
	tzs := strings.Split(q, ",")
	res := make(map[string]string)

	tnow := time.Now()

	encoder := json.NewEncoder(w)

	if len(tzs) <= 1 {
		loc, err := time.LoadLocation(q)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "invalid timezone: ", q)

			return
		}

		res["current_time"] = tnow.In(loc).String()

		w.Header().Add("Content-Type", "application/json")
		encoder.Encode(res)

		return
	}

	for _, tz := range tzs {
		loc, err := time.LoadLocation(tz)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "invalid timezone", tz)

			return
		}

		res[tz] = tnow.In(loc).String()
	}

	w.Header().Add("Content-Type", "application/json")
	encoder.Encode(res)
}
