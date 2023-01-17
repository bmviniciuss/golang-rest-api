CREATE TABLE IF NOT EXISTS comments (
    id uuid,
    slug text,
    author text,
    body text
);

INSERT INTO comments (id, slug, author, body)
values ('58a6f632-b943-4178-b683-8956ada70b92', 'how-to-train-your-dragon', 'Alex', 'First!');
