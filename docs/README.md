# daily-secretary-bot — Шпаргалки и заметки

## Git

Создать ветку и перейти:
git checkout -b reinit-project

Стратегия: одна ветка = одна задача/фича.
Примеры: feature/onboarding, feature/reminders, fix/tz-parsing, chore/logging.

Коммиты (маленькие и осмысленные):
feat: add viper+pflag config (flags>env>file>defaults)
feat: init project skeleton
fix: correct timezone validation
chore: update dependencies

Базовые команды:
git status
git add .
git commit -m "feat: init project skeleton"
git push origin reinit-project

## Создание папок

PowerShell (Windows):
"config","domain","flows","repo","scheduler","transport\telegram" | ForEach-Object { mkdir "internal\$_" } 

mkdir cmd\bot

mkdir configs

mkdir docs

Bash (Linux/macOS/Git Bash):
mkdir -p cmd/bot internal/{config,domain,flows,repo,scheduler,transport/telegram} configs docs

## Конфиги

Порядок: defaults < file < ENV < flags.

Файл: configs/config.yaml (опционален).
ENV с префиксом DSB_.
Флаги имеют самый высокий приоритет.

Примеры запуска:
go run ./cmd/bot
export DSB_TELEGRAM_TOKEN=123:ABC
go run ./cmd/bot --config ./configs/config.yaml
go run ./cmd/bot --config ./configs/config.yaml --env prod --tz Europe/Madrid --summary weekly

Ключевые поля:
- telegram_token
- app_env: dev|prod
- log_path
- default_tz
- summary_freq: off|daily|weekly

## ENV

Примеры:
DSB_TELEGRAM_TOKEN=123:ABC
DSB_APP_ENV=prod
DSB_DEFAULT_TZ=Europe/Madrid
DSB_LOG_PATH=logs/app.log
DSB_SUMMARY_FREQ=daily

PowerShell:
$env:DSB_TELEGRAM_TOKEN="123:ABC"

cmd.exe:
set DSB_TELEGRAM_TOKEN=123:ABC

Bash:
export DSB_TELEGRAM_TOKEN=123:ABC

## Структура проекта

cmd/bot/main.go
internal/
config/
domain/
flows/
repo/
scheduler/
transport/telegram/
configs/
config.yaml
docs/
README.md

## Чек-лист коммита

[ ] Код компилируется: go build ./...
[ ] Нет случайных изменений (git status)
[ ] Сообщение коммита короткое и понятное
[ ] Нет секретов (токены только в ENV)
[ ] Консистентные имена файлов и пакетов

## Соглашения по веткам

feature/<кратко> — новая фича
fix/<кратко> — багфикс
chore/<кратко> — обслуживание
docs/<кратко> — документация

Примеры:
feature/onboarding
feature/settings
fix/timezone-bug
chore/ci
docs/notes

## Go-команды

    go mod tidy
    go build ./...
    go test ./...
    go run ./cmd/bot
