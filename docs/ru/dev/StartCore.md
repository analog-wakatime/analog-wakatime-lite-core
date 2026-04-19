# Для работы с ядром вам необходимо

---
1. Установить последнюю версию GO Version 1.26.0 (<https://go.dev/dl/>).
2. Клонировать репозиторий на вашу локальную машину с помощью следующей команды:

```bash 
git clone https://github.com/analog-wakatime-lite-core/analog-wakatime-lite-core.git
cd analog-wakatime-lite-core
```

3. Установить необходимые зависимости с помощью следующей команды:

```bash
go mod tidy
```

4. Собрать проект с помощью следующей команды:

```bash
go build -o analog-wakatime-lite-core
```

5. Запустить ядро с помощью следующей команды:

```bash
./analog-wakatime-lite-core
```