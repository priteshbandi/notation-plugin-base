package plugin

// GetMetadataRequest contains the parameters passed in a get-plugin-metadata
// request.
type GetMetadataRequest struct {
	PluginConfig map[string]string `json:"pluginConfig,omitempty"`
}

// GetMetadataResponse provided by the plugin.
type GetMetadataResponse struct {
	Name                      string       `json:"name"`
	Description               string       `json:"description"`
	Version                   string       `json:"version"`
	URL                       string       `json:"url"`
	SupportedContractVersions []string     `json:"supportedContractVersions,omitempty"`
	Capabilities              []Capability `json:"capabilities"`
}
