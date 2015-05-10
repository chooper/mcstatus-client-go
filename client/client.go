
package client

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
)

type Api struct {
    Uri     string
}

type PlayerList []string

func Connect(uri string) (*Api) {
    return &Api{
        Uri:    uri,
    }
}

func (api *Api) PlayersOnline(server string) (PlayerList, error) {
    uri := api.Uri + "/playersonline/" + server

    response, err := http.Get(uri)
    if err != nil {
        log.Print(err)
        return PlayerList{}, err
    }

    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Print(err)
        return PlayerList{}, err
    }

    var players PlayerList
    if err := json.Unmarshal([]byte(body), &players); err != nil {
        log.Print(err)
        return PlayerList{}, err
    }

    return players, nil
}

