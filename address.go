package main

import "encoding/json"

type Address struct {
	Code         string `json:"-"`
	State        string `json:"-"`
	City         string `json:"-"`
	Neighborhood string `json:"-"`
	Street       string `json:"-"`
}

func (a *Address) UnmarshalJSON(data []byte) error {
	var temp map[string]string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	a.Code = temp["cep"]

	if state, ok := temp["state"]; ok {
		a.State = state
	} else if state, ok := temp["uf"]; ok {
		a.State = state
	}

	if city, ok := temp["city"]; ok {
		a.City = city
	} else if city, ok := temp["localidade"]; ok {
		a.City = city
	}

	if neighborhood, ok := temp["neighborhood"]; ok {
		a.Neighborhood = neighborhood
	} else if neighborhood, ok := temp["bairro"]; ok {
		a.Neighborhood = neighborhood
	}

	if street, ok := temp["street"]; ok {
		a.Street = street
	} else if street, ok := temp["logradouro"]; ok {
		a.Street = street
	}

	return nil
}
