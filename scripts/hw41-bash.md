# 1) Немного усовершенствуем скрипт:
```bash
$ nano hw41-first.sh
#!/usr/bin/env bash
a=1
echo a=$a
b=2
echo b=$b
c=a+b
echo c=$c
d=$a+$b
echo d=$d
e=$(($a+$b))
echo e=$e
Запустим его в bash и прокомментируем результаты:
$ chmod +x hw41-first.sh
$ ./hw41-first.sh
a=1 # переменной a было неявно присвоено целочисленное 1
b=2 # переменной b было неявно присвоено целочисленное 2
c=a+b # переменной с была присвоена строка a+b, т.к. в скрипте не было раскрытия переменных с помощью $
d=1+2 # переменной d была неявно присвоена строка, т.к. нет команды выполнить арифметическое действие
e=3 # переменной e было неявно присвоено целочисленне 3, т.к. была выполнена арифметическая операция с помощью конструкции $(( ))
```

# 2) Для того, чтобы скрипт смог завершиться, нужно добавить выход из цикла:
```bash
while ((1==1))
do
curl https://localhost:4757
if (($? != 0))
then
date >> curl.log
else
break # выход из цикла, когда сервис стал доступен
fi
done
```

# 3) Для проверки доступности ip воспользуемся утилитой netcat(nc) с опциями:
-v: подробный вывод  
-z: сканировать только прослушивающих демонов, не отправляя им никаких данных  
-w1: таймаут 1 сек, если сервис не отвечает дольше, считаем, что он недоступен  

```bash
#!/usr/bin/env bash
ip_to_check=(192.168.0.1 173.194.222.113 87.250.250.242)
counter=1
while ((1==1))
do
  for address in ${ip_to_check[@]}
  do
    echo `date`: `nc -vzw1 $address 80 2>&1` >> avail_ip.log
  done
  if [ "$counter" -eq 5 ]
  then
    break
  fi
((counter++))
sleep 1
done
```

# 4) При удачной проверке ip netcat выводит сообщение:
"Connection to 192.168.0.1 80 port [tcp/smtp] succeeded!"  
Если в выводе netcat нет подстроки "succeeded", будем выходить из бесконечного цикла.

```bash
#!/usr/bin/env bash
ip_to_check=(192.168.0.1 173.194.222.113 87.250.250.242)
succ='succeeded'
while ((1==1))
do
  for address in ${ip_to_check[@]}
  do
    ip_avail=`nc -vzw1 $address 80 2>&1`
    echo `date`: $ip_avail >> avail_ip.log
    if [[ "$ip_avail" == *"$succ"* ]]; then
      continue
    else
      break 2
    fi
  done
sleep 1
done
```