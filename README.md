gitversion
==========

Include git commit hash in golang file

usage
=====
gitversion -i pathToRepository -o pathToVersionFile -p packageName

-i defaults to .<br>
-o defaults to version.go<br>
-p defaults to version<br>

template
========
package %s

var GIT_COMMIT_HASH = "%s"
