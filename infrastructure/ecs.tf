# ECS Cluster
resource "aws_ecs_cluster" "ecs_cluster" {
  name = var.ecs_cluster_name
}

# ECS Service
resource "aws_ecs_service" "app_service" {
  depends_on = [aws_lb.app_alb]

  name            = var.app_name
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.app_task.arn
  launch_type     = "FARGATE"
  desired_count   = 1

  network_configuration {
    subnets          = [aws_subnet.private_1.id, aws_subnet.private_2.id]
    security_groups  = [aws_security_group.app_sg.id]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.app_tg.arn
    container_name   = var.app_name
    container_port   = var.app_port
  }
}

# Secrets
data "aws_secretsmanager_secret" "movie_search_secrets" {
  name = "movie-search-secrets"
}

data "aws_secretsmanager_secret_version" "movie_search_secrets_version" {
  secret_id = data.aws_secretsmanager_secret.movie_search_secrets.id
}

# ECS Task Definition
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
          name  = "POSTGRES_HOST"
          value = aws_db_instance.postgres.endpoint
        },
        {
          name  = "POSTGRES_PORT"
          value = var.environment_variables["POSTGRES_PORT"]
        },
        {
          name  = "POSTGRES_DATABASE"
          value = var.environment_variables["POSTGRES_DATABASE"]
        },
        {
          name  = "POSTGRES_USERNAME"
          value = var.environment_variables["POSTGRES_USERNAME"]
        },
        {
          name  = "REDIS_HOST"
          value = aws_elasticache_cluster.redis.cache_nodes[0].address
        },
        {
          name  = "REDIS_PORT"
          value = var.environment_variables["REDIS_PORT"]
        },
        {
          name  = "REDIS_DB"
          value = var.environment_variables["REDIS_DB"]
        },
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
          valueFrom = "${data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn}:GOOGLE_CLIENT_ID::"
        },
        {
          name      = "GOOGLE_CLIENT_SECRET"
          valueFrom = "${data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn}:GOOGLE_CLIENT_SECRET::"
        },
        {
          name      = "SESSION_SECRET"
          valueFrom = "${data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn}:SESSION_SECRET::"
        },
        {
          name      = "POSTGRES_PASSWORD"
          valueFrom = "${data.aws_secretsmanager_secret_version.movie_search_secrets_version.arn}:POSTGRES_PASSWORD::"
        }
      ]
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
