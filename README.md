# needed env variables (add to a .env):

JELLYFIN_USER_ID=[user you want to set it to played for]
JELLYFIN_API_KEY=[a jellyfin api key}
JELLYFIN_BASE_URL=https://jellyfin.domain.tld

# run options:

## docker

simply run the docker image:
```bash
docker run -d simoneklundh/jellyfin-watch-updater:1.0
```
in the same path as your .env file

## using go

clone the project
install go
cd into the source directory 
go run .