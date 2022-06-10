# Домашнее задание к занятию "08.04 Создание собственных modules"

<details>
  <summary>Задание</summary>

## Подготовка к выполнению
1. Создайте пустой публичных репозиторий в любом своём проекте: `my_own_collection`
2. Скачайте репозиторий ansible: `git clone https://github.com/ansible/ansible.git` по любому удобному вам пути
3. Зайдите в директорию ansible: `cd ansible`
4. Создайте виртуальное окружение: `python3 -m venv venv`
5. Активируйте виртуальное окружение: `. venv/bin/activate`. Дальнейшие действия производятся только в виртуальном окружении
6. Установите зависимости `pip install -r requirements.txt`
7. Запустить настройку окружения `. hacking/env-setup`
8. Если все шаги прошли успешно - выйти из виртуального окружения `deactivate`
9. Ваше окружение настроено, для того чтобы запустить его, нужно находиться в директории `ansible` и выполнить конструкцию `. venv/bin/activate && . hacking/env-setup`

## Основная часть

Наша цель - написать собственный module, который мы можем использовать в своей role, через playbook. Всё это должно быть собрано в виде collection и отправлено в наш репозиторий.

1. В виртуальном окружении создать новый `my_own_module.py` файл

<details>
  <summary>2. Наполнить его содержимым:</summary>

```python
#!/usr/bin/python

# Copyright: (c) 2018, Terry Jones <terry.jones@example.org>
# GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)
from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

DOCUMENTATION = r'''
---
module: my_test

short_description: This is my test module

# If this is part of a collection, you need to use semantic versioning,
# i.e. the version is of the form "2.5.0" and not "2.4".
version_added: "1.0.0"

description: This is my longer description explaining my test module.

options:
    name:
        description: This is the message to send to the test module.
        required: true
        type: str
    new:
        description:
            - Control to demo if the result of this module is changed or not.
            - Parameter description can be a list as well.
        required: false
        type: bool
# Specify this value according to your collection
# in format of namespace.collection.doc_fragment_name
extends_documentation_fragment:
    - my_namespace.my_collection.my_doc_fragment_name

author:
    - Your Name (@yourGitHubHandle)
'''

EXAMPLES = r'''
# Pass in a message
- name: Test with a message
  my_namespace.my_collection.my_test:
    name: hello world

# pass in a message and have changed true
- name: Test with a message and changed output
  my_namespace.my_collection.my_test:
    name: hello world
    new: true

# fail the module
- name: Test failure of the module
  my_namespace.my_collection.my_test:
    name: fail me
'''

RETURN = r'''
# These are examples of possible return values, and in general should use other names for return values.
original_message:
    description: The original name param that was passed in.
    type: str
    returned: always
    sample: 'hello world'
message:
    description: The output message that the test module generates.
    type: str
    returned: always
    sample: 'goodbye'
'''

from ansible.module_utils.basic import AnsibleModule


def run_module():
    # define available arguments/parameters a user can pass to the module
    module_args = dict(
        name=dict(type='str', required=True),
        new=dict(type='bool', required=False, default=False)
    )

    # seed the result dict in the object
    # we primarily care about changed and state
    # changed is if this module effectively modified the target
    # state will include any data that you want your module to pass back
    # for consumption, for example, in a subsequent task
    result = dict(
        changed=False,
        original_message='',
        message=''
    )

    # the AnsibleModule object will be our abstraction working with Ansible
    # this includes instantiation, a couple of common attr would be the
    # args/params passed to the execution, as well as if the module
    # supports check mode
    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    # if the user is working with this module in only check mode we do not
    # want to make any changes to the environment, just return the current
    # state with no modifications
    if module.check_mode:
        module.exit_json(**result)

    # manipulate or modify the state as needed (this is going to be the
    # part where your module will do what it needs to do)
    result['original_message'] = module.params['name']
    result['message'] = 'goodbye'

    # use whatever logic you need to determine whether or not this module
    # made any modifications to your target
    if module.params['new']:
        result['changed'] = True

    # during the execution of the module, if there is an exception or a
    # conditional state that effectively causes a failure, run
    # AnsibleModule.fail_json() to pass in the message and the result
    if module.params['name'] == 'fail me':
        module.fail_json(msg='You requested this to fail', **result)

    # in the event of a successful module execution, you will want to
    # simple AnsibleModule.exit_json(), passing the key/value results
    module.exit_json(**result)


def main():
    run_module()


if __name__ == '__main__':
    main()
```

