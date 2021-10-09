#!/bin/bash

# This script installs Robocup Soccer SImulator dependencies to run tests
# RCSSMonitor is not installed as this is a headless installation for automated testing
# Tested on Ubuntu 20.04

apt-get update
apt-get -y install tzdata
apt-get install -y wget \
    git \
    g++ \
    build-essential \
    automake \
    flex \
    bison \
    libboost-all-dev \
    libfontconfig1-dev \
    libaudio-dev \
    libxt-dev \
    libsm-dev \
    libice-dev \
    libxi-dev \
    libxrender-dev \
    libcppunit-dev \
    qt5-default \
    pyqt5-dev-tools

curl -L https://github.com/rcsoccersim/rcssserver/releases/download/rcssserver-16.0.0/rcssserver-16.0.0.tar.gz | tar -zxp
cd rcssserver-16.0.0 && ./configure && make && make install && cd ..

curl https://mirrors.xtom.com/osdn//rctools/51941/librcsc-4.1.0.tar.gz | tar -zxp
cd librcsc-4.1.0 && ./configure && make -j8 && make install && cd ..

# Need to add /usr/local/lib to ld config path or the server won't run
echo "/home/local/lib" >/etc/ld.so.conf.d/localLib.conf
ldconfig
