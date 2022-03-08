# static
Build fast, lightweight blogs using only markdown.

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
theme: theme file
```

You can then start writing your blog posts by creating new files ending with `.md`.
A blog post file contains 3 areas: title, date, and the markdown content.

```md
Your blog post title
March 07, 2022 - 21:40

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex:
- Duis
- Excepteur sint occaecat
- Culpa qui officia deserunt mollit
```

* You must specify a date with this exact format or `static` won't be able to parse it
* You must leave a new line between the date and the markdown content.
* You can name your post files anything, just make sure they end with `.md`.
* static will generate a slug from the title you provide in the first line.

After following all the necessary steps, you could simply just run `static build`, static will go through the files and generate a `dist/` folder, you can then host it somewhere of your choice.

## Arguments
```
  -config string
    	config file path (default "config.yml")
  -out string
    	directory path where the generated files will be saved (default "dist")
```
You can then use something like: `static build -config /my-custom-path/my-conf.yml -out /somewhere`.

## Themes
A theme file is simply just a tar archive that contains 5 files:

- index.tmpl
- head.tmpl
- footer.tmpl
- article.tmpl
- style.css

`.tmpl` files are golang templates, you might need to take a look at the `default_theme` to get a good understanding.
