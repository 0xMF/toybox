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
    $ git branch 0xMF_dev
    ```

  2. Switch to branch 0xMF_dev

    ```
    $ git checkout 0xMF_dev
    ```

  3. Work on branch and make all fresh changes on the branch. Some of these changes would include:
     squash, fixup, with their associated rebase.

  4. Finally when all is done on the branch and it is ready to merge; switch to master and merge:

    ```
    $ git checkout master
    $ git merge 0xMF_dev
    ```


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

<!--
# vim: spell:ft=markdown:tw=100:nonu:nowrap:colorcolumn=0
-->
