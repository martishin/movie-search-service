# AWS Region (modify if needed)
AWS_REGION=us-east-1

# Terraform commands
init:
	@echo "Initializing Terraform..."
	terraform init

plan:
	@echo "Planning Terraform changes..."
	terraform plan -var="aws_region=$(AWS_REGION)"

apply:
	@echo "Applying Terraform changes..."
	terraform apply -auto-approve -var="aws_region=$(AWS_REGION)"

destroy:
	@echo "Destroying infrastructure..."
	terraform destroy -auto-approve -var="aws_region=$(AWS_REGION)"

format:
	@echo "Formatting Terraform files..."
	terraform fmt -recursive

validate:
	@echo "Validating Terraform files..."
	terraform validate

output:
	@echo "Fetching Terraform outputs..."
	terraform output

force-redeploy:
	@echo "Forcing ECS redeployment..."
	aws ecs update-service --cluster movie-search-cluster --service movie-search-server --force-new-deployment

balancer-address:
	terraform output alb_dns

.PHONY: init plan apply destroy format validate output force-redeploy
