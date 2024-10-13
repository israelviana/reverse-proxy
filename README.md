# Reverse Proxy Project with Blocked IPs

This project implements a **reverse proxy** in Go that interacts with a PostgreSQL database to block specific IPs.

## How to Run the Project

To run the project locally using Docker and Docker Compose, follow the instructions below:

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Steps

1. Clone this repository to your local environment.
2. Navigate to the project directory.
3. Run the following command in the terminal to start the services:

   ```bash
   docker-compose up --build
   
This will:

Build the Go application.
  - Start the PostgreSQL database.
  - Run migrations to create the blocked_ips table in the database.
  - Start the reverse proxy that will block specific IPs from accessing the service.
    
### Migrations:
  - During the build process, migrations are automatically executed to create the blocked_ips table in the PostgreSQL database. This table stores the IPs that are not allowed to access     the proxy.

### Blocked IPs
The reverse proxy has a list of blocked IP addresses, which are stored in the PostgreSQL database under the blocked_ips table. These IP addresses are inserted into the database during initialization, and any requests coming from these IPs will be denied by the proxy.

By default, the following IPs are blocked:

192.168.1.1
10.0.0.5
172.16.254.10
192.168.0.10
10.1.2.3
