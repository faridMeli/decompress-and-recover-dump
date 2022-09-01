package model

// Brick the fields that represent a brick
type Brick struct {
	ID                 string         `json:"id"`
	CollaboratorScopes []string       `json:"collaborator_scopes,omitempty"`
	Variants           []BrickVariant `json:"variants,omitempty"`
}

type BrickVariant struct {
	ID       string                 `json:"id" validate:"required"`
	Ordinal  *int                   `json:"ordinal,omitempty"`
	UiType   string                 `json:"ui_type" validate:"required"`
	Data     map[string]interface{} `json:"data,omitempty"`
	Support  Support                `json:"support,omitempty"`
	Services []string               `json:"services,omitempty"`
}
