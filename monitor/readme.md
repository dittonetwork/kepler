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

### 3. Steps to Import Dashboard

1. **Access Grafana**

   Open your browser and go to [http://localhost:3000](http://localhost:3000).

2. **Import Dashboard**
   1. In Grafana's left sidebar click 🟌 **Dashboards**.
   2. Click on **Create dashboard** → **Import a dashboard**.
   3. Choose **Upload JSON file**.
   4. Navigate to the following path

      ```
      /path/to/kepler-repo/monitor/grafana/dashboard.json
      ```
      Replace `/path/to/kepler-repo` with your actual repository path:

3. **Configure Data Source**
   
   In the import screen:
     1. Under the **Prometheus** dropdown, select your Prometheus data source.
     2. Click **Import**.
     3. A dashboard named "Ditto" should appear.
