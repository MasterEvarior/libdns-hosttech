// Package hosttech implements methods for manipulating Hosttech.ch DNS records with the libdns interfaces.
// Manipulation is achieved with the Hosttech API at https://api.ns1.hosttech.eu/api/documentation/#/.
package hosttech

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/libdns/libdns"
)

// Provider facilitates DNS record manipulation with Hosttech.ch.
type Provider struct {
	APIToken string `json:"api_token,omitempty"`
}

// The URL for the Hosttech API connection
const apiHost = "https://api.ns1.hosttech.eu/api/user/v1"

// GetRecords lists all the records in the zone.
func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	reqURL := fmt.Sprintf("%s/zones/%s/records", apiHost, zone)

	returnValue, err := p.makeApiCall(ctx, http.MethodGet, reqURL, nil)

	//If there's an error return an empty slice
	if err != nil {
		return []libdns.Record{}, err
	}

	return returnValue, nil
}

// AppendRecords adds records to the zone. It returns all records that were added.
// If an error occurs while records are being added, the already successfully added records will be returned along with an error.
func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	reqURL := fmt.Sprintf("%s/zones/%s/records", apiHost, zone)

	successfullyAppendedRecords := []libdns.Record{}
	for _, record := range records {

		hosttechRecord, err := LibdnsRecordToHosttechRecordWrapper(record)
		if err != nil {
			return nil, err
		}

		bodyBytes, err := json.Marshal(hosttechRecord)
		if err != nil {
			return nil, err
		}

		resp, err := p.makeApiCall(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
		if err != nil {
			return successfullyAppendedRecords, err
		}

		successfullyAppendedRecords = append(successfullyAppendedRecords, resp...)
	}

	return successfullyAppendedRecords, nil
}

// SetRecords sets the records in the zone, either by updating existing records or creating new ones.
// It returns the updated records.
// TODO
func (p *Provider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	successfullyUpdatedRecords := []libdns.Record{}
	for _, record := range records {

		bodyBytes, err := json.Marshal(record)

		if err != nil {
			return nil, err
		}

		reqURL := fmt.Sprintf("%s/zones/%s/records/%s", apiHost, zone, record.ID)
		resp, err := p.makeApiCall(ctx, http.MethodPut, reqURL, bytes.NewReader(bodyBytes))

		if err != nil {
			return successfullyUpdatedRecords, err
		}

		successfullyUpdatedRecords = append(successfullyUpdatedRecords, resp...)
	}

	return successfullyUpdatedRecords, nil
}

// DeleteRecords deletes the records from the zone. It returns the records that were deleted.
// If an error occurs while records are being deleted, the already successfully deleted records will be returned along with an error.
func (p *Provider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	successfullyDeletedRecords := []libdns.Record{}
	for _, record := range records {
		reqUrl := fmt.Sprintf("%s/zones/%s/records/%s", apiHost, zone, record.ID)
		_, err := p.makeApiCall(ctx, http.MethodDelete, reqUrl, nil)

		if err != nil {
			return successfullyDeletedRecords, err
		}

		successfullyDeletedRecords = append(successfullyDeletedRecords, record)
	}

	return successfullyDeletedRecords, nil
}

func (p *Provider) makeApiCall(ctx context.Context, httpMethod string, reqUrl string, body io.Reader) ([]libdns.Record, error) {
	req, err := http.NewRequestWithContext(ctx, httpMethod, reqUrl, body)
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

	if resp.StatusCode < 200 && 300 >= resp.StatusCode {
		return []libdns.Record{}, fmt.Errorf("call to API was not successful, returned the status code '%s'", resp.Status)
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
