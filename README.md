Нужно создать переменую окружения CONFIG_PATH - путь yaml файла
Структура  yaml
# config/local.yaml
env: "local" # Окружение - local, dev или prod
storage_path: "postgres://postgres:postgres@localhost:5432/order" # файл, в котором будет храниться наша БД
http_server: # конфигурация нашего http-сервера
  address: "localhost:9000"
  timeout: 4s
  idle_timeout: 30s
clusterID : "test-cluster"
clientID : "test-client"
natsURL : "nats://localhost:4222"

Для получения order по order_uid используйте строку http://localhost:9000/order?id={order_uid}, где {order_uid} нужно написать order_id который вы хотите получить
Для клиетской части использовалось postman
