
# LVCS(Little Version Control System)
LVCS(Little Version Control System) is a versatile version control system designed to simplify code management and collaboration. With its intuitive command-line interface built in golang cobra, users can efficiently track, stage, and commit changes, navigate branches, and visualize commit history. Along with the version control functionalities, user will also receive a random cute ASCII art everytime the user enters commands. 

## Operations
can be found in:
* command list: https://github.com/Serenity0204/LVCS/blob/master/internal/ui/commands/list.txt
* command list with detail: https://github.com/Serenity0204/LVCS/blob/master/internal/ui/commands/detail.txt

## Install
make sure to have go install on your machine
* go install github.com/Serenity0204/LVCS@v1.0.0


## Uninstall
* windows: del %GOPATH%\bin\LVCS.exe
* macOS/Linux: rm $GOPATH/bin/LVCS


## Design
Can be found in https://github.com/Serenity0204/LVCS/blob/master/design.txt


## Features

- Allows users to init/dump the .lvcs directory
- Allows users to track/untrack files
- Allows users to view the file content by doing hashObject(file path to OID) and catFile(OID to file content)
- Allows users to do various operations of branching(CRUD all supported)
- Allowes users to do version commits in a tree fashioned(utlized nary tree to achieve this functionality)
- Allowes users to log the commit history either by version or log out all of them


