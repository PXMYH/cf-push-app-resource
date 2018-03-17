package concourse

// Source -
type Source struct {
	API      string `json:"api"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSOToken string `json:"sso-token"`

	Org   string   `json:"org"`
	Space string   `json:"space"`
	Apps  []string `json:"apps"`

	SkipSSLValidation bool `json:"skip-ssl-validation"`

	Debug bool `json:"debug"`
	Trace bool `json:"trace"`
}

// Version - { "app name": (serialized filters.AppEvent), ... }
type Version map[string]string

// MetadataPair -
type MetadataPair struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	URL      string `json:"url"`
	Markdown bool   `json:"markdown"`
}

type OutParams struct {
  Status string `json:"status"`
  Media []string `json:"media,omitempty"`
}

type OutRequest struct {
  Source Source `json:"source"`
  Params OutParams `json:"params"`
}

type OutResponse struct {
  Version Version `json:"version"`
  Metadata []MetadataPair `json:"metadata"`
}
