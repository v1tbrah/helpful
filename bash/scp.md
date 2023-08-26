#### Скопировать файл с удаленного хоста на локальный компьютер:

scp [имяУдаленногоХоста]:[путьКФайлуНаУдаленномХосте] [путьКудаСохранятьНаЛокальномКомпьютере]

	scp remoteServeHost:/path/onRemoteServer/to/fileName.json /path/onLocalServer/to

---
#### Скопировать файл с локального компьютера на удаленный хост:

scp [путьКФайлуНаЛокальномКомпьютере] [имяПользователяНаУдаленномХосте]@[имяУдаленногоХоста]:[путьКДиректорииКудаСохранять]

	scp /path/onLocalServer/to/fileName.json useOnRemoteServer@remoteServeHost:/path/onRemoteServer/to
  
---






