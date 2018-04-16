GODO_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export PROMPT_COMMAND='if [ "$(id -u)" -ne 0 ]; then echo "$(date)__GODO_SPLIT__$(pwd)__GODO_SPLIT__$(history 1 | cut -d" " -f5-)" >> ~/.godo.log; fi'

function j() {
  "$GODO_DIR"/bin/client directory "$@"
}

function c() {
  "$GODO_DIR"/bin/client command "$@"
}

function godo_server() {
  "$GODO_DIR"/bin/server ~/.godo.log
}
