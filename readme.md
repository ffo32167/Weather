Код в процессе разделения на пакеты
=====================

Основная цель данного приложения это обучение. Поэтому оно имеет следующие особенности:
* многие места упрощены
* основная программа имеет два режима работы: GRPC сервер и консольное приложение
* сторонние пакеты подбирались по принципу «что первое под руку подвернулось»

При этом некоторые важные стороны Go остались неиспользованными:
* работа с каналами
* работа с контекстом
* работа с ошибками
* работа с БД


СТРУКТУРА
=====================
Приложение состоит из двух частей:
HTTP сервер позволяет пользователю логиниться, выбирать параметры запроса GRPC и возвращает через браузер результат пользователю в виде файла.

GRPC сервер получает данные с сайтов, кеширует их и отдает в обработанном виде.
