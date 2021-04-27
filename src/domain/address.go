package domain

import (
	"errors"
	"fmt"
	"log"

	"strings"

	"net/url"
)

type Address struct {
	u *url.URL
}

func (a *Address) String() string {
	return a.u.String()
}
func NewAddress(rawurl string) (Address, error) {
	//check that scheme is added or not
	//todo:It should be check domain name with regex but it is not concern of project for now.
	//I assume that all domains are correct.
	if rawurl == "" {
		return Address{}, errors.New("Raw url can not be empty!")
	}
	if !strings.HasPrefix(rawurl, "http") {
		rawurl = fmt.Sprintf("http://%v", rawurl)
	}

	u, err := url.Parse(rawurl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return Address{}, errors.New("Invalid URL")

	}

	return Address{u: u}, nil
}

func NewAddressList(domains []string) []Address {
	alist := make([]Address, 0)
	for _, domain := range domains {
		addr, err := NewAddress(domain)
		if err != nil {
			log.Println(err)
			continue
		}
		alist = append(alist, addr)
	}
	return alist
}
