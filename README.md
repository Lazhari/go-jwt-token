# Structure of JWT

{Header Base 64}.{payload Base 64}.{signature Base 64}

## Header

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

## Payload can carry the claims

```json
{
  "email": "",
  "Issuer": "course"
}
```

## Signature

Signature computed from the Header, Payload and a secret

- Signature can not be reversible

# Database

## User table

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email  TEXT NOT NULL UNIQUE,
  password TEXT Not NULL
);
```
