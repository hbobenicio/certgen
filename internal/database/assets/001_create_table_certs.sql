CREATE TABLE IF NOT EXISTS ca (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    pkey_pem TEXT NOT NULL,  -- PEM encoded 4096 bit RSA private key
    cert_pem TEXT NOT NULL   -- PEM encoded X.509 Certificate
) STRICT;

-- CREATE TABLE IF NOT EXISTS certificates (
--     id INTEGER PRIMARY KEY,
--     subject_name TEXT NOT NULL
-- ) STRICT;
