#!/bin/sh

SERVER_NAME=JABASOFT-DS
USER=jan
TARGET_DIR=/volume1/docker/mymusic

MP3DB_SRC=./../mp3db/
WEB_DIST=./../web/dist/
PASSWORD=$(pass /home/JABASOFT-DS/jan)

cd ../web

yarn build

cd -

sshpass -p $PASSWORD rsync -avuz --progress --delete \
  -e "ssh -i ~/.ssh/rsync-key"  \
  --exclude="log/"              \
  --exclude="data/"             \
  --exclude="certificates/"     \
  ./ $USER@$SERVER_NAME:${TARGET_DIR}/

sshpass -p $PASSWORD rsync -avuz --progress          \
  -e "ssh -i ~/.ssh/rsync-key"  \
  ${MP3DB_SRC} $USER@$SERVER_NAME:${TARGET_DIR}/mp3db

sshpass -p $PASSWORD rsync -avuz --progress          \
  -e "ssh -i ~/.ssh/rsync-key"  \
  ${WEB_DIST} $USER@$SERVER_NAME:${TARGET_DIR}/public

sshpass -p $PASSWORD ssh -i ~/.ssh/rsync-key ${USER}@${SERVER_NAME} "cd ${TARGET_DIR}; ./docker-build.sh"

sshpass -p $PASSWORD ssh -i ~/.ssh/rsync-key ${USER}@${SERVER_NAME} "cd ${TARGET_DIR}; ./docker-run.sh"

