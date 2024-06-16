![Alt text](<WhatsApp Image 2024-06-15 at 12.05.02.jpeg>)
# Astrology Services

Astrology Services is a Go-based application that provides various astrology-related services, including Kundali generation, Ashtakvarga chart service, Dasha predictions, free astrology reports, remedies, dosha analysis, and Kundali matching.

## Table of Contents

- [Features](#features)
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
3. Install dependencies:
```
go mod tidy
```
4. Run the application:
```
go run cmd/main.go
```
5. Use below launch.json configuration in vs code to start debugging

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/main.go",
            "env": {
                "DATABASE_URL": "root:password@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local",
                "SERVER_ADDRESS": ":8080",
                "GPT_API_KEY": "your_openai_api_key"
            }
        }
    ]
}

```

Usage
The application provides a REST API to interact with various astrology services. Use tools like curl or Postman to test the endpoints.

Endpoints
Kundali Generation
```
GET /kundali/:id
```

Ashtakvarga Chart Service
```
GET /ashtakvarga/:id
```

Dasha Predictions
```
GET /dasha/:id
```
Free Astrology Reports
```
GET /report/general/:id
GET /report/planetary/:id
GET /report/vimshottari/:id
GET /report/yoga/:id
```
Remedies
```
GET /remedy/rudraksha/:id
GET /remedy/gemstone/:id
```
Dosha Analysis
```
GET /dosha/:id
```
Kundali Matching
```
GET /match/:id
```


Contributing - 
Contributions are welcome! Please fork the repository and create a pull request.

Fork the Project - 
Create your Feature Branch (git checkout -b feature/AmazingFeature)

Commit your Changes - 
(git commit -m 'Add some AmazingFeature')

Push to the Branch - 
(git push origin feature/AmazingFeature)

Open a Pull Request