# .zprofile
# ----------

KEY=$(/usr/bin/grep "^KEY" $HOME/.secrets/all_my_secrets|cut -d"=" -f2)

  ### START-Keychain ###

/usr/bin/env keychain $KEY -q
source $HOME/.keychain/${HOST}-sh

  ### End-Keychain ###

