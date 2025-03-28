# Movie Search Service
Example service for the ["Technical Skills for Product Managers"](https://verbetcetera.com/tech-skills-course) course.  
Website with information about movies, that uses React.js on the frontend, Go on the backend, and Terraform to provision AWS infrastructure.    

<p>
  <img width="300" src="https://github.com/user-attachments/assets/edf425e3-dc2c-4006-a7ed-3a15d5f0d8e0" />
  <img width="300" src="https://github.com/user-attachments/assets/b8ffe155-b31b-4cf1-9578-71973a54f083" />
  <img width="300" src="https://github.com/user-attachments/assets/c05b16a1-6a6e-446b-b808-bbe50e6f97b8" />
</p>

You can check the live version [here](https://ms.martishin.com/)!  

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
* User login using password or OAuth
* Browsing a list of movies
* Viewing information for an individual movie
* Ability to like a movie and see a list of liked movies
* Searching movies
* Fetching the movie's poster image from TMDB and trailer from YouTube
* APIs for Movies CRUD operations
* PostgresSQL migrations
* Using Docker to containerize the application and for local testing
* Using Grafana Allure for observability and log collection
* Provisioning AWS infrastructure using Terraform:
  * ECS cluster and Task creation
  * Creation and configuration of a load balancer (ALB) 
  * PostgreSQL (RDS) and Redis (Elasticache) provisioning
  * Setting up a network for the service
  * Configuring security groups and policies
