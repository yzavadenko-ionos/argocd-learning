# Go Monitoring App

A simple Go web server dockerized with Docker Compose, including Prometheus and Grafana for monitoring.

## Usage

1. Ensure Docker and Docker Compose are installed.

2. Run the application:
   ```bash
   docker-compose up --build
   ```

3. Access the services:
   - App: http://localhost:3002
   - Prometheus: http://localhost:9090
   - Grafana: http://localhost:3000 (username: admin, password: admin)

## Troubleshooting

- If port 3002 is in use, change the port mapping in docker-compose.yml.
- Ensure Docker daemon is running.


## Docker Registry Secret
   ```bash
      kubectl create secret docker-registry regcred \
      --docker-server=docker.io \
      --docker-username=<your-username> \
      --docker-password=<your-access-token> \
      --docker-email=<your-email>
   ```