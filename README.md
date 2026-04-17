# GREEN-API Web UI

Web-интерфейс для работы с GREEN-API WhatsApp API. Приложение предоставляет удобный HTML интерфейс для тестирования и использования методов GREEN-API.

## Возможности

- **Параметры подключения**: Поля для ввода `idInstance` и `ApiTokenInstance`
- **Методы API**:
  - `getSettings` - Получение настроек инстанса
  - `getStateInstance` - Получение состояния инстанса
  - `sendMessage` - Отправка текстовых сообщений через WhatsApp
  - `sendFileByUrl` - Отправка файлов по URL
- **Отображение ответов**: Поле с ответами API в формате JSON
- **Современный UI**: Чистый, адаптивный дизайн

## Требования

- Go 1.25 или выше
- Fiber v2 web framework

## Установка

1. Клонируйте репозиторий:
```bash
git clone git@github.com:Pavel-Vinogradov/green-api-web-ui.git
cd green-api-web-ui
```

2. Установите зависимости:
```bash
go mod download
```

## Запуск приложения

### Локальная разработка

1. Запустите сервер:
```bash
go run cmd/app/main.go
```

2. Откройте браузер и перейдите по адресу:
```
http://localhost:9090
```

### Сборка для продакшена

1. Соберите исполняемый файл:
```bash
go build -o green-api-web-ui cmd/app/main.go
```

2. Запустите приложение:
```bash
./green-api-web-ui
```

## Docker

### Сборка и запуск через Docker

```bash
# Сборка образа
docker build -t green-api-web-ui .

# Запуск контейнера
docker run -d -p 9090:9090 --name green-api-web-ui green-api-web-ui
```

### Запуск через Docker Compose

```bash
# Запуск
docker-compose up -d

# Остановка
docker-compose down

# Просмотр логов
docker-compose logs -f
```

### Настройка окружения

Создайте файл `.env`:
```bash
HTTP_CLIENT_BASE_URL=https://api.green-api.com
HTTP_CLIENT_TIMEOUT=5s
SERVER_PORT=9090
```

## Использование

1. **Создание инстанса GREEN-API**:
   - Зарегистрируйтесь на [GREEN-API](https://green-api.com)
   - Создайте новый инстанс в дашборде
   - Отсканируйте QR-код для подключения WhatsApp

2. **Получение учетных данных**:
   - Скопируйте `idInstance` из дашборда GREEN-API
   - Скопируйте `ApiTokenInstance` из дашборда GREEN-API

3. **Использование веб-интерфейса**:
   - Введите `idInstance` и `ApiTokenInstance` в поля параметров подключения
   - Нажимайте кнопки методов для тестирования различных API вызовов
   - Просматривайте ответы в поле ответов

## Архитектура проекта

Приложение следует принципам чистой архитектуры:

```
green-api-web-ui/
├── cmd/
│   ├── app/
│   │   └── main.go              # Точка входа
│   └── cli/
│       └── app.go               # Инициализация приложения
├── internal/
│   ├── config/                  # Конфигурация
│   │   ├── config.go
│   │   ├── http_client_config.go
│   │   ├── laoder.go
│   │   └── server.go
│   ├── container/               # DI контейнер
│   │   └── container.go
│   ├── delivery/
│   │   └── http/
│   │       └── routes/
│   │           └── routes.go   # Роуты
│   ├── handler/                 # HTTP хендлеры
│   │   ├── api_handler.go
│   │   └── home_handler.go
│   ├── infrastructure/
│   │   └── http/
│   │       └── client.go       # HTTP клиент для Green API
│   ├── interfaces/
│   │   └── green_api/
│   │       └── interface.go     # Интерфейсы use case
│   └── usecase/
│       └── green_api/
│           └── usecase.go      # Бизнес-логика
├── public/                      # Статические файлы
│   ├── index.html
│   ├── css/
│   └── js/
├── configs/
│   └── config.yaml             # Конфигурация приложения
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## API эндпоинты

Приложение предоставляет следующие API эндпоинты:

- `POST /api/getSettings` - Получить настройки инстанса
- `POST /api/getStateInstance` - Получить состояние инстанса
- `POST /api/sendMessage` - Отправить текстовое сообщение
- `POST /api/sendFileByUrl` - Отправить файл по URL

## Деплой на Timeweb

### Подготовка

1. Установите Docker на сервере
2. Скопируйте файлы проекта на сервер
3. Настройте `.env` файл с нужными параметрами

### Деплой

```bash
# Клонирование репозитория
git clone <repository-url>
cd green-api-web-ui

# Запуск через docker-compose
docker-compose up -d

# Проверка статуса
docker-compose ps
```

### Настройка Nginx (опционально)

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:9090;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

## Устранение проблем

### Ошибки подключения
- Убедитесь, что инстанс GREEN-API активен и авторизован
- Проверьте правильность `idInstance` и `ApiTokenInstance`
- Проверьте интернет-соединение

### Сообщения не отправляются
- Убедитесь в правильности формата ID чата (например, `79001234567@c.us`)
- Проверьте, что ваш номер WhatsApp подключен к инстансу
- Убедитесь, что состояние инстанса `authorized`

### Проблемы с конфигурацией
- Проверьте, что файл `configs/config.yaml` существует
- Убедитесь, что viper находит конфигурационный файл
- Проверьте логи при запуске приложения

## Лицензия

Этот проект создан как тестовое задание для GREEN-API.

## Документация GREEN-API

Для получения дополнительной информации о GREEN-API посетите:
- [Официальная документация](https://green-api.com/en/docs/)
- [Справочник API](https://green-api.com/en/docs/api/)