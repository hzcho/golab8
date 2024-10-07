create table if not exists users(
    id bigserial primary key,
    name varchar(255) not null,
    age int not null
);

create table if not exists accounts(
    id bigserial primary key,
    login varchar(255) not null,
    pass_hash text not null
);

CREATE TABLE IF NOT EXISTS admins (
    account_id BIGINT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    PRIMARY KEY (account_id)
);

INSERT INTO accounts (id, login, pass_hash) VALUES (1, 'admin', '$2a$10$scpUOQjVXsF7T96p2Re0m.pveN/pDcGw21QmdPfg4tYf/6s.jC8Me');

INSERT INTO admins (account_id) VALUES (1);
