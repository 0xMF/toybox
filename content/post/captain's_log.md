Captain's Log
=============

2015-Apr-21 | Learning
----------------------

How do we learn (specifically programming or system administration) in the age of the Internet and access to instant solutions to the problem we
face at the moment. The solutions are available for free on sites that are sometimes supported
through ad-revenue.

Despite the common think that people no longer need a 'human teacher', I think, being a human
teacher myself, the need for a human teacher is all the more necessary. There are still so many
social, psychological, and yes even spiritual, aspects to the learning process that can never be
fully replaced with technology or AI.

2015-Apr-21 | Write, Edit, and Re-edit
--------------------------------------

A pattern of working I'm thinking of using is as follows:

  * write in morning
  * edit in the evening
  * re-edit the next day

I'll keep my tinkering around as before, but at least I'll mentally have prepared a map of my
daily workflow.

Let's work with this for now.

2015-Apr-20 | Habits and Communities
------------------------------------

Since January of this year, I tried to make at least one GitHub push per day on something technical.
It was possible, initially, to wake up each morning and write something that covered activities of
the previous day but problems started when the subject matter I kept tinkering around with was not
fixed on one particular area. Sometimes it would be programming, other times it would be shell
scripting, still other times it was pushing my dotfiles to GH. On days I had something to push,
progress seemed effortless. Days rolled into weeks and I was really happy with my progress. Then,
slowly, the ideas and experimentation started drying up, I began to run out things to push to GH
each morning. So I started allowing myself to put off that morning activity until I got back from
work that evening. This pattern of working sometimes in the day, or at night, continued for
several weeks, and I kept swinging between day and night pushes to GitHub.

Then one morning I woke up and realized I had not pushed to GH the previous day. My work habits had
ruined my workflow. I was mentally stuck too, because immediately after that 'missed day' I happened
to miss yet another day a week later.

