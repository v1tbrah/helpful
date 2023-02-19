# find

## examples
* find file "main.go" in current directory
  * `find . -type f -name "main.go"`
* find files who start "main" case-sensitive in current directory
  * `find . -type f -name "main*"`
* find files who start "main" NO case-sensitive in current directory
  * `find . -type f -iname "main*"`
* find directory "mydir" in current directory
  * `find . -type d -name "mydir"`
* find files with perm 0664 in current directory
  * `find . -type f -perm 0664`
* find files with size grater then 10MB in current directory
  * `find . -type f -size +10M`
