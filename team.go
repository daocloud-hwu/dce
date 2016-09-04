package dce

import (
    "encoding/json"
    "fmt"
)

type Team struct {
    ID      string   `json:"Id"`
    Members []string `json:"Members"`
    Name    string   `json:"Name"`
}

func (c *Client) CreateTeam(name string) (*Team, error) {
    type Input struct {
        Name     string `json:"Name"`
    }

    input := new(Input)
    input.Name = name

    inbody, err := json.Marshal(input)
    if err != nil {
        return nil, err
    }

    status, outbody, _, err := c.do("POST", "/api/teams", nil, inbody)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Team)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) GetTeam(id string) (*Team, error) {
    status, outbody, _, err := c.do("GET", fmt.Sprintf("/api/teams/%s", id), nil, nil)
    if err != nil {
        return nil, err
    }
    if status/100 != 2 {
        return nil, fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    result := new(Team)
    if err := json.Unmarshal(outbody, result); err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) AddTeamMember(teamId, account string) error {
    team, err := c.GetTeam(teamId)
    if err != nil {
        return err
    }

    team.Members = append(team.Members, account)

    inbody, err := json.Marshal(team)
    if err != nil {
        return err
    }

    status, outbody, _, err := c.do("PATCH", fmt.Sprintf("/api/teams/%s", teamId), nil, inbody)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) DeleteTeamMember(teamId, account string) error {
    team, err := c.GetTeam(teamId)
    if err != nil {
        return err
    }

    var i int
    for i = 0; i < len(team.Members); i++ {
        if team.Members[i] == account {
            break
        }
    }

    if i == len(team.Members) {
        return nil
    }

    team.Members = append(team.Members[0:i], team.Members[i+1:]...)

    inbody, err := json.Marshal(team)
    if err != nil {
        return err
    }   
    
    status, outbody, _, err := c.do("PATCH", fmt.Sprintf("/api/teams/%s", teamId), nil, inbody)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}

func (c *Client) DeleteTeam(id string) error {
    status, outbody, _, err := c.do("DELETE", fmt.Sprintf("/api/teams/%s", id), nil, nil)
    if err != nil {
        return err
    }
    if status/100 != 2 {
        return fmt.Errorf("Status code is %d, reason %s", status, outbody)
    }

    return nil
}
