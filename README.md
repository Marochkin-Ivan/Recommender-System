# Recommender-System
## Архитектура проекта
![image](https://user-images.githubusercontent.com/57828271/170818365-e008917d-6083-470d-9df8-eceabc877310.png)

## Алгоритм рекомендации 
#### (на примере некоторой *Киры*, у которой трансфер в *Волгограде*):
1. находит всех пользователей, которые посещали те же достопримечательности в разных городах, что и *Кира*;
2. у "единомышленников" *Киры* ищет достопримечательности, которые они посетели в *Волгограде*;
3. проверияет, посещала ли *Кира* эти достопримечательности:
 * если нет, то прогоняет через фильтр по времени и добавляет к списку рекомендаций
 * если посещала, то эта достопримечательность не добавляется к списку рекомендаций
4. проверяет, пуст ли список рекомендаций:
 * если не пуст, то выводит их
 * если пуст, то подгружает из БД и выводит все достопримечательности из *Волгограда*

#### Итог: *Кира* довольна и может посмотреть, например, Родину-мать и погулять по парку ЦПКиО, пока несколько часов ждет следующего самолета в Дубай.

## Полезная информация:
- Основной сервер в директории **/Recommender-Service**
- Микросервис с алгоритмом рекомендации в директории **/Recommender-Algo**
- В проекте полность настроена поддержка Swagger для быстрого и наглядного тестирования (для обоих микросервисов)