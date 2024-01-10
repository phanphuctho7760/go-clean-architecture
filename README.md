
<br />
<div align="center">
  <h3 align="center">Go Clean Architecture</h3>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#deployment">Deployment</a></li>
  </ol>
</details>

# About The Project

This API is for **Simple Demo Clean Architecture in Golang** application.

# Built With

[![Go][Golang]][Golang-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Getting Started

## Prerequisites

1. Install go. You can download the Golang in this [page](https://go.dev/doc/install). You should install version 1.20
2. Install Postgres database. You can download the Postgres in this [page](https://www.postgresql.org/download/). You should install version 14.1

## Installation

### Via `go`

1. You run this command to install packages
   ```sh
   go mod download
   ```
2. Create `.env` file from `.env.example` file.
3. You run this command to start
   ```sh
   go run app/entry/api/main.go
   ```
4. You run this command to start
   ```
   curl --location 'http://localhost:9999/api/v1/migrate' \
   --header 'Content-Type: application/x-www-form-urlencoded' \
   --data-urlencode 'key=yourkeyhere'
   ```
5. You run this command to start
   ```   
   curl --location 'http://localhost:9999/api/v1/user' \
   --header 'Content-Type: application/json' \
   --data '{
       "user_name": "yourusername"
   }'
   ```

### Via `docker`

1. Run by docker
   ```sh
   docker-compose up
   ```
2. Then you can do from step 4 of `Run via Go`

[//]: # (# Deployment)
[//]: # (We setup Gitlab CI/CD for this project. If you merge into develop branch, it will run unit test, linting and deploy automatically.)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://go.dev/doc/
