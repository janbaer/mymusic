mkdir -p ${GOPATH}/src/github.com/janbaer

MP3DB_GOPATH=${GOPATH}/src/github.com/janbaer/mp3db
CURRENTDIR=$(pwd)

if [ ! -L ${MP3DB_GOPATH} ]; then
  echo "Create symlink from ${CURRENTDIR} to ${MP3DB_GOPATH}"
  ln -s ${CURRENTDIR} ${MP3DB_GOPATH}
fi

mkdir -p data log

cd ${MP3DB_GOPATH}

dep ensure

GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o bin/mp3db main.go

cd ${CURRENTDIR}

