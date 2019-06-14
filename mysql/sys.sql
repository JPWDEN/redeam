USE sys;

CREATE TABLE Books (
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    publisher VARCHAR(25),
    publish_date DATE,
    rating INT default 0,
    status BOOLEAN NOT NULL
);
