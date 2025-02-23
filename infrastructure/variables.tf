variable "aws_region" {
  default = "us-east-1"
}

variable "ecs_cluster_name" {
  default = "movie-search-cluster"
}

variable "app_name" {
  default = "movie-search-server"
}

variable "app_port" {
  default = 8100
}

variable "environment_variables" {
  type = map(string)
  default = {
    PORT                  = "8100"
    GOOGLE_CALLBACK_URL   = "https://ms.martishin.com/auth/callback?provider=google"
    REDIRECT_URL          = "https://ms.martishin.com/"
    SESSION_COOKIE_DOMAIN = ".martishin.com"
    ENV                   = "production"
  }
}

variable "secrets" {
  type    = list(string)
  default = ["GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "SESSION_SECRET"]
}
