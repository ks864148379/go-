package cg

import (
	"errors"
	"encoding/json"

	"ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient)AddPlayer(player *Player) error {
	b, err := jsonMarshal(*player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addplayer", string(b))

	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient)RemovePlayer(name string) error {
	ret, _ := client.Call("RemovePlayer", name)

	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient)ListPlayer(name string) (ps[] *Player, err error) {
	retsp, _ := client.Call("listPlayer", name)

	if ret.Code == "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return 
}

func (client *CenterClient)Broadcast(message string) error {
	m := &Message{Content:message}

	b, err := json.Marshal(m)

	if err != nil {
		return err
	}

	resp, _ := client.Call("broadcast", string(b))

	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}