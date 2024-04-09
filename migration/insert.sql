INSERT INTO accounts (user, currency, description, balance)
VALUES ("test", "$", "Test account", 75);

INSERT INTO transactions (user, id, date, object, amount)
VALUES ("test", 1, "2020-10-01", "Pocket money", 50);

INSERT INTO transactions (user, id, date, object, amount)
VALUES ("test", 2, "2020-10-03", "Book", -10);

INSERT INTO transactions (user, id, date, object, amount)
VALUES ("test", 3, "2020-10-04", "Sandwich", -5);
