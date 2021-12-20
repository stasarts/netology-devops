# Домашнее задание к занятию "4.3. Языки разметки JSON и YAML"


## Обязательная задача 1
Мы выгрузили JSON, который получили через API запрос к нашему сервису:
```
    { "info" : "Sample JSON output from our service\t",
        "elements" :[
            { "name" : "first",
            "type" : "server",
            "ip" : 7175 
            }
            { "name" : "second",
            "type" : "proxy",
            "ip : 71.78.22.43
            }
        ]
    }
```
Нужно найти и исправить все ошибки, которые допускает наш сервис  
### Ответ:  
0) С точки зрения синтаксиса json ошибки нет, но будет лучше, если после двоеточия соблюдать пробел перед значением, как во всем файле. (2 строка)
1) С точки зрения синтаксиса json ошибки нет, но не понятно, что это за адрес. (5 строка) Адрес в формате ipv4 следует заключать в кавычки (представлять как строку).
2) Не хватает разделителя-запятой между элементами массива "elements". (6 строка)
3) Ключи должны быть строками и обернуты в кавычки. (9 строка)
4) Значение для ip в 9 строке нужно обернуть в кавычки (сделать строкой), т.к. такой формат для значений не поддерживается в json.  
Исправленный json:
```json
    { "info" : "Sample JSON output from our service\t",
        "elements" : [
            { "name" : "first",
            "type" : "server",
            "ip" : 7175 
            },
            { "name" : "second",
            "type" : "proxy",
            "ip" : "71.78.22.43"
            }
        ]
    }
```

## Обязательная задача 2
В прошлый рабочий день мы создавали скрипт, позволяющий опрашивать веб-сервисы и получать их IP. К уже реализованному функционалу нам нужно добавить возможность записи JSON и YAML файлов, описывающих наши сервисы. Формат записи JSON по одному сервису: `{ "имя сервиса" : "его IP"}`. Формат записи YAML по одному сервису: `- имя сервиса: его IP`. Если в момент исполнения скрипта меняется IP у сервиса - он должен так же поменяться в yml и json файле.

### Ваш скрипт:
```python
#!/usr/bin/env python3
import socket
import time
import json
import yaml

hosts = {'drive.google.com': '0.0.0.1', 'mail.google.com': '0.0.0.2', 'google.com':'0.0.0.3'} # первоначальная инициализация

while True:
    for host in hosts:
        current = hosts[host]
        check = socket.gethostbyname(host)
        if current != check:
            print(f'[ERROR] {host} IP mismatch: {current} {check}')
            with open ('./service.json', 'w') as servicejson:
                json.dump({host : check}, service.json)
            with open ("./services.yaml", "w") as services_yaml:
                yaml.dump(hosts, services_yaml)
            hosts[host] = check
        else:
            print(f'{host} {current}')
            with open ('./service.json', 'w') as service.json:
                json.dump({host : current}, service.json)
            with open ("./services.yaml", "w") as services_yaml:
                yaml.dump(hosts, services_yaml)				
    time.sleep(3)
```

### Вывод скрипта при запуске при тестировании:
```bash
$ ./hw43-second.py
[ERROR] drive.google.com IP mismatch: 0.0.0.1 173.194.222.194
[ERROR] mail.google.com IP mismatch: 0.0.0.2 216.58.210.165
[ERROR] google.com IP mismatch: 0.0.0.3 216.58.209.206
drive.google.com 173.194.222.194
mail.google.com 216.58.210.165
google.com 216.58.209.206
drive.google.com 173.194.222.194
mail.google.com 216.58.210.165
google.com 216.58.209.206
drive.google.com 173.194.222.194
mail.google.com 216.58.210.165
google.com 216.58.209.206
```

### json-файл(ы), который(е) записал ваш скрипт:
```bash
cat ./services.json
```
```json
{"drive.google.com": "173.194.222.194", "mail.google.com": "216.58.210.165", "google.com": "216.58.209.206"}
```

### yml-файл(ы), который(е) записал ваш скрипт:
```bash
cat services.yaml
```
```yaml
drive.google.com: 173.194.222.194
google.com: 216.58.209.206
mail.google.com: 216.58.210.165
```
