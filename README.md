
# LVCS(Little Version Control System)
LVCS(Little Version Control System) is a versatile version control system designed to simplify code managements. With its intuitive command-line interface built in golang cobra, users can efficiently track, stage, and commit changes, navigate branches, and visualize commit history. Along with the version control functionalities, user will also receive a random cute ASCII art everytime the user enters commands. 

## Operations
can be found in:
* command list: https://github.com/Serenity0204/LVCS/blob/master/internal/ui/commands/list.txt
* command list with detail: https://github.com/Serenity0204/LVCS/blob/master/internal/ui/commands/detail.txt

## Install
make sure to have go install on your machine
```
go install github.com/Serenity0204/LVCS@v1.0.5
```

## Example Usage
Go to the directory you want to use LVCS
```
LVCS init
LVCS stage add "file 0"
LVCS commit
LVCS commit tree
LVCS stage add "file 1" "file 2"
LVCS commit fresh
LVCS stage
LVCS stage add "file 1" "file 2"
LVCS stage untrack "file 2"
LVCS commit current
LVCS commit switch v0
LVCS stage
LVCS commit
LVCS commit tree
LVCS branch
LVCS log
LVCS commit lca v1 v2
LVCS restore v0
LVCS dump
```

## Uninstall
* Windows:
```
del %GOPATH%\bin\LVCS.exe
```
* MacOS/Linux: 
```
rm $GOPATH/bin/LVCS
```
deleting source files at
```
C:\Users\your username\go\pkg\mod\github.com\!serenity0204
```
for windows

## Features

- Allows users to init/dump the .lvcs directory
- Allows users to track/untrack files
- Allows users to view the file content by doing hashObject(file path to OID) and catFile(OID to file content)
- Allows users to do various operations of branching(CRUD all supported)
- Allows users to do version commits in a tree fashioned(utilized nary tree to achieve this functionality)
- Allows users to log the commit history either by versions or logging out all of them
- Allows users to restore the snapshots of the recorded commit

## Design
Can be found in https://github.com/Serenity0204/LVCS/blob/master/design.txt
