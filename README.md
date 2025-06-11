# JWT Authentication API in Go

A secure RESTful API implementing JWT (JSON Web Token) authentication using Go, Gin, and GORM.

## Features

- User registration and login
- JWT token-based authentication
- Protected routes with middleware
- Secure password hashing
- Database integration with GORM
- Environment variable configuration
- Cookie-based token storage

## Prerequisites

- Go 1.16 or higher
- PostgreSQL (or your preferred database)
- Make sure you have Go modules enabled

## Project Structure

```
.
├── controllers/         # Request handlers
├── initializers/        # Database and environment setup
├── middleware/          # Authentication middleware
├── models/             # Database models
├── router/             # Route definitions
├── errorhandlers/      # Custom error handling
├── main.go            # Application entry point
└── go.mod             # Go module file
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/jwtgolang.git
cd jwtgolang
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=5432
SECRET=your_jwt_secret_key
```

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### 1. Sign Up
- **URL**: `/signup`
- **Method**: `POST`
- **Headers**: 
  - Content-Type: application/json
- **Body**:
```json
{
    "Email": "user@example.com",
    "Password": "your_password"
}
```
- **Response**: Success message on user creation

### 2. Login
- **URL**: `/login`
- **Method**: `POST`
- **Headers**: 
  - Content-Type: application/json
- **Body**:
```json
{
    "Email": "user@example.com",
    "Password": "your_password"
}
```
- **Response**: JWT token in cookie and success message

### 3. Validate Token
- **URL**: `/validate`
- **Method**: `GET`
- **Headers**: 
  - Cookie: Authorization=<token>
  - OR
  - Authorization: Bearer <token>
- **Response**: Success message if token is valid

## Security Features

- Password hashing using bcrypt
- JWT token expiration (10 days)
- Secure cookie settings
- Protected routes with middleware
- Input validation
- Error handling

## Database Schema

### User Model
```go
type User struct {
    gorm.Model
    Email    string `gorm:"unique"`
    Password string
}
```

## Testing with Postman

1. **Sign Up**
   - Create a new POST request to `http://localhost:8080/signup`
   - Set Content-Type header to application/json
   - Add the request body with email and password

2. **Login**
   - Create a new POST request to `http://localhost:8080/login`
   - Set Content-Type header to application/json
   - Add the request body with email and password
   - The response will include a cookie with the JWT token

3. **Validate**
   - Create a new GET request to `http://localhost:8080/validate`
   - Include the Authorization cookie or Bearer token
   - If valid, you'll receive a success message

## Error Handling

The API includes comprehensive error handling for:
- Invalid JSON input
- Database errors
- Authentication failures
- Token validation errors
- Password hashing errors

## Dependencies

- github.com/gin-gonic/gin
- github.com/golang-jwt/jwt
- gorm.io/gorm
- golang.org/x/crypto/bcrypt

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

Your Name - [Your GitHub Profile]

## Acknowledgments

- Gin Web Framework
- GORM
- JWT-Go 