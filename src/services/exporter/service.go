package exporter

import "fmt"

//This structer basiclly prints datas CLI
//This service gives a lot of opportunity to export somewhere
//for example you can implement interface for file export, json export etc.
//and with this design you can unit test more easily.
//Todo:I did not implement unit test for this struct because it is so basic.
type Service interface {
	Print(domain string, md5Bytes [16]byte) error
}

type service struct {
}

func (s *service) Print(domain string, md5Bytes [16]byte) error {
	fmt.Printf("%v - %x \n", domain, md5Bytes)
	return nil
}

func NewCliExporter() Service {
	return &service{}
}
