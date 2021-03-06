# Simple REST API with GO-GIN
Simple REST API for resident

## Tech Stack :
- Go 1.15
- Gin Gonic Framework
- PostgreSQL
- Gorm
- JWT
- Sentry
- Fetch API from rajaongkir.com for Province and City data
- redis for cache data from multiple join db query
- gocron for background task
- smtp mail for send mail

### ERD :
![alt text](https://github.com/agussaputran/testcasethree-residentAPI/blob/main/images/readme_image.png?raw=true)

### API Documentation :
<https://intip.in/PSZJ>

### Environment Variables :
```
APP=

DB=
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
DB_SSL=
DB_TIMEZONE=

JWT_SECRET=
RAJAONGKIR_APIKEY=
SENTRY_DSN=

MAIL_SMTP_HOST=
MAIL_SMTP_PORT=
MAIL_EMAIL=
MAIL_PASSWORD=
```