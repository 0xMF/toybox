# .bash_profile
# --------------------

KEY=$(/usr/bin/grep "^KEY" secrets/bash_profile|cut -d"=" -f2)

# Get PID of present ssh-agent and that stored on file
SSH_AGENT_PID=`/usr/bin/tail -1 agent.sh|/usr/bin/cut -d";" -f1|/usr/bin/cut -d"=" -f2`
AGENT_PID=`ps ax|/usr/bin/grep ssh-agent|/usr/bin/grep -v grep|cut -d" " -f1`

# compare, if different it means file version is old, so start ssh-agent
if [ "$AGENT_PID" !=  "$SSH_AGENT_PID" ]; then
  start_ssh-agent 
  ssh-add $KEY
fi 

# call .bashrc for the rest of the stuff
if [ -f ~/.bashrc ]; then
	. ~/.bashrc
fi