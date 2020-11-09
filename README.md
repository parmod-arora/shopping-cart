## How to run  

1. Git clone this project `git clone git@github.com:parmod-arora/shopping-cart.git shopping-cart`
2. `cd shopping-cart`
3. Build project `docker-compose build`
4. Run project `docker-compose up`

**Tech Stack**

- Golang 
- Postgres
- Golang Migrate
- SqlBoiler ( ORM library)
- Gorilla MUX (HTTP router)
- Docker

###Run test locally
Choose env local and run `make test`
```
#COMPOSE = docker-compose -f docker-compose.yml
COMPOSE = docker-compose -f docker-compose-local.yml
```

