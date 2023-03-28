package bigquery

import (
	"fmt"
	"net/url"
	"strings"
)

// ConfigFromConnString will return the Config structures
func ConfigFromConnString(in string) (*Config, error) {
	cfg := &Config{}
	if strings.HasPrefix(in, "bigquery://") {
		u, err := url.Parse(in)
		if err != nil {
			return nil, fmt.Errorf("invalid connection string: %s (%s)", in, err.Error())
		}
		v, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			return nil, fmt.Errorf("invalid connection string: %s (%s)", in, err.Error())
		}
		cfg.ProjectID = u.Host
		cfg.Location = strings.Trim(u.Path, "/")
		cfg.DatasetID = v.Get("dataset")
		cfg.ApiKey = v.Get("apiKey")
		cfg.Credentials = v.Get("credentials")
		return cfg, nil
	} else {
		// Nope, bad prefix
		return nil, fmt.Errorf("invalid prefix, expected bigquery:// got: %s", in)
	}
}
