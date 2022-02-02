# Домашнее задание к занятию "5.3. Введение. Экосистема. Архитектура. Жизненный цикл Docker контейнера"

---

## Задача 1

Сценарий выполения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность:
запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.

### Ответ

Создадим приветственную страничку Нетологии.
```shell
$ nano index.html
```

```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```

Соберем форк официального образа Nginx с приветственной страничкой Нетологии.  
Dockerfile:
```shell
FROM nginx
COPY index.html /usr/share/nginx/html/index.html
```
```shell
$ sudo docker build -t avt0m8/netology-nginx:v0.1 .
```

Проверим функциональность. Запустим контейнер командой:
```shell
$ sudo docker run --rm -d --name netology-hello -p 8080:80 avt0m8/netology-nginx:v0.1
```
Приветственная страничка должна открываться по адресу: http://127.0.0.1:8080.

Отправим созданный образ в Docker-Hub:
```shell
$ sudo docker login -u avt0m8
Password:
Login Succeeded
$ sudo docker push avt0m8/netology-nginx:v0.1
```

Ссылка на образ: https://hub.docker.com/r/avt0m8/netology-nginx

---

## Задача 2

Посмотрите на сценарий ниже и ответьте на вопрос:
"Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

Сценарий:

- Высоконагруженное монолитное java веб-приложение;
- Nodejs веб-приложение;
- Мобильное приложение c версиями для Android и iOS;
- Шина данных на базе Apache Kafka;
- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;
- Мониторинг-стек на базе Prometheus и Grafana;
- MongoDB, как основное хранилище данных для java-приложения;
- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.

### Ответ

