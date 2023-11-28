package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"franchises-system/internal/constants"
	"franchises-system/internal/domain/entity"
	"franchises-system/internal/domain/ports/postgres/interfaces"
	"franchises-system/internal/utils/http"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type WebInfo interface {
	GetWebInfo(ctx context.Context, url string) (entity.WebInfo, error)
}

type webInfoApp struct {
	http http.HttpClient
	repo interfaces.Repository
}

func NewWebInfoApp(http http.HttpClient, repo interfaces.Repository) WebInfo {
	return &webInfoApp{
		http,
		repo,
	}
}

func (app *webInfoApp) GetWebInfo(ctx context.Context, url string) (entity.WebInfo, error) {
	ssllabs := entity.SSLLabsResponse{}
	result := whoisparser.WhoisInfo{}
	for {
		resp, err := app.http.Get(ctx, fmt.Sprintf(constants.Ssllabs, url))
		if err != nil {
			return entity.WebInfo{}, err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return entity.WebInfo{}, err
		}

		if err = json.Unmarshal(body, &ssllabs); err != nil {
			return entity.WebInfo{}, err
		}

		if ssllabs.Status != "READY" {
			continue
		} else {
			break
		}
	}

	for {
		raw, err := whois.Whois(url)
		if err != nil {
			continue
		}

		result, err = whoisparser.Parse(raw)
		if err != nil {
			return entity.WebInfo{}, err
		}
		break
	}

	return entity.WebInfo{
		Protocol:    ssllabs.Protocol,
		TraceRoutes: ssllabs.Endpoints.GetServerNames(),
		Domain: entity.Domain{
			CreatedAt:       *result.Domain.CreatedDateInTime,
			ExpiredAt:       *result.Domain.ExpirationDateInTime,
			RegistrantName:  result.Registrant.Name,
			RegistrantEmail: result.Registrant.Email,
		},
	}, nil

}
