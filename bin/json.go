package main

import (
	"encoding/json"

	"github.com/gotamer/hdkeys"
)

func GetAllJson(no uint32) (string, error) {
	Info.Println("no: ", no)
	type Keys struct {
		Nostr   *hdkeys.Nostr
		Bitcoin *hdkeys.Bitcoin
	}

	var out = new(Keys)
	var err error

	out.Nostr, err = km.GetNostr(no)
	if err != nil {
		return "", err
	}

	out.Bitcoin, err = hdkeys.GetBitcoinFromWif(out.Nostr.Wif)
	if err != nil {
		return "", err
	}

	jsonbyte, err := json.Marshal(&out)
	if err != nil {
		return "", err
	}
	return string(jsonbyte), nil
}

func GetNostrJson(no uint32) (string, error) {

	nostr, err := km.GetNostr(no)
	if err != nil {
		return "", err
	}

	jsonbyte, err := json.Marshal(&nostr)
	if err != nil {
		return "", err
	}
	return string(jsonbyte), nil
}

func GetBitcoinFromWifJson(wif string) (string, error) {

	btc, err := hdkeys.GetBitcoinFromWif(wif)
	if err != nil {
		return "", err
	}

	jsonbyte, err := json.Marshal(&btc)
	if err != nil {
		return "", err
	}
	return string(jsonbyte), nil
}
