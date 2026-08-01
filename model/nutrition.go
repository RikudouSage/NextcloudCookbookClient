package model

import "encoding/json"

type Nutrition struct {
	Type                  string `json:"@type,omitzero"`
	Calories              string `json:"calories,omitzero"`
	CarbohydrateContent   string `json:"carbohydrateContent,omitzero"`
	CholesterolContent    string `json:"cholesterolContent,omitzero"`
	FatContent            string `json:"fatContent,omitzero"`
	FiberContent          string `json:"fiberContent,omitzero"`
	ProteinContent        string `json:"proteinContent,omitzero"`
	SaturatedFatContent   string `json:"saturatedFatContent,omitzero"`
	ServingSize           string `json:"servingSize,omitzero"`
	SodiumContent         string `json:"sodiumContent,omitzero"`
	SugarContent          string `json:"sugarContent,omitzero"`
	TransFatContent       string `json:"transFatContent,omitzero"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent,omitzero"`
}

func (receiver *Nutrition) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "[]" { // god bless unsanitized PHP's json_encode
		*receiver = Nutrition{}
		return nil
	}

	if bytes == nil || string(bytes) == "null" {
		*receiver = Nutrition{}
		return nil
	}

	type nutrition Nutrition

	var value nutrition
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}

	*receiver = Nutrition(value)
	return nil
}
