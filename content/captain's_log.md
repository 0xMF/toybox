Captain's Log
=============

2016-Feb-25 | GitHub contributions
----------------------------------

Doing `git push -f` to master does not show up in the contributions for that day if the
changes had a fixup. In other words, if you have the following sequence of commits:

```
$ git commit --fixup ...
$ git rebase ...
$ git push -f upstream/master
```

then GitHub will not count that contribution for that day, even if I am the owner of the repo I
pushed to. This seems weird but maybe it's just me.  

2016-Feb-24 | broken vim
------------------------

While trying to sync two .vim repos, some settings broke. I'll need to dig deep to find out 
what broke and why. 

(update: 2016-Feb-25) I quickly found out why. I had an abbreviation expand that ruined my
local variable expansion so after correcting that, all seems to be fine.

(update: 2016-Feb-26) Added a few keymappings. 
(update: 2016-Feb-27) Took the leap and jumped into using dotfiles/repo/vim. This means all
subsequent changes will be tracked in the repo.

2016-Feb-19 | git squashing and shared folder symlinks
------------------------------------------------------

I liked this blog post explanation of [automatically squashing commits in git]
(https://robots.thoughtbot.com/autosquashing-git-commits). The central idea is to use

```
$ git commit --fixup :/earlier commit
```
with

```
$ git rebase --interactive --autosquash
```

to automatically squash older commits where a typo or minor mistakes had been made. The
```:/earlier commit``` detects which most recent commit message matches the text "earlier commit"
and issues a fixup for that commit so a subsequent rebase later would do a fixup on that commit
message. Setting 

```
$ git config --global rebase.autosquash true
```

automatically makes every rebase an autosquash.


In other news, running Debian guest on Windows brings up peculiar issues when using shared folders.
For example I had to use 

```
C:\> VBoxManage setextradata VM_NAME VBoxInternal2/SharedFoldersEnableSymlinksCreate/SHARE_NAME 1
```

based on [this VirtualBox bug and workaround](https://www.virtualbox.org/ticket/10085).
Note: VM_NAME and SHARE_NAME need to be changed appropriately.


2016-Feb-17 | Virtual Network Computing
---------------------------------------

First time I've tried out Virtual Network Computing. I used [TigerVNC](http://tigervnc.org/).
A couple of points to keep in mind when using VNC:

* A VNC server needs to be installed on remote
* A VNC client on local
* The remote VNC server should allow vnc server packets across
* Start the vnc share to the local desktop (unprivileged user)
* When connecting, with TigerVNC, in full screen mode use F8 to exit.

More details later...

2016-Feb-15 | Copying Debian Packages
-------------------------------------

Tips to [copy installed packages from source to target](http://qref.sourceforge.net/Debian/reference/ch-package.en.html)

```
# source 
# dpkg --get-selections  > debian.installed_packages
# debconf-get-selections > debian.debconf_selections

# target
# dselect update
# debconf-set-selections < debian.debconf_selections
# dpkg --set-selections  < debian.installed_packages
# apt-get -u dselect-upgrade    # or dselect install
```

2016-Feb-03 | VirtualBox on USB
-------------------------------

This is a very helpful post to making a [bootable USB run as a VM in VirtualBox](http://www.howtogeek.com/187721/how-to-boot-from-a-usb-drive-in-virtualbox/). The basic command for this activity is:

```
VBoxManage internalcommands createrawvmdk -filename C:\usb.vmdk -rawdisk \\.\PhysicalDrive#
```

The PhysicalDrive# is got from Disk Management (Windows Key + R and type ```diskmgmt.msc```). Once the
vmdk file is created (it will be <1K), go through the steps of making a VM, do not make a separate
VM disk instead point the disk to ```C:\usb.vmdk```.

Remember all process shown above require Administrator access including running the VM off the USB.


2015-June-17 | Learning continuously
------------------------------------

One of the hardest things for me to do is to write code, small pieces each day and post (even if it
does not work) to GitHub. There are several factors for this:

  - cannot focus on a language or a project for a sufficient duration of time
  - if I do, then leaving it and getting back into it is hard. Very hard
  - distractions from newer, fancier, cooler projects tends to derail 

Having written this now, let me think over it for a couple of days.

2015-June-17 | Win 8.1 local and online accounts
------------------------------------------------

I hate being logged with 'online only' accounts such as Microsoft's Online Accounts for Win 8.1.
Recently being forced to use Skype meant I reached for the Skype Metro app, which meant I was forced
to upgrade my local account to Microsoft's Online Account. Only much later, I found a way to convert
Win 8.1 [Online to Local
Account](http://www.nextofwindows.com/windows-8-1-how-to-convert-windows-live-account-to-local-account/)

2015-June-16 | solved problems
------------------------------

Most programming related problems are solved by looking at code that is similar to another problem
but will work with small changes on the existing one. This means when learning something new, that
database or memory map of howto do stuff needs to be learned and until that is done, everything will
be trickier than need be because the mental models aren't in place (yet).


2015-June-15 | git Protocol
---------------------------

A great list of steps for [maintaining a git repo](https://github.com/thoughtbot/guides/tree/master/protocol/git) especially the steps involved in writing a feature, the sequence of steps is as follows:

```shell
git checkout master
git pull
git checkout -b <feature-branch-name>

# get new changes (if any)
git fetch origin
git rebase origin/master

# stage and commit
git add --all
git status
git commit --verbose
(git rebase -i orgin/master)

# create remote branch
git push --force origin <feature-branch-name>

# merge locally and push remote
git log origin/master..<branch-name>
git diff --stat origin/master
git checkout master
git merge <branch-name> --ff-only
git push

# delete remote and local branches
git push origin --delete <feature-branch-name>
git branch --delete <feature-branch-name>
```


2015-June-14 | Solarized
------------------------

Finally discovered and started using [Solarized=dark](http://ethanschoonover.com/solarized), and got
it installed on all my regular workflow. Very good in all the lighting conditions I use so far.
These are the setting values I needed for dark:

  Background RGB:   0,  43,  54
  Foreground RGB: 147, 161, 161

A complete list of all Solarized values, for both [light and dark are here](http://ethanschoonover.com/solarized#the-values)

2015-June-09 | From tmux-Cygwin pane ssh into FreeBSD local VM
--------------------------------------------------------------

So much FreeBSD! Now I'm running FreeBSD in a VirtualBox image on Windows. I even managed to
configure it as a Scheduled Task to run whenever I login with these settings

```
Settings for Task Scheduler:
- Edit Trigger Tab:
    * Begin the tast: On connection to user session
    * Connection from local computer
- Action Tab:
    * Start a program: (calls script 'freebsd.cmd' see below)
- Settings Tab (enable these, disable all else):
    * Allow task to be run on demand
    * If the running task does not end when requested, force it to stop
    * If the task is already running, then the following rule applies:
      Do not start another instance

freebsd.cmd
@c:\path\\to\Oracle\VirtualBox\VBoxManage startvm "name of FreeBSD vm" --type headless

```

This should now set me up for a full blown Unix environment I am used to within my Windows. Will
test and update if anything changes.

2015-June-09 | Raspberry Pi as router using FreeBSD
---------------------------------------------------

Ideas come unexpectedly and the I feel pressurized to try it out. As my existing FreeBSD PC-based
router is on it's last legs and I was looking at my Raspberry Pi doing nothing. I wondered if
FreeBSD was supported on Raspberry Pi. It is and has been since [Jan 2013](https://www.raspberrypi.org/freebsd-is-here/#comments) and is supported in [FreeBSD/Arm Tier 2](https://www.freebsd.org/platforms/arm.html).

A quick scan also reveal an [interesting blog post](http://qcktech.blogspot.ca/2012/08/raspberry-pi-as-router.html) with [associated video showing proof of concept](https://www.youtube.com/watch?v=l-_J2po3i0U).

Will update with progress as I step through this setup over the next few days.

2015-June-08 | Old example and .Net
-----------------------------------

I use .Net, which means I use C# and VB.Net. Earlier today, I thought I had written code for
something using C# but it turned out my code was actually written in VB.Net. I was searching for it
in the wrong place!

The point of this observation is the seamless nature of .Net. I was so sure I solved the problem in
C# that it did not occur to me (until much later) that my code could have been in an entirely
different language!

2015-June-07 | updating cygwin64 and tmux
-----------------------------------------

I got tmux to work on cygwin. Yaay! Sequence of steps:

```shell
apt-cyg install tmux
apt-cyg install mintty
mintty
tmux
```

Here are some relevant .tmux.conf settings I use

```shell

# use screen style prefix key
set -g prefix C-a
unbind C-b
bind C-a send-prefix

# no status bar
set -g status off

set -g default-terminal "xterm-256color"
set -g default-command "${SHELL}"

# avoid selecting by mouse so we can use copy paste
set-option -g mouse-select-pane off

# select window
bind '"' choose-window

# split windows like vim
# vim's definition of a horizontal/vertical split is reversed from tmux's
bind s split-window -v
bind v split-window -h

# move around panes with hjkl, as one would in vim after pressing ctrl-w
bind h select-pane -L
bind j select-pane -D
bind k select-pane -U
bind l select-pane -R

# resize panes like vim
# feel free to change the "1" to however many lines you want to resize by, only
# one at a time can be slow
bind < resize-pane -L 1
bind > resize-pane -R 1
bind - resize-pane -D 1
bind + resize-pane -U 1

# vi-style controls for copy mode
setw -g mode-keys vi
```

I had a small issue with git pushing to GitHub after I started using tmux, msysgit was behaving
strangely, so I decided to `apt-cyg install git` and run with that (I had to rename my existing
msysgit folder to avoid path conflicts). With that out of the way, git (from Cygwin) worked. The
good part of the Cygwin port is git run within the Unix-like context unlike mysgit (which does use
Windows pathnames and a few Windows-like settings). If this push goes through then git worked!


2015-June-06 | updating cygwin64 and Term::ReadKey
--------------------------------------------------

After updating cygwin64 via `apt-cyg` I noticed my `Term::ReadKey` Perl module stopped
working and so had to reinstall making sure the right dependencies, extra in place:

```shell
apt-cyg install gcc-core gcc-g++ libcrypt-devel openssh-devel make
```

and later install the Perl module via

```
cpan install Term:ReadKey
```

2015-June-05 | pandoc for static content
----------------------------------------

pandoc is useful for generating static content, for example:

```shell
pandoc -s -S --toc -c pandoc.css  B_adding_pages.md -o B_adding_pages.html
```

The only problem I've noticed is after generating the HTML files, and opening them directly from
a browser, images do not get loaded so I would need a simple webserver (or since I'm already on
Windows just use IIS). This would be the next step.

2015-June-03 | parsing HN
--------------------------

An initial  parse of Hacker News:

```shell
curl "https://news.ycombinator.com"\
 |sed -e 's/<[b-zA-Z\/][^>]*>//g'\
 |grep " *[0-9][0-9]"\
 |sed 's/<a id=[^>]*>//'
```

2015-May-27 | Contributions on GH
---------------------------------

The contributions graph on my GH profile page is addictive. I've been trying to make at least one,
some days small, change to something that is publicly available to get a checkbox on my contribution
for that day. Here's what I discovered today: opening a pull request on a forked repo acts as
a contribution but subsequent updates to master on the forked repo no longer count as contributions.
This means the only time GH will view my contribution on a forked repo is if I open a PR each time,
which is pretty difficult to do on a daily basis don't you think? Hmm....

2015-May-24 | Cygwin ls colors
------------------------------

One nitpick I have with cygwin `LS_COLORS` is it shows permissions as other's writable. On a machine
that only I use permissions of Windows and Linux differences don't matter much to me, however, I had
a problem finding which value it was that wasn't being changed. I kept thinking it was `di` for
directory, until I got it was `ow`. Here are my `LS_COLORS`

```shell
LS_COLORS='rs=0:di=01;33:ln=01;36:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:'\
'su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=01;31;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arc=01;31:'\
'*.arj=01;31:*.taz=01;31:*.lha=01;31:*.lz4=01;31:*.lzh=01;31:*.lzma=01;31:*.tlz=01;31:*.txz=01;31:'\
'*.tzo=01;31:*.t7z=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.dz=01;31:*.gz=01;31:*.lrz=01;31:*.lz=01;31:'\
'*.lzo=01;31:*.xz=01;31:*.bz2=01;31:*.bz=01;31:*.tbz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:'\
'*.rpm=01;31:*.jar=01;31:*.war=01;31:*.ear=01;31:*.sar=01;31:*.rar=01;31:*.alz=01;31:*.ace=01;31:'\
'*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.cab=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:'\
'*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:'\
'*.tiff=01;35:*.png=01;35:*.svg=01;35:*.svgz=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:'\
'*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.webm=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:'\
'*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:'\
'*.avi=01;35:*.fli=01;35:*.flv=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:'\
'*.cgm=01;35:*.emf=01;35:*.axv=01;35:*.anx=01;35:*.ogv=01;35:*.ogx=01;35:*.aac=00;36:*.au=00;36:'\
'*.flac=00;36:*.m4a=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:'\
'*.ra=00;36:*.wav=00;36:*.axa=00;36:*.oga=00;36:*.spx=00;36:*.xspf=00;36:';
export LS_COLORS
```

Along the way, I discovered

```shell
dircolors -b
dircolors -p
```


2015-May-23 | Color and text settings
-------------------------------------

I've become accustomed to a few settings:

- Colour scheme (mystic):
  * background: #082040 (8,32,64)
  * foreground: #eeeeec (238,238,236)

- Text
  * font: Source Code Pro Medium: 12pt

This is a start, obviously, but at least I'm hoping I wont go crazy trying to achieve the colour
combinations I am used to on the various desktops I use.

2015-May-23 | AsciiDoc (LaTeX) on Cygwin
----------------------------------------

Successfully imported shell scripts to generate PDFs with my custom LaTeX settings under Cygwin. It
was a pain to get some dblatex macros to work with my settings so I commented them out. Fortunately
for the work I'm doing it doesn't seem to have made a difference and the test files I used to
recreate the PDF documents worked fine.


2015-May-22 | pandoc and MiKTeX
-------------------------------

Yup, I gave in and installed MiKTeX, so now there is no need for me to boot into that Ubuntu VM (at
least for the next several weeks) as all my tools work on Windows.


2015-May-21 | Entity Framework Exceptions
-----------------------------------------

Gah! EF problems with scaffolding. Tired. Nuff said.


2015-May-20 | More love for pandoc
----------------------------------

Here's the batch file I use for creating revealjs slides

```batchfile
@echo off

@set REPOS=path\to\repo
@set DATA=%REPOS%\pandoc
@set CSS_THEME=%DATA%\css\some_theme.css
@set SLIDE_TYPE=revealjs

@set EXTENSIONS=markdown+^
definition_lists+^
escaped_line_breaks+^
example_lists+^
fancy_lists+^
fenced_code_attributes+^
fenced_code_blocks+^
grid_tables+^
line_blocks+^
markdown_in_html_blocks+^
multiline_tables+^
pandoc_title_block+^
pipe_tables+^
raw_html+^
simple_tables+^
startnum+^
table_captions+^
yaml_metadata_block

pandoc -f %EXTENSIONS% -c %CSS_THEME% -t %SLIDE_TYPE% --self-contained -o %~n1.html %*
```

2015-May-16 | More love for pandoc
----------------------------------

I started to use pandoc regularly, even started to learn it's own markdown syntax. Configured vim
with vim-pandoc syntax highlighting. The depth and breadth of output that can be generated from one
(pandoc markdown) source document is incredible. I am using

  * revealjs to make slides
  * odt/docx for creating docx/odt files from source
  * pdf/epub (haven't tried this yet but will do so soon)

The only regret I have this far is the limited support pandoc uses for AsciiDoc input. In fact there
is no input support for AsciiDoc files (as is), however, pandoc can write to AsciiDoc but I have not
tried that yet.

2015-May-16 | reveal.js and pandoc
----------------------------------

Spent a lot time over reveal.js and pandoc. reveal.js certainly simplifies presentations without
needing external apps. Got a browser with JavaScript enabled, then you've got a presentation. Here
are some shortcuts for reveal.js

  Up/Down/Right/Left  arrows work as expected
  . toggles           dark/light on current slide
  ESC                 brings up the (alternate) navigation window
  F                   toggles full screen
  s                   toggles speaker notes

Other goodies:

  - The given sample stylesheets are easy to customize too.
  - Appending ?print-pdf puts all the slides into one long webpage which can be printed to PDF in
    Chrome using settings: Layout:Landscape; Margins:None; Pages:All 

2015-May-15 | reveal.js and pandoc
----------------------------------

WHOAH! Now, I know why the cool kids use reveal.js and pandoc. Sadly, I am still figuring out the
syntax needed to create the slides that reveal.js runs and when I figure that you, it will be my
next post.

2015-May-13 | nice_social and IE11
----------------------------------

This is what I did to get [nice_social](https://github.com/matigo/nice_social) to work with
IE11 on Windows:

  * Clone the repo (assume ```C:\nice_social``` or ```C:\inetpub\nice_social```)
  * [Create a virtual directory](https://support.microsoft.com/en-us/kb/172138) using [IIS Manager](https://msdn.microsoft.com/en-us/library/bb763170(v=vs.140).aspx)
  * You might have to configure nice_social to work from a port other than 80 if that has already
    been used by another site (assume ```8080```)
  * Right click on the folder, go to the **Security** tab and add user ```IIS_IUSRS```.
  * Give the following permission to ```IIS_IUSRS```
    - Read
    - Read and execute
    - List folder contents
  * Open up IE11 and browse ```http:\\localhost:8080```

Caveats

  * Avoid creating an IIS virtual directory within a VirtualBox shared sub-folder. So, if you
    have a VirtualBox shared folder ```C:\share\repos``` that you use to share with other operating
    systems do not make ```c:\share\repos\nice_social``` an IIS website otherwise additional steps to
    configure your set up maybe involved. You could, however, clone nice_social within a sub-folder
    of ```C:\inetpub``` and configure ```C:\inetpub\nice_social``` as remote of
    ```C:\share\repos\nice_social``` thereby using your existing setup with multiple operating
    sytems and having a Windows test machine available for testing nice_social with IE11.

Upon further testing I discovered another problem. My current setup is as follows:

```
c:\share\repos\nice_social  <= dev folder shared with Ubuntu VM 
c:\inetpub\nice_social      <= windows production

# C:\share\repos\nice_social
git remote -v
origin  git@github.com:0xMF/nice_social.git (fetch)
origin  git@github.com:0xMF/nice_social.git (push)
upstream        git@github.com:matigo/nice_social.git (fetch)
upstream        git@github.com:matigo/nice_social.git (push)
win     c:\inetpub\nice_social (fetch)
win     c:\inetpub\nice_social (push)
```

The trouble is: I could not do a ```git push win``` from within ```c:\share\repos\nice_social``` to
```c:\inetpub\nice_social``` until I did a [git config --local receive.denyCurrentBranch
updateInstead]((http://stackoverflow.com/questions/2816369/git-push-error-remote-rejected-master-master-branch-is-currently-checked)
when in ```c:\inetpub\nice_social```

To be sure I also added my user to have permissions (Full Control) in ```c:\inetpub\nice_social```.
More testing is still needed to check whether the additional (full control) permissions were really
necessary.


2015-May-11 | Texapp on Windows
-------------------------------

I successfully got Texapp (0.6.10) to install and work correctly on Windows. Here's what I did:

  1. installed Cygwin
  2. from the Cygwin setup itself I installed Perl (5.14):
    * perl
    * perl-base
    * perl-pods
    * Perl-Term-ReadKey
    * Perl-WWW-Curl
    * Perl-Win32
    * Perl-YAML
    * Perl-JSON-PP
    * Perl-CPAN-\*
  3. from CPAN I updated and reloaded cpan then I installed Term::ReadLine::TTYtter
  4. installed Texapp

Not all of the above might be needed but after fouling up my Perl installation once, the only way to
recover was to recreate another Cygwin install. Please note, I also have msysgit installed, which
comes with it's own Cygwin Perl (5.8) and utilities needed to get git running. So, yes, this means
there are two cygwin installations.

The biggest downside to this process was the audio bell beeping, which I could only get rid of by
turning the [Default Bell sound off on
Windows](http://www.7tutorials.com/how-disable-system-beep-windows-7). 

2015-May-08 | git on Windows
----------------------------

After a long time, I am back to using Windows for development. This means an upgrade to:

  - git 
  - vim
  - conemu
  - cygwin

If this push goes through then I know all my settings on Windows worked correctly. I had to jump
through a lot of hoops since my existing ssh keys did not work so I had to recreate new ones and
upload them to GitHub. The command to generate the keys was influenced by [stribika](https://stribika.github.io/2015/01/04/secure-secure-shell.html) and it was similar to:

ssh-keygen -t rsa -b 4096 -o -a 100 -C (your email address to associate with the key)

The other problem was getting my private key to be remembered, instead of having to add it each
time. The way around that was solved with a helpful blog post from [Ryan Lanciaux](http://ryanlanciaux.github.io/blog/2014/05/15/running-ssh-agent-on-windows/). I installed cygwin and started up bash, then

```shell
eval `ssh-agent`
ssh-add
```

If you're seeing this, then you know it worked! That did, but it brought up another host of
problems, namely every new shell that got created started a new instance of ssh-agent. The solution
around this, eventually came [from GitHub itself](https://help.github.com/articles/working-with-ssh-key-passphrases/#platform-windows). 

There was a slight tweak I had to make to remove superfluous errors, like these:

```shell
bash: $'\r': command not found
bash: $'\r': command not found
bash: $'\r': command not found


# in the given functions I changed
agent_load_env() {
    . "$env" >/dev/null
}

# to ignore error output
agent_load_env() {
    . "$env" 2>/dev/null
}

```




2015-May-07 | Reflections
-------------------------

Starting a new file called [reflections](reflections.asciidoc) based on bullet point lessons I
learned recently.


2015-Apr-29 | VisualStudio on Ubuntu
------------------------------------

This is a simple test of running VS in Ubuntu after getting to know that it was released at //build2015. Wow.
Currently testing the .git features....and it works! *WOW!!*

2015-Apr-27 | Janitorial Duties
-------------------------------

Some amount of janitorial work is necessary to have an efficient workflow. Alas, I discovered this
too late. I'm hoping I can rectify this going forward.

2015-Apr-23 | remote gvim
-------------------------

Today's tip is tremendously useful yet surprisingly simple. To open several files in an already open
instance of gvim use:

```shell
gvim --remote-tab-silent `find $HOME/.*rc`

# if you have the following alias in your shell rc file
alias gv='gvim --remote-tab-silent'

# then, this is enough
gv `find $HOME/.*rc`

```

Neat. Huh?

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

<!--
# vim: spell:ft=markdown:tw=100:nonu
-->
