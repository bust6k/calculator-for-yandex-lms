﻿## ЧТО ДЕЛАЕТ:

● принимает post запрос от
пользователя где хранится какое то арифметическое выражение(калькулятор поддерживает только базовые операции ```/,*,-,+```)


## КАК РАБОТАЕТ:

● проект подключает необходимые файлы.

● проект копирует код из финальной задачи 0 спринта

● код реализовывет обработчик калькулятора по url    ```http://localhost:8080/api/v1/calculate```


● обработчик калькулятора проверяет что запрос является именно post

● обработчик делает проверку того,получилось ли задекодировать тело ответа

● потом, если декодирование получилось неверным или если что то произошло,например ошибка в выражении(```код 422```) то обработчик выдает ошибку

● дальше обработчик конвертирует ответ в string,записывая что запрос 
выполнен успешно

● дальше если запрос выполнил все проверки то обработчик энкодирует его в json



## ПРИМЕРЫ:

код 200 появляется при успешном ответе,пример:

● ```result := "2+2*7"```


код 422 появляется при ошибке в выражении или если тело запроса неверное. пример:

● ```result := "3+q+0"```


код 500 появляется если проблемы с сервеом или какие то другие ошибки,пример:

● ```result := "3+0+500"```

//при этом сервер не может обработать даже правильный запрос если возникает  какая то непридвиденная ошибка


## КАК ЗАПУСТИТЬ ПРОЕКТ: 

● если вы не заете как склонировать проект,воспользуйтесь командой git clone

 ● в моем примере для копирования проекта нужно ввести команду:
 
 ```bash
git clone git@github.com:bust6k/bust6k-calculator-for-yandex-lms.git
```
если у вас не работает,попробуйте скопировать данную команду:
```bash
git clone https://github.com/bust6k/calculator-for-yandex-lms

```


● для того чтобы запустить проект сначала перейдите в директорию Main_Directories(cd Main_Directories),а потом в  CMD(cd CMD)

● потом перейти в директорию Main(cd Main)

● потом запустите команду ```go run .```

для запуска теста перейдите в директорию Tests 

и введите команду ```go test .```

## КАК ЗАПУСТИТЬ И ЗАЙТИ НА САМ СЕРВЕР
● запустите проект

● зайдите в браузер

● введите в поисковой строке 
```http://localhost:8080/api/v1/calculate```

● Если вас приводит на сервер другого человека,попробуйте сначала ввести ```http://localhost:8080/```

● потом добавьте к url ```/api/v1/calculate```

●На худой конец,если даже это не помогло,введите в поисковую строку просто ```localhost:8080```

●потом зайдите на первый попавшиеся сайт

●дальше нажмите на ```open http://localhost:8080```

●После того как вы зашли на сервер добавьте к url ```/api/v1/calculate```

●Если ничего из вышеперечисленного не помогло,обращайтесь в лс ```@java_shcolo``` чтобы персонально обсудить вопрос

●Поздравляю,вы зашли на локальный сервер,теперь вы можете тестировать post запросы

## КАК ТЕСТИРОВАТЬ ПРОГРАММУ

●скачайте git bash(или просто зайдите в него если уже есть)
●если у вас линукс то в терминале введите команду ```sudo apt install git```
● запустите проект

● введите команду  если хотите проверить веб сайи на код 200
 
```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'```

●если хотите посмотреть веб сайт на код 422:

```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2 </"}'```

●Если вы имеете postman вы можете сделать запрос более просто:
 зайдите в postman

нажмите на "new request"

в открывшемся меню измените тип запроса на post
в url замените пустоту на:  ```http://localhost:8080/api/v1/calculate```

дальше нажмите на body

выберете raw

там введите: ``` {"expression" : "введи твое выражение"}```

● Поздравляю,вы сделали запрос на сервер

##  ИНФОРМАЦИЯ О ПАКЕТАХ:



● ```CMD```-пакет,в котором хранится  пакет ```Main```(там главный файл маин) а также калькулятор и его обработчик(имеется ввиду http обработчик)

● ```Tests```-пакет с тестом для обработчика

● ```Main``` - пакет где хранится главный файл


## КОМАНДЫ В ТЕРМИНАЛЕ,КОТОРЫЕ ВАМ МОГУТ ПРИГОДИТЬСЯ(Windows)

```dir``` - команда,для просмотра файлов  и каталогов в текущем каталоге

```cd (название каталога) ``` - команда для перехода в следущий каталог

```more (название файла)```- команда позволяющая посмотреть файл

## КОМАНДЫ В ТЕРМИНАЛЕ,КОТОРЫЕ ВАМ МОГУТ ПРИГОДИТЬСЯ(Linux)
```ls``` - команда,для просмотра файлов  и каталогов в текущем каталоге

```cd (название каталога) ``` - команда для перехода в следущий каталог

```more (название файла)```- команда позволяющая посмотреть файл

## КОМАНДЫ В ТЕРМИНАЛЕ,КОТОРЫЕ ВАМ МОГУТ ПРИГОДИТЬСЯ(Mac Os)
```ls``` - команда,для просмотра файлов  и каталогов в текущем каталоге

```cd (название каталога) ``` - команда для перехода в следущий каталог


```open (название файла)```- команда позволяющая посмотреть файл








