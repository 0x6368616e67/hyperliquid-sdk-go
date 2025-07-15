package hyperliquid

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"reflect"

	"github.com/0x6368616e67/hyperliquid/api"
)

// Client represent a RPC Client
type Client struct {
	httpConn *httpConn
	wsConn   *wsConn
	apiURL   string
}

// Dial connects a client to the given URL.
func Dial(rpcurl string, wscurl string) (*Client, error) {
	return DialContext(context.Background(), rpcurl, wscurl)
}

// Dial connects a client to the given URL and given context
func DialContext(ctx context.Context, apiURL string, wsURL string) (client *Client, err error) {
	_, err = url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	client = &Client{}

	client.httpConn = newHTTPConn(apiURL)
	client.apiURL = apiURL
	if wsURL != "" {
		client.wsConn = newWSConn(wsURL)
		client.wsConn.Connect()
	}
	return
}

func (cli *Client) request(ctx context.Context, apiPath string, param interface{}, result interface{}) (err error) {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("result parameter must be pointer or nil interface: %v", result)
	}
	var resp []byte

	respBody, err := cli.httpConn.postJSON(ctx, apiPath, param)
	if err != nil {
		return
	}
	defer respBody.Close()

	resp, err = io.ReadAll(respBody)
	if err != nil {
		return err
	}
	log.Printf("body:%s", string(resp))
	json.Unmarshal(resp, result)
	return nil
}

// Retrieve mids for all coins
func (cli *Client) AllMids(ctx context.Context, dex string) (mids map[string]string, err error) {
	param := api.AllMidsRequest{
		Type: "allMids",
		Dex:  dex,
	}
	mids = make(map[string]string)
	err = cli.request(ctx, "info", param, &mids)
	if err != nil {
		return nil, err
	}
	return
}

func (cli *Client) Metadata(ctx context.Context, dex string) (coins []api.AssetMetadata, err error) {
	param := api.MetadataRequest{
		Type: "meta",
		Dex:  dex,
	}
	type WrapMetadata struct {
		Universe []api.AssetMetadata `json:"universe"`
	}
	var data WrapMetadata

	err = cli.request(ctx, "info", param, &data)
	if err != nil {
		return
	}
	coins = append(coins, data.Universe...)
	return
}

func (cli *Client) SubscribeAllMids() (err error) {
	msg := struct {
		Type string `json:"type"`
		//	Dex  string `json:"dex,omitempty"`
	}{
		Type: "allMids",
		//	Dex:  dex,
	}
	cli.wsConn.Subscribe(msg)
	return
}
