https://go.dev/doc/tutorial/

#

### go run .
Запуск программы

### go help 
Справка по командам

### go mod tidy
скачивает все зависимости

### go mod init {название проекта(адрес)}/{название модуля}
инициализация модуля

### go mod edit -replace example.com/greetings=../greetings
замена пути для локальной разработки

### go build
сборка программы в локальную папку

### go install
сборка и установка зависимостей в папку go/bin

### go list -f '{{.Target}}'
в папке с main запускать для получения пути к исполняемому файлу

Add the Go install directory to your system's shell path
export PATH=$PATH:/path/to/your/install/directory

example: export PATH=$PATH:/Users/gently/go/bin

Run your application by simply typing its name "hello" in the terminal

