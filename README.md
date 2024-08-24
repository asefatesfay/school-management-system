# School Management System

## Getting Started

First, run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Basic System Architecture

+------------------+                    +------------------+  
|                  |                    |                  |  
|  React + Next.js | <-- REST API ----> |  Golang + go-chi |  
|    Frontend      |                    |     Backend      |  
|                  |                    |                  |  
+--------+---------+                    +--------+---------+  
         |                                       |  
         |                                       |  
         v                                       v  
+--------+---------+                    +--------+---------+  
|                  |                    |                  |  
|   CloudFront     |                    |  Kubernetes      |  
|     (CDN)        |                    |    Service       |  
|                  |                    |                  |  
+--------+---------+                    +--------+---------+  
         |                                       |  
         |                                       |  
         v                                       v  
+--------+---------+                    +--------+---------+  
|                  |                    |                  |  
|     S3 Bucket    |                    |  Kubernetes      |  
|  (Static Assets) |                    | Deployment for   |  
|                  |                    |  Backend         |  
+--------+---------+                    +--------+---------+  
         |                                       |  
         |                                       |  
         v                                       v  
+--------+---------+                    +--------+---------+  
|                  |                    |                  |  
|  Kubernetes      |                    |  AWS Load        |  
| Deployment for   |                    |  Balancer/ALB    |  
|  Frontend        |                    |                  |  
|                  |                    |                  |  
+--------+---------+                    +--------+---------+  
         |                                       |  
         |                                       |  
         v                                       v  
+--------+---------+                    +--------+---------+  
|                  |                    |                  |  
|   AWS Load       |                    |  AWS RDS         |  
|  Balancer/ALB    |                    |  (PostgreSQL)    |  
|                  |                    |                  |  
+--------+---------+                    +--------+---------+  
         |                                       
         |                                         
         v                                         
+--------+---------+                     
|                  |                           
|  AWS EKS         |                    
|  (Kubernetes)    |                    
|                  |                    
+------------------+                
 
### Frontend (React + Next.js)
* **Components:** User interfaces for students, teachers, parents, and admins.
* **State Management**: Context API or Redux to manage the state.
* **Routing:** Next.js routing for different pages.
* **API Calls**: Axios or Fetch for making API requests to the backend.
* **Static Assets**: Deployed to an S3 bucket.

### CloudFront (CDN)
* **Distribution:** Serves static assets from S3, caching them globally for low-latency access.
* **Edge Locations:** Ensures fast content delivery by caching content at edge locations around the world.
  
### S3 Bucket (Static Assets)
* **Storage:** Stores static assets like images, CSS, and JavaScript files.
* **Integration:** Works with CloudFront to serve static assets efficiently.
  
### Backend (Golang + go-chi)
* **Router:** Defines the RESTful API endpoints.
* **Handlers:** Functions to handle the API requests.
* **Models:** Database models representing the tables.
* **Middleware:** For authentication, authorization, logging, etc.
* **Services:** Business logic and database interaction.
* **Utilities:** Helper functions for tasks like JWT generation, password hashing, etc.

### Database (PostgreSQL on AWS RDS)
* **Database Instance:** Managed PostgreSQL database hosted on AWS RDS.
* **Data Storage:** Stores all application data including user information, student records, teacher data, etc.
* **Security:** Use RDS security groups