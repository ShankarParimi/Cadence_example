package worker

import (
	"OnBoardingPOC/internal/common"
	"OnBoardingPOC/internal/workflow/admin"
	"go.uber.org/cadence/worker"
)

func StartWorkers(h *common.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, admin.ApplicationName, workerOptions)
}
