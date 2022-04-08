# Golang base

This is a project template for future Golang projects of my own.

## Features

* Simple environment variables picking and setting.
* Basic Docker configuration file to build images for production.
* Testing frameworks added.
* Basic HTTP server configuration on port 8080, with status handler.

## Getting started

### 

### Build the image
```
docker build . -t ariel17/base
```

### Using environment variables file
Add keys to `.env` file:
```
MY_SECRET_KEY1=v4lu3!#
```

Make Docker pick them as follows:
```
docker run --env-file .env ariel17/base
```