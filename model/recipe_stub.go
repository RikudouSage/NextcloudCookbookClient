package model

import "go.chrastecky.dev/nextcloud-cookbook/cookbook/types"

type RecipeStub struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Keywords            types.CSVSlice `json:"keywords"`
	CreatedDate         types.APITime  `json:"dateCreated"`
	ModifiedDate        types.APITime  `json:"dateModified"`
	ImageURL            *types.APIURL  `json:"imageUrl"`
	ImagePlaceholderURL *types.APIURL  `json:"imagePlaceholderUrl"`
}
