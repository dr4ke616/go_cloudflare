package go_cloudflare

import (
	"errors"
	"fmt"
)

type Record struct {
	Id       string `json:"rec_id"`
	Domain   string `json:"zone_name"`
	Name     string `json:"display_name"`
	FullName string `json:"name"`
	Value    string `json:"content"`
	Type     string `json:"type"`
	Priority string `json:"prio"`
	Ttl      string `json:"ttl"`
}

type RecordsResponse struct {
	Response struct {
		Recs struct {
			Records []Record `json:"objs"`
		} `json:"recs"`
	} `json:"response"`
	Result  string `json:"result"`
	Message string `json:"msg"`
}

type RecordResponse struct {
	Response struct {
		Rec struct {
			Record Record `json:"obj"`
		} `json:"rec"`
	} `json:"response"`
	Result  string `json:"result"`
	Message string `json:"msg"`
}

type UpdateRecord struct {
	Type     string
	Name     string
	Content  string
	Ttl      string
	Priority string
}

func (c *Client) RetrieveAllRecords(domain string) (*RecordsResponse, error) {

	params := make(map[string]string)
	params["z"] = domain

	req, err := c.NewRequest(params, "POST", "rec_load_all")
	if err != nil {
		return nil, err
	}

	resp, err := checkResponse(c.HttpHandler.Do(req))
	if err != nil {
		return nil, fmt.Errorf("Errror retrieving records", err)
	}

	records := new(RecordsResponse)
	err = decodeBody(resp, records)
	if err != nil {
		return nil, fmt.Errorf("Error decoding record response", err)
	}

	return records, nil
}

func (c *Client) RetrieveARecord(domain string, id string) (*Record, error) {

	records, err := c.RetrieveAllRecords(domain)
	if err != nil {
		return nil, err
	}

	return records.FindRecord(id)
}

func (r *RecordsResponse) FindRecord(id string) (*Record, error) {

	if r.Result == "error" {
		return nil, fmt.Errorf("API Error: %s", r.Message)
	}

	objs := r.Response.Recs.Records
	notFoundErr := errors.New("Record not found")

	if len(objs) < 0 {
		return nil, notFoundErr
	}

	for _, v := range objs {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, notFoundErr
}

func (r *RecordResponse) GetRecord() (*Record, error) {

	if r.Result == "error" {
		return nil, fmt.Errorf("API Error: %s", r.Message)
	}
	return &r.Response.Rec.Record, nil
}

func (c *Client) UpdateRecord(domain string, id string, opts *UpdateRecord) error {

	params := make(map[string]string)
	params["z"] = domain
	params["id"] = id
	params["type"] = opts.Type
	params["name"] = opts.Name
	params["content"] = opts.Content
	params["prio"] = opts.Priority
	params["ttl"] = opts.Ttl

	req, err := c.NewRequest(params, "POST", "rec_edit")
	if err != nil {
		return err
	}

	resp, err := checkResponse(c.HttpHandler.Do(req))
	if err != nil {
		return fmt.Errorf("Error updating record: %s", err)
	}

	recordResp := new(RecordResponse)
	err = decodeBody(resp, &recordResp)
	if err != nil {
		return fmt.Errorf("Error parsing record response: %s", err)
	}

	_, err = recordResp.GetRecord()
	if err != nil {
		return err
	}

	// The request was successful
	return nil
}
