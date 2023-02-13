package main

import (
	"fmt"
	"github.com/libdns/libdns"
	"strconv"
	"time"
)

type HosttechRecord interface {
	toLibdnsRecord() libdns.Record
	fromLibdnsRecord(record libdns.Record)
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
		TTL:   time.Duration(a.TTL * 1000000000),
	}
}

func (a AAAARecord) fromLibdnsRecord(record libdns.Record) {
	a.Name = record.Name
	a.IPV6 = record.Value
	a.TTL = durationToIntSeconds(record.TTL)
	a.Comment = generateComment()
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
		TTL:   time.Duration(a.TTL * 1000000000),
	}
}

func (a ARecord) fromLibdnsRecord(record libdns.Record) {
	a.Name = record.Name
	a.IPV4 = record.Value
	a.TTL = durationToIntSeconds(record.TTL)
	a.Comment = generateComment()
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
		TTL:   time.Duration(c.TTL * 1000000000),
	}
}

func (c CNAMERecord) fromLibdnsRecord(record libdns.Record) {
	c.Name = record.Name
	c.Cname = record.Value
	c.TTL = durationToIntSeconds(record.TTL)
	c.Comment = generateComment()
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
		TTL:      time.Duration(m.TTL * 1000000000),
		Priority: m.Pref,
	}
}

func (m MXRecord) fromLibdnsRecord(record libdns.Record) {
	m.OwnerName = record.Name
	m.TTL = durationToIntSeconds(record.TTL)
	m.Name = record.Value
	m.Pref = record.Priority
	m.Comment = generateComment()
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
		TTL:   time.Duration(n.TTL * 1000000000),
	}
}

func (n NSRecord) fromLibdnsRecord(record libdns.Record) {
	n.OwnerName = record.Name
	n.TargetName = record.Value
	n.TTL = durationToIntSeconds(record.TTL)
	n.Comment = generateComment()
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
		TTL:   time.Duration(t.TTL * 1000000000),
	}
}

func (t TXTRecord) fromLibdnsRecord(record libdns.Record) {
	t.Name = record.Name
	t.Text = record.Value
	t.TTL = durationToIntSeconds(record.TTL)
	t.Comment = generateComment()
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
		TTL:   time.Duration(t.TTL * 1000000000),
	}
}

func (t TLSARecord) fromLibdnsRecord(record libdns.Record) {
	t.Name = record.Name
	t.Text = record.Value
	t.TTL = durationToIntSeconds(record.TTL)
	t.Comment = generateComment()
}

func durationToIntSeconds(duration time.Duration) int {
	return int(duration.Minutes())
}

func generateComment() string {
	return fmt.Sprintf("This record was created with libdns at %s UTC", time.Now().UTC().Format(time.DateTime))
}
