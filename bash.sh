GODO_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export PROMPT_COMMAND='if [ "$(id -u)" -ne 0 ]; then echo "$(date) $(pwd) $(history 1 | cut -d" " -f4-)" >> ~/.godo.log; fi'

function j() {
  "$GODO_DIR"/bin/client directory "$@"
}

function c() {
  "$GODO_DIR"/bin/client command "$@"
}

function godo_server() {
  "$GODO_DIR"/bin/server ~/.godo.log
}
