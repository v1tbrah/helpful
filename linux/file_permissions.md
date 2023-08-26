# category
* owner (shortcut in commands 'u')
* group (shortcut in commands 'g')
* others (shortcut in commands 'o')
## commands
* change owner `chown [newUserName] [fileName]`, for example `chown user2 file1`
* change group `chgrp [newGroupName] [fileName]`, for example `chgrp group2 file1`
# permissions
* read (r) - 4
* write (w) - 2
* execution (x) - 1

| perm  | num |
|-------|-----|
| r     | 4   |
| w     | 2   |
| x     | 1   |
| r+x   | 5   |
| r+w   | 6   |
| r+w+x | 7   |
| ---   | 0   |

## commands
* change permissions `chmod [permToOwner|permToGroup|permToOthers]`, for example
  `chmod 664 file1` - r+w to owner, r+w to group, r to others
* print verbose permissions `getfacl [fileName]`, for example `getfacl file1`
* set additional permissions `setfacl [category]:[categoryMember]:[perm] [fileName]`, for example
    * `setfacl u:user1:r-- file1`
    * `setfacl g:group1:rw- file1`
* delete additional permissions `setfacl -b [fileName]`, for example `setfacl -b file1`
# links
* https://www.youtube.com/watch?v=BVnsrkL6ODw
