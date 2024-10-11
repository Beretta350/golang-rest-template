-- Create the 'users' table if it doesn't exist
CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()), -- UUID in MySQL
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert the first user (admin) if it doesn't already exist
INSERT INTO users (username, password)
VALUES ('admin', '$2a$12$FEcwl6m6XDfKM9grMoaVTOi0a45oRf1/FJNzzYeQhreLM3oKXL11G')
ON DUPLICATE KEY UPDATE username = username; -- Prevents duplicate inserts