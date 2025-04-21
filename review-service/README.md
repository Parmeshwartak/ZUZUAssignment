# 🏨 Review Ingestion Microservice

This Go-based microservice pulls `.jl` review files from AWS S3, validates/parses the reviews, and stores them in PostgreSQL.

## ✅ Features
- Parses JSONL (.jl) reviews from S3
- Validates input, logs malformed lines
- Stores in PostgreSQL (extensible schema)
- Tracks processed files (idempotency)
- Dockerized for deployment
- Cron-compatible

## 🔧 Requirements
- Go 1.20+
- Docker
- PostgreSQL or MySQL

## 🛠️ Setup

```bash
git clone git@github.com:Parmeshwartak/ZUZUAssignment.git
cd review-service
cp .env.example .env

