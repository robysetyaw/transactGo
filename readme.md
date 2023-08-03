# TransactGo

# TransactGo - Simple Transaction API

TransactGo is a simple REST API for handling transactions between users and merchants. This API allows users to register, login, and perform transactions, and merchants to manage their businesses. 

## How to Use

#### System Requirements
- Go version 1.16 or higher

#### Steps
1. Clone this repository to your local directory using the git command:
```bash
git clone https://github.com/username/TransactGo.git
```
2. Navigate into the TransactGo directory
```bash
cd TransactGo
```
3. Run the application
```bash
go run app/main.go
```

## Endpoints

### User Registration and Login

- Register: `POST /users`
- Login: `POST /login`

### User Management

- Get User: `GET /users/:username`
- Update User: `PUT /users/:username`
- Delete User: `DELETE /users/:username`

### Merchant Management

- Create Merchant: `POST /merchants`
- Get Merchant: `GET /merchants/:id`
- Update Merchant: `PUT /merchants/:id`
- Delete Merchant: `DELETE /merchants/:id`
- Get All Merchants: `GET /merchants`

### Transaction Management

- Create Transaction: `POST /transactions`
- Get Transaction: `GET /transactions/:id`
- Get All Transactions: `GET /transactions`


## Architecture
TransactGo is built using the Clean Architecture principles. These principles ensure separation of concerns where the software is divided into several circles with specific responsibilities:

Handlers: Define how the HTTP request for a specific route will be handled.
Middleware: Handle common tasks across different handlers like authentication.
Models: Define the basic structure of data.
Repositories: Handle the data layer, provide methods to interact with the data source.
Services: Encapsulate business logic.
This design makes our application:

Independent of UI, Database, Frameworks, and External agencies.
Testable in isolation.
Easier to manage and understand as it's organized around the business logic.

### Folder Structure / Folder Layers

/TransactGo
├── app
│   ├── handler
│   │   ├── user_handler.go
│   │   ├── merchant_handler.go
│   │   └── history_handler.go
│   ├── middleware
│   │   └── auth.go
│   ├── model
│   │   ├── response
│   │   │   ├── template_response.go
│   │   ├── account_model.go
│   │   ├── user_model.go
│   │   ├── merchant_model.go
│   │   └── transaction_model.go
│   ├── repository
│   │   ├── account_repository.go
│   │   ├── user_repository.go
│   │   ├── merchant_repository.go
│   │   └── transaction_repository.go
│   ├── service
│   │   ├── account_service.go
│   │   ├── user_service.go
│   │   ├── merchant_service.go
│   │   └── transaction_service.go
│   └── main.go
├── data
│   ├── accounts.json
│   ├── users.json
│   ├── merchants.json
│   └── transactions.json
├── .gitignore
├── README.md
└── go.mod
