resource "aws_db_subnet_group" "postgres_subnet_group" {
  name       = "postgres-subnet-group"
  subnet_ids = [aws_subnet.private_1.id, aws_subnet.private_2.id]
}

resource "aws_security_group" "postgres_sg" {
  name        = "postgres-sg"
  description = "Allow inbound PostgreSQL access from ECS tasks"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.app_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "postgres" {
  identifier             = "movie-search-db"
  engine                 = "postgres"
  engine_version         = "17.4"
  instance_class         = "db.t3.micro"
  allocated_storage      = 20
  storage_type           = "gp3"
  publicly_accessible    = false
  vpc_security_group_ids = [aws_security_group.postgres_sg.id]
  db_subnet_group_name   = aws_db_subnet_group.postgres_subnet_group.name
  multi_az               = false
  skip_final_snapshot    = true

  username = "postgres"
  password = jsondecode(data.aws_secretsmanager_secret_version.movie_search_secrets_version.secret_string)["POSTGRES_PASSWORD"]
}

output "postgres_endpoint" {
  value = aws_db_instance.postgres.endpoint
}
