variable "default_region" {
  type = string
}

variable "project" {
  type = string
}

variable "db_password" {
  type      = string
  sensitive = true
}

variable "db_socket" {
  type      = string
  sensitive = true
}

variable "openai_api_key" {
  type      = string
  sensitive = true
}
