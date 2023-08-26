# users
### create user
`useradd [username]`, for example `useradd user1`
* __need set passwd for new user__ `passwd [username]`, for example `passwd user1`
* print default settings for new users `useradd -D`
* file with default settings for new users __/etc/default/useradd__
* there are more settings in file __/etc/login.defs__
### read users

`cat /etc/passwd`

`v1tbrah:x:1000:1000:Viktor,,,:/home/v1tbrah:/bin/bash`

`:` here is separator

| login   | passwd | uid  | gid  | comments  | home dir      | shell     |
|---------|--------|------|------|-----------|---------------|-----------|
| v1tbrah | x      | 1000 | 1000 | Viktor,,, | /home/v1tbrah | /bin/bash |

there are hashed passwords in file __/etc/shadow__

### delete user
`userdel [username]`, for example `userdel user1`
use `userdel -r [username]` for delete user and delete his work dir, private group, etc

