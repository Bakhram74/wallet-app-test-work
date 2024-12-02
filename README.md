

Напишите приложение, которое по REST принимает запрос вида
POST api/v1/wallet
{
valletId: UUID,
operationType: DEPOSIT or WITHDRAW,
amount: 1000 }
после выполнять логику по изменению счета в базе данных также есть возможность получить баланс кошелька
GET api/v1/wallets/{WALLET_UUID}
стек:
Golang
Postgresql
Docker
Обратите особое внимание проблемам при работе в конкурентной среде (1000 RPS по одному кошельку). Ни один запрос не должен быть не обработан (50Х error)
приложение должно запускаться в докер контейнере, база данных тоже, вся система должна подниматься с помощью docker-compose
Необходимо покрыть приложение тестами
Решенное задание залить на гитхаб, предоставить ссылку Переменные окружения должны считываться из файла config.env
Все возникающие вопросы по заданию решать самостоятельно, по своему усмотрению.




