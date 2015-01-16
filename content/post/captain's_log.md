Captain's Log
=============

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
