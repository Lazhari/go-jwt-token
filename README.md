# JWT Token Auth with Go

## User table

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email  TEXT NOT NULL UNIQUE,
  password TEXT Not NULL
);
```
