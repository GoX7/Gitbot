# GitHub Analyzer Bot 📊

GitHub Analyzer Bot is a Telegram bot for analyzing GitHub profiles. It provides quick access to user statistics, repositories, activity, and commits directly inside Telegram.  

GitHub Analyzer Bot — это Telegram-бот для анализа профилей GitHub. Он позволяет получать статистику о пользователях, их репозиториях, активности и коммитах прямо в Telegram.  

---

## Features / Функционал

| Command / Команда | Description (EN)                         | Описание (RU)                                | Example / Пример       |
| ----------------- | ---------------------------------------- | -------------------------------------------- | ---------------------- |
| `/start`          | Start interaction with the bot           | Начало работы с ботом                        | `/start`               |
| `/about`          | Information about the bot                | Информация о боте                            | `/about`               |
| `/profile`        | Get GitHub user profile                  | Получить профиль пользователя GitHub         | `/profile torvalds`    |
| `/search`         | Search for GitHub users                  | Поиск пользователей GitHub                   | `/search openai`       |
| Callback buttons  | Navigation                               | Навигация                                    | Inline buttons in chat |

---

## Key Features / Основные особенности

* GitHub profile **search and analysis**  
  Поиск и **анализ профилей GitHub**
* Repository statistics: **information, starts, forks**  
  Статистика репозиториев: **информация, звезды, форки**
* Telegram **callback buttons** for navigation  
  **Инлайн-кнопки** Telegram для удобной навигации
* **PostgreSQL** for storing user data  
  **PostgreSQL** для хранения данных пользователей
* **Redis** for sessions  
  **Redis** для сессий
* **Docker** and **Makefile** for easy deployment  
  **Docker** и **Makefile** для удобного деплоя

---

## Tech Stack / Технологии

* **Golang** (core language / основной язык)  
* **Telegram Bot API**  
* **PostgreSQL** (Database / база данных)  
* **Redis** (session / сесии)  
* **Docker / Docker Compose**  
* **Makefile** (automation / автоматизация)  

---

## Quick Start / Быстрый старт

1. Clone the repository / Клонируем репозиторий:  

```bash
git clone https://github.com/YourUsername/tgbot.git
cd tgbot
````

2. Configure `.env` file / Настройте `.env` файл:

```env
BOT_TOKEN=your_telegram_token
GITHUB_TOKEN=your_github_token
```

3. Run with Docker / Запуск через Docker:

```bash
docker-compose up --build
```

4. Or run with makefile / или через makefile
```bash
make build
```

---

## Logging / Логирование

All requests and errors are logged for easy debugging and monitoring.
Все запросы и ошибки логируются для удобного дебага и мониторинга.

---

## License / Лицензия

MIT
