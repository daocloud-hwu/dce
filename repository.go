package dce

import (
    "encoding/json"
    "fmt"
)

type RepoMetadata struct {
    Scopes []string `json:"Scopes"`
    Visibility bool `json:"Visibility"`
}

type Repository struct {
    Name             string        `json:"Name"`
    Namespace        string        `json:"Namespace"`
    RegistryName     string        `json:"RegistryName"`
    Registry         string        `json:"Registry"`
    FullName         string        `json:"FullName"`
    ShortDescription string        `json:"ShortDescription"`
    LongDescription  string        `json:"LongDescription"`
    Metadata         *RepoMetadata `json:"Metadata"`
}

type RepoTag struct {
    Name         string `json:"Name"`
    IsDownloaded bool   `json:"IsDownloaded"`
    UpdatedAt    int64  `json:"UpdatedAt"`
}

func (c *Client) ListRepository(namespace string) ([]*Repository, error) {
    url := fmt.Sprintf("/api/registries/buildin-registry/repositories/%s", namespace)
    status, outbody, _, err := c.do("GET", url, nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    var result []*Repository
    if err := json.Unmarshal(outbody, &result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) ListRepoTag(namespace, name string) ([]*RepoTag, error){
    url := fmt.Sprintf("/api/registries/buildin-registry/repositories/%s/%s/tags", namespace, name)
    status, outbody, _, err := c.do("GET", url, nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    var result []*RepoTag
    if err := json.Unmarshal(outbody, &result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) GetRepository(namespace, name string) (*Repository, error) {
    url := fmt.Sprintf("/api/registries/buildin-registry/repositories/%s/%s", namespace, name)
    status, outbody, _, err := c.do("GET", url, nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Repository)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) CreateRepository(namespace, name, sDes, lDes string) (*Repository, error) {
    type Input struct {
        Name             string `json:"Name"`
        ShortDescription string `json:"ShortDescription"`
        LongDescription  string `json:"LongDescription"`
    }

    input := new(Input)
    input.Name = name
    input.ShortDescription = sDes
    input.LongDescription = lDes

    inbody, err := json.Marshal(input)
    if err != nil {
        return nil, err
    }

    url := fmt.Sprintf("/api/registries/buildin-registry/repositories/%s", namespace)
    status, outbody, _, err := c.do("POST", url, nil, inbody)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Repository)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) UpdateRepository(namespace, name, sDes, lDes string) error {
    // TBD
    return nil
}

func (c *Client) DeleteRepository(namespace, name string) error {
    url := fmt.Sprintf("/api/registries/buildin-registry/repositories/%s/%s", namespace, name)
    status, outbody, _, err := c.do("DELETE", url, nil, nil)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}
