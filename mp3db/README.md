# MP3DB

This is a small console application which provides the ability to query mp3 files with using title
or artist names.

Before you can query the dabase , you have to import the mp3 files.

## Usage

The following commands are supported

- import <directory> creates a new database and imports all mp3 files into it
- update <directory> updates the database with new or changed files from the passed directory
- cleanup goes through all records in an existing database and checks if the file still exists. If
  not, the record will be removed
- search <searchterm> searches for the given searchterm in the following fields: artist, title,
  album. You can specify explicitely the field you want to search with using the field name as flag.
  For example `search --artist Abba`
- versions Shows you, wich version you've installed on your computer.
- serve --port=8082 provides a ready to use Rest server for searching songs from a web application
  The following routes are available:
  - GET /songs?q=<searchTerm>&artist Searches for all songs that matches the given searchTerm. In no field
  are passed as filter it will search within **artist**, **title** and **album** fields
  - GET /songs/%id Return the song by the given id
  - PUT /songs/%id Writes the songs that was posted in the body to the database and updates also the
    tags in the underlying mp3 file
  - DELETE /songs/%id Deletes the song by the given id from the database and as well from the disk
  - GET /songs/%id/content Returns the content of the song with the given id for streaming on the
    web client

## Dependency management

Since version 1.11 of Go it's no longer necessary to put your go code strictly into a subfolder of
the GOPATH folder. Because of the new module system, it's now possible to have the code somewhere
else.

Everything around the new Golang module system is described
[here](https://github.com/golang/go/wiki/Modules).

But to summarize it, all you have todo is to delete any file of a previously used dependency
management system and as well also the vendor folder. After it you just have to enter

```
go mod init github.com/janbaer/mp3db
```

This will create a **go.mod** file, where all the dependencies will be saved.

Now just enter **go build** and Golang will automatically all the dependencies it found from your
imports. In case you want to update a dependency afterward to us a newer version, just enter `go get
-u` and it'll update the **go.mod** file. To install a specific version, you just have to the the
specific version at the end of your command.

```
go get github.com/gorilla/mux@v1.6.2
```

Just for to case you still want to save your external dependencies as part of your source code
within a vendor folder, you have to enter `go mod vendor`

## External dependencies

For the usage of go-taglib it's necessary to install the static **taglib** libraries before. So you
have to enter `sudo apt install libtagc0-dev` before.
