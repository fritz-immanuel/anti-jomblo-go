# Dealls Technical Test (Case Study)
Welcome to the Dealls Technical Test backend service! This project, written in Golang, simulates the functionality of a Tinder-like dating application. This program was developed specifically for a tech test case study conducted by Dealls.


## Run Locally

#### Prerequisites:
- Go installed on your machine. You can download it from the [official Go website](https://go.dev/dl/).
- MySQL server is installed on your machine, and running. [Download here](https://dev.mysql.com/downloads/mysql/)
- Redis already set up on your machine. [Download here](https://redis.io/downloads/)

#### 1. Clone the repository
```bash
  git clone https://github.com/fritz-immanuel/anti-jomblo-go
```

#### 2. Navigate to the project directory:
```bash
  cd anti-jomblo-go
```

#### 3. Install dependencies:
```bash
  go mod tidy
```

#### 4. Run the SQL script in 'db_init.sql' from ***top to bottom***

#### 5. Make sure you have the latest .env file. See below for a more detailed .env information.

#### 6. Run the program:
```bash
go run main.go
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file
```
{
    "SERVER_NAME": "ANTI JOMBLO API Local",
    "APPLICATION_NAME":"ANTI JOMBLO API",

    "PORT_APPS": ":9034",
    "APP_URL": "http://localhost:9034",

    "CONF_ENV_LOCATION": "",

    "DB_CONNECTION_STRING":"<mysql_user>:<mysql_pass>@(localhost:3306)/anti_jomblo?parseTime=true",

    "ACTIVE_WORKER": "0",

    "APPLICATION_VERSION_PATH":".",
    "ANDROID_POS_APP_MINIMUM_VERSION": "1.0.0",
    "IOS_POS_APP_MINIMUM_VERSION": "1.0.0",

    "REDIS_ADDR": "localhost:6379",
    "REDIS_TIME_OUT": "259200",
    "REDIS_DB": "0",
    "REDIS_PASSWORD": "",

    "JWT_TIME_OUT": "259200",

    "VULTR_SECRET_KEY": "",
    "VULTR_ACCESS_KEY": "",
    "VULTR_HOSTNAME": "https://sgp1.vultrobjects.com",
    "VULTR_BUCKET": "",
    "VULTR_REGION": "ap-southeast-1",

    "WHITELISTED_IPS": "0.0.0.0"
}
```


## Tech Stack

**Server:** Golang

**Database:** MySQL


## Optimizations

I recognize that while this program is functional, there's always room for improvement. As evidenced by the provided `.env` file, I've included specific variables aimed at enhancing image handling capabilities. With further dedication and refinement, this backend service has the potential to evolve into an even more robust and efficient solution, catering to a wider array of user needs and requirements.

