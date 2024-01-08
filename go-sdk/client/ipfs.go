package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ipfs/kubo/client/rpc"
	"io"

	"log"
	"net/http"
)

type IpfsClient struct {
	Ipfs *rpc.HttpApi
}

func NewIpfsClient(connection string) *IpfsClient {

	api, err := rpc.NewURLApiWithClient(connection, &http.Client{})
	if err != nil {
		log.Fatal(err)
	}
	return &IpfsClient{
		Ipfs: api,
	}
}

func (cli *IpfsClient) add(data []byte) (map[string]string, error) {
	reader := bytes.NewBuffer(data)
	resp, err := cli.Ipfs.Request("add").FileBody(reader).Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, *resp.Error
	}
	return readJSON(resp.Output)

}

func (cli *IpfsClient) AddAndPublish(data []byte) (map[string]string, error) {
	resp, err := cli.add(data)
	if err != nil {
		log.Fatal("Could not store data", err)
		return nil, err
	}
	return resp, nil
	//return cli.publish(resp["Hash"], "") //todo: find a way to get faster response with ipns
}
func (cli *IpfsClient) StoreObject(data []byte) ([]byte, error) {
	// store data in io reader
	reader := bytes.NewBuffer(data)
	resp, err := cli.Ipfs.Request("object/put").FileBody(reader).Send(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.Error != nil {
		log.Fatal(*resp.Error)
		return nil, err
	}
	all, err := io.ReadAll(resp.Output)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return all, nil
}

func (cli *IpfsClient) publish(hash string, key string) (map[string]string, error) {
	//bind ipns to hash
	req := cli.Ipfs.Request("name/publish", hash)
	if key != "" {
		req = req.Option("key", key)
	}
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, *resp.Error
	}
	return readJSON(resp.Output)
}

func (cli *IpfsClient) Retrieve(uri string, ipns bool) (map[string]string, error) {
	var hash string
	if ipns {
		resolverJson, err := cli.Resolve(uri)
		if err != nil {
			log.Fatal("Could not resolve ipnsHash", err)
			return nil, err
		}
		// path gets returned as /ipfs/<hash> get hash from path
		path := resolverJson["Path"]
		hash = path[6:]
	} else {
		hash = uri
	}

	resp, err := cli.Ipfs.Request("cat", hash).Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, *resp.Error
	}

	return readJSON(resp.Output)

}
func (cli *IpfsClient) RetrieveRaw(ipfsUri string) ([]byte, error) {

	resp, err := cli.Ipfs.Request("cat", ipfsUri).Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, *resp.Error
	}

	return readRaw(resp.Output)

}

func (cli *IpfsClient) Update(ipnsHash string, data []byte) (map[string]string, error) {
	resp, err := cli.add(data)
	if err != nil {
		log.Fatal("Could not store data", err)
		return nil, err
	}

	return cli.publish(resp["Hash"], ipnsHash)

}
func (cli *IpfsClient) Resolve(hash string) (map[string]string, error) {
	hash = "/ipns/" + hash
	resp, err := cli.Ipfs.Request("resolve", hash).Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, *resp.Error
	}

	return readJSON(resp.Output)

}

func readJSON(resp io.ReadCloser) (map[string]string, error) {
	all, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	jsonRes, err := unMarshallJson(all)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func readRaw(resp io.ReadCloser) ([]byte, error) {
	all, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func unMarshallJson(data []byte) (map[string]string, error) {
	respDict := make(map[string]string)
	err := json.Unmarshal(data, &respDict)
	if err != nil {
		log.Fatal("Could not unmarshal data", err)
		return nil, err
	}
	return respDict, nil
}
