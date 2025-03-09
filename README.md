# 🏷️ Go URL Shortener

A simple **URL Shortener API** built with **Golang**, **Gin**, **PostgreSQL**, and **Redis**.  
It allows users to generate short URLs, store them in a database, and retrieve the original long URLs.

---

## 🚀 Features

- **Shorten URLs** 🔗
- **Redirect to Original URL** 🌍
- **Rate Limiting** 🛑 (Prevents spam)
- **Redis Caching** ⚡ (Faster lookups)
- **Swagger API Docs** 📖
- **Dockerized Deployment** 🐳

---

## 📦 Tech Stack

- **Go** (Gin Framework) 🏗️
- **PostgreSQL** (Database) 🗄️
- **Redis** (Caching) 🚀
- **Docker & Docker Compose** 🐳
- **Swagger** (API Documentation) 📄

---

## 🛠️ Installation & Setup

### **1️⃣ Clone the Repository**

```sh
git clone https://github.com/mohamedelbalshy/go-url-shortener.git
cd go-url-shortener
```

2️⃣ Configure Environment Variables

Copy the .env.example file and update your settings:

cp .env.example .env

3️⃣ Run with Docker Compose

docker-compose up --build

✅ The API will run on: http://localhost:8080
📌 API Endpoints
1️⃣ Shorten a URL

POST /api/v1/url/shorten

Request Body (JSON):

{
"long_url": "https://www.google.com"
}

Response:

{
"status": "success",
"message": "URL shortened successfully",
"data": {
"short_url": "abc123",
"long_url": "https://www.google.com"
}
}

2️⃣ Redirect to Original URL

GET /api/v1/url/{short_url}

🔗 Example: http://localhost:8080/api/v1/url/abc123
👉 Redirects to: https://www.google.com
📖 Swagger Documentation

Once the server is running, access API docs:
📌 Swagger UI:
👉 http://localhost:8080/swagger/index.html
🛡️ Rate Limiting

This API uses Redis-based rate limiting.
Default: 10 requests per minute per IP.
Configured in .env:

RATE_LIMIT=10
RATE_LIMIT_WINDOW=60

🔍 Project Structure

go-url-shortener/
│── modules/
│ ├── url/ # URL Module
│ │ ├── url_controller.go # Handles API requests
│ │ ├── url_service.go # Business logic
│ │ ├── url_repository.go # Database interactions
│ │ ├── url_model.go # Data model
│── middlewares/ # Middleware (Logging, Rate Limiting)
│── utils/ # Utility functions (Logging, Responses)
│── database/ # Database and Redis connection
│── routes/ # API Route Management
│── main.go # Entry point
│── Dockerfile # Docker Build File
│── docker-compose.yml # Docker Compose Config
│── README.md # Project Documentation

🤝 Contributing

    Fork the project
    Create a new branch:

git checkout -b feature-xyz

Commit your changes:

git commit -m "Added feature XYZ"

Push to GitHub:

    git push origin feature-xyz

    Submit a Pull Request (PR)

🔥 Next Steps

    ✅ Deploy to AWS/GCP
    ✅ Add JWT Authentication
    ✅ Implement a Frontend UI

📝 License

This project is MIT Licensed. Feel free to use and modify it! 🚀

---

## **🚀 Next Steps**

Would you like to:
1️⃣ **Set up CI/CD with GitHub Actions?**  
2️⃣ **Deploy this app to AWS or DigitalOcean?**  
3️⃣ **Add unit tests for controllers and services?**

Let me know! 🔥
