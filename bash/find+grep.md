# find + grep combinations

## examples
* find files who start "main" NO case-sensitive in current directory and find "som" case-sensitive
    * `find . -type f -iname "main*" -exec grep "som" {} +`

