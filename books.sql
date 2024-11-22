CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

INSERT INTO books (title, author, quantity, stock)
VALUES
    ('In Search of Lost Time', 'Marcel Proust', 2, 2),
    ('The Great Gatsby', 'F. Scott Fitzgerald', 5, 5),
    ('War and Peace', 'Leo Tolstoy', 6, 6);