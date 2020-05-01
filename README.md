### Dev Log
[![CircleCI](https://circleci.com/gh/automoto/devlog.svg?style=svg)](https://circleci.com/gh/automoto/devlog)

*The software development log to reflect and capture your thoughts, progress and TODOs.*

Devlog's goal is to help you create a "Development Log" to reflect and generate notes after a coding session.

Devlog generates a simple markdown file to a directory you specify which gives you alot of flexibility. You can save your log files into a git repository or to a cloud file service directory like
dropbox, google drive, or one drive for automated syncing and backup.

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

#### Install Using Go

If you already have an updated verison of go lang, installing via go is easy:
`go get -u github.com/automoto/devlog/`


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

```
devlog
2019/09/02 22:00:32 Successfully saved dev log to directory: /home/dev/null
```
