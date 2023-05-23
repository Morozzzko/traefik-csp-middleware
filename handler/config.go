package handler

import (
	"regexp"

	"github.com/morozzzko/traefik-csp-middleware/httputil"
)

type nonceGenerator func() []byte

// Config holds the plugin configuration.
type Config struct {
	LastModified   bool                      `json:"lastModified" toml:"lastModified" yaml:"lastModified"`
	Placeholder    string                    `json:"placeholder" toml:"placeholder" yaml:"placeholder" default:"DhcnhD3khTMePgXw"`
	NonceGenerator nonceGenerator            `json:"-" toml:"-" yaml:"-"`
	LogLevel       int8                      `json:"logLevel" toml:"logLevel" yaml:"logLevel"`
	Monitoring     httputil.MonitoringConfig `json:"monitoring" toml:"monitoring" yaml:"monitoring"`
}

type rewrite struct {
	regex         *regexp.Regexp
	generateNonce nonceGenerator
}
