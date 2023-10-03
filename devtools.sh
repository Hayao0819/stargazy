#!/bin/sh

build_always=true
script_path=$(cd "$(dirname "$0")" || exit 1; pwd)
tools_path="$script_path/sg-tools"

{ [ -e "${tools_path}" ] && ! "$build_always"; } || {
    go build -o "$tools_path" "$script_path/tools/." 
}

"$tools_path" "$@"
