package main

import (
	"flag"
	"fmt"
	"takehome/src/domain"
	"takehome/src/services/exporter"
	"takehome/src/services/requestManager"
	"takehome/src/services/requester"
)

var DefaultWorkerNum = 10

func main() {
	//Getting worker number and domains
	workerNum := flag.Int("parallel", 10, "Number of parallel jobs to run")
	flag.Parse()
	DomainList := flag.Args()
	//check worker num
	if *workerNum <= 0 {
		*workerNum = 1
	}
	//check domain list
	if len(DomainList) < 1 {
		fmt.Println("There is no domain please check your input !!")
		return
	}

	AddressesList := domain.NewAddressList(DomainList)
	//Services
	ExportService := exporter.NewCliExporter()
	WorkerService := requester.NewRequesterService()
	RequestManager := requestManager.NewRequestManager(WorkerService, ExportService)

	RequestManager.Start(*workerNum, AddressesList)
}
