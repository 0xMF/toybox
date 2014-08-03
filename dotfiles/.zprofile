# .zprofile
# ----------

KEY=$(/usr/bin/grep "^KEY" $HOME/.secrets/all_my_secrets|cut -d"=" -f2)

  ### START-Keychain ###

# run keychain in quiet mode, and set SSH_AGENT_PID/SSH_AGENT_SOCK in shells
/usr/bin/env keychain $KEY -q
source $HOME/.keychain/${HOST}-sh

  ### End-Keychain ###

