# static
Build fast and lightweight blogs using markdown.

## Features
- Fast
- Minimalist
- RSS feed
- Easy to setup
- Highlighted code blocks

## Install
Go 1.16+:
```
go install github.com/rramiachraf/static@latest
```

## Usage
Create a directory to host your blog posts, and in there create a `config.yml` with the following fields:

```yaml
title: blog title
description: blog description
url: example.com #must be provided if you want RSS
```

You can then start writing your blog posts by creating new files ending with `.md`.  
A blog post file contains 2 areas seperated with an empty line, metadata and the post content.

```md
TITLE Your blog post title
DATE 07/03/2022 21:40

Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex:
- Duis
- Excepteur sint occaecat
- Culpa qui officia deserunt mollit
```

* You must specify a date with the format `DD/MM/YYYY HH:MM` or `static` won't be able to parse it.
* You can name your post files anything, just make sure they end with `.md`.
* static will generate a slug from the title you provide in the title.

After following all the necessary steps, you could simply just run `static`, static will go through the files and generate the `./dist` folder with all the necessary files to host your blog.

## Flags
```
  -config string
    	config file path (default "config.yml")
  -out string
    	directory path where the generated files will be saved (default "dist")
```
So, you can do something like: `static -config /my-custom-path/my-conf.yml -out /somewhere`.
