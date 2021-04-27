package requester

import (
	"crypto/md5"
	"io/ioutil"
	"net/http"
	"takehome/src/domain"
)

//This is Request Service
//Basically doing same job as http.get()
//We could use directly http package but if we decide to change request system, I need to change every http function in my business logic
//thats why sperated from business logic

type Service interface {
	//returns md5 of body and error
	Get(URL domain.Address) ([16]byte, error)
}

type service struct {
}

func (s *service) Get(URL domain.Address) ([16]byte, error) {
	//there is no specific request rule that's why it uses default http transport
	res, err := http.Get(URL.String())
	if err != nil {
		return [16]byte{}, err
	}
	defer res.Body.Close()
	//We Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return [16]byte{}, err
	}

	return md5.Sum(body), nil
}

func NewRequesterService() Service {
	return &service{}
}
