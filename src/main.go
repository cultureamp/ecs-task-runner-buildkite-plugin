package main

import (
	"context"
	"os"

	"ecs-task-runner/buildkite"
	"ecs-task-runner/plugin"
)

func main() {
	ctx := context.Background()
	agent := &buildkite.Agent{}
	fetcher := plugin.EnvironmentConfigFetcher{} //TODO: Is this the BK-specific env-vars fetcher?
	examplePlugin := plugin.ExamplePlugin{}

	err := examplePlugin.Run(ctx, fetcher, agent)

	if err != nil {
		buildkite.LogFailuref("plugin execution failed: %s\n", err.Error())
		os.Exit(1)
	}
}
