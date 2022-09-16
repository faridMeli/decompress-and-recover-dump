package model

type Shortcut struct {
	ID        string      `json:"id"`
	Scopes    []string    `json:"scopes"`
	Audiences []string    `json:"audiences"`
	Tags      []string    `json:"tags"`
	Status    string      `json:"status"`
	Variants  []Variant   `json:"variants"`
	Overwrite []Overwrite `json:"overwrite"`
}

type VariantContent struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Link  string `json:"link"`
}

type VariantCreation struct {
	ID      string         `json:"id"`
	Content VariantContent `json:"content"`
	Support Support        `json:"support,omitempty"`
}
type Support []SupportSpec
type ClientSupportSpec struct {
	ID           uint64 `json:"id"`
	SinceVersion string `json:"since_version,omitempty"` // TODO: Add serilizable version
}

type Variant struct {
	VariantCreation
	Status string `json:"status"`
}
type SupportSpec struct {
	Sites        []string            `json:"sites"`
	Clients      []ClientSupportSpec `json:"clients"`
	Audiences    []string            `json:"audiences,omitempty"`
	HasCompanion *bool               `json:"has_companion,omitempty"`
	IsAuth       *bool               `json:"is_auth,omitempty"`
}

type Overwrite struct {
	OverwriteCreation
	Status string `json:"status"`
}
type OverwriteContent struct {
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Link  string `json:"link,omitempty"`
}

type OverwriteCreation struct {
	ID      string           `json:"id"`
	Content OverwriteContent `json:"content"`
	Support Support          `json:"support,omitempty"`
}
