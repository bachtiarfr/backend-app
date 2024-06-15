CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    premium BOOLEAN DEFAULT FALSE
);

CREATE TABLE swipes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    profile_id INTEGER NOT NULL,
    action VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    bio TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (username, email, password, premium) VALUES
    ('user1', 'user1@example.com', '$2a$12$e0MYzXyjpJS7Pd0RVvHwHe/JgI5gQS.s3N2uF9dGeerQozZ9sue8.', FALSE),
    ('user2', 'user2@example.com', '$2a$12$e0MYzXyjpJS7Pd0RVvHwHe/JgI5gQS.s3N2uF9dGeerQozZ9sue8.', FALSE),
    ('premiumUser', 'premium@example.com', '$2a$12$e0MYzXyjpJS7Pd0RVvHwHe/JgI5gQS.s3N2uF9dGeerQozZ9sue8.', TRUE);

INSERT INTO swipes (user_id, profile_id, action) VALUES
    (1, 2, 'like'),
    (2, 1, 'pass');

INSERT INTO profiles (user_id, bio, created_at, updated_at)
VALUES
    (1, 'Bio for user1', NOW(), NOW()),
    (2, 'Bio for user2', NOW(), NOW());