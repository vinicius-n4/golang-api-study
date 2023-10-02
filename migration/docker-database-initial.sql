create table items(
    id serial primary key,
    name varchar,
    document varchar
);

INSERT INTO items(name,document) VALUES
    ('Vinicius', '09876543210'),
    ('Nogueira', '98765432109'),
    ('Costa', '87654321098');
