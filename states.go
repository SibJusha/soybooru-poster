package main

import (
	"encoding/json"
	"os"
)

const stateFile = "soybooru_state.json"

func loadState() (State, error) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return State{LastMaxID: 0}, nil
		}
		return State{}, err
	}
	var s State
	err = json.Unmarshal(data, &s)
	return s, err
}

func saveState(s State) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}
