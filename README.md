# Jellyfin Watch Updater

## What is it?

Jellyfin Watch Updater is a simple go script/docker container which adds a "lastPlayedDate" to media that has been watched but hasn't had the variable set.   
This is only necessary if 1. you're using a jellyfin client that doesn't properly set that variable (Blink, findroid, probably more) and 2. you're using something like the media cleaner plugin to delete videos that you have watched.

## How do I run it?

### needed env variables (add to a .env):

JELLYFIN_USER_ID=[user you want to set it to played for]  
JELLYFIN_API_KEY=[a jellyfin api key] 
JELLYFIN_BASE_URL=https://jellyfin.domain.tld  

### JELLYFIN_USER_ID

to get the JELLYFIN_USER_ID you can go to your jellyfin server on the web and go to the user page. the userId is in the url. You can also get this using the jellyfin api.

### run options:

#### docker (recommended for continuous use of the clients that don't set it properly)
for latest image release, check the tags or build yourself from source
https://hub.docker.com/r/simoneklundh/jellyfin-watch-updater/tags

The docker image runs the go script once an hour on the hour.
As well as once on boot 

simply run the docker image
with your own environment variables
if you're using an .env file you can do the following:
```bash
docker run -d --env-file .env simoneklundh/jellyfin-watch-updater:2.0
```
or if you want individual environment variables:

```bash
docker run -d -e JELLYFIN_USER_ID=[your id here] -e JELLYFIN_API_KEY=[your api key here] -e JELLYFIN_BASE_URL=[your jellyfin url here] simoneklundh/jellyfin-watch-updater:2.0
```
#### using go (not recommended, but still completely usable)

This simply runs the script once 

clone the project  
install go  
cd into the source directory   
go run .  

