ДЛЯ ЧЕГО ЭТО
=====================
Веб-приложение позволяет сравнивать несколько месяцев погоды в разных городах России. Например сравнить погоду в июне-августе в Москве и Волгодонске.

СТРУКТУРА
=====================
Приложение состоит из двух частей:
1. HTTP сервер позволяет пользователю пройти идентификацию, выбрать параметры запроса и получить через браузер результат в виде csv или json файла.

2. GRPC сервер получает данные с сайтов, кэширует их и отдает в обработанном виде.


Основная цель данного приложения это обучение. 
=====================
Поэтому оно имеет следующие особенности:
* многие места упрощены
* сторонние пакеты подбирались по принципу «что первое под руку подвернулось»(например logrus)
