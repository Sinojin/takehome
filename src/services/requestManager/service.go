package requestManager

import (
	"log"
	"sync"
	"takehome/src/domain"
	"takehome/src/services/exporter"
	"takehome/src/services/requester"
)

type Service interface {
	Start(WorkerNum int, DomainList []domain.Address)
}

type service struct {
	worker   requester.Service
	exporter exporter.Service
}

func (s *service) Start(WorkerNum int, DomainList []domain.Address) {
	//Check worker number for tasks
	//if worker number bigger than task number, change worker number.
	taskNum := len(DomainList)
	if WorkerNum > taskNum {
		WorkerNum = taskNum
	}
	//watcher controls workers
	watcher := make(chan struct{}, WorkerNum)

	//Wait until all tasks are finished
	var wg sync.WaitGroup
	wg.Add(taskNum)

	for i := 0; i < taskNum; i++ {
		//after the limit of channel will block for loop
		//and it will wait workers to finish their jobs
		watcher <- struct{}{}
		//Go function GO ! :)
		go func(url domain.Address) {
			//in any case this function must work except fatal error
			//in this logic, I am thinking best case everytime.
			//todo: it can be more reliable.
			defer func() {
				//one of waiting group is done.
				wg.Done()
				//this take one empty struct from watcher pipeline
				<-watcher
			}()
			md5bytes, err := s.worker.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			s.exporter.Print(url.String(), md5bytes)
		}(DomainList[i])
	}
	wg.Wait()
}

func NewRequestManager(worker requester.Service, exporter exporter.Service) Service {
	return &service{
		worker:   worker,
		exporter: exporter,
	}
}
