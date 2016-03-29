Robin
-----

This is a list of things I did to fix Robin Client for Android to the
way I wanted it to work. The list is in reverse chronological order to
make the more recent work done appear at the top

- forked [scruffyfox/Robin-Client](https://github.com/scruffyfox/Robin-Client) to [0xMF/Robin-Client](https://github.com/0xMF/Robin-Client) and added the following into .git/config
- added the following to .git/config

```
[remote "origin"]
        url = git@github.com:0xMF/Robin-Client.git
        fetch = +refs/heads/*:refs/remotes/origin/*
[remote "upstream"]
        url = git://github.com/scruffyfox/Robin-Client.git
        fetch = +refs/heads/*:refs/remotes/upstream/*
```

- discovered I need gradle.
