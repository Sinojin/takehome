package requestManager

import (
	"crypto/md5"
	"sync"
	"takehome/src/domain"
	"takehome/src/services/exporter"
	"takehome/src/services/requester"
	"testing"
	"time"
)

//this implementations is mocking struct to apply unit test
type MockServiceWorker struct {
	Count          int
	AppliedByOrder []string
	mu             *sync.Mutex
}

func (m *MockServiceWorker) Get(URL domain.Address) ([16]byte, error) {

	if URL.String() == "http://first.com" {
		time.Sleep(1 * time.Second)
		// return md5.Sum([]byte("this is test mocking")), nil
	} else if URL.String() == "http://second.com" {
		time.Sleep(500 * time.Millisecond)
	}
	//we lock function because in async task it can throw fatal.
	m.mu.Lock()
	m.Count++
	m.AppliedByOrder = append(m.AppliedByOrder, URL.String())
	m.mu.Unlock()
	return md5.Sum([]byte("this is test mocking")), nil
}

func (m *MockServiceWorker) checkOrder(list []string) bool {
	for index, element := range list {
		if element != m.AppliedByOrder[index] {
			return false
		}
	}
	return true
}
func (m *MockServiceWorker) checkAppledTask(NumberOfTask int) bool {
	return NumberOfTask == m.Count
}

func newMockServiceWorker() *MockServiceWorker {
	return &MockServiceWorker{mu: &sync.Mutex{}, AppliedByOrder: make([]string, 0), Count: 0}
}

// func newExporterMock()
type mockExporter struct {
}

func (m *mockExporter) Print(domain string, md5Bytes [16]byte) error {
	return nil
}

func Test_service_Start(t *testing.T) {
	rawDomains := []string{"http://first.com", "http://second.com", "http://first.com"}
	asyncOrder := []string{"http://second.com", "http://first.com", "http://first.com"}

	DomainList := domain.NewAddressList(rawDomains)
	type fields struct {
		worker   requester.Service
		exporter exporter.Service
	}
	type args struct {
		WorkerNum  int
		DomainList []domain.Address
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedOrder []string
	}{
		{name: "Request manager worked sync", fields: fields{worker: newMockServiceWorker(), exporter: &mockExporter{}}, args: args{WorkerNum: 1, DomainList: DomainList}, expectedOrder: rawDomains},
		{name: "Request manager worked async", fields: fields{worker: newMockServiceWorker(), exporter: &mockExporter{}}, args: args{WorkerNum: len(DomainList), DomainList: DomainList}, expectedOrder: asyncOrder},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				worker:   tt.fields.worker,
				exporter: tt.fields.exporter,
			}
			s.Start(tt.args.WorkerNum, tt.args.DomainList)
			a := tt.fields.worker.(*MockServiceWorker)
			if !a.checkOrder(tt.expectedOrder) {
				t.Errorf("Task doesn't work properly, order is wrong Expected: %v, Actual: %v", tt.expectedOrder, a.AppliedByOrder)
			}
		})
	}
}
