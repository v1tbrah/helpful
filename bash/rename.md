# find + grep combinations

## examples
* rename recursive files "ReadMe.txt" to "README.md"
  * find -name "ReadMe.txt" -type f -exec rename 's/ReadMe.txt/README.md/' '{}' \;