Или возьмите данное наполнение из [статьи](https://docs.ansible.com/ansible/latest/dev_guide/developing_modules_general.html#creating-a-module).

</details>

3. Заполните файл в соответствии с требованиями ansible так, чтобы он выполнял основную задачу: module должен создавать текстовый файл на удалённом хосте по пути, определённом в параметре `path`, с содержимым, определённым в параметре `content`.
4. Проверьте module на исполняемость локально.
5. Напишите single task playbook и используйте module в нём.
6. Проверьте через playbook на идемпотентность.
7. Выйдите из виртуального окружения.
8. Инициализируйте новую collection: `ansible-galaxy collection init my_own_namespace.my_own_collection`
9. В данную collection перенесите свой module в соответствующую директорию.
10. Single task playbook преобразуйте в single task role и перенесите в collection. У role должны быть default всех параметров module
11. Создайте playbook для использования этой role.
12. Заполните всю документацию по collection, выложите в свой репозиторий, поставьте тег `1.0.0` на этот коммит.
13. Создайте .tar.gz этой collection: `ansible-galaxy collection build` в корневой директории collection.
14. Создайте ещё одну директорию любого наименования, перенесите туда single task playbook и архив c collection.
15. Установите collection из локального архива: `ansible-galaxy collection install <archivename>.tar.gz`
16. Запустите playbook, убедитесь, что он работает.
17. В ответ необходимо прислать ссылку на репозиторий с collection

## Необязательная часть

1. Используйте свой полёт фантазии: Создайте свой собственный module для тех roles, что мы делали в рамках предыдущих лекций.
2. Соберите из roles и module отдельную collection.
3. Создайте новый репозиторий и выложите новую collection туда.

Если идей нет, но очень хочется попробовать что-то реализовать: реализовать module восстановления из backup elasticsearch.

</details>

<details>
  <summary>Ответ</summary>

<details>
  <summary>  Подготовка к выполнению </summary>

1. Создадим пустой публичный репозиторий в любом своём проекте: `example-ansible-collection`.

`OK`

2. Скачаем репозиторий ansible: `git clone https://github.com/ansible/ansible.git`.
```shell
~/ansible$ git clone https://github.com/ansible/ansible.git
Cloning into 'ansible'...
remote: Enumerating objects: 578196, done.
remote: Counting objects: 100% (336/336), done.
remote: Compressing objects: 100% (257/257), done.
remote: Total 578196 (delta 143), reused 208 (delta 52), pack-reused 577860
Receiving objects: 100% (578196/578196), 202.72 MiB | 171.00 KiB/s, done.
Resolving deltas: 100% (388020/388020), done.
```

3. Зайдем в директорию ansible: `cd ansible`.
```shell
~/ansible$ cd ansible/
~/ansible/ansible$ ll
total 144
drwxrwxr-x 14 stasarts stasarts  4096 июн  9 14:31 ./
drwxrwxr-x  9 stasarts stasarts  4096 июн  9 14:10 ../
drwxrwxr-x  4 stasarts stasarts  4096 июн  9 14:31 .azure-pipelines/
drwxrwxr-x  2 stasarts stasarts  4096 июн  9 14:31 bin/
drwxrwxr-x  3 stasarts stasarts  4096 июн  9 14:31 changelogs/
-rw-rw-r--  1 stasarts stasarts   202 июн  9 14:31 .cherry_picker.toml
-rw-rw-r--  1 stasarts stasarts 35148 июн  9 14:31 COPYING
drwxrwxr-x  6 stasarts stasarts  4096 июн  9 14:31 docs/
drwxrwxr-x  3 stasarts stasarts  4096 июн  9 14:31 examples/
drwxrwxr-x  8 stasarts stasarts  4096 июн  9 14:31 .git/
-rw-rw-r--  1 stasarts stasarts    23 июн  9 14:31 .gitattributes
drwxrwxr-x  3 stasarts stasarts  4096 июн  9 14:31 .github/
-rw-rw-r--  1 stasarts stasarts  2806 июн  9 14:31 .gitignore
drwxrwxr-x  7 stasarts stasarts  4096 июн  9 14:31 hacking/
drwxrwxr-x  3 stasarts stasarts  4096 июн  9 14:31 lib/
drwxrwxr-x  2 stasarts stasarts  4096 июн  9 14:31 licenses/
-rw-rw-r--  1 stasarts stasarts  2200 июн  9 14:31 .mailmap
-rw-rw-r--  1 stasarts stasarts  4922 июн  9 14:31 Makefile
-rw-rw-r--  1 stasarts stasarts  2529 июн  9 14:31 MANIFEST.in
drwxrwxr-x  4 stasarts stasarts  4096 июн  9 14:31 packaging/
-rw-rw-r--  1 stasarts stasarts   100 июн  9 14:31 pyproject.toml
-rw-rw-r--  1 stasarts stasarts  5697 июн  9 14:31 README.rst
-rw-rw-r--  1 stasarts stasarts   838 июн  9 14:31 requirements.txt
-rw-rw-r--  1 stasarts stasarts  2476 июн  9 14:31 setup.cfg
-rw-rw-r--  1 stasarts stasarts  1116 июн  9 14:31 setup.py
drwxrwxr-x  9 stasarts stasarts  4096 июн  9 14:31 test/
```

4. Создадим виртуальное окружение: `python3 -m venv venv`.

`OK`

5. Активируем виртуальное окружение: `. venv/bin/activate`. Дальнейшие действия производятся только в виртуальном окружении.
```shell
~/ansible/ansible$ . venv/bin/activate
(venv) ~/ansible/ansible$ 
```

6. Установим зависимости `pip install -r requirements.txt`.
```shell
(venv) ~/ansible/ansible$ pip install -r requirements.txt
Collecting jinja2>=3.0.0
  Downloading Jinja2-3.1.2-py3-none-any.whl (133 kB)
     |████████████████████████████████| 133 kB 298 kB/s 
Collecting PyYAML>=5.1
  Downloading PyYAML-6.0-cp38-cp38-manylinux_2_5_x86_64.manylinux1_x86_64.manylinux_2_12_x86_64.manylinux2010_x86_64.whl (701 kB)
     |████████████████████████████████| 701 kB 514 kB/s 
Collecting cryptography
  Downloading cryptography-37.0.2-cp36-abi3-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (4.1 MB)
     |████████████████████████████████| 4.1 MB 377 kB/s 
Collecting packaging
  Downloading packaging-21.3-py3-none-any.whl (40 kB)
     |████████████████████████████████| 40 kB 469 kB/s 
Collecting resolvelib<0.9.0,>=0.5.3
  Downloading resolvelib-0.8.1-py2.py3-none-any.whl (16 kB)
Collecting MarkupSafe>=2.0
  Downloading MarkupSafe-2.1.1-cp38-cp38-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (25 kB)
Collecting cffi>=1.12
  Downloading cffi-1.15.0-cp38-cp38-manylinux_2_12_x86_64.manylinux2010_x86_64.whl (446 kB)
     |████████████████████████████████| 446 kB 495 kB/s 
Collecting pyparsing!=3.0.5,>=2.0.2
  Downloading pyparsing-3.0.9-py3-none-any.whl (98 kB)
     |████████████████████████████████| 98 kB 444 kB/s 
Collecting pycparser
  Downloading pycparser-2.21-py2.py3-none-any.whl (118 kB)
     |████████████████████████████████| 118 kB 708 kB/s 
Installing collected packages: MarkupSafe, jinja2, PyYAML, pycparser, cffi, cryptography, pyparsing, packaging, resolvelib
Successfully installed MarkupSafe-2.1.1 PyYAML-6.0 cffi-1.15.0 cryptography-37.0.2 jinja2-3.1.2 packaging-21.3 pycparser-2.21 pyparsing-3.0.9 resolvelib-0.8.1
```

7. Запустим настройку окружения `. hacking/env-setup`.
```shell
(venv) stasarts@stasarts-laptop:~/ansible/ansible$ . hacking/env-setup
running egg_info
creating lib/ansible_core.egg-info
writing lib/ansible_core.egg-info/PKG-INFO
writing dependency_links to lib/ansible_core.egg-info/dependency_links.txt
writing entry points to lib/ansible_core.egg-info/entry_points.txt
writing requirements to lib/ansible_core.egg-info/requires.txt
writing top-level names to lib/ansible_core.egg-info/top_level.txt
writing manifest file 'lib/ansible_core.egg-info/SOURCES.txt'
reading manifest file 'lib/ansible_core.egg-info/SOURCES.txt'
reading manifest template 'MANIFEST.in'
warning: no files found matching 'SYMLINK_CACHE.json'
warning: no previously-included files found matching 'docs/docsite/rst_warnings'
warning: no previously-included files found matching 'docs/docsite/rst/conf.py'
warning: no previously-included files found matching 'docs/docsite/rst/index.rst'
warning: no previously-included files found matching 'docs/docsite/rst/dev_guide/index.rst'
warning: no previously-included files matching '*' found under directory 'docs/docsite/_build'
warning: no previously-included files matching '*.pyc' found under directory 'docs/docsite/_extensions'
warning: no previously-included files matching '*.pyo' found under directory 'docs/docsite/_extensions'
warning: no files found matching '*.ps1' under directory 'lib/ansible/modules/windows'
warning: no files found matching '*.yml' under directory 'lib/ansible/modules'
warning: no files found matching 'validate-modules' under directory 'test/lib/ansible_test/_util/controller/sanity/validate-modules'
writing manifest file 'lib/ansible_core.egg-info/SOURCES.txt'

Setting up Ansible to run out of checkout...

PATH=/home/stasarts/ansible/ansible/bin:/home/stasarts/ansible/ansible/venv/bin:/home/stasarts/yandex-cloud/bin:/home/stasarts/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin:/usr/local/go/bin:/opt/sonar-scanner-4.7.0.2747-linux/bin:/opt/apache-maven-3.8.5/bin:/home/stasarts/.local/bin
PYTHONPATH=/home/stasarts/ansible/ansible/test/lib:/home/stasarts/ansible/ansible/lib
MANPATH=/home/stasarts/ansible/ansible/docs/man:/usr/local/man:/usr/local/share/man:/usr/share/man

Remember, you may wish to specify your host file with -i

Done!

```

9. Все шаги прошли успешно - выйдем из виртуального окружения `deactivate`.
```shell
(venv) ~/ansible/ansible$ deactivate
~/ansible/ansible$
```

10. Окружение настроено, для того чтобы запустить его, нужно находиться в директории `ansible` и выполнить конструкцию `. venv/bin/activate && . hacking/env-setup`.

`OK`

</details>

<details>
  <summary>  Основная часть </summary>

Наша цель - написать собственный module, который мы можем использовать в своей role, через playbook. Всё это должно быть собрано в виде collection и отправлено в наш репозиторий.

1. В виртуальном окружении создадим новый `my_own_module.py` файл (create_file.py).
```shell
(venv) ~/ansible/ansible$ touch lib/ansible/modules/create_file.py
```

<details>
  <summary>2. Наполнить его содержимым:</summary>

```python
#!/usr/bin/python

# Copyright: (c) 2018, Terry Jones <terry.jones@example.org>
# GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)
from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

DOCUMENTATION = r'''
---
module: my_test

short_description: This is my test module

# If this is part of a collection, you need to use semantic versioning,
# i.e. the version is of the form "2.5.0" and not "2.4".
version_added: "1.0.0"

description: This is my longer description explaining my test module.

options:
    name:
        description: This is the message to send to the test module.
        required: true
        type: str
    new:
        description:
            - Control to demo if the result of this module is changed or not.
            - Parameter description can be a list as well.
        required: false
        type: bool
# Specify this value according to your collection
# in format of namespace.collection.doc_fragment_name
extends_documentation_fragment:
    - my_namespace.my_collection.my_doc_fragment_name

author:
    - Your Name (@yourGitHubHandle)
'''

EXAMPLES = r'''
# Pass in a message
- name: Test with a message
  my_namespace.my_collection.my_test:
    name: hello world

# pass in a message and have changed true
- name: Test with a message and changed output
  my_namespace.my_collection.my_test:
    name: hello world
    new: true

# fail the module
- name: Test failure of the module
  my_namespace.my_collection.my_test:
    name: fail me
'''

RETURN = r'''
# These are examples of possible return values, and in general should use other names for return values.
original_message:
    description: The original name param that was passed in.
    type: str
    returned: always
    sample: 'hello world'
message:
    description: The output message that the test module generates.
    type: str
    returned: always
    sample: 'goodbye'
'''

from ansible.module_utils.basic import AnsibleModule


def run_module():
    # define available arguments/parameters a user can pass to the module
    module_args = dict(
        name=dict(type='str', required=True),
        new=dict(type='bool', required=False, default=False)
    )

    # seed the result dict in the object
    # we primarily care about changed and state
    # changed is if this module effectively modified the target
    # state will include any data that you want your module to pass back
    # for consumption, for example, in a subsequent task
    result = dict(
        changed=False,
        original_message='',
        message=''
    )

    # the AnsibleModule object will be our abstraction working with Ansible
    # this includes instantiation, a couple of common attr would be the
    # args/params passed to the execution, as well as if the module
    # supports check mode
    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    # if the user is working with this module in only check mode we do not
    # want to make any changes to the environment, just return the current
    # state with no modifications
    if module.check_mode:
        module.exit_json(**result)

    # manipulate or modify the state as needed (this is going to be the
    # part where your module will do what it needs to do)
    result['original_message'] = module.params['name']
    result['message'] = 'goodbye'

    # use whatever logic you need to determine whether or not this module
    # made any modifications to your target
    if module.params['new']:
        result['changed'] = True

    # during the execution of the module, if there is an exception or a
    # conditional state that effectively causes a failure, run
    # AnsibleModule.fail_json() to pass in the message and the result
    if module.params['name'] == 'fail me':
        module.fail_json(msg='You requested this to fail', **result)

    # in the event of a successful module execution, you will want to
    # simple AnsibleModule.exit_json(), passing the key/value results
    module.exit_json(**result)


def main():
    run_module()


if __name__ == '__main__':
    main()
```

Или возьмите данное наполнение из [статьи](https://docs.ansible.com/ansible/latest/dev_guide/developing_modules_general.html#creating-a-module).

</details>

`OK`

<details>
  <summary>3. Заполним файл в соответствии с требованиями ansible так, чтобы он выполнял основную задачу: module должен создавать текстовый файл на удалённом хосте по пути, определённом в параметре `path`, с содержимым, определённым в параметре `content`.</summary>

```shell
#!/usr/bin/python

# Copyright: (c) 2018, Terry Jones <terry.jones@example.org>
# GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)
from __future__ import (absolute_import, division, print_function)

__metaclass__ = type

DOCUMENTATION = r'''
---
module: create_file

short_description: This is my test module for file creation

version_added: "1.0.0"

description: This is my netology practice test module for file creation.

options:
    path:
        description: This is relative path to file should be created or rewritten to send to the test module.
        required: true
        type: str
    content:
        description:
            - Content to be written to target file.
            - Default content is empty. If not provided, target file content would be erased.
        required: false
        type: str
# Specify this value according to your collection
# in format of namespace.collection.doc_fragment_name
extends_documentation_fragment:
    - avt0m8.netology_test_module_create_file.my_doc_create_file

author:
    - Stas Arts (@stasarts)
'''

EXAMPLES = r'''
# Pass with empty file
- path: Test with a path only
  avt0m8.netology_test_module_create_file.create_file:
    path: ./create_file_test.txt

# Pass with fulfilled file
- name: Test with created/rewritten and fulfilled file
  my_namespace.my_collection.my_test:
    path: ./create_file_test.txt
    content: "outstanding content"
'''

RETURN = r'''
# These are examples of possible return values, and in general should use other names for return values.
original_message:
    description: The original content param that was passed in.
    type: str
    returned: always
    sample: 'outstanding content'
message:
    description: The output message that the test module generates.
    type: str
    returned: always
    sample: 
        - 'file was created'
        - 'file was rewritten'
'''

import os
from ansible.module_utils.basic import AnsibleModule


# Check if file exists
def file_exists(path):
    if os.path.exists(path):
        return True
    else:
        return False


# Check is the file content as provided
def check_content(path, content):
    with open(path, 'r') as file:
        file_content = file.read()
    if file_content == content:
        return True
    else:
        return False


# Write content to file
def write_content_to_file(path, content):
    with open(path, 'w') as file:
        file.write(content)
    return True


def run_module():
    # define available arguments/parameters a user can pass to the module
    module_args = dict(
        path=dict(type='str', required=True),
        content=dict(type='str', required=False, default=' ')
    )

    result = dict(
        changed=False,
        original_message='',
        message=''
    )

    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    if module.check_mode:
        if not os.path.exists(module.params['path']):
            result['changed'] = True
        module.exit_json(**result)

    if not file_exists(module.params['path']):
        write_content_to_file(module.params['path'], module.params['content'])
        result['changed'] = True
        result['original_message'] = "File {path} was successfully created".format(path=module.params['path'])
        result['message'] = 'file was created'
    else:
        if check_content(module.params['path'], module.params['content']):
            result['changed'] = False
            result['original_message'] = "File {path} exists with the same content".format(path=module.params['path'])
            result['message'] = 'file exists'
        else:
            write_content_to_file(module.params['path'], module.params['content'])
            result['changed'] = True
            result['original_message'] = "File {path} was successfully rewritten".format(path=module.params['path'])
            result['message'] = 'file was rewritten'

    module.exit_json(**result)


def main():
    run_module()


if __name__ == '__main__':
    main()

```
</details>

`OK`

4. Проверим module на исполняемость локально.
* Создадим payload.json:
```json
{
    "ANSIBLE_MODULE_ARGS": {
        "path": "create_file_test.txt",
        "content": "outstanding content"
    }
}
```

* Создание файла:
```shell
(venv) ~/ansible/ansible$ cat create_file_test.txt
cat: create_file_test.txt: No such file or directory
(venv) ~/ansible/ansible$ python -m ansible.modules.create_file payload.json

{"changed": true, "original_message": "File ./create_file_test.txt was successfully created", "message": "file was created", "invocation": {"module_args": {"path": "./create_file_test.txt", "content": "outstanding content"}}}
(venv) ~/ansible/ansible$ cat create_file_test.txt
outstanding content
```

* Запуск с уже имеющимся файлом:
```shell
(venv) ~/ansible/ansible$ python -m ansible.modules.create_file payload.json

{"changed": false, "original_message": "File ./create_file_test.txt exists with the same content", "message": "file exists", "invocation": {"module_args": {"path": "./create_file_test.txt", "content": "outstanding content"}}}

```

* Перезапись файла с неправильным контентом:
```shell
(venv) ~/ansible/ansible$ echo abrakadabra > create_file_test.txt 
(venv) ~/ansible/ansible$ cat create_file_test.txt 
abrakadabra
(venv) ~/ansible/ansible$ python -m ansible.modules.create_file payload.json

{"changed": true, "original_message": "File ./create_file_test.txt was successfully rewritten", "message": "file was rewritten", "invocation": {"module_args": {"path": "./create_file_test.txt", "content": "outstanding content"}}}
(venv) ~/ansible/ansible$ cat create_file_test.txt 
outstanding content
```

5. Напишем single task playbook и используем module в нём.
* playbooks/site.yml:
```shell
---
- name: Test file creation
  hosts: localhost
  tasks:
    - name: Create file
      create_file:
        path: ./create_file_test.txt
        content: "outstanding content"
```

6. Проверим через playbook на идемпотентность.
* Запуск с --check.
```shell
(venv) ~/ansible/ansible$ ll playbooks/
total 12
drwxrwxr-x  2 stasarts stasarts 4096 июн 10 23:59 ./
drwxrwxr-x 16 stasarts stasarts 4096 июн 10 23:52 ../
-rw-rw-r--  1 stasarts stasarts  177 июн 10 19:41 site.yml
(venv) stasarts@stasarts-laptop:~/ansible/ansible$ ansible-playbook playbooks/site.yml --check
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test file creation] *********************************************************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************************************************************
ok: [localhost]

TASK [Create file] ****************************************************************************************************************************************************************************************
changed: [localhost]

PLAY RECAP ************************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

(venv) ~/ansible/ansible$ ll playbooks/
total 12
drwxrwxr-x  2 stasarts stasarts 4096 июн 10 23:59 ./
drwxrwxr-x 16 stasarts stasarts 4096 июн 10 23:52 ../
-rw-rw-r--  1 stasarts stasarts  177 июн 10 19:41 site.yml
```

* Создание файла:
```shell
(venv) ~/ansible/ansible$ ll playbooks/
total 12
drwxrwxr-x  2 stasarts stasarts 4096 июн 10 23:59 ./
drwxrwxr-x 16 stasarts stasarts 4096 июн 10 23:52 ../
-rw-rw-r--  1 stasarts stasarts  177 июн 10 19:41 site.yml
(venv) stasarts@stasarts-laptop:~/ansible/ansible$ ansible-playbook playbooks/site.yml
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test file creation] *********************************************************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************************************************************
ok: [localhost]

TASK [Create file] ****************************************************************************************************************************************************************************************
changed: [localhost]

PLAY RECAP ************************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

(venv) ~/ansible/ansible$ cat playbooks/create_file_test.txt 
outstanding content
```

* Запуск с уже имеющимся файлом:
```shell
(venv) ~/ansible/ansible$ ansible-playbook playbooks/site.yml
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test file creation] *********************************************************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************************************************************
ok: [localhost]

TASK [Create file] ****************************************************************************************************************************************************************************************
ok: [localhost]

PLAY RECAP ************************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


```

* Перезапись файла с неправильным контентом:
```shell
(venv) ~/ansible/ansible$ echo abrakadabra > playbooks/create_file_test.txt 
(venv) ~/ansible/ansible$ cat playbooks/create_file_test.txt 
abrakadabra
(venv) ~/ansible/ansible$ ansible-playbook playbooks/site.yml
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test file creation] *********************************************************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************************************************************
ok: [localhost]

TASK [Create file] ****************************************************************************************************************************************************************************************
changed: [localhost]

PLAY RECAP ************************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

(venv) ~/ansible/ansible$ cat playbooks/create_file_test.txt 
outstanding content
```

7. Выйдем из виртуального окружения.
```shell
(venv) ~/ansible/ansible$ deactivate
~/ansible/ansible$
```

8. Инициализируем новую collection: `ansible-galaxy collection init my_own_namespace.my_own_collection`.
```shell
~/ansible/example-ansible-collection$ ansible-galaxy collection init avt0m8.netology_test_module_create_file 
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
- Collection avt0m8.netology_test_module_create_file was created successfully
```

9. В данную collection перенесем свой module в соответствующую директорию.
```shell
~/ansible/example-ansible-collection$ ll avt0m8/netology_test_module_create_file/plugins/modules
total 12
drwxrwxr-x 2 stasarts stasarts 4096 июн 11 02:35 ./
drwxrwxr-x 3 stasarts stasarts 4096 июн 11 01:41 ../
-rw-rw-r-- 1 stasarts stasarts 4077 июн 10 23:49 create_file.py
```

10. Single task playbook преобразуем в single task role и перенесем в collection. У role должны быть default всех параметров module.
* Инициализируем role:
```shell
~/ansible/example-ansible-collection/avt0m8/netology_test_module_create_file/roles$ ansible-galaxy init create_file_role
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
- Role create_file_role was created successfully
```

* avt0m8/netology_test_module_create_file/roles/create_file_role/tasks/main.yml:
```yaml
---
# tasks file for create_file_role
- name: Test file creation
  tasks:
    - name: Create file module test
      avt0m8.netology_test_module_create_file.create_file:
        path: "{{ path }}"
        content: "{{ content }}"
```

* avt0m8/netology_test_module_create_file/roles/create_file_role/defaults/main.yml:
```yaml
---
# defaults file for create_file_role
  path: ./create_file_test.txt
  content: "outstanding content"
```

11. Создадим playbook для использования этой role.
* site.yml:
```yaml
---
- hosts: all
  collections:
    - avt0m8.netology_test_module_create_file
  roles:
    - create_file_role
```

* group_vars/all/vars.yml:
```yaml
---
path: ./prod_create_file.txt
content: "prod content"
```

* inventory/prod.yml:
```yaml
---
  all:
      hosts:
        localhost:
            ansible_connection: local
```

12. Заполним всю документацию по collection, выложим в свой репозиторий, поставим тег `1.0.0` на этот коммит.

[Ссылка на репозиторий с коллекцией.](https://github.com/stasarts/example-ansible-collection)

13. Создадим .tar.gz этой collection: `ansible-galaxy collection build` в корневой директории collection.
```shell
~/ansible/example-ansible-collection/avt0m8/netology_test_module_create_file$ ansible-galaxy collection build
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
Created collection for avt0m8.netology_test_module_create_file at /home/stasarts/ansible/example-ansible-collection/avt0m8/netology_test_module_create_file/avt0m8-netology_test_module_create_file-1.0.0.tar.gz

```

14. Создадим ещё одну директорию (example-ansible-module), перенесите туда single task playbook и архив c collection.
```yaml
~/ansible/example-ansible-module$ ll
total 36
drwxrwxr-x  6 stasarts stasarts 4096 июн 11 03:35 ./
drwxrwxr-x 10 stasarts stasarts 4096 июн 11 03:31 ../
-rw-rw-r--  1 stasarts stasarts 5367 июн 11 03:29 avt0m8-netology_test_module_create_file-1.0.0.tar.gz
drwxrwxr-x  8 stasarts stasarts 4096 июн  9 14:09 .git/
drwxrwxr-x  4 stasarts stasarts 4096 июн  8 20:16 group_vars/
drwxrwxr-x  2 stasarts stasarts 4096 июн  8 23:41 inventory/
-rwxrwxr-x  1 stasarts stasarts   39 июн 11 03:35 site.yml*
```

15. Установим collection из локального архива: `ansible-galaxy collection install <archivename>.tar.gz`
```shell
~/ansible/example-ansible-module$ ansible-galaxy collection install avt0m8-netology_test_module_create_file-1.0.0.tar.gz -p ./collections
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.
Starting galaxy collection install process
[WARNING]: The specified collections path '/home/stasarts/ansible/example-ansible-module/collections' is not part of the configured Ansible collections paths
'/home/stasarts/.ansible/collections:/usr/share/ansible/collections'. The installed collection will not be picked up in an Ansible run, unless within a playbook-adjacent collections directory.
Process install dependency map
Starting collection install process
Installing 'avt0m8.netology_test_module_create_file:1.0.0' to '/home/stasarts/ansible/example-ansible-module/collections/ansible_collections/avt0m8/netology_test_module_create_file'
avt0m8.netology_test_module_create_file:1.0.0 was installed successfully
```

16. Запустим playbook, убедимся, что он работает.
```shell
~/ansible/example-ansible-module$ ansible-playbook -i inventory/prod.yml site.yml
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a
rapidly changing source of code and can become unstable at any point.

PLAY [all] ************************************************************************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************************************************************
ok: [localhost]

TASK [avt0m8.netology_test_module_create_file.create_file_role : Create file module test] *****************************************************************************************************************
changed: [localhost]

PLAY RECAP ************************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

```

```shell
~/ansible/example-ansible-module$ ll
total 40
drwxrwxr-x  6 stasarts stasarts 4096 июн 11 04:06 ./
drwxrwxr-x 10 stasarts stasarts 4096 июн 11 03:31 ../
-rw-rw-r--  1 stasarts stasarts 5367 июн 11 03:29 avt0m8-netology_test_module_create_file-1.0.0.tar.gz
drwxrwxr-x  3 stasarts stasarts 4096 июн 11 03:36 collections/
drwxrwxr-x  8 stasarts stasarts 4096 июн 11 03:39 .git/
drwxrwxr-x  3 stasarts stasarts 4096 июн 11 03:39 group_vars/
drwxrwxr-x  2 stasarts stasarts 4096 июн 11 03:54 inventory/
-rw-rw-r--  1 stasarts stasarts   12 июн 11 03:54 prod_create_file.txt
-rwxrwxr-x  1 stasarts stasarts  109 июн 11 03:49 site.yml*
```

```shell
~/ansible/example-ansible-module$ cat prod_create_file.txt 
prod content
```

17. Ссылки на репозитории с collection и playbook.

[Ссылка на репозиторий с коллекцией.](https://github.com/stasarts/example-ansible-collection)

[Ссылка на репозиторий с playbook, использующую установленную коллекцию.](https://github.com/stasarts/example-ansible-module)

</details>

</details>



---