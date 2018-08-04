# MP3DB

This is a small console application which provides the ability to query mp3 files with using title
or artist names.

Before you can query the dabase , you have to import the mp3 files.

## Usage

The following commands are supported

- import <directory> creates a new database and imports all mp3 files into it
- search <searchterm> searches for the given searchterm in the following fields: artist, title,
  album. You can specify explicitely the field you want to search with using the field name as flag.
  For example `search --artist Abba`
- versions Shows you, wich version you've installed on your computer.
- serve --port=8082 provides a ready to use Rest server for searching songs from a web application

## Dependency management

This project is using **dep** fro the dependency management. To learn more about the tool, visit
[https://golang.github.io/dep/].

At first you've to install **dep**. You can do this with *pacman* or just install from the source
with `go get -d -u github.com/golang/dep`.

To add a new dependency manually, you have to enter 

```
dep ensure -add github.com/foo/bar github.com/baz/quux
```

In case you forgot to add a dependency before you used it, you can just execute `dep ensure`
afterwards. This will download the dependency to the vendor folder and update the **Gopkg.lock**

In case you don't want to checkin the *vendor* folder, you also need to enter `dep ensure` to download
all packages to the *vendor*, that will be created automatically from dep.

