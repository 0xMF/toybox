# .zprofile
# ----------

KEY=$(/usr/bin/grep "^KEY" $HOME/.secrets/all_my_secrets|cut -d"=" -f2)
LOG=~/agent.zsh

  ### START-Keychain ###

/usr/bin/env keychain $KEY 2> $LOG
/usr/bin/grep 'Found existing' $LOG > /dev/null

# display newly added otherwise remove log
[ $? -ne 0 ] && cat $LOG || rm $LOG

  ### End-Keychain ###

