resource "google_secret_manager_secret" "heatmap-creator_db_socket" {
  secret_id = "heatmap-creator-db-socket"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "heatmap-creator_db_socket_v1" {
  secret      = google_secret_manager_secret.heatmap-creator_db_socket.name
  secret_data = var.db_socket
}

resource "google_secret_manager_secret_iam_member" "cloud_run_heatmap-creator_db_socket" {
  project   = var.project
  secret_id = google_secret_manager_secret.heatmap-creator_db_socket.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.heatmap-creator_cloud_run.email}"
}

resource "google_secret_manager_secret" "heatmap-creator_openai_api_key" {
  secret_id = "heatmap-creator-openai-api-key"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "heatmap-creator_openai_api_key_v1" {
  secret      = google_secret_manager_secret.heatmap-creator_openai_api_key.name
  secret_data = var.openai_api_key
}

resource "google_secret_manager_secret_iam_member" "cloud_run_heatmap-creator_openai_api_key" {
  project   = var.project
  secret_id = google_secret_manager_secret.heatmap-creator_openai_api_key.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.heatmap-creator_cloud_run.email}"
}
