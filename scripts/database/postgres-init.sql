-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the 'users' table if it doesn't exist
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert the first user (admin) if it doesn't already exist
INSERT INTO users (username, password)
VALUES ('admin', '$2a$12$FEcwl6m6XDfKM9grMoaVTOi0a45oRf1/FJNzzYeQhreLM3oKXL11G');