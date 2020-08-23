package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/medicinas/information-api/pkg/country"
	"github.com/medicinas/information-api/pkg/entity"
	"github.com/prometheus/common/log"
	"net/http"
)

func findAllCountries(service country.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func addCountry(service country.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error trying to add a country"
		var c *entity.Country
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Infoln("Error trying to deserialize the body from the add country operation")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.Save(c)
		if err != nil {
			log.Infoln("Error trying to save from the add country operation")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		c.ID = id
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(c); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeBookmarkHandlers make url handlers
func MakeCountryHandlers(r *mux.Router, n negroni.Negroni, service country.UseCase) {
	r.Handle("/v1/countries", n.With(
		negroni.Wrap(findAllCountries(service)),
	)).Methods("GET", "OPTIONS").Name("findAllCountries")

	r.Handle("/v1/countries", n.With(
		negroni.Wrap(addCountry(service)),
	)).Methods("POST", "OPTIONS").Name("addCountry")

	/*r.Handle("/v1/countries", n.With(
		negroni.Wrap(updateCountry(service)),
	)).Methods("PUT", "OPTIONS").Name("updateCountry")

	r.Handle("/v1/countries/{id]", n.With(
		negroni.Wrap(findCountry(service)),
	)).Methods("GET", "OPTIONS").Name("findCountry")

	r.Handle("/v1/countries/{id]", n.With(
		negroni.Wrap(deleteCountry(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteCountry")*/
}
