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

variable "domain" {
  default = "ms-api.martishin.com"
}

variable "docker_image" {
  default = "100381574725.dkr.ecr.us-east-1.amazonaws.com/movie-search/server:latest"
}

variable "environment_variables" {
  type = map(string)
  default = {
    PORT                  = "8100"
    POSTGRES_DATABASE     = "moviesearch"
    POSTGRES_USERNAME     = "postgres"
    REDIS_PORT            = "6379"
    REDIS_DB              = "0"
    GOOGLE_CALLBACK_URL   = "https://ms-api.martishin.com/auth/callback?provider=google"
    REDIRECT_URL          = "https://ms.martishin.com/"
    SESSION_COOKIE_DOMAIN = ".martishin.com"
    ENV                   = "production"
    ALLOY_HOST            = "127.0.0.1:8100"
    LOGS_PATH             = "/var/log/movie-search.log"
  }
}

variable "secrets" {
  type    = list(string)
  default = ["GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "SESSION_SECRET"]
}
