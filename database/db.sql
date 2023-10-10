CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(15) UNIQUE,
    CONSTRAINT name_length CHECK (CHAR_LENGTH(first_name) >= 3 AND CHAR_LENGTH(last_name) >= 3),
    CONSTRAINT phone_validation CHECK (phone_number REGEXP '^[+]\\d{12}$')
);