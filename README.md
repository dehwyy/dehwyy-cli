### dehwyy CLI

## Prerequests:
- Go with configured GOPATH and GOROOT
____
## To get started:
Type in terminal:
```zsh
  go install github.com/dehwyy/dehwyy-cli@latest
  dehwyy-cli --help
```

## Commands
### Print
Funny command that change the input string to something different:
```zsh
  dehwyy-cli print -r hello dehwyy
  # would output hELLO DEHWYY
  # in this case r-flag means reverseCapitalize

  # for more information
  dehwyy-cli print --help
```
### System
Command that make an easy access to information about current PC:
```zsh
  dehwyy-cli system homedir
  # in my case output is "/home/dehwyy/"

  # for more information
  dehwyy-cli system
  # either
  dehwyy-cli system --help 
```
### Runner
Runner is a terminalScriptRunner, it's easy to manage usage of several commands by using just one:
```zsh
  dehwyy-cli runner -add onStartOs "code ." "easyeffects" "gnome-tweaks"
  # it will create command with key "onStartOs" which will execute 3 command continuously

  # to run command
  dehwyy-cli runner --exec onStartOs

  # to delete command
  dehwyy-cli runner --del onStartOs

  # for more information
  dehwyy-cli runner --help
```


_____
### Optionally, you can create alias like 'dh' or 'dehwyy'
In my case, I'm using zsh terminal so I have to ```sudo nano ~/.zshrc``` and write at the end of the file ```alias dehwyy="dehwyy-cli"```
