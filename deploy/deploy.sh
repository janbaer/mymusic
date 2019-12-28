#!/bin/sh

SERVER_NAME=JABASOFT-DS
USER=jan
TARGET_DIR=/volume1/docker/mymusic

MP3DB_SRC=./../mp3db/
WEB_DIST=./../web/dist/

cd ../web

yarn build

cd -

rsync -avuz --progress --delete \
  -e "ssh -i ~/.ssh/rsync-key"  \
  --exclude="log/"              \
  --exclude="data/"             \
  --exclude="certificates/"     \
  ./ $USER@$SERVER_NAME:${TARGET_DIR}/

rsync -avuz --progress          \
  -e "ssh -i ~/.ssh/rsync-key"  \
  ${MP3DB_SRC} $USER@$SERVER_NAME:${TARGET_DIR}/mp3db

rsync -avuz --progress          \
  -e "ssh -i ~/.ssh/rsync-key"  \
  ${WEB_DIST} $USER@$SERVER_NAME:${TARGET_DIR}/public

ssh -i ~/.ssh/rsync-key ${USER}@${SERVER_NAME} "cd ${TARGET_DIR}; ./docker-build.sh"

ssh -i ~/.ssh/rsync-key ${USER}@${SERVER_NAME} "cd ${TARGET_DIR}; ./docker-run.sh"

