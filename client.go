package dce

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
)

type Client struct {
    Host        string
    AccessToken string
}

func (c *Client) do(method, url string, header map[string]string, body []byte) (int, []byte, map[string]string, error) {
    var reader io.Reader = nil
    if body != nil {
        reader = bytes.NewBuffer(body)
    }

    client := &http.Client{}
    client.Transport = &http.Transport{DisableKeepAlives: true}

    req, err := http.NewRequest(method, fmt.Sprintf("http://%s%s", c.Host, url), reader)
    if err != nil {
        return 0, nil, nil, err
    }

    for k, v := range header {
        req.Header.Set(k, v)
    }

    req.Header.Set("Content-Type", "application/json")
    if c.AccessToken != "" {
        req.Header.Set("X-DCE-Access-Token", c.AccessToken)
    }

    res, err := client.Do(req)
    if err != nil {
        return 0, nil, nil, err
    }
    defer res.Body.Close()

    outbody, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return 0, nil, nil, err
    }

    var resHeader map[string]string = nil
    for k, _ := range res.Header {
        if resHeader == nil {
            resHeader = make(map[string]string)
        }
        resHeader[k] = res.Header.Get(k)
    }

    return res.StatusCode, outbody, resHeader, nil
}
