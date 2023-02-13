package main

import (
	"encoding/json"
	"fmt"
	"github.com/libdns/libdns"
)

type HosttechResponseWrapper struct {
	Data []HosttechRecordWrapper `json:"data"`
}

type HosttechRecordWrapper struct {
	value HosttechRecord
}

func (h HosttechRecordWrapper) toLibdnsRecord() libdns.Record {
	return h.value.toLibdnsRecord()
}

func (h HosttechRecordWrapper) fromLibdnsRecord(record libdns.Record) {
	h.value.fromLibdnsRecord(record)
}

func (h *HosttechRecordWrapper) UnmarshalJSON(b []byte) error {
	var base Base
	err := json.Unmarshal(b, &base)
	if err != nil {
		return err
	}
	switch base.Type {
	case "AAAA":
		record := AAAARecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "A":
		record := ARecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "NS":
		record := NSRecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "CNAME":
		record := CNAMERecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "MX":
		record := MXRecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "TXT":
		record := TXTRecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "TLSA":
		record := TLSARecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	default:
		err = fmt.Errorf(`record type "%s" is not supported"`, base.Type)
	}
	if err != nil {
		return err
	}

	return nil
}
