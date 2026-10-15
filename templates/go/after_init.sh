#!/bin/bash
# This script receives the following params
# - prjname
# - gitremote
# - description
local prjname="${1}"
local gitremote="${2}"
"git@github.com:slaxor/${prjname}.git"
local gosrc="${GOPATH}/src/github.com/slaxor"
#
set -x
go mod init "git.t5l.de/slaxor/${prjname}"
go mod tidy
git init
git add .
git commit -m "Initial commit"
git remote add origin "it@git.t5l.de/slaxor/${prjname}.git"
# git push -u origin master
