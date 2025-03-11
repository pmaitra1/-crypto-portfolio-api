
# Crypto Portfolio API

This is a Go (Golang) based API for managing a user's cryptocurrency portfolio. The API allows users to register, login, and manage their portfolio items securely with JWT (JSON Web Tokens) for authentication. The backend uses PostgreSQL for data storage and GORM as the ORM.

## Features

- User Registration
- User Login with JWT token generation
- Add, Get, Update, and Delete Portfolio Items
- Secure authentication using JWT tokens

## Prerequisites


- **Postman** or any other tool to make API requests (optional, but recommended).
- **JWT token**: You will need to log in to get a JWT token that will authenticate your requests.
- **Go** installed (for local setup): If you want to run the API locally, you need to have Go installed on your machine.

**Install Go**: [Install Go](https://golang.org/dl/)

- **MySQL** or any database you are using for your application. Make sure you set up a local database for development.


## Running Locally
To run this API locally, follow the steps below:

**Prerequisites**

 - Go (Golang) installed on your machine
 - PostgreSQL installed or use a cloud database (e.g., Heroku PostgreSQL)
 ### Step 1: Clone the Repository
 ```bash
 git clone https://github.com/pmaitra1/-crypto-portfolio-api
```

### Step 1: Install Dependencies
Ensure you have **Go** installed. If not, you can download and install Go from the official site: [Install Go](https://golang.org/dl/).

Once Go is installed, run the following command to install the necessary dependencies for the project:
```bash
go mod tidy
```
### Step 2: Configure the Database

For local development, you'll need a **MySQL** database running on your machine.

-   Create a new database for this project (e.g., `crypto_portfolio`).
### Step 3:Set Up the `.env` File
Create a `.env` file in the root directory of your project and configure the following settings:
```
PORT=8080
DATABASE_URL=your_database_url
JWT_SECRET_KEY=your_secret_key
```
-   **PORT**: This is the port your local server will listen on. The default is `8080`, but you can change it if needed.
-   **DATABASE_URL**: The connection URL for your database (e.g., PostgreSQL, MySQL).
-   **JWT_SECRET_KEY**: A secret key to sign your JWT tokens. Use a secure, random key for production, but for local development, a simple string is fine.

Example for **DATABASE_URL** (PostgreSQL):
```plaintext
DATABASE_URL=postgres://username:password@localhost:5432/crypto_portfolio?sslmode=disable
```
### Step 4: Run the Application
```bash
go run .
```
The API will be available at `http://localhost:8080`.


## API Endpoints

### 1. Register User
**Endpoint**: `POST /register`  
**Description**: Registers a new user.

**Request Body**:
```json
{
  "username": "your_username",
  "password": "your_password"
}
```
**Response**:

```json
{
  "message": "User registered successfully"
}
```
### 2. Login User
**Endpoint**: `POST /login`  
**Description**: Logs in an existing user and returns a JWT token.

**Request Body**:
```json
{
  "username": "your_username",
  "password": "your_password"
}
```
**Response**:

```json
{
  "token": "your_jwt_token",
  "user_id": 1
}
```
### 3. Add an Asset
**Endpoint**: `POST /portfolio`  
**Description**: Adds a new asset to the portfolio.

**Headers**:
```makefile
Authorization: Bearer <JWT_TOKEN>
```


**Request Body**:
```json
{
  "name": "Bitcoin",
  "amount": 1.5,
  "price": 45000,
  "user_id": 1
}
```
**Response**:

```json
{
  "id": 1,
  "name": "bitcoin",
  "amount": 1.5,
  "price": 45000,
  "user_id": 1,
  "created_at": "2025-03-12T00:00:00Z",
  "updated_at": "2025-03-12T00:00:00Z"
}
```
### 4. Get Portfolio Item
**Endpoint**: `GET /portfolio/:id`  
**Description**: Retrieves a portfolio item by ID.

**Headers**:
```makefile
Authorization: Bearer <JWT_TOKEN>
```

**Response**:

```json
{
  "id": 1,
  "name": "bitcoin",
  "amount": 1.5,
  "price": 45000,
  "user_id": 1,
  "created_at": "2025-03-12T00:00:00Z",
  "updated_at": "2025-03-12T00:00:00Z"
}
```
### 5. Update an Asset
**Endpoint**: `PUT /portfolio/:id`  
**Description**: Updates an existing portfolio asset by ID.

**Headers**:
```makefile
Authorization: Bearer <JWT_TOKEN>
```


**Request Body**:
```json
{
  "amount": 2.0,
  "price": 46000
}
```
**Response**:

```json
{
  "id": 1,
  "name": "bitcoin",
  "amount": 2.0,
  "price": 46000,
  "user_id": 1,
  "created_at": "2025-03-12T00:00:00Z",
  "updated_at": "2025-03-12T00:00:00Z"
}
```

### 6. Delete an Asset
**Endpoint**: `DELETE /portfolio/:id`  
**Description**: Deletes an asset by its ID.

**Headers**:
```makefile
Authorization: Bearer <JWT_TOKEN>
```

**Response**:

```json
{
"message": "Asset deleted"
}
```

## Authentication
All routes that require a user to be logged in need a JWT token in the `Authorization` header. The token is generated after the user successfully logs in, and it needs to be included in the headers for the endpoints where the user is expected to be authenticated.

**Example**:
```http
Authorization: Bearer <JWT_TOKEN>
```

## Error Handling

 - **401 Unauthorized**: The token is missing, invalid, or expired.
 -   **403 Forbidden**: The user does not have permission to perform the action.
-   **404 Not Found**: The asset or user does not exist.
-   **400 Bad Request**: Invalid input or missing data in the request.

## Security Considerations

-   Ensure that the **JWT token** is stored securely.
-   Always use **HTTPS** for secure communication

## Conclusion

This API allows users to manage their crypto portfolio by adding, viewing, updating, and deleting assets. Users can authenticate using JWT tokens, and all requests are secured. Follow the steps in this documentation to get started and interact with the API.
