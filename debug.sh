#!/bin/sh

script_path=$(cd "$(dirname "$0")" || exit 1; pwd)
tools_path="$script_path/sg-tools"

[ -e "${tools_path}" ] || {
    go build -o "$tools_path" "$script_path/tools/." 
}

"$tools_path" "$@"
