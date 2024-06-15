# Astrology Services

Astrology Services is a Go-based application that provides various astrology-related services, including Kundali generation, Ashtakvarga chart service, Dasha predictions, free astrology reports, remedies, dosha analysis, and Kundali matching.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Setup](#setup)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- Kundali generation (for North Indian and South Indian styles)
- Ashtakvarga chart service
- Dasha predictions (Vimshottari)
- Free astrology reports (general, planetary, Vimshottari Dasha, yoga)
- Remedies (Rudraksha suggestions, gemstone recommendations)
- Dosha analysis (Mangalik, Kalsarpa, Sadesati)
- Kundali matching

## Project Structure

astrology-services/
├── cmd/
│ └── main.go
├── config/
│ └── config.go
├── internal/
│ ├── domain/
│ │ ├── models/
│ │ │ ├── kundali.go
│ │ │ ├── ashtakvarga.go
│ │ │ ├── dasha.go
│ │ │ ├── report.go
│ │ │ ├── remedy.go
│ │ │ ├── dosha.go
│ │ │ └── match.go
│ │ └── repositories/
│ │ └── repository.go
│ ├── usecase/
│ │ ├── kundali_usecase.go
│ │ ├── ashtakvarga_usecase.go
│ │ ├── dasha_usecase.go
│ │ ├── report_usecase.go
│ │ ├── remedy_usecase.go
│ │ ├── dosha_usecase.go
│ │ └── match_usecase.go
│ ├── delivery/
│ │ ├── http/
│ │ │ ├── handler.go
│ │ │ ├── kundali_handler.go
│ │ │ ├── ashtakvarga_handler.go
│ │ │ ├── dasha_handler.go
│ │ │ ├── report_handler.go
│ │ │ ├── remedy_handler.go
│ │ │ ├── dosha_handler.go
│ │ │ └── match_handler.go
│ ├── repository/
│ │ ├── mysql/
│ │ │ ├── mysql_kundali_repository.go
│ │ │ ├── mysql_ashtakvarga_repository.go
│ │ │ ├── mysql_dasha_repository.go
│ │ │ ├── mysql_report_repository.go
│ │ │ ├── mysql_remedy_repository.go
│ │ │ ├── mysql_dosha_repository.go
│ │ │ └── mysql_match_repository.go
│ └── di/
│ └── di.go
├── pkg/
│ └── gpt/
│ └── gpt_service.go
├── .env
└── go.mod


## Setup

### Prerequisites

- Go 1.18+
- MySQL

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/astrology-services.git
   cd astrology-services
   ```
2. Create a .env file in the root directory with the following content:
   ```
    env
    Copy code
    DATABASE_URL="username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    SERVER_ADDRESS=":8080"
    GPT_API_KEY="your_openai_api_key"
   ```
3. 
Install dependencies:
```
go mod tidy
```
4. 
Run the application:
```
go run cmd/main.go
```

Usage
The application provides a REST API to interact with various astrology services. Use tools like curl or Postman to test the endpoints.

Endpoints
Kundali Generation

GET /kundali/:id
Ashtakvarga Chart Service

GET /ashtakvarga/:id
Dasha Predictions

GET /dasha/:id
Free Astrology Reports

GET /report/general/:id
GET /report/planetary/:id
GET /report/vimshottari/:id
GET /report/yoga/:id
Remedies

GET /remedy/rudraksha/:id
GET /remedy/gemstone/:id
Dosha Analysis

GET /dosha/:id
Kundali Matching

GET /match/:id


Contributing - 
Contributions are welcome! Please fork the repository and create a pull request.

Fork the Project
Create your Feature Branch (git checkout -b feature/AmazingFeature)
Commit your Changes (git commit -m 'Add some AmazingFeature')
Push to the Branch (git push origin feature/AmazingFeature)
Open a Pull Request