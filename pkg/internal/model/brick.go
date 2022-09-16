package model

// Brick the fields that represent a brick
type Brick struct {
	ID                 string         `json:"id"`
	CollaboratorScopes []string       `json:"collaborator_scopes,omitempty"`
	Variants           []BrickVariant `json:"variants,omitempty"`
}

type BrickVariant struct {
	BrickDTO
	Support  Support  `json:"support,omitempty"`
	Services []string `json:"services,omitempty"`
}

type BrickDTO struct {
	ID      string                 `json:"id"`
	Ordinal int                    `json:"ordinal,omitempty"`
	UiType  string                 `json:"ui_type,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
