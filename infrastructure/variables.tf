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

variable "alloy_config" {
  default = <<EOT
prometheus.scrape "movie_search_metrics" {
	targets = [{
		__address__ = sys.env("ALLOY_HOST"),
	}]
	forward_to      = [prometheus.remote_write.grafanacloud.receiver]
	job_name        = "movie-search"
	scrape_interval = "60s"

	basic_auth {
		username = sys.env("ALLOY_USERNAME")
		password = sys.env("ALLOY_PASSWORD")
	}
}

prometheus.remote_write "grafanacloud" {
	external_labels = {
		env = sys.env("ENV"),
	}

	endpoint {
		url = sys.env("GRAFANA_CLOUD_PROMETHEUS_URL")

		basic_auth {
			username = sys.env("GRAFANA_CLOUD_USERNAME")
			password = sys.env("GRAFANA_CLOUD_API_KEY")
		}

		queue_config { }

		metadata_config { }
	}
}

loki.source.file "movie_search_logs" {
  targets    = [ {
    __path__ = sys.env("LOGS_PATH"),
  } ]
  forward_to = [loki.relabel.structured_logs.receiver]
  tail_from_end = true
  file_watch {
    min_poll_frequency = "1s"
    max_poll_frequency = "5s"
  }
}

loki.relabel "structured_logs" {
  forward_to = [loki.write.grafana_loki.receiver]

  rule {
    source_labels = ["time"]
    target_label  = "__time__"
    action        = "replace"
  }

  rule {
    source_labels = ["level"]
    target_label  = "level"
  }

  rule {
    source_labels = ["msg"]
    target_label  = "message"
  }

  rule {
    target_label  = "service_name"
    replacement   = "movie-search-service"
  }
}

loki.write "grafana_loki" {
  endpoint {
    url = sys.env("LOKI_URL")

    basic_auth {
      username = sys.env("LOKI_USERNAME")
      password = sys.env("LOKI_API_KEY")
    }
  }
}
EOT
}
