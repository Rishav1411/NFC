CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(128),
    reg_number CHAR(9) UNIQUE,
    phone_number CHAR(13) UNIQUE,
    CONSTRAINT name_length CHECK (CHAR_LENGTH(user_name) >= 3),
    CONSTRAINT phone_validation CHECK (phone_number REGEXP '^[+]\\d{12}$'),
    CONSTRAINT reg_validation CHECK (reg_number REGEXP '^\\d{2}[A-Z]{3}\\d{4}$')
);