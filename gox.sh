#!/usr/bin/env bash

#创建 $GOPATH/src/golang.org/x 目录
mkdir -p $GOPATH/src/github.com/golang
mkdir -p $GOPATH/src/golang.org

if [ ! -d $GOPATH/src/golang.org/x ]
then
    ln -s $GOPATH/src/github.com/golang $GOPATH/src/golang.org/x
fi

cd $GOPATH/src/golang.org/x
echo "Installed package"
ls
echo "------------------------------------------------------------------------"

packages=(
    "glog" "image" "perf" "snappy" "term" "sync" 
    "winstrap" "cwg" "leveldb" "text" "net" 
    "build" "protobuf" "dep" "sys" "crypto" 
    "gddo" "tools" "scratch" "proposal" "mock" 
    "oauth2" "freetype" "debug" "mobile" "gofrontend" 
    "lint" "appengine" "geo" "review" "arch" 
    "vgo" "exp" "time"
)
for name in ${packages[*]}
do
    echo "${name} installing..."
    url=github.com/golang/${name}
    go get -v ${url}
done
