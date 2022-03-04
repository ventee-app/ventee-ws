## Ventee Websockets server

A simple Websockets server for [Ventee](https://github.com/ventee-app/ventee-mobile) mobile application

Stack: [Golang](https://go.dev), [Gorilla Websocket](https://github.com/gorilla/websocket)

DEV: http://localhost:9099

HEROKU: https://ventee-ws.herokuapp.com

### Deploy

```shell script
git clone https://github.com/ventee-app/ventee-ws
cd ./ventee-ws
gvm use 1.16
go get
```

### Environment variables

The `.env` file is required (**not** required for [Heroku](https://www.heroku.com)), see [.env.example](./.env.example) for details

### Launch

```shell script
go run ./main.go
```

[AIR](https://github.com/cosmtrek/air) can be used as well

### Heroku

The `release` branch is auto-deployed to [Heroku](https://www.heroku.com)

### License

[MIT](./LICENSE.md)
