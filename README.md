### Dev Log
[![CircleCI](https://circleci.com/gh/automoto/devlog.svg?style=svg)](https://circleci.com/gh/automoto/devlog)

*The software development log to reflect and capture your thoughts, progress and TODOs.*

Devlog's goal is to help you create a "Development Log" to reflect and generate notes during and after a coding session.

<!-- MarkdownTOC -->

- What is a Development Log?
- What is Devlog?
- Install
	- Install Using Go
- Using Devlog
- Configure
	- Configuration using command line options:
- Customizing Questions and Content of Devlog
- Building Locally

<!-- /MarkdownTOC -->


#### What is a Development Log?


#### What is Devlog?

Devlog generates a time stamped, simple, "development log" file to a directory you specify. This gives us alot of flexibility, you can save your devlog files into a git repository or to a cloud file service directory like dropbox, google drive, or one drive for automated syncing and backup.

Devlog prioritizes:
- Open standards over closed. Keep your notes in markdown files that can be queried for easily in a directory, not locked into some vendors service or custom formatting standards.
- Ease of use and simplicity. This is not meant to be a complex static content generator. It strives to be an easy to configure and create log files to fill out in a text editor of your choice.

#### Install
The easiest way to install devlog is to use the provided installation script
``` sh
wget https://raw.githubusercontent.com/automoto/devlog/master/scripts/get-devlog.sh
# you're welcome to examine the script, it just grabs the latest release from github for your OS and installs it
sh get-devlog.sh
```

If you prefer to install it yourself, you can get the latest release binary directly from github https://github.com/automoto/devlog/releases and extract it to your `/usr/local/bin`.

##### Install Using Go

If you already have an updated verison of go lang, installing via go is easy:
`go get -u github.com/automoto/devlog/`


#### Using Devlog
Ideally you will run Devlog during or towards the end of your development session and it will be filled out with your notes and answers to the questions when your finished. The "Notes" seciton at the bottom is great for TODOs or other information you want to keep. [Here is an example of a filled out devlog](https://gist.github.com/automoto/15e037d40258df1b8c2394ba1bae2c07). 

Devlog is designed to automate generating a development log markdown file for you to fill out in your favorite text editor. Once you have installed it, just type `devlog` and it will generate a time stamped markdown document. 

```shell
# Generate a markdown document log in the current directory
devlog

# Output the log document to your terminal instead of a document
devlog -p 'stdout'

```

By default, it returns the path of the document created so you can input this into a text editor easily. Here is a way you can generate a dev log document and open it in some popular text editors: 

```shell

# create a new Devlog file and open it in SublimeText
subl `./devlog | tail -n 1`

# create a new Devlog file and open it in vim
vim `./devlog | tail -n 1`

# create a new Devlog file and open it in nano
nano `./devlog | tail -n 1`
```


#### Configure
By default `devlog` will generate a file in the current directory unless specify the directory via setting the environment variable `DEVLOG_DIR`.

*Configuration using environment variables:*

Set the directory to save devlog files to:
```
export DEVLOG_DIR="/home/your_username/your_directory"
```

You can override this by temporarily setting the value when calling devlog

```
DEVLOG_DIR="/home/your_username/other_directory" devlog
```

##### Configuration using command line options:

You can also pass in configurations via command line options. Command line options take precedence over configurations set via environment variables.
```
devlog -p "/home/your_username/your_directory" -c "/home/your_username/your_directory/custom-config.yaml" 
```
To view all the possible command line options, just pass in the `-h` command line option for help e.g. `devlog -h` 

#### Customizing Questions and Content of Devlog
You can customize the sections of your Devlog by creating a `.yaml` and specifying your custom questions and sections there.
``` yaml
questions:
  - "My Custom Question One?"
  - "My Custom Question Two?"
other_section:
  - "TODO"
```
Then pass in your configuration file to devlog:
```
devlog -p your_custom_questions_file.yaml
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
