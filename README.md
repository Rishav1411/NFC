# NFC Payment Application APIs

This document provides an overview of the NFC Payment Application APIs.

## API Endpoints

### Register a New User

- **Path:** `/sign_up/`
- **Description:** Register a new user in the Redis database.
- **HTTP Method:** POST
- **Input Parameters:**
  - `name` (string): Name of the user (3 to 50 characters).
  - `phone` (string): Phone number in the format '+123456789012'.
  - `reg` (string): Registration format, e.g., '21BCE0980'.
- **Responses:**
  - `201`: User successfully registered in Redis.
  - `400`: Invalid user data.
  - `409`: User already exists.
  - `500`: Server error.

### Login to the Application

- **Path:** `/login/`
- **Description:** Used by the user to log into the application.
- **HTTP Method:** POST
- **Input Parameters:**
  - `phone` (string): Phone number in the format '+123456789012'.
- **Responses:**
  - `200`: OTP sent successfully.
  - `400`: Invalid user data.
  - `404`: User doesn't exist.
  - `500`: Server error.

### Verify OTP

- **Path:** `/otp/`
- **Description:** Used to verify OTP for user authentication.
- **HTTP Method:** POST
- **Input Parameters:**
  - `phone` (string): Phone number in the format '+123456789012'.
  - `otp` (string): OTP (4 digits).
- **Responses:**
  - `200`: JWT token successfully generated.
  - `201`: JWT token successfully generated.
  - `400`: Invalid user data.
  - `403`: Invalid or expired OTP.
  - `500`: Server error.

For detailed information about each API endpoint, their request parameters, and responses, refer to the OpenAPI specification.
