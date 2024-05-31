<p align="center">
  <img src="https://raw.githubusercontent.com/muhrifqii/curium-go-fiber/master/assets/cm_pertable.png" />
</p>
<h1 align="center"><i>Curium</i> GoFiber (with steroids)</h1>
<p align="center">
  <a href="#">
    <img src="https://img.shields.io/badge/go-1.22.3-blue"/>
  </a>
  <a href="https://github.com/muhrifqii/curium-go-fiber/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/muhrifqii/curium-go-fiber"/>
  </a>
</p>
<p align="center">
  <b>Curium (CM)</b> is a man-made metallic element heavier than anything found naturally on Earth. Because it's radioactive, it's used very carefully in research and in some special tools, like power sources for long space missions. In this case, a long space mission on software engineering project.
</p>

---

### Uncle Bob's Clean Architecture Golang Project Template with Fiber Framework

This project provides a Golang template using the Fiber framework and adheres to Uncle Bob's Clean Architecture principles. This approach promotes loose coupling and testability by separating your application logic into distinct layers. It includes integrations with various tools for development, monitoring, and logging.

## Project Structure
This project has 4 Domain layer :

- Domain Layer
- Repository Layer
- Usecase Layer
- Presentation Layer

## Key Features
- Fiber Framework: High-performance and minimalist web framework for building APIs.
- PostgreSQL: Relational database for storing application data. 
    -  sqlx is used for db queries
- Redis: In-memory data store for caching and fast data access.
- Monitoring Stack:
    - Promtail: Log collector for scraping container logs.
    - Loki: Log storage backend for scalable log management.
    - Grafana: Visualization tool for creating dashboards and analyzing metrics.
- Logging:
    - Zap: Powerful logging library with structured logging capabilities.
    - Lumberjack: Log rotation tool for managing log files.
- JSON Parsing:
    - Sonic: High-performance JSON parser for efficient data handling.
- Development and Build:
    - Makefile: Automates build tasks for easier project management.
    - docker-compose: Manages dependencies and runs services within Docker containers.
    - air: Hot reloading tool for faster development cycles.

## Getting Started

### Run The Application
```bash
# Clone into your workspace
$ git clone git@github.com:muhrifqii/curium-go-fiber.git

#move to project
$ cd curium-go-fiber

# copy the example.env to .env
$ cp .sample.env .env

# Run the application
$ make up

# The hot reload will run locally and the infrastructure will run on docker

# Execute using httpie request in new terminal session (or using any other tools like curl or wget)
$ http GET localhost:8080/up-up-and-ready
```

## Contribution:
We welcome contributions to this project. Feel free to submit pull requests with improvements or additional features.