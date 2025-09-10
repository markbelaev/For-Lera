# CUTE - Telegram бот для жидких моментов 💖

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![Telegram](https://img.shields.io/badge/Telegram-Bot-2CA5E0?style=for-the-badge&logo=telegram)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Clean-FF6B6B?style=for-the-badge)

## 🏗️ Архитектура проекта

Проект реализован с использованием **чистой архитектуры** (Clean Architecture) и следует лучшим практикам Go:

### 🎯 Слои архитектуры:
- **`internal/config`** - конфигурация приложения и загрузка переменных окружения
- **`internal/handlers`** - обработчики Telegram команд и callback'ов  
- **`internal/data`** - доменные данные и бизнес-логика
- **`cmd/bot/main.go`** - точка входа и инициализация приложения

### 🔧 Ключевые принципы:
- **Разделение ответственности** - каждый модуль отвечает за свою зону
- **Инверсия зависимостей** - высокоуровневые модули не зависят от низкоуровневых
- **Тестируемость** - архитектура позволяет легко писать unit-тесты
- **Масштабируемость** - легко добавлять новые функции и команды

## ✨ Возможности

### 💌 Случайные комплименты
Милые и теплые сообщения, которые поднимут настроение

### ✨ Мотивационные цитаты  
Вдохновляющие фразы для поднятия мотивации

### 💫 Игривые сообщения
Флиртующие сообщения (18+)

### ❤️ Интерактивное меню
Удобные кнопки для быстрой навигации

### 📊 Логирование
Детальное логирование всех действий бота

### 🐳 Docker поддержка
Легкий деплой через контейнеры

## 🚀 Быстрый старт

### Предварительные требования
- **Go 1.21+** или **Docker**
- Telegram бот токен от [@BotFather](https://t.me/BotFather)

### 1. Получение токена бота
1. Напишите [@BotFather](https://t.me/BotFather) в Telegram
2. Используйте команду `/newbot`
3. Следуйте инструкциям и получите токен
4. Сохраните токен для следующего шага

### 2. Настройка окружения
Создайте файл `.env` в корне проекта:

```env
# Telegram Bot Token from @BotFather
TOKEN_BOT=your_telegram_bot_token_here
# Optional: Bot settings
BOT_NAME=CUTE
BOT_ADMIN_ID=your_telegram_id
