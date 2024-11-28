package plugin_test

import (
	"os"
	"testing"

	"ecs-task-runner/plugin"

	"github.com/stretchr/testify/assert"
)

func TestFailOnMissingEnvironment(t *testing.T) {
	var config plugin.Config
	fetcher := plugin.EnvironmentConfigFetcher{}

	t.Setenv("BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE", "")
	os.Unsetenv("BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE")

	err := fetcher.Fetch(&config)

	assert.NotNil(t, err, "fetch should error")
}

func TestFetchConfigFromEnvironment(t *testing.T) {
	var config plugin.Config
	fetcher := plugin.EnvironmentConfigFetcher{}

	t.Setenv("BUILDKITE_PLUGIN_EXAMPLE_GO_PARAMETER_NAME", "test-parameter")
	t.Setenv("BUILDKITE_PLUGIN_EXAMPLE_GO_SCRIPT", "hello-world")

	err := fetcher.Fetch(&config)

	assert.Nil(t, err, "fetch should not error")
	assert.Equal(t, config.ParameterName, "test-parameter", "fetched message should match environment")
	assert.Equal(t, config.Script, "hello-world", "fetched message should match environment")
}
