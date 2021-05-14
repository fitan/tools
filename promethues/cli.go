package promethues

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"strings"
)

type Client struct {
	HttpClient *resty.Client
}

func NewClient(url string, debug bool) *Client {
	return &Client{HttpClient: resty.New().SetDebug(debug).SetHostURL(url)}
}

// QueryRange 范围查询
func (c *Client) QueryRange(query, startTime, endTime, step, timeout string) (*resty.Response, error) {
	parmas := map[string]string{
		"query": query,
		"start": startTime,
		"end":   endTime,
		"step":  step,
	}
	if timeout != "" {
		parmas["timeout"] = timeout
	}
	return c.HttpClient.R().SetQueryParams(parmas).Get("/api/v1/query_range")
}

// InstantQuery 时间点查询
func (c *Client) InstantQuery(query, time, timeout string) (*resty.Response, error) {
	parmas := map[string]string{}
	parmas["query"] = query
	if time != "" {
		parmas["time"] = time
	}
	if timeout != "" {
		parmas["timeout"] = timeout
	}
	return c.HttpClient.R().SetQueryParams(parmas).Get("/api/v1/query")
}

var InstantIPF = func(data []byte) (map[string]string, error) {
	it := InstantType{}
	err := json.Unmarshal(data, &it)
	if err != nil {
		return nil, err
	}
	if it.Status != "success" {
		return nil, errors.New(it.Error)
	}
	m := make(map[string]string, len(it.Data.Result))
	for _, v := range it.Data.Result {
		m[strings.Split(v.Metric.Instance, ":")[0]] = v.Value[1].(string)
	}
	return m, nil
}

var RangeIPF = func(data []byte) (map[string][][]interface{}, error) {
	rt := RangeType{}
	err := json.Unmarshal(data, &rt)
	if err != nil {
		return nil, err
	}
	if rt.Status != "success" {
		return nil, errors.New(rt.Error)
	}
	m := make(map[string][][]interface{}, len(rt.Data.Result))
	for _, v := range rt.Data.Result {
		m[strings.Split(v.Metric.Instance, ":")[0]] = v.Values
	}
	return m, nil
}

func (c *Client) BatchInstant(query string, f func(data []byte) (map[string]string, error)) (map[string]string, error) {
	res, err := c.InstantQuery(query, "", "")
	if err != nil {
		return nil, err
	}
	return f(res.Body())
}

func (c *Client) BatchInstantByIP(query string) (map[string]string, error) {
	return c.BatchInstant(query, InstantIPF)
}

func (c *Client) BatchRange(query, start, end, step string, f func(data []byte) (map[string][][]interface{}, error)) (map[string][][]interface{}, error) {
	res, err := c.QueryRange(query, start, end, step, "")
	if err != nil {
		return nil, err
	}
	return f(res.Body())
}

func (c *Client) BatchRangeByIP(query, start, end, step string) (map[string][][]interface{}, error) {
	return c.BatchRange(query, start, end, step, RangeIPF)
}
