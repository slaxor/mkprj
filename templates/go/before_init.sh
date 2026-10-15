#!/bin/bash
# This script receives the following params
# - prjname
# - gitremote
# - description
local prjname="${1}"
local srcname="${2:-${prjname}}"
local gitremote="git@github.com:slaxor/${prjname}.git"
local gosrc="${GOPATH}/src/github.com/slaxor"
#   local prjpath="${gosrc}/${prjname}"
berg repo create --default-branch master --private public \
  --name ${prjname} \
  --description "${prjname} is not yet properly described"
  git clone "${gitremote}" ${gosrc}
mkdir -p "${prjpath}/cmd"

