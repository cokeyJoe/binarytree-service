# README
1. Запуск приложения
    Выполнить `go run cmd/main.go`
2. Путь к файлу
    Если необходимо указать путь к файлу с []int, используе флаг -path
    `go run cmd/main.go -path /some/path/to/file.json`
3. Структура файла
    Файл, содержащий начальные данные, должен представлять собой JSON вида
    `{ "ints": [1,2,3]}`
    Где ints - массив целых чисел.
4. Кол-во значений
    Минимальное необходимое количество целых чисел в массиве = 30.