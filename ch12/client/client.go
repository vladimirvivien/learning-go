package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/vladimirvivien/learning-go/ch12/vector"
)

type vecClient struct {
	svcAddr string
	client  *http.Client
}

func newVecClient(addr string) *vecClient {
	return &vecClient{
		svcAddr: addr,
		client:  &http.Client{},
	}
}

func (c *vecClient) add(vec0, vec1 vector.SimpleVector) (vector.SimpleVector, error) {
	uri := c.svcAddr + "/vec/add"

	// encode params
	var body bytes.Buffer
	params := []vector.SimpleVector{vec0, vec1}
	if err := json.NewEncoder(&body).Encode(&params); err != nil {
		return []float64{}, err
	}
	req, err := http.NewRequest("POST", uri, &body)
	if err != nil {
		return []float64{}, err
	}

	// send request
	resp, err := c.client.Do(req)
	if err != nil {
		return []float64{}, err
	}
	defer resp.Body.Close()

	// handle response
	var result vector.SimpleVector
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []float64{}, err
	}
	return result, nil
}

func main() {
	c := newVecClient("http://127.0.0.1:4040")
	result, err := c.add(vector.New(1, 2), vector.New(3, 4))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
