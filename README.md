# Simply 

This is an example repository of how you can build a golang rest api project with wire for dependency injection and mvc pattern.

It has simples dependencies

- [Gin](https://gin-gonic.com/)
- [Wire](https://github.com/google/wire)
- [Prometheuse](https://github.com/prometheus)
- [Entgo](https://github.com/ent/ent)
- [Viper](https://github.com/spf13/viper)

# Get started

```
git clone https://github.com/iunary/simply
```

## Setup dependencies

```
go get ./...
```

## Run the project

```
make run
```

Or using [air](https://github.com/cosmtrek/air) for live reload

```
air
```

## Endpoints examples

- Check users list example [http://localhost:8888/users](http://localhost:8888/users)

- Check health check endpoint [http://localhost:8888/health](http://localhost:8888/health)

- Check prometheuse metrics endpoint [http://localhost:8888/metrics](http://localhost:8888/metrics)




