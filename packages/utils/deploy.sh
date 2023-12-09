#!/bin/sh

# 1. Fetch the latest code from remote
git pull -f origin main

# 2. Install dependencies
pnpm install

# 3. Build step that compiles code, bundles assets, etc.
turbo build

# 4. Kill all screens
pkill screen

# 5. Make new screen
screen -S nuxtjs

# 6. Serve page
turbo serve