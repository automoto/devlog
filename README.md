### Dev Log
[![CircleCI](https://circleci.com/gh/automoto/devlog.svg?style=svg)](https://circleci.com/gh/automoto/devlog)

*The software development log to reflect and capture your thoughts, progress and TODOs.*

Devlog's goal is to help you create a "Development Log" to reflect and generate notes during and after a coding session.

<!-- MarkdownTOC autolink="true" -->

- [What is a Development Log?](#what-is-a-development-log)
- [What is Devlog?](#what-is-devlog)
- [Install](#install)
	- [Install Using Go](#install-using-go)
- [Using Devlog](#using-devlog)
- [Configure](#configure)
	- [Configuration using command line options:](#configuration-using-command-line-options)
- [Customizing Questions and Content of Devlog](#customizing-questions-and-content-of-devlog)
- [Building Locally](#building-locally)
- [Contributing](#contributing)

<!-- /MarkdownTOC -->


#### What is a Development Log?

A development log is a like a software development journal that you fill out during and after a coding session. It's great for reflecting on how coding sessions went, doing brain dumps and documenting TODOs.

Filling out these type of "development logs" and notes about coding sessions is inspired by the [note taking practices of prolific Doom/VR developer John Carmack](https://news.ycombinator.com/item?id=12575501). This can also be useful when paired with a "shutdown routines" at the end of the session of intense work to help us mentally disconnect from our work. More about "shutdown routines" [1](https://www.calnewport.com/blog/2009/06/08/drastically-reduce-stress-with-a-work-shutdown-ritual/)[2](https://www.calnewport.com/blog/2012/08/02/work-less-to-work-better-my-experiments-with-shutdown-routines/).

#### What is Devlog?

Devlog generates a customizable, time stamped, simple, "development log" markdown file to a directory you specify. This gives us alot of flexibility, you can save your devlog files into a git repository or to a cloud file service directory like dropbox, google drive, or one drive for automated syncing and backup.

Devlog prioritizes:
- Open standards over closed. Keep your notes in markdown files that can be queried for easily in a directory, not locked into some vendors service or custom formatting standards.
- Ease of use and simplicity. This is not meant to be a complex static content generator. It strives to be an easy to configure and create customizable markdown files to fill out in a text editor of your choice.

#### Install
The easiest way to install devlog is to use the provided installation script
``` sh
wget https://raw.githubusercontent.com/automoto/devlog/master/scripts/get-devlog.sh
# you're welcome to examine the script, it just grabs the latest release from github for your OS and installs it
sh get-devlog.sh
```

If you prefer to install it yourself, you can get the latest release binary directly from github https://github.com/automoto/devlog/releases and extract it to your `/usr/local/bin`.

##### Install Using Go

If you already have an updated version of go lang, installing via go is easy:
`go get -u github.com/automoto/devlog/`


#### Using Devlog
Ideally you will run Devlog during or towards the end of your development session and it will be filled out with your notes and answers to the questions when your finished. The "Notes" section at the bottom is great for TODOs or other information you want to keep. [Here is an example of a filled out devlog](https://gist.github.com/automoto/15e037d40258df1b8c2394ba1bae2c07). 

Devlog is designed to automate generating a development log markdown file for you to fill out in your favorite text editor. Once you have installed it, just type `devlog` and it will generate a time stamped markdown document. 

```shell
# Generate a markdown document log in the current directory
devlog

# Output the log document to your terminal instead of a document
devlog -p 'stdout'

```

By default, it returns the path of the document created so you can input this into a text editor easily. Here is a way you can generate a markdown document and open it in some popular text editors: 

```shell

# create a new Devlog file and open it in SublimeText
subl `./devlog | tail -n 1`

# create a new Devlog file and open it in vim
vim `./devlog | tail -n 1`

# create a new Devlog file and open it in nano
nano `./devlog | tail -n 1`
```


#### Configure
By default `devlog` will generate a markdown file in the current directory unless you specify the directory via setting environment variables or through command line options you set.

*Configuration using environment variables:*

Set the directory to save devlog files to:
```
export DEVLOG_DIR="/home/your_username/your_directory"
```

Set the default template for your devlog files content:
```
export DEVLOG_CONTENT="/home/your_username/your_directory/custom.gohtml"
```


You can override this by temporarily setting the value when calling devlog

```
DEVLOG_DIR="/home/your_username/other_directory" devlog
```

##### Configuration using command line options:

You can also pass in configurations via command line options. Command line options take precedence over configurations set via environment variables.
```
devlog -p "/home/your_username/your_directory" -t "/home/your_username/your_directory/custom.gohtml" 
```
To view all the possible command line options, just pass in the `-h` command line option for help e.g. `devlog -h` 

#### Customizing Questions and Content of Devlog
You can customize the sections of your Devlog by creating a `.gohtml` file and specifying your custom content there.
``` gohtml
### Development Log
*created: {{.FormattedCurrentTime}}*

##### Notes

**TODO**

- [ ]
- [ ]
```
Then pass in your configuration file to devlog:
```
devlog -t your_custom_questions_file.gohtml
```

#### Building Locally

*Prerequisites:*
  - go lang version 1.13

Simply clone this repository and run the following command to build the binary:
```shell
make build
```

This will create a binary locally you can run commands against already, like so:

`./devlog`

Build and copy the binary to your local bin to access the CLI anywhere. It will likely prompt you for a password since it's needed to install things to your `/usr/local/bin`.

```shell
make install
```

Now you can run the command `devlog` from anywhere to generate a new devlog file:

``` shell
devlog
2019/09/02 22:00:32 Successfully saved dev log to directory: /home/dev/null
```

#### Contributing

The `Makefile` has most of the commands you need to build, lint, and run the tests locally. Fork and create a pull-request if you wish to contribute a fix or improvement to devlog. A contributor will build your PR in circle-ci and review the change. Please include tests and information in your pull requests description about the change.