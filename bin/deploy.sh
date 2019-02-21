#!/usr/bin/env bash

ROOM_DIR=/home/deliang/www

cd ${ROOM_DIR}/hi_come_on/store-fe
git pull
git checkout dev-store

npm run prod
git add *
git commit -m "chore: build dev prod"
git push

cp dist/index.html ${ROOM_DIR}/xcx_wx/resources/views/web/home.blade.php
echo success!!!
