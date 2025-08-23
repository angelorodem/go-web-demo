-- Insert mock users
INSERT INTO users (username, email, password_hash) VALUES
  ('alice', 'alice@example.com', 'hash_alice_123'),
  ('bob', 'bob@example.com', 'hash_bob_123'),
  ('charlie', 'charlie@example.com', 'hash_charlie_123');

-- Insert mock posts
INSERT INTO posts (id, user_id, title, content) VALUES
  (1, (SELECT id FROM users WHERE username = 'alice'), 'Hello World', 'This is Alice''s first post.'),
  (2, (SELECT id FROM users WHERE username = 'alice'), 'Another Day', 'Alice writes about her day.'),
  (3, (SELECT id FROM users WHERE username = 'bob'), 'Bob''s Thoughts', 'Bob shares his thoughts.'),
  (4, (SELECT id FROM users WHERE username = 'charlie'), 'Charlie''s Intro', 'Charlie introduces himself.');
