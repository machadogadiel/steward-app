#!/bin/sh

# 1. Fetch the latest code from remote
cd ~/Projects/steward-app && git pull -f origin master

# 2. Install dependencies
pnpm install

# 3. Build step that compiles code, bundles assets, etc.
turbo build

# 4. Kill all screens
screen -ls | grep 'nuxt' | awk '{print $1}' | xargs -I % -t screen -X -S % qui
pkill -9 -f 'node .output/server/index.mjs'
screen -ls | grep 'nuxt' | awk '{print $1}' | xargs -I % -t screen -X -S % quit
pkill -9 -f 'node .output/server/index.mjs'

# 5. Make new screen
screen -S nuxt

# 6. Serve page
turbo serve