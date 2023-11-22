# WebsiteHealthChecker

## Description/Описание
Простая программа для проверки доступности веб-сайта. Работает через командную строку.

A simple program for checking website availability. Operates through the command line.

## Usage/Использование
Программу можно собрать из исходного кода в исполняемый файл, для этого по очереди нужно выполнить следующие шаги:

1. Открыть командную строку и ввести команду `git clone https://github.com/ebonibor/WebsiteHealthChecker.git`
2. Перейти в каталог с проектом, используя команду `cd`
3. Находясь в корневом каталоге проекта, ввести команду `go build`. После этого появится исполняемый файл с именем WebsiteHealthChecker

4. Запустите файл: `./WebsiteHealthChecker --u ya.ru`
5. Наблюдайте результат

The program can be compiled from source code into an executable file. Follow these steps in order:

1. Open the command line and enter the command `git clone https://github.com/ebonibor/WebsiteHealthChecker.git`
2. Navigate to the project directory using the `cd` command
3. While in the root directory of the project, enter the command `go build`. After this, an executable file named WebsiteHealthChecker will be created

4. Run the file: `./WebsiteHealthChecker --u google.com`
5. Observe the result

## Available Flags/Доступные флаги
Для указания URL сайта используйте флаг `--u`

Для указания порта подключения (необязательно) используйте флаг `--p`

To specify the URL of the site, use the `--u` flag

To specify the connection port (optional), use the `--p` flag

## Example/Пример
`./WebsiteHealthChecker --u ya.ru`
