# Про проект

Данный небольшой проект был создан в учебных целях, для взаимодействия с goroutines в языке go.

Проект представляет собой сортировку, данные берутся из Json файла.

# Принцип работы сортировки

Сортировка осуществляется следующими шагами:

- Считываются данные из файла [500kb.json](./500kb.json)
- Массив делится на чанки в зависимости от того, сколько goroutines было выставлено
- Происходит сортировка отдельных чанков в своей goroutine
- Происходит слияние чанков в один единый массив, который также отсортировывается