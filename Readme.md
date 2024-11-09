# GoWizard

This is my attempt to migrate my current bash file to create `devcontainer` enviroments to Golang

The original project [containerwizard](https://github.com/dabiddo/containerwizard) is still maintained because I have not finished migrating all the templates to this project, and any new templates will be added for the time being, but will eventually become deprecated in order to maintain this new GoWizard repo.

## Requisites
This new golang repo uses charm.sh libraries to make the TUI more friendly and colorful, so if you clone this repo, make sure to install the packages needed.

### Build
Once you have cloned the repo, installed the packages, you can build your own `gowizard` binary, and move the executable to `/usr/local/bin` , I recommend making it executable first `sudo chmod +x gowizard` and the move it to the `bin` folder

```bash
go mod tidy

go build -o gowizard
```