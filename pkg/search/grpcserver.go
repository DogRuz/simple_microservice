package search

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"simple_microservice/pkg/api"
	"strings"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.GetInformationRequest) (*api.GetInformationResponse, error) {
	res, _ := http.Get(fmt.Sprintf(URL, req.GetInn()))
	body, _ := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile(`Попробуйте изменить поисковой запрос`)
	oops := search(string(body), re)
	if oops == "Попробуйте изменить поисковой запрос" {
		return &api.GetInformationResponse{}, nil
	}
	re = regexp.MustCompile(`ИНН \d+ ?`)
	inn := strings.Replace(search(string(body), re), "ИНН ", "", 1)
	re = regexp.MustCompile(`КПП \d+ ?`)
	kpp := strings.Replace(search(string(body), re), "КПП ", "", 1)
	re = regexp.MustCompile(`[А-Я][а-я]+ ? [А-Я][а-я]* [А-Я][а-я]*`)
	fio := search(string(body), re)
	re = regexp.MustCompile(`&quot;.+&quot; `)
	company := strings.Replace(search(string(body), re), "&quot;", "", 2)
	return &api.GetInformationResponse{Inn: inn, Kpp: kpp, Fio: fio, Company: company}, nil
}
