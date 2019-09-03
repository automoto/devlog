### Dev Log
*The software development log to reflect and capture your thoughts, progress and TODOs.*

Devlog's goal is to help you create a "Development Log" to reflect and generate notes after a coding session.

Devlog generates a simple markdown file to a directory you specify which gives you alot of flexibility. You can save your log files into a git repository or to a cloud file service directory like
dropbox, google drive, or one drive for automated syncing and backup.

Devlog prioritizes:
- Open standards over closed. Keep your notes in markdown files that can be queried for easily in a directory, not locked into some vendors service or custom formatting standards.
- Ease of use and simplicity. This is not meant to be a complex static content generator. It strives to be an easy to configure and create log files to fill out in a text editor of your choice.

#### Configure
By default `devlog` will generate a file in the current directory unless specify the directory via setting the environment variable `DEVLOG_DIR`.

*examples:*

Set the directory to save devlog files to:
```
export DEVLOG_DIR="/home/your_username/your_directory"
```

You can override this by temporarily setting the value when calling devlog

```
DEVLOG_DIR="/home/your_username/other_directory" devlog
```

#### Build and Install
simply clone this repository and run the following command to build the binary:
```shell
go build -o devlog main.go output.go util.go util_interactive.go
```

This will create a binary locally you can run commands against already, like so:

`./devlog`

Copy the binary to your local bin to access the CLI anywhere.

`sudo cp devlog /usr/local/bin`

Now you can run the command `devlog` from anywhere to generate a new devlog file:

```
devlog
2019/09/02 22:00:32 Successfully saved dev log to directory: /home/dev/null
```