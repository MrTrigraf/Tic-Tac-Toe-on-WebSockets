<div align="center">

[![ru](https://img.shields.io/badge/lang-ru-orange?labelColor=blue)](https://github.com/MrTrigraf/Tic-Tac-Toe-on-WebSockets/blob/main/README.md)
[![golang](https://img.shields.io/badge/-1.25%2B-orange?logo=goland&labelColor=blue)](https://go.dev/)
[![gorilla/websocket](https://img.shields.io/badge/gorilla%2Fwebsocket-v1.5.3-orange?labelColor=blue)](https://github.com/gorilla/websocket)

<h1> Tic-Tac-Toe on WebSockets </h1>

</div>

Проект разработан в целях изучения протокола WebSocket, многопользовательская игра "Крестики-Нолики" (Tic-Tac-Toe) на языке Go с
использованием WebSockets для real-time коммуникации между клиентом и сервером.
Архитектура приложения построена по паттерну MVC (Model-View-Controller).



## Установка

### Скачивание проекта и установка зависимости 

``` bash
# git clone project
git clone https://github.com/gorilla/websocket.git

# open the project
cd ./Tic-Tac-Toe-on-WebSockets/

# download dependence
go mod download
```

### Запуска сервера

``` bash
# open server
cd ./cmd/server/

# start go file server
go run ./
```

### Запуска клиента

``` bash
# open server
cd ./cmd/client/

# start go file client
go run ./
```



## Требование

- GO 1.25.6 или выше

## Зависимости

|Пакет|Описание|
|-----|--------|
|[gorilla/websocket](https://github.com/gorilla/websocket)|Websocket-клиент и сервер для Go|

## Структура проекта

``` text 
Tic-Tac-Toe-on-WebSockets/
├── cmd/
│   ├── client/ # точка входа клиента игры
│   └── server/ # точка входа сервера игры
├── internal/
│   ├── controllers/ #  транспортный слой
│   │   ├── http/ # ручки http
│   │   └── websocket/ # реализация websocket
│   ├── domain/ # структуры приложения
│   ├── usecase/ # слой логики игры
│   │   ├── game/
│   │   └── player/
│   └── views/
├── pkg/
│   ├── config/ # файлы конфигурации проекта
│   └── logger/ # файлы логирования 
├── test/
└── README.md
```