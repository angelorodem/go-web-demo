DELETE FROM posts
WHERE id IN (1, 2, 3, 4);

DELETE FROM users
WHERE email IN (
  'alice@example.com',
  'bob@example.com',
  'charlie@example.com'
);
