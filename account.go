package dce

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
)

type Account struct {
    IsAdmin bool   `json:"IsAdmin"`
    IsLdap  bool   `json:"IsLdap"`
    Name    string `json:"Name"`
    Email   string `json:"Email"`
}

func (c *Client) CreateAccount(name, password, email string, isAdmin bool) (*Account, error) {
    type Input struct {
        Name     string `json:"Name"`
        Password string `json:"Password"`
        Email    string `json:"Email"`
        IsAdmin  string `json:"IsAdmin"`
        Role     string `json:"Role"`
    }

    input := new(Input)
    input.Name = name
    input.Password = password
    input.Email = email
    input.IsAdmin = "no"
    input.Role = "no_access"

    if isAdmin {
        input.IsAdmin = "yes"
        input.Role = "full_control"
    }

    inbody, err := json.Marshal(input)
    if err != nil {
        return nil, err
    }

    status, outbody, _, err := c.do("POST", "/api/accounts", nil, inbody)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Account)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) DeleteAccount(name string) error {
    status, outbody, _, err := c.do("DELETE", fmt.Sprintf("/api/accounts/%s", name), nil, nil)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) Login(username, password string) error {
    header := make(map[string]string)
    header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))

    status, outbody, _, err := c.do("POST", "/api/login", header, nil)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    type Output struct {
        AccessToken string `json:"AccessToken"`
    }

    result := new(Output)
    if err := json.Unmarshal(outbody, result); err != nil {
        return err
    }

    c.AccessToken = result.AccessToken
    return nil
}
