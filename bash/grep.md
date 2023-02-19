# grep

## examples
* find "some" in file "main.go" case-sensitive
  * `grep "some" main.go`
* find "some" in files who start "main" case-sensitive
  * `grep "some" main*`
* find "some" in files who start "main" NO case-sensitive
  * `grep -i "some" main*`
* find "some" in files who start "main" case-sensitive with string number
  * `grep -n "some" main*`
* print "hello" to STDIN and find "lo" case-sensitive
  * `echo "hello" | grep "lo"`
