package model

type Layout struct {
	ID       string          `json:"id"`
	Status   string          `json:"status"`
	Variants []LayoutVariant `json:"variants"`
}

type LayoutVariant struct {
	ID      string        `json:"id"`
	Units   []string      `json:"units"`
	Support LayoutSupport `json:"support,omitempty"`
	Status  string        `json:"status"`
}

type LayoutSupport []LayoutSupportSpec

type LayoutSupportSpec struct {
	Sites       []string                `json:"sites"`
	Clients     []ClientSupportSpec     `json:"clients"`
	Characters  []string                `json:"characters,omitempty"`
	Audiences   []string                `json:"audiences,omitempty"`
	Experiments []ExperimentSupportSpec `json:"experiments,omitempty"`
}

type ExperimentSupportSpec struct {
	ID      string `json:"id"`
	Variant string `json:"variant"`
}
