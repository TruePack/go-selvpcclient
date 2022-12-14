package projects

import (
	"encoding/json"
)

// CreateOpts represents options for the project Create request.
type CreateOpts struct {
	// Name sets the name for a new project.
	Name string `json:"-"`
}

// MarshalJSON implements custom marshalling method for the CreateOpts.
func (opts *CreateOpts) MarshalJSON() ([]byte, error) {
	// Return create options with only name and auto_quotas parameters if quotas
	// parameter hadn't been provided.
	return json.Marshal(&struct {
		Name string `json:"name"`
	}{
		Name: opts.Name,
	})
}

// UpdateOpts represents options for the project Update request.
type UpdateOpts struct {
	// Name represents the name of a project.
	Name string `json:"name,omitempty"`

	// CustomURL is a public url of a project that can be set by a user.
	CustomURL *string `json:"custom_url,omitempty"`

	// Theme represents project theme settings.
	Theme *ThemeUpdateOpts `json:"theme,omitempty"`
}

// ThemeUpdateOpts represents project theme options for the Update request.
type ThemeUpdateOpts struct {
	// Color is a hex string with a custom background color.
	Color *string `json:"color,omitempty"`

	// Logo contains url for the project custom header logotype.
	Logo *string `json:"logo,omitempty"`
}
