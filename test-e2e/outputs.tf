output "foo" {
  description = "A single ouptput"
  value       = "bar"
}

output "json" {
  description = "A multiline output"
  value       = <<EOF
  {
    "foo": "bar",
    "baz": "woz"
  }
EOF
}

output "sensitive" {
  description = "A sensitive ouptput to test masking"
  value       = "secret"
  sensitive   = true
}
