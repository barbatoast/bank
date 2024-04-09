create table if not exists accounts (
    user         text,
    currency     text,
    description  text,
    balance      float
);

create table if not exists transactions (
    user   text,
    id     integer,
    date   text,
    object text,
    amount float
);
