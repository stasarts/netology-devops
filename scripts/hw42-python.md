# 1) Добавим в скрипт print(c), чтобы видеть результат.
```bash
$ nano hw42-first.py
#!/usr/bin/env python3
a = 1
b = '2'
c = a + b
print(c)
```
Выполним:
```bash
$ chmod +x hw42-first.py
$ ./hw42-first.py
Traceback (most recent call last):
  File "./hw42-first.py", line 4, in <module>
    c = a + b
TypeError: unsupported operand type(s) for +: 'int' and 'str'
```
Т.о. в текущем виде возникает ошибка, т.к. переменные a и b - разных типов, 'int' и 'str' соответственно.

**Как получить для переменной c значение 12?**  
Нужно привести переменную a к типу str (либо изначально задать ее как str):
```bash
$ nano hw42-first.py
#!/usr/bin/env python3
a = 1
b = '2'
c = str(a) + b
print(c)
$ ./hw42-first.py
12
```

**Как получить для переменной c значение 3?**  
Нужно привести переменную b к типу int (либо изначально задать ее как int):
```bash
$ nano hw42-first.py
#!/usr/bin/env python3
a = 1
b = '2'
c = a + int(b)
print(c)
$ ./hw42-first.py
3
```

# 2) Доработаем скрипт:
* переменная is_change нигде не используется
* чтобы выводились все измененные файлы, нужно убрать выход из цикла "break"
* чтобы получить абсолютный путь до измененного файла, воспользуемся os.path.abspath()
* об изменении файлов в git может говорить и статус renamed, добавим его проверку
**Окончательный скрипт.**
```python
#!/usr/bin/env python3

import os

bash_command = ["cd ~/hw42-python/sysadm-homeworks/", "git status"]
abs_path = os.path.abspath(os.path.expanduser(os.path.expandvars(bash_command[0].replace('cd ', ''))))
result_os = os.popen(' && '.join(bash_command)).read()
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(abs_path+'/'+prepare_result)
    if result.find('renamed') != -1:
        prepare_result = result.replace('\trenamed:    ', '')
        print(abs_path+'/'+prepare_result)
```
**Тестирование.**  
Изменим несколько файлов в локальном репозитории, посмотрим `git status`:
```bash
~/netology/sysadm-homeworks$ git status
On branch master
Your branch is up to date with 'origin/master'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   01-intro-01/README.md
        renamed:    README.md -> README_renamed.md
```
Запустим скрипт:
```bash
$ ~/hw42-second.py
/home/vagrant/netology/sysadm-homeworks/01-intro-01/README.md
/home/vagrant/netology/sysadm-homeworks/README.md -> README_renamed.md
```

