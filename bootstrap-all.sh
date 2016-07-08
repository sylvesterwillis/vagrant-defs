#!/usr/bin/env bash
cd /home/vagrant
add-apt-repository ppa:ubuntu-lxc/lxd-stable
apt-get update

#install git and bison and gvm dependencies
apt-get install -y git-core
apt-get install -y bison
apt-get install -y curl git mercurial make binutils bison gcc build-essential

# install gvm and setup golang
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source /root/.gvm/scripts/gvm

# Have to install go1.4 because of switch from C to Go compiler in 1.5
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.6.2
gvm use go1.6.2

gopath=/home/vagrant/godev

# setup gopath
mkdir $gopath
export GOPATH=$gopath
echo "GOPATH=$gopath" >> /home/vagrant/.bashrc
source /home/vagrant/.bashrc

# install nsq
wget https://s3.amazonaws.com/bitly-downloads/nsq/nsq-0.3.8.linux-amd64.go1.6.2.tar.gz
tar xvfz nsq-0.3.8.linux-amd64.go1.6.2.tar.gz
sudo mv nsq-0.3.8.linux-amd64.go1.6.2/bin/* /usr/local/bin

# get go-nsq library
go get -u -v github.com/bitly/go-nsq