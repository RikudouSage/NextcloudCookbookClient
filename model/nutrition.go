package model

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
