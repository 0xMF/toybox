Captain's Log
=============

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
