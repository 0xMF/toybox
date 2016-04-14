My git use
==========

As I learn git, with its associated complexities, the time has come to write what I learn in git
in one place simply so I can go back to it later.

Initialization
---------------

There are two ways to do this:

  * start a repo from scratch
  * copy (in git parlance the term clone is used) an existing repo

##### Creates a new repo
```
$ git init repo_name
```

##### Clones an existing repo
```
$ git clone origin_repo_name
```

Branches
--------

A problem with working on master all the time is, after a bit, the git log gets ugly because many
small commits could've been merged into fewer larger and more complete ones at the end of the day
(session). ```git rebase``` works but that can be problematic when using ```git push -f``` simply to
clean up a master (```git push -f``` rewrites history and that can be problematic when changes made
to master have been shared-with-others).

A better solution to this problem is to have branches and use those branches for daily work; then
when branches commit logs look ok, merge that change onto master. This approach has the benefit of:

  * working on a branch without worrying about ruining a perfectly good master commit log
  * merging changes with master when done to continue that pristine commit log

The steps when using branches are:

  1. Create a branch

    ```
    $ git branch wip
    ```

  2. Switch to branch wip

    ```
    $ git checkout wip
    ```

  3. Work on branch and make all fresh changes on the branch. Some of
     these changes would include: squash, fixup, with their associated
     rebase.

  4. Finally when all is done on the branch and it is ready to merge,
     there are two strategies to choose from:

    - merge with a linear commit history on master (my preference)

    ```
    $ git checkout wip  # done in step 2. above
    $ git rebase master # alternatively: git rebase master origin/master
    $ git checkout master
    $ git merge wip
    $ git branch -d wip
    ```

    - merge keeping branch commit history

    ```
    $ git checkout master
    $ git merge wip
    $ git branch -d wip
    ```

    The only difference between the two approaches is doing a ```git
    rebase master``` prior to merging with master.

    The advantage of having a linear commit history on master is every
    commit on master would be a condensed commit of one or more squashed
    commits from its branches. There are no tramlines when viewing
    commit history which makes the project commit history readable. This
    approach is suited for single-developer or teams preferring linear
    commit history on master.

    The disadvantage of a linear commit history is the branch commit
    history is lost so this strategy might not work for teams or
    situations where maintiaining a record of branch history is
    important.


Remotes
-------

Here is the way I've setup my remotes

  1. Create an empty repo on GitHub, say toybox
  2. Clone this repo locally using (I have this setup on another drive)

    ```
    $ git clone --bare git@github.com:0xMF/toybox
    Cloning into bare repository 'toybox.git'...
    remote: .....more stuff...
    ```

  3. Clone the bare repo from step 2 in another directory on same drive

    ```
    $ git clone toybox.git toybox
    Cloning into 'toybox'...
    done.
    ```

  4. Clone the bare repo into another location, in my case I used ```$HOME:/repos/toybox```
  5. Add an upstream in second clone to github

    ```
    $ git remote add upstream git@github.com:0xMF/toybox
    ```

  6. Check remotes in repo

    ```
    $ git remote -v
    origin  /home/mark/share/repos/bare/toybox.git (fetch)
    origin  /home/mark/share/repos/bare/toybox.git (push)
    upstream        git@github.com:0xMF/toybox (fetch)
    upstream        git@github.com:0xMF/toybox (push)
    ```

What's the advantage of doing this? Heh! Right now I don't know but intuitively I've seen many posts
suggest a similar structure. As I become aware of these advantages, I'll update this file.


Pull/Merge/Rebase
-----------------

Having a simple, well documented, and linear git history is a prized
goal and it should be something to work towards as that makes
understanding the project easier, however, when sharing a repo across
machines I noticed using `git push` to post to GitHub and then using
`git pull` to pull the same repo but on another machine meant creating
non-linear git histories. This happened because `git pull` is in fact
two steps: `git fetch` to bring the new content and `git merge` to merge
that content. The `git merge` adds a branch merge into the git commit
history of that repo.

Needless to say, this isn't an ideal situation and the way out of it is
to avoid using `git pull` but use `git fetch` followed by `git rebase`
to add the new changes from the GitHub remote repo onto the local repo
(that wasn't the originating local repo). I created two shell aliases
for doing this:

```
alias gfr='git fetch; git rebase remotes/origin/master'
alias gpr='git pull --rebase'
```

Using either alias now ensures both local repos on different machines
keep a linear commit history.

There is a caveat, obviously, to this method because there might be
times when it is preferable to have multiple branched commit histories
appear (to tract changes made by different people or different features)
and in such cases the following workflow would be recommended

1. `git fetch`
2. `git diff`
3. `git merge remotes/origin/master`

Alternatively, when simply wanting to avoid the hassle of doing the
above three step process and trusting the merge without questions...then
in that case it would be fine to default to `git pull`


Editing commit messages
-----------------------

So many times I've got into the situation of pushing changes, yes even
to remote, and then realizing my mistake/typos/etc in the commit message
and that bothers me. Unfortunately there are two conflicting demands
going on when I realize the commit message(s) that I've already shared
on GitHub:

  1. The desire to fix the commit message
  2. The problem with a shared remote going out of sync

Git gives me several strategies to edit the incorrect/badly worded
commit message:

  1. For the most recent commit message: `$ git commit --amend`
  2. For a change earlier in the commit history: `$ git rebase -i` which
     is followed by:
      * pick - keep commit and commit message as is
      * fixup - keep commit but reuse an existing commit message as is
      * squash - keep commit but throw away the commit message
      * reword - keep comit but reword (edit) it's commit message

After editing the commit message, the problem of sharing the changes
still remains because the new SHA-1 that were created due to these
changes are no longer consistent with the remote repo which others are
using, so pushing to remote no longer works correctly unless I do a `git
push -f` but that is problematic for anyone who might have already got
the old (badly worded) commit history prior to me making those changes.
This would mean they would need to re-download the new changes that I
made after I used `git push -f`. Thus far this has been only me working
with a remote repo but as many people start contributing to a common
shared repo, the concept of using a remote branches for each developer
and then using `git push -f` **only** on that branch and not on `master`
is the way out of this situation.

The problem does not go away as there might still be the case where
after merging from contributors the git commit history still has typos
and other errors...in which case `git branch` to create a branch that is
synced with master; followed by `git revert`, on master, to revert
changes as if they never happened to master; followed by `git
cherry-pick/git patch` to get the commits in with the changes with the
commit log written the way they were wanted is the way out of this mess.

Not an easy task at all, still with a bit of understanding what to do
followed by doing this a coupled of times makes the workflow easier.

Unexplored Edge Cases
---------------------

(*consider this section wip*) A few edge case issues that I had using git but
have not yet found an aceptable workaround:

  - HOWTO filter out lines from a file rather than the file itself, has
    [this solution from
    StackOverflow](http://stackoverflow.com/questions/6557467/can-git-ignore-a-specific-line)
    my immediate workaround for that solution was to ignore the file
    completely with:
    
        git update-index --assume-unchanged [filename]
        git update-index --no-assume-unchanged [filename]

  - HOWTO fix commit tramlines in master

        git checkout -B branch_name <commit>
        ...
        git rebase master
        git checkout master
        git merge branch_name

<!--
# vim: spell:ft=markdown:nonu:nowrap:colorcolumn=0
-->
