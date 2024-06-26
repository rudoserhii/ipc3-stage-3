# MINI SHOP API APP
## InternPulse Cohort 3, Stage 3 Project

## Table of Contents

- [MINI SHOP API APP](#mini-shop-api-app)
  - [InternPulse Cohort 3, Stage 3 Project](#internpulse-cohort-3-stage-3-project)
  - [Table of Contents](#table-of-contents)
  - [1. Introduction](#1-introduction)
  - [2. API Documentation](#2-api-documentation)
    - [2.1 How to Call the API](#21-how-to-call-the-api)
  - [3. Setting Up and Running the API (Locally and otherwise)](#3-setting-up-and-running-the-api-locally-and-otherwise)
    - [3.1 Docker Compose Setup](#31-docker-compose-setup)
    - [3.2 Run via Executable](#32-run-via-executable)
  - [4. Some Additional Notes](#4-some-additional-notes)

---

## 1. Introduction

This project is a simple REST compliant shop API written in Go.

## 2. API Documentation

This API is documented [here](https://ips3.obi.ninja/api/v1/docs/index.html) via Swagger Docs.

### 2.1 How to Call the API

The API can be accessed via HTTP requests. It exposes endpoints for various CRUD operations.

Sample API base URL: `http://example.com/api`

**Current [Active]** base URL: `https://ips3.obi.ninja/api/v1`

## 3. Setting Up and Running the API (Locally and otherwise)

### 3.1 Docker Compose Setup

I've created and attached a Docker compose file containing instructions for building and running a Docker image for this API. This file should be enough to quickly get the App up and running.

At the project root run:
   
```sh
docker compose up --build
```

The API should now be accessible at `http://localhost:8081`

### 3.2 Run via Executable
You can head over to the [Release page](https://github.com/obiMadu/ipc3-stage-3/releases) to download the binary for your operating system, and execute in your terminal.


## 4. Some Additional Notes

- This repository contains Github Actions workflows for Continuous Integration.