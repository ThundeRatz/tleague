#!/bin/bash

if [ $# -eq 0 ]; then
    BRANCH="master"
else
    BRANCH=$1
fi

if [ $# -eq 2 ]; then
    FOLDER=$2
else
    FOLDER="current"
fi

cd ThunderLeague

git checkout ${BRANCH}
git pull

./bootstrap
./configure
make -j8

rm -rf ~/teams/${FOLDER}
mkdir -p ~/teams/${FOLDER}

cp -R src/start.sh src/sample_player src/sample_coach configurations/*.conf src/formations-* ~/teams/${FOLDER}/
