GODO_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export PROMPT_COMMAND='if [ "$(id -u)" -ne 0 ]; then echo "$(date) $(pwd) $(history 1 | cut -d" " -f4-)" >> ~/.godo.log; fi'

function j() {
  "$GODO_DIR"/cmd/client/main directory "$@"
}

function c() {
  "$GODO_DIR"/cmd/client/main command "$@"
}

function godo_server() {
  "$GODO_DIR"/cmd/server/main ~/.godo.log
}
