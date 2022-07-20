package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Neumann88/currency-grpc/api"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/PuerkitoBio/goquery"
)

type CurrencyHander struct {
	api.UnimplementedCurrencyServiceServer
}

func NewCurrencyHandler() *CurrencyHander {
	return &CurrencyHander{}
}

func (a *CurrencyHander) GetCurrency(ctx context.Context, _ *emptypb.Empty) (*api.CurrencyResponse, error) {
	res, err := http.Get("https://markets.businessinsider.com/currencies")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var records []*api.Currency
	// loop over rows
	doc.Find(".row-hover").Each(func(i int, s *goquery.Selection) {
		children := s.Children()

		// select
		name := children.First()
		ctr := name.Next()
		dsc := ctr.Next()
		chP := dsc.Next().Next()
		prc := chP.Next()
		last := prc.Next()

		// filter
		name = name.Children().First()
		dsc = dsc.Children().First()
		chP = chP.Children().First()

		lastTemp := last.Children().Last()

		if lastTemp.Text() != "" {
			last = lastTemp
		}

		n := strings.TrimSpace(name.Text())
		c := strings.TrimSpace(ctr.Text())
		d := strings.TrimSpace(dsc.Text())
		ch := strings.TrimSpace(chP.Text())
		p := strings.TrimSpace(prc.Text())
		l := strings.TrimSpace(last.Text())

		r := &api.Currency{
			Name:        n,
			Country:     c,
			Description: d,
			ChangeP:     ch,
			Rate:        p,
			LastUpdate:  l,
		}

		records = append(records, r)
	})

	return &api.CurrencyResponse{
		Date: time.Now().Format("2006-01-02 15:04:05"),
		Data: records,
	}, nil
}
