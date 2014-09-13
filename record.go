package go_cloudflare

import (
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
