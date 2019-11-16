package boot

import (
	"OnBoardingPOC/internal/common"
	"OnBoardingPOC/internal/controller/admin"
	"OnBoardingPOC/internal/controller/user"
	"OnBoardingPOC/internal/worker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var h common.SampleHelper

func Init(mode string) {
	if mode == "worker" {
		h := common.GetCommonHelper()
		h.SetupServiceConfig()
		worker.StartWorkers(h)
		// The workers are supposed to be long running process that should not exit.
		// Use select{} to block indefinitely for samples, you can quit by CMD+C.
		select {}
	} else {
		router := mux.NewRouter().StrictSlash(true)
		api := router.PathPrefix("/api/v1").Subrouter()
		//define routes for admin and User
		api.HandleFunc("/admin", admin.Controller)
		api.HandleFunc("/User", user.Controller)
		log.Fatal(http.ListenAndServe(":8082", router))
		//Initialise the Worker For Cadence
	}
}
