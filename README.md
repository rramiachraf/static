# static
Build fast and lightweight blogs using markdown.

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
footer: footer text
# theme field is optional, if not specified static will fallback to the default theme.
theme: theme file
```

You can then start writing your blog posts by creating new files ending with `.md`.  
A blog post file contains 3 areas: title, date, and the body written in markdown.

```md
Your blog post title
March 07, 2022 - 21:40

Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex:
- Duis
- Excepteur sint occaecat
- Culpa qui officia deserunt mollit
```

* You must specify a date with this exact format or `static` won't be able to parse it
* You must leave an empty new line between the date and the markdown content.
* You can name your post files anything, just make sure they end with `.md`.
* static will generate a slug from the title you provide in the title (first line).

After following all the necessary steps, you could simply just run `static build`, static will go through the files and generate the `/dist` folder with all the necessary files to host your blog.

## Flags
```
  -config string
    	config file path (default "config.yml")
  -out string
    	directory path where the generated files will be saved (default "dist")
```
So, you can something like: `static build -config /my-custom-path/my-conf.yml -out /somewhere`.

## Themes
A theme file is simply a tarball that contains 5 files:

- index.tmpl
- head.tmpl
- footer.tmpl
- post.tmpl
- style.css

`.tmpl` files are golang templates, you might need to take a look at the `classic` theme to get a good understanding.
