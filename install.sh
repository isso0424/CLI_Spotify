SCRIPT_DIR=$(cd $(dirname $0); pwd)
function writeToPath() {
  if [ $(echo "$SHELL") = $(which "zsh") ]; then
    cat ~/.zsh_profile | grep "export PATH=\$PATH:$SCRIPT_DIR/bin" || echo "export PATH=\$PATH:$SCRIPT_DIR/bin" >> ~/.zsh_profile
  elif [ $(echo "$SHELL") = $(which "bash") ]; then
    cat ~/.bash_profile | grep "export PATH=\$PATH:$SCRIPT_DIR/bin" || echo "export PATH=\$PATH:$SCRIPT_DIR/bin" >> ~/.bash_profile
  fi
}
function build() {
  mkdir "$SCRIPT_DIR/bin"
  go build -o bin/spotify_CLI
  writeToPath
  echo "install successful"
}
cd $(dirname $0)
if [ -e $(which go) ]; then
  build
else
  if [ -e $(which pacman) ]; then
    sudo pacman -S go
    build
  elif [ -e $(which apt) ]; then
    sudo apt install go
    build
  else
    echo "install failed"
    return
  fi
fi

if [[ -e $XDG_CONFIG_HOME ]]; then
  configDir="$XDG_CONFIG_HOME/spotify_CLI"
else
  configDir="$HOME/.config/spotify_CLI"
fi
mkdir -p "$configDir"
touch "$configDir/config"

echo -n "Please input clientID (If you don't have clientID? You can get here!(https://developer.spotify.com/dashboard/applications):"
read clientID
echo -n "Please input secretID:"
read secretID

echo "SPOTIFY_ID: $clientID" > "$configDir/config"
echo "SPOTIFY_SECRET: $secretID" >> "$configDir/config"
echo "... and you have to set to 'http://localhost:8888/callback' redirect url in spotify dashboard"
