package dce

import (
    "encoding/json"
    "fmt"
)

type License struct {
    TimeLeft            int    `json:"TimeLeft"`
    IsFeatureLimit      bool   `json:"IsFeatureLimit"`
    CpuLeft             int    `json:"CpuLeft"`
    MessageType         string `json:"MessageType"`
    HasLicense          bool   `json:"HasLicense"`
    IsOverCpu           bool   `json:"IsOverCpu"`
    CanEnterDashboard   bool   `json:"CanEnterDashboard"`
    Message             string `json:"Message"`
    PromoteEnterLicense bool   `json:"PromoteEnterLicense"`
}

func (c *Client) SetAuth() error {
    type Input struct {
        Method string `json:"Method"`
    }

    input := new(Input)
    input.Method = "managed"

    inbody, err := json.Marshal(input)
    if err != nil {
        return err
    }

    status, outbody, _, err := c.do("PATCH", "/api/settings/auth", nil, inbody)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) GetLicenseKey() (*License, error) {
    status, outbody, _, err := c.do("GET", "/api/license/check", nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(License)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) SetLicenseKey(key string) error {
    type Input struct {
        Key string `json:"key"`
    }

    input := new(Input)
    input.Key = key

    inbody, err := json.Marshal(input)
    if err != nil {
        return err
    }

    status, outbody, _, err := c.do("PUT", "/api/settings/license/key", nil, inbody)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}
