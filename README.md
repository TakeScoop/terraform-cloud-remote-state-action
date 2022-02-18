# terraform-cloud-remote-state-action

GitHub Action to fetch outputs from a Terraform Cloud workspace

## Usage

Given a workspace output `foo` in the Terraform Cloud organization `org` and workspace name `name`

```hcl
output "foo" {
  value = "bar
}
```

```yml
- name: Run action
  id: action
  uses: takescoop/terraform-cloud-remote-state-action@v0.1
  with:
    workspace: name
    organization: org
    token: ${{ secrets.terraform_cloud_token }}
- name: Test string
  run: test "${{ steps.action.outputs.foo }}" == "bar"
```
