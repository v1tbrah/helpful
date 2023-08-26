# groups
### create group
`groupadd [groupname]`, for example `groupadd group1`

### read groups
`cat /etc/group`

`sambashare:x:135:v1tbrah,user2`

`:` here is separator

| group name | passwd | gid  | users in group |
|------------|--------|------|----------------|
| sambashare | x      | 1000 | v1tbrah,user2  |

there are hashed passwords in file __/etc/gshadow__
### delete group
`groupdel [groupname]`, for example `groupdel group1`
