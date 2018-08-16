# MyMusic

This repository contains the source code for my personal music management project.
Its for my personal needs and it does only what's important for my personal workflows.

When anyone is interested to use it, feel free to fork it, but I don't guarantee for anything and I
also don't take care on any comptibility aspects.

The project contains three different parts.

1. Web - contains the source for the website. It's using Parcel.js for bundling and Preact.js
   together with Bulma.css implementing the web app.

2. MP3B - This is the backend and it's written in GO. So it's a console application that provides
   commands for importing a complete directory tree with MP3 files into a Sqlite database.
   These database can be updated or cleaned later when the directory tree has changed.
   And of course you can search for titles with the command line.
   But MP3DB provides also a webserver with a REST interface for searching in the database.
   This REST interface is used by the web application.

3. Deploy - This is for deploying the whole project to my Synology-DS, build and run it there
   with **docker-compose**. It's a really great feature that Synology supports Docker, so I've my own
   server at home for private things.

For more explanations read the README files in the subdirectories.
