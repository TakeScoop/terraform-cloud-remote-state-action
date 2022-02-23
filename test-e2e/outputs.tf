output "foo" {
  description = "A single ouptput"
  value       = "bar"
}

output "int" {
  description = "An int"
  value       = 2
}

output "float" {
  description = "A floating point number"
  value       = 3.14
}

output "bool" {
  description = "A boolean"
  value       = true
}

output "map" {
  description = "A map"
  value = tomap({
    foo = {
      bar = "baz"
    }
  })
}

output "json" {
  description = "A JSON output"
  value = jsonencode({
    foo = "bar"
    baz = "woz"
  })
}

output "set" {
  description = "A set"
  value       = toset(["foo", "bar"])
}

output "list" {
  description = "A list"
  value       = tolist(["foo", "bar"])
}

output "null" {
  description = "null"
  value       = null
}

output "multiline" {
  description = "A multiline output"
  value       = join("\n", ["multi", "line", "output"])
}

output "sensitive" {
  description = "A sensitive ouptput to test masking"
  value       = "secret"
  sensitive   = true
}
