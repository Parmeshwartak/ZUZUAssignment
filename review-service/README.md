# 🏨 Review Ingestion Microservice

This Go-based microservice pulls `.jl` review files from AWS S3, validates/parses the reviews, and stores them in PostgreSQL.

## ✅ Features
- Parses JSONL (.jl) reviews from S3
- Validates input, logs malformed lines
- Stores in PostgreSQL (extensible schema)
- Tracks processed files (idempotency)
- Dockerized for deployment
- Cron-compatible

## ✅ Architechture diagram
┌────────────┐     ┌──────────────┐      ┌───────────────┐
│   AWS S3   │ ──▶│ Review Loader│ ──▶ │ PostgreSQL DB │
│ .jl Files  │     │ Microservice │      │   (Relational)│
└────────────┘     └──────────────┘      └───────────────┘
                         ▲
                         │
                 ┌────────────┐
                 │Scheduler / │
                 │  Cron Job  │
                 └────────────┘

## 🔧 Requirements
- Go 1.20+
- Docker
- PostgreSQL or MySQL

## 🛠️ Setup

```bash
git clone git@github.com:Parmeshwartak/ZUZUAssignment.git
cd review-service
cp .env

### **build and run the repo **
docker build -t review-service .
docker run --network=host --env-file .env review-service
for postgres==>
docker run --name my-postgres   -e POSTGRES_USER=postgres   -e POSTGRES_PASSWORD=password   -e POSTGRES_DB=reviews   -p 5432:5432  -d postgres


