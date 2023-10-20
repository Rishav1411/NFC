# NFC Payment Application APIs

## Overview

The NFC Payment Application APIs provide a set of endpoints for building a secure NFC payment application. These APIs allow users to perform actions such as registration, login, OTP verification, wallet creation, funds transfer, and accessing transaction history.

## Endpoints

### 1. Register a New User

- **POST /sign_up**
  - Registers a new user in the Redis database.

### 2. Login

- **POST /login**
  - Allows users to log into the application and sends an OTP for verification.

### 3. Verify OTP

- **POST /otp**
  - Used to verify OTP for user authentication and generate a JWT token.

### 4. Create Wallet

- **GET /wallet**
  - Allows users to create a wallet.

### 5. Transfer Amount

- **POST /wallet/transfer**
  - Enables users to transfer an amount to another person.

### 6. Transaction History

- **GET /wallet/history**
  - Provides access to the transaction history.

Please refer to the API documentation for detailed request and response formats.

For detailed information about each API endpoint, their request parameters, and responses, refer to the OpenAPI specification.
## Running the NFC Payment Application with Docker

To run the NFC Payment Application using Docker, follow these steps:

1. **Install Docker:** If you don't have Docker installed on your system, you can download and install it from the [official Docker website](https://www.docker.com/get-started).

2. **Pull the Docker Image:** Use the following command to pull the NFC Payment Application Docker image from the public repository:

   ```bash
   docker pull dockerhub-siva0310/nfc:v0.2
