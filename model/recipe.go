package model

import (
	"net/url"

	"github.com/sosodev/duration"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/types"
)

type Recipe struct {
	ID                  string             `json:"id,omitzero"`
	Type                string             `json:"@type,omitzero"`
	Name                string             `json:"name"`
	Keywords            types.CSVSlice     `json:"keywords,omitempty"`
	CreatedDate         types.APITime      `json:"dateCreated,omitzero"`
	ModifiedDate        types.APITime      `json:"dateModified,omitzero"`
	ImageURL            *types.APIURL      `json:"imageUrl,omitempty"`
	ImagePlaceholderURL *types.APIURL      `json:"imagePlaceholderUrl,omitempty"`
	PreparationTime     *duration.Duration `json:"prepTime,omitempty"`
	CookTime            *duration.Duration `json:"cookTime,omitempty"`
	TotalTime           *duration.Duration `json:"totalTime,omitempty"`
	Description         string             `json:"description"`
	URL                 *types.APIURL      `json:"url,omitempty"`
	ImageURL2           *types.APIURL      `json:"image,omitempty"`
	Servings            uint               `json:"recipeYield"`
	Category            string             `json:"recipeCategory"`
	Tools               []string           `json:"tool,omitempty"`
	Ingredients         []string           `json:"recipeIngredient,omitempty"`
	Instructions        []string           `json:"recipeInstructions,omitempty"`
	Nutrition           Nutrition          `json:"nutrition,omitzero"`
}

func (receiver *Recipe) SetBaseURL(baseURL *url.URL) error {
	return populateAbsoluteUrls(receiver, baseURL)
}
