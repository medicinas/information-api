package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/medicinas/information-api/api/handler"
	"github.com/medicinas/information-api/config"
	"github.com/medicinas/information-api/pkg/country"
	"github.com/medicinas/information-api/pkg/metric"
	"github.com/medicinas/information-api/pkg/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	Dbdriver := "postgres"
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.DbHost, config.DbPort, config.DbUser, config.DbName, config.DbPassword)
	db, err := gorm.Open("post", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	countryRepository := country.NewRepository(db)
	countryUseCase := country.NewUseCase(countryRepository)

	r := mux.NewRouter()

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)

	handler.MakeCountryHandlers(r, n, countryUseCase)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	svr := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