# 3) Доработаем скрипт:
* добавим обработку исключения: если нет аргумента (путь до гит репозитория), считать рабочим каталогом текущий
* добавим проверку result_os на пустоту (если выполнение `git status` заканчивается ошибкой - директория не является гит репозиторием)
**Окончательный скрипт.**  
```python
#!/usr/bin/env python3

import os
import sys

try:
    workingDirectory = sys.argv[1]
except IndexError: # IndexError - индекс не входит в диапазон элементов
    workingDirectory = os.getcwd() # os.getcwd() - текущая рабочая директория

bash_command = [f"cd {workingDirectory}", "git status"]
abs_path = os.path.abspath(os.path.expanduser(os.path.expandvars(bash_command[0].replace('cd ', ''))))
result_os = os.popen(' && '.join(bash_command)).read()

if not result_os: # проверка result_os на пустоту
    print('Not a git repository!\nNavigate to git repository\nOr run with path to git repository as arg')
    exit()

for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(abs_path+'/'+prepare_result)
    if result.find('renamed') != -1:
        prepare_result = result.replace('\trenamed:    ', '')
        print(abs_path+'/'+prepare_result)
```
**Тестирование.**  
Посмотрим `git status`:
```bash
~/netology/sysadm-homeworks$ git status
On branch master
Your branch is up to date with 'origin/master'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   01-intro-01/README.md
        modified:   02-git-01-vcs/README.md
        renamed:    README.md -> README_renamed.md
```
Запустим скрипт с указанием пути до гит репозитория и без.
```bash
~/netology/sysadm-homeworks$ ~/hw42-python/hw42-third.py ~/netology/sysadm-homeworks/
/home/vagrant/netology/sysadm-homeworks/01-intro-01/README.md
/home/vagrant/netology/sysadm-homeworks/02-git-01-vcs/README.md
/home/vagrant/netology/sysadm-homeworks/README.md -> README_renamed.md
```
```bash
~/netology/sysadm-homeworks$ ~/hw42-python/hw42-third.py
/home/vagrant/netology/sysadm-homeworks/01-intro-01/README.md
/home/vagrant/netology/sysadm-homeworks/02-git-01-vcs/README.md
/home/vagrant/netology/sysadm-homeworks/README.md -> README_renamed.md
```
```bash
~/netology/sysadm-homeworks$ ~/hw42-python/hw42-third.py ~/netology
fatal: not a git repository (or any of the parent directories): .git
Not a git repository!
Navigate to git repository
Or run with path to git repository as arg
```
Выйдем из директории с гит репозиторием:
```bash
~/netology/sysadm-homeworks$ cd ..
~/netology$ ~/hw42-python/hw42-third.py
fatal: not a git repository (or any of the parent directories): .git
Not a git repository!
Navigate to git repository
Or run with path to git repository as arg
```
```bash
~/netology$ ~/hw42-python/hw42-third.py ~/netology/sysadm-homeworks/
/home/vagrant/netology/sysadm-homeworks/01-intro-01/README.md
/home/vagrant/netology/sysadm-homeworks/02-git-01-vcs/README.md
/home/vagrant/netology/sysadm-homeworks/README.md -> README_renamed.md
```

# 4) При написании скрипта будем использовать:
* socket.gethostbyname(hostname) - переводит имя хоста в формат адреса IPv4
* time.sleep(s) - приостановить выполнение программы на заданное количество секунд

**Окончательный скрипт.**
```python
#!/usr/bin/env python3
import socket
import time

hosts = {'drive.google.com': '0.0.0.1', 'mail.google.com': '0.0.0.2', 'google.com':'0.0.0.3'} # первоначальная инициализация

while True:
    for host in hosts:
        current = hosts[host]
        check = socket.gethostbyname(host)
        if current != check:
            print(f'[ERROR] {host} IP mismatch: {current} {check}')
            hosts[host] = check
        else:
            print(f'{host} {current}')
    time.sleep(3)
```
**Тестирование.**
```bash
$ ~/hw42-python/hw42-fourth.py
[ERROR] drive.google.com IP mismatch: 0.0.0.1 142.251.1.194
[ERROR] mail.google.com IP mismatch: 0.0.0.2 142.250.150.83
[ERROR] google.com IP mismatch: 0.0.0.3 142.251.1.139
drive.google.com 142.251.1.194
[ERROR] mail.google.com IP mismatch: 142.250.150.83 142.250.150.18
[ERROR] google.com IP mismatch: 142.251.1.139 142.251.1.113
drive.google.com 142.251.1.194
mail.google.com 142.250.150.18
google.com 142.251.1.113
drive.google.com 142.251.1.194
mail.google.com 142.250.150.18
google.com 142.251.1.113
[ERROR] drive.google.com IP mismatch: 64.233.165.194 173.194.73.194
mail.google.com 64.233.161.83
google.com 64.233.162.102
drive.google.com 173.194.73.194
mail.google.com 64.233.161.83
google.com 64.233.162.102
drive.google.com 173.194.73.194
mail.google.com 64.233.161.83
[ERROR] google.com IP mismatch: 64.233.162.102 173.194.222.101
drive.google.com 173.194.73.194
mail.google.com 64.233.161.83
[ERROR] google.com IP mismatch: 173.194.222.101 173.194.222.139
```
