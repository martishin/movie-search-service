data "aws_secretsmanager_secret" "movie_search_secrets" {
  name = "movie-search-secrets"
}

data "aws_secretsmanager_secret_version" "movie_search_secrets_version" {
  secret_id = data.aws_secretsmanager_secret.movie_search_secrets.id
}

variable "docker_image" {
  default = "100381574725.dkr.ecr.us-east-1.amazonaws.com/movie-search/server:latest"
}

resource "aws_ecs_task_definition" "app_task" {
  family                   = var.app_name
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([
    {
      name  = var.app_name
      image = var.docker_image
      portMappings = [
        {
          containerPort = var.app_port
        }
      ],
      environment = [
        {
          name  = "GOOGLE_CALLBACK_URL"
          value = var.environment_variables["GOOGLE_CALLBACK_URL"]
        },
        {
          name  = "REDIRECT_URL"
          value = var.environment_variables["REDIRECT_URL"]
        },
        {
          name  = "SESSION_COOKIE_DOMAIN"
          value = var.environment_variables["SESSION_COOKIE_DOMAIN"]
        },
        {
          name  = "ENV"
          value = var.environment_variables["ENV"]
        },
        {
          name  = "PORT"
          value = var.environment_variables["PORT"]
        }
      ],
      secrets = [
        {
          name      = "GOOGLE_CLIENT_ID"
          valueFrom = data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn
        },
        {
          name      = "GOOGLE_CLIENT_SECRET"
          valueFrom = data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn
        },
        {
          name      = "SESSION_SECRET"
          valueFrom = data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn
        }
      ],
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = "/ecs/${var.app_name}"
          awslogs-region        = "us-east-1"
          awslogs-stream-prefix = "ecs"
        }
      }
    }
  ])
}
