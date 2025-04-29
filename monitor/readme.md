# Kepler Node Monitoring Setup Guide

This document provides step-by-step instructions to enable monitoring of your local Kepler node using Prometheus and Grafana.

## Prerequisites
- Docker and Docker Compose installed
- Kepler chain node running locally
- Basic familiarity with terminal commands

## Configuration Steps

### 1. Configure Kepler Telemetry
1. Navigate to Kepler config directory
```bash
cd ~/.kepler/config
```
2. Edit app.toml
```yml
[telemetry]
enabled = true
prometheus-retention-time = 180
```
3. Edit config.toml
```yml
prometheus = true
```

### 2. Start Monitoring Stack
1. Navigate to monitoring directory in Kepler repo:
```bash
cd /path/to/kepler-repo/monitor
```
2. Start Docker containers:
```bash
docker-compose up -d
```
3. Verify services are running:
```bash
docker-compose ps
```
4. Login to grafana with `admin/admin` and set your own password
