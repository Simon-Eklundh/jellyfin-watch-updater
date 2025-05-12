# needed env variables (add to a .env):

JELLYFIN_USER_ID=[user you want to set it to played for]
JELLYFIN_API_KEY=[a jellyfin api key}
JELLYFIN_BASE_URL=https://jellyfin.domain.tld

# run options:

## docker
for latest image release, check the tags or build yourself from source
https://hub.docker.com/r/simoneklundh/jellyfin-watch-updater/tags

The docker image runs the go script once an hour on the hour.

simply run the docker image:
```bash
docker run -d simoneklundh/jellyfin-watch-updater:2.0
```
in the same path as your .env file

## using go

This simply runs the script once 

clone the project
install go
cd into the source directory 
go run .

