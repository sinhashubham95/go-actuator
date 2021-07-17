package actuator

import (
	"fmt"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/flags"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/valyala/fasthttp"
	"log"
)

// ListenAndServe is used to bring the actuator endpoints up and running in the parent's goroutine itself
func ListenAndServe(config *models.Config) {
	serve()
}

// Serve is used to bring the actuator endpoints up and running.
func Serve(config *models.Config) {
	go serve()
}

func serve() {
	err := fasthttp.ListenAndServe(fmt.Sprintf("%s%d", commons.Colon, flags.Port()), router)
	if err != nil {
		// note that actuator is just a feature which is good to have for most applications
		// so even if the application does not start we won't crash the application
		log.Printf("Error starting the actuator endpoints %v.", err)
	}
}
