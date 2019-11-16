package user

import (
	"OnBoardingPOC/internal/common"
	"OnBoardingPOC/internal/contracts"
	"OnBoardingPOC/internal/workflow/user"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	h := common.GetCommonHelper()
	switch r.Method {
	case "GET":
		response, err := user.StartWorkflow(context.Background(), h)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Something Wrong!"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	case "POST":
		body, er := ioutil.ReadAll(r.Body)
		if er != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		var request contracts.Request
		if !json.Valid(body) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		inputJson, err := json.Marshal(request.Data)
		response, err1 := user.Process(request.WorkflowId, string(inputJson), h)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
