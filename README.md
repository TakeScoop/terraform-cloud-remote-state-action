# terraform-cloud-outputs-action

GitHub Action to fetch outputs from a Terraform Cloud workspace

## Usage

Given a workspace output `foo` in the Terraform Cloud organization `org` and workspace name `name`

```hcl
output "foo" {
  value = "bar
}
```

```yml
- uses: takescoop/terraform-cloud-outputs-action@v0.1
  id: action
  with:
    workspace: name
    organization: org
    token: ${{ secrets.terraform_cloud_token }}
- name: Test string
  run: test "${{ steps.action.outputs.foo }}" == "bar"
```
