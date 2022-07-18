insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario@email.com", "$2a$10$EngtaIx/.yRXGh0V3QgsXewZ3pxrbNgI4mz4TOxpjx7z7ihCKHE8u"),
("Usuario 2", "usuario_2", "usuario2@email.com", "$2a$10$EngtaIx/.yRXGh0V3QgsXewZ3pxrbNgI4mz4TOxpjx7z7ihCKHE8u"),
("Usuario 3", "usuario_3", "usuario3@email.com", "$2a$10$EngtaIx/.yRXGh0V3QgsXewZ3pxrbNgI4mz4TOxpjx7z7ihCKHE8u"),
("Usuario 4", "usuario_4", "usuario4@email.com", "$2a$10$EngtaIx/.yRXGh0V3QgsXewZ3pxrbNgI4mz4TOxpjx7z7ihCKHE8u");

insert into seguidores (usuario_id, seguidores_id)
values
(1, 2),
(1, 3),
(1, 4),
(2, 1),
(2, 4),
(3, 1),
(4, 2),
(4, 3);