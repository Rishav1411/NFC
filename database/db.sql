CREATE TABLE IF NOT EXISTS users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(128) NOT NULL,
    reg_number CHAR(9) NOT NULL UNIQUE,
    phone_number CHAR(13) NOT NULL UNIQUE,
    CONSTRAINT name_length CHECK (CHAR_LENGTH(user_name) >= 3),
    CONSTRAINT phone_validation CHECK (phone_number REGEXP '^[+]\\d{12}$'),
    CONSTRAINT reg_validation CHECK (reg_number REGEXP '^\\d{2}[A-Z]{3}\\d{4}$')
);

CREATE TABLE IF NOT EXISTS wallet (
    wallet_id INT AUTO_INCREMENT PRIMARY KEY,
    balance INT DEFAULT 100,
    user_id INT NOT NULL UNIQUE,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS transactions(
    trasaction_id INT AUTO_INCREMENT PRIMARY KEY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    amount INT NOT NULL,
    transaction_date DATE DEFAULT (CURRENT_DATE),
    transaction_time TIME DEFAULT (CURRENT_TIME),
    FOREIGN KEY (sender_id) REFERENCES wallet(wallet_id),
    FOREIGN KEY (receiver_id) REFERENCES wallet(wallet_id),
    CONSTRAINT transaction_validation CHECK (sender_id != receiver_id)
);
