Работа в Windows Terminal.  
Хост машина (stasarts@avt0m8): Win 10 c wsl2 (Ubuntu20.04).  
Виртуальная машина (vagrant@Ubuntu20-04): Vagrant-VirtualBox-Ubuntu20.04.

# 1) Создадим виртуальную машину.
```bash
stasarts@avt0m8:~$ vagrant init bento/ubuntu-20.04
stasarts@avt0m8:~$ nano Vagrantfile
```
```ruby
Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-20.04"
  config.vm.hostname = "Ubuntu20-04"
  config.vm.network "forwarded_port", guest: 443, host: 443
  config.vm.provider :virtualbox do |vb|
    vb.name = "Ubuntu20-04"
      vb.memory = 2048
      vb.cpus = 2
  end
end
```
```bash
stasarts@avt0m8:~$ vagrant up
stasarts@avt0m8:~$ vagrant ssh
```
Подготовим систему.
```bash
vagrant@Ubuntu20-04:~$ apt update
vagrant@Ubuntu20-04:~$ apt upgrade
```
# 2) Установим firewall ufw и откроем порты 22 и 443. Настроим firewall.
```bash
vagrant@Ubuntu20-04:~$ apt install ufw
vagrant@Ubuntu20-04:~$ ufw allow 22
vagrant@Ubuntu20-04:~$ ufw allow 443
vagrant@Ubuntu20-04:~$ ufw enable
```
# 3) Установим hashicorp vault.
- Add the HashiCorp GPG key.  
```bash
vagrant@Ubuntu20-04:~$ curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
```
- Add the official HashiCorp Linux repository.
```bash
vagrant@Ubuntu20-04:~$ apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
```
- Update and install.  
```bash
vagrant@Ubuntu20-04:~$ apt-get update && sudo apt-get install vault
```
- Verifying the Installation.
```bash
vagrant@Ubuntu20-04:~$ vault

Usage: vault <command> [args]

Common commands:
    read        Read data and retrieves secrets
    write       Write data, configuration, and secrets
    delete      Delete secrets and configuration
    list        List data or secrets
    login       Authenticate locally
    agent       Start a Vault agent
    server      Start a Vault server
    status      Print seal and HA status
    unwrap      Unwrap a wrapped secret

Other commands:
    audit          Interact with audit devices
    auth           Interact with auth methods
    debug          Runs the debug command
    kv             Interact with Vault's Key-Value storage
    lease          Interact with leases
    monitor        Stream log messages from a Vault server
    namespace      Interact with namespaces
    operator       Perform operator-specific tasks
    path-help      Retrieve API help for paths
    plugin         Interact with Vault plugins and catalog
    policy         Interact with policies
    print          Prints runtime configurations
    secrets        Interact with secrets engines
    ssh            Initiate an SSH session
    token          Interact with tokens
```
# 4) Cоздадим центр сертификации по инструкции и выпустим сертификат для использования его в настройке веб-сервера nginx (срок жизни сертификата - месяц).
Подготовим систему.
```bash
vagrant@Ubuntu20-04:~$ apt install jq
```
Поправим конфигурационный файл для Vault
```bash
vagrant@Ubuntu20-04:~$ sudo nano /etc/vault.d/vault.hcl
# отключаем ssl
# HTTP listener
listener "tcp" {
  address = "127.0.0.1:8200"
  tls_disable = 1
}

# HTTPS listener
#listener "tcp" {
#  address       = "0.0.0.0:8200"
#  tls_cert_file = "/opt/vault/tls/tls.crt"
#  tls_key_file  = "/opt/vault/tls/tls.key"
#}
```
Инициализируем сервис с сохранением ключей в файл /etc/vault.d/init.file (для удобства, не безопасно).
```bash
vagrant@Ubuntu20-04:~$ vault operator init -n 5 -t 3 | sudo tee /etc/vault.d/init.file
Unseal Key 1: X4b0zlKFiudDbpIy+YZCV3okEzwzCLh0K3CLrGtkEejY
Unseal Key 2: ooZFWzyq0eLytLh6iLqnTp/VzSQCaQQRb+LxvFW+z3tW
Unseal Key 3: L4L1uSMQXKItgqEJkHAzIp7TQCjy8KeGeVI2F2BbPVyD
Unseal Key 4: kl6IH/c5giSj1VzNEoxAHIXErlA5aMN1DXwketG5fefu
Unseal Key 5: kvh5nPazCrCdZXtfS7mkyMqyHxhlte9FhTw9vSjZrQqT

Initial Root Token: s.op4wLup0V0P8UmcTBBgpvP6O

Vault initialized with 5 key shares and a key threshold of 3. Please securely
distribute the key shares printed above. When the Vault is re-sealed,
restarted, or stopped, you must supply at least 3 of these keys to unseal it
before it can start servicing requests.

Vault does not store the generated master key. Without at least 3 keys to
reconstruct the master key, Vault will remain permanently sealed!

It is possible to generate new unseal keys, provided you have a quorum of
existing unseal keys shares. See "vault operator rekey" for more information.
```
где:  
-n (-key-share) — Количество общих ключей, на которые нужно разделить сгенерированный главный ключ. Это количество «ключей распечатки», которое нужно сгенерировать.  
-t (-key-threshold) — Количество общих ключей, необходимых для восстановления главного ключа. Это должно быть меньше или равно -key-share.  
Добавим переменные среды:
```bash
vagrant@Ubuntu20-04:~$ export VAULT_ADDR=http://127.0.0.1:8200
vagrant@Ubuntu20-04:~$ export VAULT_TOKEN=s.op4wLup0V0P8UmcTBBgpvP6O
```
Распечатываем хранилище.
```bash
vagrant@Ubuntu20-04:~$ vault operator unseal X4b0zlKFiudDbpIy+YZCV3okEzwzCLh0K3CLrGtkEejY
vagrant@Ubuntu20-04:~$ vault operator unseal ooZFWzyq0eLytLh6iLqnTp/VzSQCaQQRb+LxvFW+z3tW
vagrant@Ubuntu20-04:~$ vault operator unseal L4L1uSMQXKItgqEJkHAzIp7TQCjy8KeGeVI2F2BbPVyD
Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
Total Shares    5
Threshold       3
Version         1.9.2
Storage Type    file
Cluster Name    vault-cluster-d0e2e8e5
Cluster ID      2d4e4ac9-2eb5-0edb-f8fe-d9899242bb27
HA Enabled      false
```
Активируем PKI тип секрета для корневого центра сертификации.
```bash
vagrant@Ubuntu20-04:~$ vault secrets enable \
    -path=pki_root_ca \
    -description="PKI Root CA" \
    -max-lease-ttl="87600h" \
    pki
Success! Enabled the pki secrets engine at: pki_root_ca/
```
Создаем корневой сертификат центра сертификации (CA). 87600h = 5 лет.
```bash
$ vault write -format=json pki_root_ca/root/generate/internal \
    common_name="Root Certificate Authority" \
    country="RU" \
    locality="Yekaterinburg" \
    street_address="Lenina" \
    postal_code="624800" \
    organization="Avt0m8 LLC" \
    ou="DevOops" \
    ttl="87600h" > pki-root-ca.json
```
Сохраняем корневой сертификат. В дальнейшем именно его надо распространять в организации и делать доверенным.
```bash
vagrant@Ubuntu20-04:~$ cat pki-root-ca.json | jq -r .data.certificate > rootCA.pem
```
Публикуем URL’ы для корневого центра сертификации.
```bash
vagrant@Ubuntu20-04:~$ vault write pki_root_ca/config/urls \
    issuing_certificates="http://127.0.0.1:8200/v1/pki_root_ca/ca" \
    crl_distribution_points="http://127.0.0.1:8200/v1/pki_root_ca/crl"
Success! Data written to: pki_root_ca/config/urls
```
Активируем PKI тип секрета для промежуточного центра сертификации.
```bash
vagrant@Ubuntu20-04:~$ vault secrets enable \
    -path=pki_int_ca \
    -description="PKI Intermediate CA" \
    -max-lease-ttl="43800h" \
    pki
Success! Enabled the pki secrets engine at: pki_int_ca/
```
Генерируем запрос на выдачу сертификата для промежуточного центра сертификации.
```bash
vagrant@Ubuntu20-04:~$ vault write -format=json pki_int_ca/intermediate/generate/internal \
    common_name="Intermediate CA" \
    country="RU" \
    locality="Yekaterinburg" \
    street_address="Lenina" \
    postal_code="624800" \
    organization="Avt0m8" \
    ou="DevOops" \
    ttl="43800h" | jq -r '.data.csr' > pki_intermediate_ca.csr
```
Отправляем полученный CSR-файл в корневой центр сертификации, получаем сертификат для промежуточного центра сертификации. 43800h = 5 лет.
```bash
vagrant@Ubuntu20-04:~$ vault write -format=json pki_root_ca/root/sign-intermediate csr=@pki_intermediate_ca.csr \
    country="RU" \
    locality="Yekaterinburg" \
    street_address="Lenina" \
    postal_code="624800" \
    organization="Avt0m8" \
    ou="DevOops" \
    format=pem_bundle \
    ttl="43800h" | jq -r '.data.certificate' > intermediateCA.cert.pem
```
Публикуем подписанный сертификат промежуточного центра сертификации.
```bash
vagrant@Ubuntu20-04:~$ vault write pki_int_ca/intermediate/set-signed certificate=@intermediateCA.cert.pem
Success! Data written to: pki_int_ca/intermediate/set-signed
```
Публикуем URL’ы для промежуточного центра сертификации.
```bash
vagrant@Ubuntu20-04:~$ vault write pki_int_ca/config/urls \
    issuing_certificates="http://127.0.0.1:8200/v1/pki_int_ca/ca" \
    crl_distribution_points="http://127.0.0.1:8200/v1/pki_int_ca/crl"
Success! Data written to: pki_int_ca/config/urls
```
Создаем роль, с помощью которой будем выдавать сертификаты.
```bash
vagrant@Ubuntu20-04:~$ vault write pki_int_ca/roles/example-dot-com \
    country="RU" \
    locality="Yekaterinburg" \
    street_address="Lenina" \
    postal_code="624800" \
    organization="Avt0m8" \
    ou="DevOops" \
    allowed_domains="example.com" \
    allow_subdomains=true \
    max_ttl="720h"
Success! Data written to: pki_int_ca/roles/example-dot-com
```
```bash
vagrant@Ubuntu20-04:~$ vault write -format=json pki_int_ca/issue/example-dot-com \
common_name="vault.example.com" \
ttl="720h" > vault.example.com.crt
```
Сохраняем сертификат в правильном формате.
```bash
vagrant@Ubuntu20-04:~$ cat vault.example.com.crt | jq -r .data.certificate > vault.example.com.crt.pem
vagrant@Ubuntu20-04:~$ cat vault.example.com.crt | jq -r .data.issuing_ca >> vault.example.com.crt.pem
vagrant@Ubuntu20-04:~$ cat vault.example.com.crt | jq -r .data.private_key > vault.example.com.crt.key
```
# 5) Установим корневой сертификат созданного центра сертификации в доверенные в хостовой системе.  
Скопируем корневой сертификат центра сертификации (CA) на хостовую машину.  
Воспользуемся плагином vagrant scp.  
```bash
stasarts@avt0m8:~$ vagrant plugin install vagrant-scp
stasarts@avt0m8:~$ vagrant scp /mnt/c/Users/stasarts/netology/coursework/rootCA.pem default:/home/vagrant/rootCA.pem
```
На хостовой машине (в Win10) добавим корневой сертификат центра сертификации в браузер Edge.  
"Settings and more" -> "Settings" -> "Privacy, search, and services" -> "Manage HTTPS/SSL certificates and settings" -> "Trusted Root Certification Authorities" -> "Import"

