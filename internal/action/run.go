package action

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-tfe"
	"github.com/sethvargo/go-githubactions"
)

func Run(inputs Inputs) error {
	ctx := context.Background()

	client, err := tfe.NewClient(&tfe.Config{
		Token:   inputs.Token,
		Address: inputs.Address,
	})
	if err != nil {
		return fmt.Errorf("failed to configure client: %w", err)
	}

	workspace, err := client.Workspaces.Read(ctx, inputs.Organization, inputs.Workspace)
	if err != nil {
		return fmt.Errorf("failed to find workspace %s/%s: %w", inputs.Organization, inputs.Workspace, err)
	}

	stateVersion, err := client.StateVersions.CurrentWithOptions(ctx, workspace.ID, &tfe.StateVersionCurrentOptions{
		Include: "outputs",
	})
	if err != nil {
		return fmt.Errorf("failed to fetch state outputs for workspace %s/%s: %w", inputs.Organization, inputs.Workspace, err)
	}

	for _, o := range stateVersion.Outputs {
		str := fmt.Sprint(o.Value)

		fmt.Printf("value: %v\n", o.Value)

		fmt.Println(o.Name, str)

		if o.Sensitive {
			githubactions.AddMask(str)
		}

		githubactions.SetOutput(o.Name, str)
	}

	return nil
}

type Inputs struct {
	Token        string
	Address      string
	Workspace    string
	Organization string
}
