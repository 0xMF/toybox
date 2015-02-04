Captain's Log
=============

2015-Feb-3 | Insert Key on Lenovo
---------------------------------

The Lenovo keyboard is missing a couple of keys:

  Fn + B = Break
  Fn + I = Insert
  Fn + P = Pause
  Fn + S = Scroll Lock


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
