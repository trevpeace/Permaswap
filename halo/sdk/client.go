package sdk

import (
	"encoding/json"
	"errors"

	hvmSchema "github.com/permadao/permaswap/halo/hvm/schema"
	"github.com/permadao/permaswap/halo/schema"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

type Client struct {
	cli *gentleman.Client
}

func NewClient(haloURL string) *Client {
	return &Client{
		cli: gentleman.New().URL(haloURL),
	}
}

func (c *Client) GetInfo() (info schema.InfoRes, err error) {
	req := c.cli.Request()
	req.AddPath("/info")

	res, err := req.Send()
	if err != nil {
		return
	}
	defer res.Close()
	if !res.Ok {
		err = errors.New(res.String())
		return
	}

	info = schema.InfoRes{}
	err = json.Unmarshal(res.Bytes(), &info)
	return
}

func (c *Client) SubmitTx(tx hvmSchema.Transaction) (everhash string, err error) {
	req := c.cli.Request()
	req.AddPath("/submit")
	req.Method("POST")
	req.Use(body.JSON(tx))

	res, err := req.Send()
	if err != nil {
		return
	}
	defer res.Close()
	if !res.Ok {
		err = errors.New(res.String())
		return
	}
	submitRes := schema.SubmitRes{}
	err = json.Unmarshal(res.Bytes(), &submitRes)
	if err != nil {
		return
	}
	return submitRes.EverHash, nil
}
