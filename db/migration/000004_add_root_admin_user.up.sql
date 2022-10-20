INSERT INTO users (username, role, hashed_password, full_name, email)
SELECT 'root.admin',
       7,
       '$2a$10$WUyw9uqyxtURqsKoNitHzeh3r6UoS0Je0mJWOtci.4UzUvs7uJxy.',
       'root.admin',
       'root.admin@e-wallet-app.com'
WHERE NOT EXISTS(
        SELECT 1 FROM users WHERE username = 'root.admin' OR email = 'root.admin@e-wallet-app.com'
    );