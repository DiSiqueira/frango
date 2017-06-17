package main

import (
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/trace"
	"github.com/disiqueira/frango/src/api/proto/search"
)

func main() {
	e := newEnv()
	http.Handle("/", requestHandler(e))
	log.Fatal(http.ListenAndServe(e.serviceAddr(), nil))
}

func requestHandler(e *env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		span := e.Tracer.SpanFromRequest(r)
		defer span.Finish()
		ctx := trace.NewContext(r.Context(), span)

		asaList, err := e.SearchClient.Asa(ctx, &search.AsaFilter{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(
			geoJSONResponse(asaList),
		)
	})
}

type response struct {
	Type    string          `json:"type"`
	AsaList *search.AsaList `json:"features"`
}

func geoJSONResponse(asaList *search.AsaList) response {
	return response{
		Type:    "AsaCollection",
		AsaList: asaList,
	}
}
