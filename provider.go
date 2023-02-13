package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/libdns/libdns"
)

// TODO: Providers must not require additional provisioning steps by the callers; it
// should work simply by populating a struct and calling methods on it. If your DNS
// service requires long-lived state or some extra provisioning step, do it implicitly
// when methods are called; sync.Once can help with this, and/or you can use a
// sync.(RW)Mutex in your Provider struct to synchronize implicit provisioning.

// Provider facilitates DNS record manipulation with <TODO: PROVIDER NAME>.
type Provider struct {
	APIToken string `json:"api_token,omitempty"`
}

func (p *Provider) getApiHost() string {
	return "https://api.ns1.hosttech.eu/api/user/v1"
}

// GetRecords lists all the records in the zone.
func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	reqURL := fmt.Sprintf("%s/zones/%s/records", p.getApiHost(), zone)

	returnValue, err := p.makeApiCall(ctx, http.MethodGet, reqURL)

	//If there's an error return an empty slice
	if err != nil {
		return []libdns.Record{}, err
	}

	return returnValue, nil
}

// AppendRecords adds records to the zone. It returns the records that were added.
func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	reqURL := fmt.Sprintf("%s/zones/%s/records", p.getApiHost(), zone)

	returnValue, err := p.makeApiCall(ctx, http.MethodPost, reqURL)

	if err != nil {
		return []libdns.Record{}, err
	}

	return returnValue, nil
}

// SetRecords sets the records in the zone, either by updating existing records or creating new ones.
// It returns the updated records.
func (p *Provider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

// DeleteRecords deletes the records from the zone. It returns the records that were deleted.
func (p *Provider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

func (p *Provider) makeApiCall(ctx context.Context, httpMethod string, reqUrl string) ([]libdns.Record, error) {
	req, err := http.NewRequestWithContext(ctx, httpMethod, reqUrl, nil)
	req.Header.Set("Authorization", "Bearer "+p.APIToken)
	req.Header.Set("Content-Type", "application/json")

	//Return nil if there's an error
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	//Return an empty slice if there's an error
	if err != nil {
		return []libdns.Record{}, err
	}

	var parsedResponse HosttechResponseWrapper
	json.NewDecoder(resp.Body).Decode(&parsedResponse)

	var libdnsRecords []libdns.Record
	for _, record := range parsedResponse.Data {
		libdnsRecords = append(libdnsRecords, record.toLibdnsRecord())
	}

	return libdnsRecords, nil
}

// Interface guards
var (
	_ libdns.RecordGetter   = (*Provider)(nil)
	_ libdns.RecordAppender = (*Provider)(nil)
	_ libdns.RecordSetter   = (*Provider)(nil)
	_ libdns.RecordDeleter  = (*Provider)(nil)
)
