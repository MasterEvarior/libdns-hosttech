package main

import (
	"encoding/json"
	"fmt"
	"github.com/libdns/libdns"
	"strconv"
	"time"
)

type HosttechRecord interface {
	toLibdnsRecord() libdns.Record
	//fromLibdnsRecord(record libdns.Record)
}

type HosttechRecordWrapper struct {
	value HosttechRecord
}

func (h HosttechRecordWrapper) toLibdnsRecord() libdns.Record {
	return h.value.toLibdnsRecord()
}

//func (h HosttechRecordWrapper) fromLibdnsRecord(record libdns.Record) {
//	h.value.fromLibdnsRecord(record)
//}

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
	case "CAA":
		record := CAARecord{}
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
	case "PTR":
		record := PTRRecord{}
		err = json.Unmarshal(b, &record)
		h.value = HosttechRecord(record)
	case "SRV":
		record := SRVRecord{}
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
		fmt.Errorf(`record type "%s" is not supported"`, base.Type)
	}
	if err != nil {
		return err
	}

	return nil
}

type Base struct {
	Id      int    `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	TTL     int    `json:"ttl,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type AAAARecord struct {
	Base
	Name string `json:"name,omitempty"`
	IPV6 string `json:"ipv6,omitempty"`
}

func (a AAAARecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(a.Id),
		Type:  a.Type,
		Name:  a.Name,
		Value: a.IPV6,
		TTL:   time.Duration(a.TTL),
	}
}

type ARecord struct {
	Base
	Name string `json:"name,omitempty"`
	IPV4 string `json:"ipv4,omitempty"`
}

func (a ARecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(a.Id),
		Type:  a.Type,
		Name:  a.Name,
		Value: a.IPV4,
		TTL:   time.Duration(a.TTL),
	}
}

type CAARecord struct {
	Base
	Name  string `json:"name,omitempty"`
	Flag  string `json:"flag,omitempty"`
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

func (c CAARecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(c.Id),
		Type:  c.Type,
		Name:  c.Name,
		Value: c.Value,
		TTL:   time.Duration(c.TTL),
	}
}

type CNAMERecord struct {
	Base
	Name  string `json:"name,omitempty"`
	Cname string `json:"cname,omitempty"`
}

func (c CNAMERecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(c.Id),
		Type:  c.Type,
		Name:  c.Name,
		Value: c.Cname,
		TTL:   time.Duration(c.TTL),
	}
}

type MXRecord struct {
	Base
	Name      string `json:"name,omitempty"`
	OwnerName string `json:"ownername,omitempty"`
	Pref      int    `json:"pref,omitempty"`
}

func (m MXRecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:       strconv.Itoa(m.Id),
		Type:     m.Type,
		Name:     m.OwnerName,
		Value:    m.Name,
		TTL:      time.Duration(m.TTL),
		Priority: m.Pref,
	}
}

type NSRecord struct {
	Base
	OwnerName  string `json:"ownername,omitempty"`
	TargetName string `json:"targetname,omitempty"`
}

func (n NSRecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(n.Id),
		Type:  n.Type,
		Name:  n.OwnerName,
		Value: n.TargetName,
		TTL:   time.Duration(n.TTL),
	}
}

type PTRRecord struct {
	Base
	Name   string `json:"name,omitempty"`
	Origin string `json:"origin,omitempty"`
}

func (p PTRRecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:       strconv.Itoa(p.Id),
		Type:     p.Type,
		Name:     p.Name,
		Value:    p.Origin,
		TTL:      time.Duration(p.TTL),
		Priority: 0,
	}
}

type SRVRecord struct {
	Base
	Service  string `json:"service,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Port     int    `json:"port,omitempty"`
	Target   string `json:"target,omitempty"`
}

func (s SRVRecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:       strconv.Itoa(s.Id),
		Type:     s.Type,
		Name:     s.Service,
		Value:    s.Target,
		TTL:      time.Duration(s.TTL),
		Priority: s.Priority,
	}
}

type TXTRecord struct {
	Base
	Name string `json:"name,omitempty"`
	Text string `json:"text,omitempty"`
}

func (t TXTRecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(t.Id),
		Type:  t.Type,
		Name:  t.Name,
		Value: t.Text,
		TTL:   time.Duration(t.TTL),
	}
}

type TLSARecord struct {
	Base
	Name string `json:"name,omitempty"`
	Text string `json:"text,omitempty"`
}

func (t TLSARecord) toLibdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(t.Id),
		Type:  t.Type,
		Name:  t.Name,
		Value: t.Text,
		TTL:   time.Duration(t.TTL),
	}
}

type HosttechResponse struct {
	Data []HosttechRecordWrapper `json:"data"`
}
