output "ecs_cluster_id" {
  value = aws_ecs_cluster.ecs_cluster.id
}

output "alb_dns" {
  value = aws_lb.app_alb.dns_name
}