While I lamented about this fact on my preferred social media platform: App.Net, [Pam
Davis](http://pamdavis.com/) kept encouraging me on. I was pleasantly surprised to learn that Pam,
too, had been doing something similar (writing a blog post each day) from January. Knowing Pam
maintained an unbroken link of making a blog post each day from January till today reinvigorated me
to look at how I did things and wonder if I could do something different.

Will I be able to continue that unbroken link from now on? I don't know and, in a sense, I've moved
past the I-must-do-something-each-day checkbox towards trying to establish a pattern that can fit
into my daily schedule at the start of the day and to have something I can roll over to the next
day, meaning, I could publish what I did today, tomorrow instead of trying to do everything on the
same day. In case you were wondering, I got that tip from Pam. It will take a few days to adjust to
this newer thought pattern but it helps knowing there are others who like me are working toward
their own goals yet are kind enough to encourage me on.

2015-Apr-19 | stingy sed
------------------------

sed is greedy, by default, to make it stingy try putting '?' before the *

```shell
# greedy match basic regular expressions
$ echo a_b_c|sed  's/\(.*\)_\(.*\)/\1/'
a_b

# non-greedy (stingy) match basic regular expressions
$ echo a_b_c|sed  's/\(.?*\)_\(.*\)/\1/'
a

```
A slight change is needed for extended regular expressions

```shell
# greedy match extended regular expressions
$ echo a_b_c|sed  -r 's/(.*)_(.*)/\1/'
a_b

# non-greedy (stingy) match extended regular expressions
$ echo a_b_c|sed  -r 's/(.\?*)_(.*)/\1/'
a

```

This works on Ubuntu and on FreeBSD. BSD's have their own conservative versions of sed that have
subtle differences with GNU sed in some respects.

2015-Apr-18 | which is where
----------------------------

I constantly keep getting confused between `/bin/grep` or `/usr/bin/grep` and other such utilities
(or shells). Why? I switch from OSes so much on a daily basis and I have my `grep` aliased to `grep
--color=always -i` but ever so often, I'd like to further process the output match from my aliased
grep into something else (like sed or awk). The trouble is due to the colorizations my output from
grep now has the ansi color codes mixed in. This always trips me up initially but I've got better
about this so I recover quickly. It's the having to do the next step that is annoying. Namely
figuring out where grep is installed on the system I'm logged in. Of course I could create another
non-color alias or have create a soft-link somewhere earlier in my path (so it's portable across
systems) but that's like covering up the problem than solving it.

The real issue is why can't /bin and /usr/bin be consistent across all Unix-based systems?

2015-Apr-17 | awk and shell
---------------------------

A peculiar problem with calling awk from a shell function, say I define a function that uses awk to
format something like this:

```shell
$bawk() { awk '{print $1",",$0}'; }
$echo How are you?|bawk Mark
How, How are you?
```
I need to change the awk in the shell function to

```shell
$ myawk() { awk -v name=$1 '{print name",",$0}'; }
$ echo How are you?|myawk Mark
Mark, How are you?
````

`awk -v` lets me create my own variables which are very useful in this situation.


2015-Apr-14 | Scripted Attacks
------------------------------

Noticed my site is now in the secondary stages of being attacked. The first is the usual set
of scripts which every site gets, these scripts have no underlying knowledge of the site they
attempt to attack. The next stage is looking for specific content based on some knowledge about
the website. This is where my site is right now. Log file activity indicates scripted attacks
are now coming in based on my site's content.

Fortunately, I'm using a jail within FreeBSD but there are still steps I could take to make my
setup a bit less vulnerable than it is at the moment. If the attack activity steps up, I might be
forced to lock down that next level of defense.

2015-Apr-14 | Scripting
-----------------------

One advantage to knowing shell scripting is the relative ease at which so many Unix utilities can be put
together to create something useful, and serves a purpose. Never mind if the intended recipient
of the purpose is me. The added bonus to my script today was that I was able to call an old and already useful script from the new script I wrote.

2015-Apr-11 | Daily Distractions
--------------------------------

Being distracted from what I want to do is a continuous battle I fight with on a daily basis. This
is made worse when I want to move ahead in some tech area but cannot either because I do not
have the skill set I need, or I am entrenched with a problem I want to solve, for example:
yesterday (Sunday) I literally spent the entire day working on Makefiles because I wanted to have
just one makefile on Linux (gmake) and BSD (bsd make). The syntax is similar but they do not even
agree on how to implement conditional statements. Finally I cracked and split the logic I was
having a hard time with into two
platform specific files. The gmake portion was much more approachable since I had already
done my time with learning gmake a long time back. BSD make, however, proved to be a completely
different beast from gmake (and nmake: Microsoft's archaic build tool). A simple task like trying
to figure out if a line was in a file or not took up all my time yesterday (all 16 hrs that is).
Finally dropped to sleep exhausted and then when I awoke at around 3:30 am (local time) I gave
it another grind through. Finally something clicked (possibly by accident, or my prayers were
finally answered) and I managed to move on from what was essentially a one-liner in a
platform-specific make file.

Why do I even bother anymore.

2015-Apr-11 | Makefiles
-----------------------

What a mess Makefiles are, especially since I want to keep one makefile on multiple systems, and,
no, Windows is not involved. This was simply GNU make and BSD make.


2015-Apr-10 | Pushing Prod
--------------------------

I was successfully able to install git in my jail without needing to install Perl in my jail environment.
Here's what I did:

* installed git in the non-jail part of the server
* copied the git binaries over
* ran ldd on the git binaries and made sure to copy the dependent libs over as well
* next I copied from /usr/local/libexec/git-core

It's not the complete install of git and keeping up-to-date will mean I manually need to repeat these
steps but I'm going to see how this works for now.

...and oh, if this get's pushed to GH then all steps have worked!

2015-Apr-9 | Broken streak
--------------------------

My longest streak of 95 days pushing something to GH, even if it was a small change, was just broken. It might
be time to becoming a paying member because the new ideas I have are best kept private.

2015-Apr-7 | Moving to ksh
--------------------------

I got frustrated with FreeBSD default shell (tcsh) handling of command line redirection, and then
after reading [Ellie Quigley's
book](http://www.pearsoned.co.uk/bookshop/detail.asp?item=100000000073401), I discovered why I got
frustrated with tcsh: it's the way tcsh does IO redirection compared to others, for example this one
was the straw that broke me:

```shell
# in sh/bash/ksh
 - command errors redirected to a file is: cmd 2>err
 - output and errors redirected to file is:  cmd > file 2>&1

# in tcsh the equivalents for each are
  - command errors redirected to a file is: (cmd > /dev/tty) >& err
  - output and errors redirected to file is: cmd >& err
```

Fortunately, @ibara has a [FreeBSD port](https://github.com/ibara/oksh) which compiled and
installed on FreeBSD and in the ezjail.


2015-Apr-5 | False positives
----------------------------

Weird false positive bug with valgrind on Ubuntu 14.04

```shell
$ g++ --version|head -1
g++ (Ubuntu 4.9.2-0ubuntu1~14.04) 4.9.2

$ cat leak.cpp
#include <iostream>
int main(int argc, char* argv[]){}
```

compiling with `g++ -std=c++0x -Wall` and running with ` valgrind --leak-check=full --show-leak-kinds=all a.out 2>leaky` gives

```shell
HEAP SUMMARY:
  in use at exit: 72,704 bytes in 1 blocks
  total heap usage: 1 allocs, 0 frees, 72,704 bytes allocated

  72,704 bytes in 1 blocks are still reachable in loss record 1 of 1
  at 0x4C2AB80: malloc (in /usr/lib/valgrind/vgpreload_memcheck-amd64-linux.so)
  by 0x4EC0DDF: ??? (in /usr/lib/x86_64-linux-gnu/libstdc++.so.6.0.21)
  by 0x4010139: call_init.part.0 (dl-init.c:78)
  by 0x4010222: _dl_init (dl-init.c:36)
  by 0x4001309: ??? (in /lib/x86_64-linux-gnu/ld-2.19.so)

LEAK SUMMARY:
  definitely lost: 0 bytes in 0 blocks
  indirectly lost: 0 bytes in 0 blocks
  possibly lost: 0 bytes in 0 blocks
  still reachable: 72,704 bytes in 1 blocks
  suppressed: 0 bytes in 0 blocks
```

Keeping my eye on this one.


2015-Apr-2 | Setting up 0xMF
----------------------------

Instantly after posting my previous edit, I decided to change the theme used by Hugo and voilà the
site renders somewhat correctly.  

Now for a few things that do not render correctly...and that might take a while.  

2015-Apr-1 | Setting up 0xMF
----------------------------

So I have a bunch of markup files, that render with Hugo, the static file generator written in Go. I
now want to move this to my production site. This involves integrating with my existing jail
setup and is running behind nginx routing.

Wow, just this simple thing alone is proving to be a challenge. Fortunately, the fallback
mechanism I setup earlier just works out of the box.

This might be resolved quickly, or who knows ... take a lot of time!

2015-Mar-31 | Setting up an ezjail
----------------------------------

Some more good advice [on setting up ezjail](https://www.secure-computing.net/wiki/index.php/FreeBSD_jails_with_ezjail) I liked the configurations for `/etc/rc.conf` and `/etc/periodic.conf`. All good settings, some of which I did not find reference to elsewhere.

2015-Mar-30 | nginx load balancing
----------------------------------

Successful at doing load balancing. These were the critical steps:

1. Create an `upstream` block in `nginx.conf`. Populate it with servers that will be handling
   requests. The [proxying configuration examples on the nginx](http://wiki.nginx.org/Configuration) are necessary and sufficient. Reading through the first few links were enough for me.

2. Create an `server` block in `nginx.conf`, for every server in the `upstream` block, particularly if using a single site configuration. Each server on the same physical server
   should be listening on ports other than port 80. It's best to have at least two servers (one
   acts as a backup) besides load balancing means sharing of load so a minimum of two would be
   expected.

3. Test and reload. That should be about it.

2015-Mar-28 | nginx blocking
----------------------------

I have been forced to learn a bit of `tcsh` scripting. This is a simple script that I wrote:

1. look at the `nginx-errors.log` file and grab all ips and the link requested
2. add these ips to the block list
3. restart `nginx`

2015-Mar-27 | nginx load balancing
----------------------------------

Next step is [load balancing using nginx](http://blog.jsdelivr.com/2013/01/nginx-load-balancing-basics.html). This will be interesting to combine nginx with proxy and other web servers.


2015-Mar-26 | FreeBSD man pages advice
--------------------------------------

Running

```shell
/usr/share/games/fortune freebsd-tips
```

and a bit of examples like `grep -C+3 "security" /usr/share/games/fortune/freebsd-tips*` helped me
unearth this treasure trove

```
Useful man pages recommended by /usr/games/fortune freebsd-tips

man boot0cfg  # repair damage to MBR caused by other OSes

man firewall  # advice on building a FreeBSD firewall
man hier      # the way FreeBSD systems are laid out

man intro     # 1:General 2:System calls 3:C libs 4:Device drivers
              # 5:File formats 6:Games 7:Misc 8 Sysadmin, 9 Kernel

man ports     # installing FreeBSD ports
man security  # advice on securing FreeBSD system
man tuning    # advice on performance tuning a FreeBSD system
```


2015-Mar-25 | jails
-------------------

Here is how I setup jails with Internet connectivity. Much of the posts deal with setting up a jail
but do not have instructions on getting Internet connectivity in the jail, so after struggling for
a bit, these instructions should help (and be reproducible):

1. Check Internet-facing ip, on iterface em0 (assuming em0 is the Internet-facing interface), with: `ifconfig em0` observe the ipv4 and ipv6 values

2. Follow instructions on creating the jail, with the caveat being to use the ip address got from step 1 using [FreeBSD Handbook](https://www.freebsd.org/doc/en/books/handbook/jails-ezjail.html). An example of such is shown in [Dan Langille's](http://dan.langille.org/2013/12/23/accessing-freebsd-jails-over-openvpn/) post on jails.

3. Copy the nameserver values from the host to jail's `/etc/resolv.conf`

2015-Mar-20 to 23 | nginx
-------------------------

Having decided to use nginx as a server on FreeBSD, these are the initial files I needed to know
about:

  /usr/local/etc/nginx => location of nginx conf files
  /usr/local/www => location of website files


Immediately noticed scripts hunting for 'admin' and other known weak links... so I decided to block
them using [access control with nginx](http://www.cyberciti.biz/faq/linux-unix-nginx-access-control-howto/) and this simple one-liner:

```shell
grep admin nginx-error.log | awk -F':' '{ print $6 }' | sed 's/ \(.*\),\(.*\)/deny \1;/' | sort >> nginx/blockips.conf
```

Not perfect, but it is a start.

Attempting to jail nginx, using ezjail but have not been too successful this far. An initial
[reference](http://blog.shatow.net/post/2013-11-27-sandboxing-php-part2.markdown) for further steps.

The biggest problem I had was to get networking going on in the jail. There was no way of
connecting to the Internet, until I realized my IP number that I was using was incorrect (it was
a local IP) when I should've been using an external facing IP. After I did that and connected the
nameserver (via /etc/resolv.conf). I was in business. [Dan Langille's](http://dan.langille.org/2013/12/23/accessing-freebsd-jails-over-openvpn/) post on jails was the tipping point for me. Now that I have Internet connectivity within the jail, I'm ok.


2015-Mar-9 | Narrowing
----------------------

I sometimes struggle with staying focused on the task at hand, because my mind and interests -
especially in tech - go far and wide. I've decided to do the following to help me stay focussed:

- Keep a pen and paper near my desk for notes for TODOs
- Occasionally when taking a break sort the notes
- Ruthlessly stay focused on the task at hand.

Let's track this over a period of time.

2015-Feb-3 | Insert Key on Lenovo
---------------------------------

The Lenovo keyboard is missing a couple of keys:

-  Fn + B = Break
-  Fn + I = Insert
-  Fn + P = Pause
-  Fn + S = Scroll Lock


2015-Jan-31 | git stop tracking
-------------------------------

Today was the first time I wanted git to stop tracking a file, so I did this:

```shell
git update-index --assume-unchanged local.bash
```

To go back to tracking the same file again, I'll use:

```shell
git update-index --no-assume-unchanged local.bash
```


2015-Jan-24 | FreeBSD links
---------------------------

Linux and BSD have different [concepts of load averages](http://undeadly.org/cgi?action=article&sid=20090715034920). Also, a good (albeit dated for an old version of FreeBSD) resource for [FreeBSD performance tuning](http://serverfault.com/questions/64356/freebsd-performance-tuning-sysctls-loader-conf-kernel)

2015-Jan-23 | vnc (update)
--------------------------

I managed to get vnc connectivity over ssh configured correctly from remmina. Big upside for
me.

A tip for all monitoring man pages installed:

```shell
man -k '*' | grep -E 'stat\((1|8)'
```


2015-Jan-22 | vnc
-----------------

For the first time ever, I've locked myself out of my remote VM. Luckily my provider let's me
access the VM through VNC. Steps to remote in via VNC is underway...and will update on how the
experience has been.

2015-Jan-19 | top
------------------
One of my FreeBSD virtual machines has rather high load averages (around 20% compared to typically
0% on other Linux and FreeBSD servers), trying to figure out what causes that load is a tricky
operation. Today I learned about `iotop` on Linux and 

```shell
top -mio -ototal 
```
on FreeBSD. This led me to discover that `devd` is probably causing the unusual load. Digging a bit
deeper, I discovered `/etc/devd.conf`. Now I have to plumb through the different options to figure
out which line from the 350+ lines is the problem. Fun times! 

2015-Jan-16 | careful with git commit -a
----------------------------------------
I usually always used `git commit -a` when making commits, but `git commit -a` is problematic when
using `git add -p` to slice and dice changes into separate commits. So the rule would be use `git
commit -a` unless using `git add -p`. 

List of git commands I've used so far (will update as I use more)

```shell

git add
git add -p

git branch

git checkout

git commit -am
git commit -m

git init

git log --graph --decorate --pretty=oneline --abbrev-commit
git log --graph --decorate --pretty=oneline --abbrev-commit --all

git pull
git push

git rebase
git reflog

git reset
git reset --hard

git status
git status -s
git whatchanged -p --abbrev-commit --pretty=medium

```


2015-Jan-15 | splitting a commit after updating master
------------------------------------------------------
Yesterday I used 

```shell
git add -p
```

to add 2 commits but I wanted the changes to be correctly split between the two commits, which did
not happen and now I am stuck. I need to rebase and re-commit but I am stuck about how to go about
this at the moment.

(_a few hours later_)

I managed to solve the problem, here's what I did:

* created a branch (`git branch mystic`)
* rebased back to where I wanted to be before the incorrect commits (`git rebased HEAD^^`)
* called `git add -p` to interactively stage the lines I wanted in each commit
* called `git commit -m` (**here's where my mistake** was last time, I had (*incorrectly*) used `git commit -am`)
* go back to branch master `git checkout master`
* merged commit `git merge mystic`
* I got a conflict so I used `git add` and `git commit` (no options) but it would have been better
  to simply edit the file and remove out the conflicting portions (that were demarkated by tags for
  master and mystic).

As a bonus, I managed to goof up again because I (foolishly, without understanding the consequences)
did a `git rebase`, which pretty much overwrote all my commits and I went back to where I was before
the merge. This meant I had to learn [`git
reflog`](http://stackoverflow.com/questions/134882/undoing-a-git-rebase).


2015-Jan-13 | mystic theme for Hugo
-----------------------------------
First step in creating [Hugo](http://gohugo.io) theme. Changed default theme license from MIT to ISC. This
will take time to shape up and I'll document each activity I do when I create the theme.
 
2015-Jan-12 | Old scripts that work
-----------------------------------
So much of my time is spent on learning new stuff that often I'm just happy when something I wrote
previously continues to work as expected, despite me not touching again for a very long time. All
too often I notice is the case with shell scripts. Once written they continue to work as expected,
and if they are helpful in their documentation on error conditions, then it's all the more better.

Here's something I discussed with a colleague at work earlier today. The notion of a sticky bit.
Sticky bit turned on a directory lets files created in the directory be deletable (if write
permission is given) only by the owner. This is useful in multi-user environments where several
users might want to create log files in a directory but you do not want users to accidentally (or
maliciously) delete each others files.

```sh
chmod +t test 
```

2015-Jan-11 | Hugo Static Site Builder 
--------------------------------------
I had to create a captain's log markdown file because Hugo does not accept AsciiDoc just yet.

# vim: spell:ft=markdown
