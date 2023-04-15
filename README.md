### Сервис обработки текста

# text-processing-service

Сервис уже задеплоен и работает по адресу http://13.235.247.96:4001.

Сервис реализует базовый функционал и принимает только GET запросы (потому что на момент прохождения экзамена тема POST запросов не усвоена). Из-за этого данные может получать только по query параметрам. В тестах выявлено, что в браузере Firefox максимально допустимая длина значения query параметров для данного сервиса - 195857 символов. В тексте не должно быть символов &, # и других символов которые могут влиять на URL.

## Доступные маршруты

### Подсчет количества символов

```
/number-of-symbols
```

* Принимает один query параметр: ``text`` строкового типа
* Возвращает количество символов в переданном тексте со статусом 200 OK
* В случае не передачи параметра ``text`` или передачи пустой строки, возвращает ошибку 400 Bad Request с телом респонса ниже:
```json

{
    "message": "Provide the 'text' param to perform calculations"
}
```

### Подсчет количества слов

```
/number-of-words
```

* Принимает один query параметр: ``text`` строкового типа
* Возвращает количество слов в переданном тексте со статусом 200 OK
* В случае не передачи параметра ``text`` или передачи пустой строки, возвращает ошибку 400 Bad Request с телом респонса ниже:

```json
{
    "message": "Provide the 'text' param to perform calculations"
}
```

### Подсчет количества символов

```
/reading-time
```
    
* Принимает один query параметр: ``text`` строкового типа
* Возвращает примерное время прочтения переданного текста в минутах со статусом 200 OK
* В случае не передачи параметра ``text`` или передачи пустой строки, возвращает ошибку 400 Bad Request с телом респонса ниже:

```json
{
    "message": "Provide the 'text' param to perform calculations"
}
```

В случае несоответствия маршрута сервис возвращает ошибку 404 NotFound с текстом:

```
404 page not found
```

## Примеры использования

* [Пример 1](http://13.235.247.96:4001/number-of-symbols?text=Lorem%20ipsum%20dolor%20sit%20amet,%20consectetur%20adipiscing%20elit.%20Morbi%20et%20magna%20velit.%20Quisque%20pulvinar%20urna%20et%20tortor%20dignissim,%20vel%20feugiat%20eros%20rhoncus.%20Mauris%20vel%20tempus%20eros,%20in%20fermentum%20ipsum.%20Suspendisse%20ut%20consequat%20justo.%20Integer%20ut%20feugiat%20tortor.%20Phasellus%20ante%20lacus,%20lacinia%20vitae%20mi%20tincidunt,%20porttitor%20vehicula%20mauris.%20Nunc%20vitae%20consequat%20nisi.)
* [Пример 2](http://13.235.247.96:4001/number-of-words?text=Lorem%20ipsum%20dolor%20sit%20amet,%20consectetur%20adipiscing%20elit.%20Morbi%20et%20magna%20velit.%20Quisque%20pulvinar%20urna%20et%20tortor%20dignissim,%20vel%20feugiat%20eros%20rhoncus.%20Mauris%20vel%20tempus%20eros,%20in%20fermentum%20ipsum.%20Suspendisse%20ut%20consequat%20justo.%20Integer%20ut%20feugiat%20tortor.%20Phasellus%20ante%20lacus,%20lacinia%20vitae%20mi%20tincidunt,%20porttitor%20vehicula%20mauris.%20Nunc%20vitae%20consequat%20nisi.)
* [Пример 3](http://13.235.247.96:4001/reading-time?text=Lorem%20ipsum%20dolor%20sit%20amet,%20consectetur%20adipiscing%20elit.%20Morbi%20et%20magna%20velit.%20Quisque%20pulvinar%20urna%20et%20tortor%20dignissim,%20vel%20feugiat%20eros%20rhoncus.%20Mauris%20vel%20tempus%20eros,%20in%20fermentum%20ipsum.%20Suspendisse%20ut%20consequat%20justo.%20Integer%20ut%20feugiat%20tortor.%20Phasellus%20ante%20lacus,%20lacinia%20vitae%20mi%20tincidunt,%20porttitor%20vehicula%20mauris.%20Nunc%20vitae%20consequat%20nisi.)
* [Пример 4](http://13.235.247.96:4001/some-other-endpoint)
