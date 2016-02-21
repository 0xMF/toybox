My git use
==========

As I learn git, with its associated complexities, the time has come to write what I learn in git
in one place simply so I can go back to it later.

Initialization
---------------

There are two ways to do this:

  * start a repo from scratch
  * copy (in git parlance the term clone is used) an existing repo

#### Creates a new repo
```
$ git init repo_name
```

#### Clones an existing repo
```
$ git clone origin_repo_name
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
