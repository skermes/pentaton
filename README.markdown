Replacement new tab page.

In order to serve anything, you'll need to set the DATABASE_URL environment
variable.  On my machine, it's

    postgres://skermes:skermes@localhost:5432/pentaton?sslmode=disable

Yours may vary.  The database schema should look like

    create table categories (
      id serial primary key,
      name varchar(100)
    )

    create table links (
      id serial primary key,
      url varchar(1000),
      name varchar(100),
      color char(6),
      position integer,
      category integer references categories (id)
    )
