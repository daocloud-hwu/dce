package dce

import (
    "encoding/json"
    "fmt"
)

type NamespaceMetadata struct {
    Visibility bool `json:"Visibility"`
}

type BriefNamespace struct {
    Name     string             `json:"Name"`
    Metadata *NamespaceMetadata `json:"Metadata"`
}

type Accessible struct {
    TeamID string `json:"TeamId"`
    Role   string `json:"Role"`
}

type Namespace struct {
    Name           string        `json:"Name"`
    Visibility     bool          `json:"Visibility"`
    AccessibleList []*Accessible `json:"AccessibleList"`
    Scopes         []string      `json:"Scopes"`
}

func (c *Client) ListNamespace() ([]*BriefNamespace, error) {
    status, outbody, _, err := c.do("GET", "/api/registries/buildin-registry/namespaces", nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    var result []*BriefNamespace
    if err := json.Unmarshal(outbody, &result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) GetNamespace(name string) (*Namespace, error) {
    namespace := fmt.Sprintf("/api/registries/buildin-registry/namespaces/%s", name)
    status, outbody, _, err := c.do("GET", namespace, nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Namespace)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) CreateNamespace(name string) (*Namespace, error) {
    type Input struct {
        Name string `json:"Name"`
    }

    input := new(Input)
    input.Name = name

    inbody, err := json.Marshal(input)
    if err != nil {
        return nil, err
    }

    status, outbody, _, err := c.do("POST", "/api/registries/buildin-registry/namespaces", nil, inbody)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Namespace)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) UpdateNamespace(name string, visibility bool) error {
    type Input struct {
        Visibility bool `json:"Visibility"`
    }

    input := new(Input)
    input.Visibility = visibility

    inbody, err := json.Marshal(input)
    if err != nil {
        return err
    }

    url := fmt.Sprintf("/api/registries/buildin-registry/namespaces/%s", name)
    status, outbody, _, err := c.do("PATCH", url, nil, inbody)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) DeleteNamespace(name string) error {
    namespace := fmt.Sprintf("/api/registries/buildin-registry/namespaces/%s", name)
    status, outbody, _, err := c.do("DELETE", namespace, nil, nil)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) CreateAccess(namespace, teamID, role string) error {
    inbody, err := json.Marshal(&Accessible{TeamID: teamID, Role: role})
    if err != nil {
        return err
    }

    url := fmt.Sprintf("/api/registries/buildin-registry/namespaces/%s/accessible-list", namespace)
    status, outbody, _, err := c.do("POST", url, nil, inbody)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) UpdateAccess(namespace, teamID, role string) error {
    return c.CreateAccess(namespace, teamID, role)
}

func (c *Client) DeleteAccess(namespace, teamID string) error {
    url := fmt.Sprintf("/api/registries/buildin-registry/namespaces/%s/accessible-list?TeamId=%s", namespace, teamID)
    status, outbody, _, err := c.do("DELETE", url, nil, nil)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}
