-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(127) NOT NULL,
    username VARCHAR(63) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar_path TEXT,
    bio VARCHAR(255),
    last_active_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE TYPE chat_type AS ENUM ('personal', 'group');

-- Chats Table
CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(63),
    type chat_type NOT NULL,  -- 'personal' or 'group'
    info TEXT,
    owner_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Chat Members Table (Many-to-Many for users in chats)
CREATE TABLE IF NOT EXISTS chat_members (
    id SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id),
    user_id INT REFERENCES users(id),
    added_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),

    UNIQUE (chat_id, user_id)
);

-- Messages Table
CREATE TABLE IF NOT EXISTS messages (
   id SERIAL PRIMARY KEY,
   sender_id INT REFERENCES users(id),
    chat_id INT REFERENCES chats(id),
    text TEXT,
    reply_to_id INT REFERENCES messages(id),
    forward_from_user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Attachments Table
CREATE TABLE IF NOT EXISTS attachments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mime_type VARCHAR(31) NOT NULL,
    size INT NOT NULL,
    uploaded_by INT REFERENCES users(id),
    unloading_dur_ms INT NOT NULL,  --  personal interest
    created_at TIMESTAMP DEFAULT NOW()
);

-- Message Attachments Table (Many-to-Many for messages and attachments)
CREATE TABLE IF NOT EXISTS message_attachments (
    message_id INT REFERENCES messages(id),
    attachment_id INT REFERENCES attachments(id),
    PRIMARY KEY (message_id, attachment_id)
);

-- Message Seen Table (Tracking message read status)
CREATE TABLE IF NOT EXISTS message_seen (
    chat_id INT REFERENCES chats(id),
    message_id INT REFERENCES messages(id),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (message_id, user_id)
);

-- Blocking Table (Tracks blocked users in chats)
CREATE TABLE IF NOT EXISTS blocking (
   id SERIAL PRIMARY KEY,
   user_id INT REFERENCES users(id),
    blocked_by INT REFERENCES users(id),
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
