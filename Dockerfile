# APPNICENAME=Stash
# APPDESCRIPTION=An organizer for your porn, written in Go
version: '3.4'
services:
  stash:
    image: stashapp/stash:latest
    restart: unless-stopped
    ports:
      - "9999:9999"
    logging:
      driver: "json-file"
      options:
        max-file: "10"
        max-size: "2m"
    environment:
      - STASH_STASH=/data/
      - STASH_GENERATED=/generated/
      - STASH_METADATA=/metadata/
      - STASH_CACHE=/cache/
    volumes:
      - /etc/localtime:/etc/localtime:ro
      ## Adjust below paths (the left part) to your liking.
      ## E.g. you can change ./config:/root/.stash to ./stash:/root/.stash
      
      ## Keep configs here.
      - ./config:/root/.stash
      ## Point this at your collection.
      - ./data:/data
      ## This is where your stash's metadata lives
      - ./metadata:/metadata
      ## Any other cache content.
      - ./cache:/cache
      ## Where to store generated content (screenshots,previews,transcodes,sprites)
      - ./generated:/generated
