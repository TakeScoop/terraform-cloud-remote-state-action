package main

import (
	"github.com/sethvargo/go-githubactions"
	"github.com/takescoop/terraform-cloud-remote-state-action/internal/action"
)

func main() {
	if err := action.Run(action.Inputs{
		Token:        githubactions.GetInput("token"),
		Address:      githubactions.GetInput("address"),
		Workspace:    githubactions.GetInput("workspace"),
		Organization: githubactions.GetInput("organization"),
	}); err != nil {
		githubactions.Fatalf("Error: %s", err)
	}
}
