provider "grafana" {
  url  = "https://your-grafana-instance.grafana.net"
  auth = "YOUR_GRAFANA_API_KEY"
}

# 🔥 Notification Channel (Slack Example)
resource "grafana_notification_policy" "slack" {
  name   = "Slack Alert Channel"
  is_default = false
  receivers = [
    {
      name = "slack",
      type = "slack",
      settings = jsonencode({
        webhook_url = "https://hooks.slack.com/services/XXXXX/YYYYY/ZZZZZ"
      })
    }
  ]
}

# 🚀 Alert Rule for High Latency (p95 > 1s for 5 minutes)
resource "grafana_alert_rule" "high_latency" {
  name           = "High Latency Alert"
  folder_uid     = "alerts-folder"
  dashboard_uid  = "fefd5litax0qoa"
  condition      = "B"
  for            = "5m"
  exec_err_state = "alerting"
  no_data_state  = "no_data"

  data {
    ref_id  = "B"
    query   = "histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le, method, route, env))"
    datasource_uid = "grafanacloud-prom"
  }

  notifications {
    uid = grafana_notification_policy.slack.id
  }
}

# 📈 Alert Rule for High Traffic (RPS > 1000)
resource "grafana_alert_rule" "high_traffic" {
  name           = "High Traffic Alert"
  folder_uid     = "alerts-folder"
  dashboard_uid  = "fefd5litax0qoa"
  condition      = "C"
  for            = "2m"
  exec_err_state = "alerting"
  no_data_state  = "no_data"

  data {
    ref_id  = "C"
    query   = "sum(rate(http_requests_total[5m]))"
    datasource_uid = "grafanacloud-prom"
  }

  notifications {
    uid = grafana_notification_policy.slack.id
  }
}

# ❌ Alert Rule for High Error Rate (>5%)
resource "grafana_alert_rule" "high_error_rate" {
  name           = "High Error Rate Alert"
  folder_uid     = "alerts-folder"
  dashboard_uid  = "fefd5litax0qoa"
  condition      = "D"
  for            = "3m"
  exec_err_state = "alerting"
  no_data_state  = "no_data"

  data {
    ref_id  = "D"
    query   = "sum(rate(http_requests_errors_total[5m])) / sum(rate(http_requests_total[5m]))"
    datasource_uid = "grafanacloud-prom"
  }

  notifications {
    uid = grafana_notification_policy.slack.id
  }
}

# ⚠️ Alert Rule for High Saturation (In-Flight Requests > 100)
resource "grafana_alert_rule" "high_saturation" {
  name           = "High In-Flight Requests Alert"
  folder_uid     = "alerts-folder"
  dashboard_uid  = "fefd5litax0qoa"
  condition      = "E"
  for            = "3m"
  exec_err_state = "alerting"
  no_data_state  = "no_data"

  data {
    ref_id  = "E"
    query   = "sum(http_requests_in_flight)"
    datasource_uid = "grafanacloud-prom"
  }

  notifications {
    uid = grafana_notification_policy.slack.id
  }
}
