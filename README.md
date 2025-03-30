# Movie Search
Movie Search is a full-stack web app for browsing and streaming movies.  
It’s built with React.js and Go and uses Terraform for infrastructure provisioning.

You can find the live version [here](https://ms.martishin.com/)!  

<p>
  <img width="250" src="https://github.com/user-attachments/assets/edf425e3-dc2c-4006-a7ed-3a15d5f0d8e0" />
  <img width="250" src="https://github.com/user-attachments/assets/b8ffe155-b31b-4cf1-9578-71973a54f083" />
  <img width="250" src="https://github.com/user-attachments/assets/c05b16a1-6a6e-446b-b808-bbe50e6f97b8" />
</p>

This project was originally built as an example full-stack service for the [Technical Skills for Product Managers](https://verbetcetera.com/tech-skills-course) course.

## 🚀 Running Locally
### Server
* Navigate to server folder: `cd server`
* Create the `.env` file from `.env.example`: `cp .env.example .env`
* Start dependencies (PostgreSQL and Redis): `make start-all`
* Run the server: `make run`
* API will be available at http://localhost:8100/
### Client
* Navigate to client folder: `cd client`
* Install dependencies `npm install`
* Start the client `npm run dev`
* UI will be available at https://localhost:5173/

## ⚙️ Features
### Frontend (React.js, Tailwind CSS, Vite)
- Responsive UI built with React.js and Tailwind CSS
- Client-side routing using React Router
- Browse a list of movies
- View detailed movie information
- Search movies by title or genre
- Like/unlike movies and view your liked list
- Fully responsive layout for all screen sizes
- Tested with Jest and Testing Library

### Backend (Go)
- RESTful API built with Go and the Chi router
- Provide CRUD operations for movies
- Authentication via OAuth and passwords using Goth and Gorilla Sessions
- PostgreSQL persistence with pgx, SQL migrations via golang-migrate
- Caching layer using Redis with go-redis
- Structured logging via Go’s `slog` package
- Monitoring and metrics collection using Prometheus
- Fetches movie posters from TMDB and trailers from YouTube

### DevOps & Infrastructure (Terraform + AWS)
- Dockerized for local development and testing
- CI/CD powered by GitHub Actions and Vercel
- Infrastructure provisioned via Terraform, including:
  - ECS (Fargate) for container orchestration
  - Application Load Balancer (ALB) with custom routing
  - PostgreSQL (RDS) and Redis (ElastiCache) setup
  - VPC, subnets, NAT gateways, and route tables
  - IAM roles, security groups, and policies
- Observability with CloudWatch for logs and alarms
