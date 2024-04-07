# Spendaro-API

## Description

This is the REST API for the Spendaro application.

**What is Spendaro?**

Spendaro is a web application that focuses on making financial management easier for families and individuals. It allows users to track their expenses, set budgets, and manage their finances in a simple and intuitive way using the "Give Every Dollar a Job" philosophy.

**What is the "Give Every Dollar a Job" philosophy?**

The "Give Every Dollar a Job" philosophy is a budgeting method that involves assigning every dollar of your income to a specific category, such as rent, groceries, or entertainment. This method helps users to be more intentional with their spending and to avoid overspending in certain areas.

[For more information, visit the Spendaro website](https://spendaro.com)

## Technologies

> This only outlines the technologies used in the API. For the full list of technologies used in the Spendaro application, visit the [Spendaro repository](https://github.com/TannerBarcelos/Spendaro)

**Core Technologies**

- Go (Golang)
- Echo (Web framework)
- GORM (ORM)

**Database**

- PostgreSQL (Database)
- Redis (Caching)

**DevOps**

- Docker (Containerization)

**Observability**

- Grafana Loki (Log aggregation - Logs generated using ZeroLog are sent to Grafana Loki)
- Grafana Tempo (Tracing - Traces generated using OpenTelemetry are sent to Grafana Tempo)
- Grafana Cloud (Monitoring - Metrics generated using Prometheus are sent to Grafana Cloud)

## Installation

1. Clone the repository

```bash
git clone
```

2. Install dependencies

```bash
go mod download
```

3. Run the application

```bash
make run-dev

# Docker
make run-dev-docker

# Docker Compose
make dev-compose

```