# 6) Установим nginx на ВМ.
```bash
vagrant@Ubuntu20-04:~$ sudo apt install nginx
```
# 7) Настроим nginx на https, используя ранее подготовленный сертификат.
```bash
vagrant@Ubuntu20-04:~$ sudo nano /etc/nginx/nginx.conf
```
```bash
user www-data;
worker_processes auto;
pid /run/nginx.pid;
include /etc/nginx/modules-enabled/*.conf;

events {
        worker_connections 768;
        # multi_accept on;
}

http {

        ##
        # Basic Settings
        ##

        sendfile on;
        tcp_nopush on;
        tcp_nodelay on;
        keepalive_timeout 65;
        types_hash_max_size 2048;
        # server_tokens off;

        # server_names_hash_bucket_size 64;
        # server_name_in_redirect off;

        include /etc/nginx/mime.types;
        default_type application/octet-stream;
        server {
            listen 443 ssl default_server;
            server_name vault.example.com;
            ssl_session_timeout 5m;
            ssl_certificate /home/vagrant/vault.example.com.crt.pem;
            ssl_certificate_key /home/vagrant/vault.example.com.crt.key;

        ##
        # SSL Settings
        ##

            ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3; # Dropping SSLv3, ref: POODLE
            ssl_prefer_server_ciphers on;

            location / {
                root /usr/share/nginx/html;
                index index.html index.htm;
            }
        }
		
        ##
        # Logging Settings
        ##

        access_log /var/log/nginx/access.log;
        error_log /var/log/nginx/error.log;

        ##
        # Gzip Settings
        ##

        gzip on;

        ##
        # Virtual Host Configs
        ##

        include /etc/nginx/conf.d/*.conf;
        include /etc/nginx/sites-enabled/*;
}
```
# 8) Откроем в браузере на хосте https адрес страницы, которую обслуживает сервер nginx.  
Добавим в файл hosts (c:\windows\system32\drivers\etc\hosts):  
```
127.0.0.1	vault.example.com
```
Откроем в браузере Edge на хостовой машине Win10 адрес: https://vault.example.com:
![](https://github.com/stasarts/netology-devops/blob/main/coursework/DevOps%26SysAdm/secure.png)  
Видим, что "соединение безопасно".  
Откроем сведения о сертификате:
![](https://github.com/stasarts/netology-devops/blob/main/coursework/DevOps%26SysAdm/cert.png)  
Видим всю цепочкус от корневого сертификата.

# 9) Создадим скрипт, который будет генерировать новый сертификат в vault.
Добавим переменные окружения:
```bash
vagrant@Ubuntu20-04:~$ echo "export VAULT_ADDR=http://127.0.0.1:8200" >> .bashrc
vagrant@Ubuntu20-04:~$ echo "export VAULT_TOKEN=s.op4wLup0V0P8UmcTBBgpvP6O" >> .bashrc
```
Напишем скрипт:
```bash
vagrant@Ubuntu20-04:~$ nano autocert.sh
```
```bash
#!/usr/bin/env bash

# распечатаем vault
vault operator unseal X4b0zlKFiudDbpIy+YZCV3okEzwzCLh0K3CLrGtkEejY
vault operator unseal ooZFWzyq0eLytLh6iLqnTp/VzSQCaQQRb+LxvFW+z3tW
vault operator unseal L4L1uSMQXKItgqEJkHAzIp7TQCjy8KeGeVI2F2BbPVyD

# выпускаем сертификат
vault write -format=json pki_int_ca/issue/example-dot-com \
common_name="vault.example.com" \
ttl="720h" > vault.example.com.crt

# Сохраняем сертификат в правильном формате.
cat vault.example.com.crt | jq -r .data.certificate > vault.example.com.crt.pem
cat vault.example.com.crt | jq -r .data.issuing_ca >> vault.example.com.crt.pem
cat vault.example.com.crt | jq -r .data.private_key > vault.example.com.crt.key

# перечитываем nginx.conf
systemctl reload nginx
```

# 10) Поместим скрипт в crontab, чтобы сертификат обновлялся какого-то числа каждого месяца в удобное время.
```bash
vagrant@Ubuntu20-04:~$ crontab -e
```
Будем запускать скрипт 15-го числа каждого месяца в 2:00.
```bash
0 2 15 * * /home/vagrant/autocert.sh
```
