output "foo" {
  description = "A single ouptput"
  value       = "bar"
}

output "json" {
  description = "A JSON output"
  value = jsonencode({
    foo = "bar"
    baz = "woz"
  })
}

output "sensitive" {
  description = "A sensitive ouptput to test masking"
  value       = "secret"
  sensitive   = true
}

output "int" {
  description = "An int"
  value       = 2
}

output "float" {
  description = "A floating point number"
  value       = 3.14
}

output "map" {
  description = "A map"
  value       = tomap({ "foo" = "bar" })
}

output "list" {
  description = "A list"
  value       = tolist(["foo", "bar"])
}

output "set" {
  description = "A set"
  value       = toset(["foo", "bar"])
}

output "null" {
  description = "null"
  value       = null
}

output "multiline" {
  description = "A multiline output"
  value       = <<EOF
    mutli
    line
    output
  EOF
}