```
Высоконагруженное монолитное java веб-приложение
```
Физический сервер, т.к. высоконагруженное, необходим физический доступ к ресурсами.
```
Nodejs веб-приложение
```
Docker контейнер, широко используется для разворачивания статичных веб-приложений.
```
Мобильное приложение c версиями для Android и iOS
```
Мобильное приложение разворачивается на ОС устройства пользователя с GUI - это несовместимо с Docker.  
С другой стороны, серверная backend часть может быть упакована в Docker контейнер.
```
Шина данных на базе Apache Kafka
```
Docker контейнер. Apache Kafka — распределённый программный брокер сообщений.  
Спроектирован как распределённая, горизонтально масштабируемая система, обеспечивающая наращивание пропускной способности как при росте числа и нагрузки со стороны источников, так и количества систем-подписчиков.  
Т.о. Docker позволит быстро масштабировать инфраструктуру.
```
Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana
```
Возможно применение как виртуальных машин, так и Docker контейнеров.
```
Мониторинг-стек на базе Prometheus и Grafana
```
Удобно разворачивать и масштабировать с помощью Docker контейнеров.
```
MongoDB, как основное хранилище данных для java-приложения
```
БД лучше размещать на физической или виртуальной машине.
Хотя у MongoDB есть официальный образ на Docker Hub. 
Думаю, что зависит от критичности сохранности данных.
```
Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry
```
В результате поиска проблемы обнаружил много примеров развертывания в Docker контейнере.  
[Например.](https://ealebed.github.io/posts/2017/gitlab-gitlab-ci-docker-registry-%D1%81-%D0%BF%D0%BE%D0%BC%D0%BE%D1%89%D1%8C%D1%8E-docker-compose/)

---

## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
- Добавьте еще один файл в папку ```/data``` на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.

### Ответ

Создадим директорию /data в текущей на хостовой машине.
```shell
stasarts@host:~/$ mkdir ./data
```

Скачаем образ centos и запустим из него контейнер в фоновом режиме, подключив к нему ./data.
```shell
stasarts@host:~/$ sudo docker pull centos
stasarts@host:~/$ sudo docker run -it --rm -d --name centos-one -v $(pwd)/data:/data centos
```

Скачаем образ debian и запустим из него контейнер в фоновом режиме, подключив к нему ./data.
```shell
stasarts@host:~/$ sudo docker pull debian
stasarts@host:~/$ sudo docker run -it --rm -d --name debian-two -v $(pwd)/data:/data debian
```

Проверим, что контейнеры работают.
```shell
stasarts@host:~/$ sudo docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED         STATUS         PORTS     NAMES
2d06fe788896   centos    "/bin/bash"   4 minutes ago   Up 4 minutes             centos-one
e07fcf25b18d   debian    "bash"        4 minutes ago   Up 4 minutes             debian-two
```

Подключимся к первому контейнеру с помощью ```docker exec``` и создадим текстовый файл любого содержания в /data.
```shell
stasarts@host:~/$ sudo docker exec -it centos-one bash
[root@2d06fe788896 /]# echo "Hello from centos container" > /data/centos-one 
[root@2d06fe788896 /]# exit
exit
```

Добавим еще один файл в папку ./data на хостовой машине.
```shell
stasarts@host:~/$ echo "Hello from hostmachine" > ./data/host
```

Подключимся во второй контейнер и отобразим листинг и содержание файлов в /data контейнера.
```shell
stasarts@host:~/$ sudo docker exec -it debian-two bash
root@e07fcf25b18d:/# ls -la /data
total 16
drwxrwxr-x 2 1000 1000 4096 Jan 30 13:36 .
drwxr-xr-x 1 root root 4096 Jan 30 13:30 ..
-rw-r--r-- 1 root root   28 Jan 30 13:34 centos-one
-rw-rw-r-- 1 1000 1000   23 Jan 30 13:36 host
root@e07fcf25b18d:/# cat /data/centos-one 
Hello from centos container
root@e07fcf25b18d:/# cat /data/host       
Hello from hostmachine
root@e07fcf25b18d:/# exit
exit
```

---

## Задача 4 (*)

Воспроизвести практическую часть лекции самостоятельно.

Соберите Docker образ с Ansible, загрузите на Docker Hub и пришлите ссылку вместе с остальными ответами к задачам.

### Ответ

Соберем Docker образ с Ansible, загрузим на Docker Hub.  
```shell
$ sudo nano Dockerfile
```
```
FROM alpine:3.14

RUN CARGO_NET_GIT_FETCH_WITH_CLI=1 && \
    apk --no-cache add \
        sudo \
        python3\
        py3-pip \
        openssl \
        ca-certificates \
        sshpass \
        openssh-client \
        rsync \
        git && \
    apk --no-cache add --virtual build-dependencies \
        python3-dev \
        libffi-dev \
        musl-dev \
        gcc \
        cargo \
        openssl-dev \
        libressl-dev \
        build-base && \
    pip install --upgrade pip wheel && \
    pip install --upgrade cryptography cffi && \
    pip install ansible==2.9.24 && \
    pip install mitogen ansible-lint jmespath && \
    pip install --upgrade pywinrm && \
    apk del build-dependencies && \
    rm -rf /var/cache/apk/* && \
    rm -rf /root/.cache/pip && \
    rm -rf /root/.cargo

RUN mkdir /ansible && \
    mkdir -p /etc/ansible && \
    echo 'localhost' > /etc/ansible/hosts

WORKDIR /ansible

CMD [ "ansible-playbook", "--version" ]
```
```shell
$ sudo docker build -t avt0m8/ansible:2.9.24 .
Sending build context to Docker daemon   2.56kB
Step 1/5 : FROM alpine:3.14
3.14: Pulling from library/alpine
97518928ae5f: Pull complete 
Digest: sha256:635f0aa53d99017b38d1a0aa5b2082f7812b03e3cdb299103fe77b5c8a07f1d2
Status: Downloaded newer image for alpine:3.14
 ---> 0a97eee8041e
Step 2/5 : RUN CARGO_NET_GIT_FETCH_WITH_CLI=1 &&     apk --no-cache add         sudo         python3        py3-pip         openssl         ca-certificates         sshpass         openssh-client         rsync         git &&     apk --no-cache add --virtual build-dependencies         python3-dev         libffi-dev         musl-dev         gcc         cargo         openssl-dev         libressl-dev         build-base &&     pip install --upgrade pip wheel &&     pip install --upgrade cryptography cffi &&     pip install ansible==2.9.24 &&     pip install mitogen ansible-lint jmespath &&     pip install --upgrade pywinrm &&     apk del build-dependencies &&     rm -rf /var/cache/apk/* &&     rm -rf /root/.cache/pip &&     rm -rf /root/.cargo
 ---> Running in 04f09e971e5e
...
...
...
Removing intermediate container 04f09e971e5e
 ---> 7d43b1327550
Step 3/5 : RUN mkdir /ansible &&     mkdir -p /etc/ansible &&     echo 'localhost' > /etc/ansible/hosts
 ---> Running in 75d51e48f4c1
Removing intermediate container 75d51e48f4c1
 ---> 0bc1c0a0df7c
Step 4/5 : WORKDIR /ansible
 ---> Running in 3a97f32d0bbe
Removing intermediate container 3a97f32d0bbe
 ---> 6449bb5152b7
Step 5/5 : CMD [ "ansible-playbook", "--version" ]
 ---> Running in e47418f806d6
Removing intermediate container e47418f806d6
 ---> a79b256da2da
Successfully built a79b256da2da
Successfully tagged avt0m8/ansible:2.9.24
```
```shell
$ sudo docker push avt0m8/ansible:2.9.24
The push refers to repository [docker.io/avt0m8/ansible]
77d2cce79d50: Pushed 
280e9c9eedd1: Pushed 
1a058d5342cc: Mounted from library/alpine 
2.9.24: digest: sha256:9689f7b758d3a595f1420d75f33de864e40b9e0518bc26d981418fbbffcfc470 size: 947
```

Ссылка на образ: https://hub.docker.com/r/avt0m8/ansible

---