package plugin

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ParameterName string `split_words:"true" required:"true"`
	Script        string `split_words:"true" required:"true"`
}

type EnvironmentConfigFetcher struct {
}

const pluginEnvironmentPrefix = "BUILDKITE_PLUGIN_EXAMPLE_GO"

func (f EnvironmentConfigFetcher) Fetch(config *Config) error {
	return envconfig.Process(pluginEnvironmentPrefix, config)
}
