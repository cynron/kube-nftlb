#!/bin/sh

# This is a script that helps you set up a fresh Debian Testing system
# running in a virtualized environment. It installs everything you need
# before testing kube-nftlb. Don't forget to run this script as root.

# Recommended Debian Testing ISO:
#    https://cdimage.debian.org/cdimage/daily-builds/daily/arch-latest/amd64/iso-cd/
#    debian-testing-amd64-netinst.iso


# 0. Change directory to /root/
cd


# 1. Set PATH env correctly
export PATH=$PATH:/usr/local/sbin:/usr/sbin:/sbin
echo 'export PATH=$PATH:/usr/local/sbin:/usr/sbin:/sbin' >> /root/.bashrc


# 2. Update packages and upgrade them
apt-get update
apt-get upgrade -y


# 3. Install Docker (v18.06.1-ce)
apt-get install -y apt-transport-https ca-certificates curl gnupg2 software-properties-common
curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
apt-get update
apt-get install -y docker-ce=18.06.1~ce~3-0~debian


# 4. Install Docker Machine (v0.16.0)
base=https://github.com/docker/machine/releases/download/v0.16.0 && curl -L $base/docker-machine-$(uname -s)-$(uname -m) >/tmp/docker-machine && install /tmp/docker-machine /usr/local/bin/docker-machine


# 5. Install kubectl (no hypervisor)
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list
apt-get update
apt-get install -y kubectl


# 6. Install Minikube (v0.30.0)
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.30.0/minikube-linux-amd64 && chmod +x minikube && cp minikube /usr/local/bin/ && rm minikube


# 7. Install Golang
apt install -y golang-go golang-golang-x-tools


# 8. Make directory where nftables/nftlb and dependencies will be installed, and install necessary tools to build them
mkdir .nft
cd .nft
apt-get install -y git bison flex binutils build-essential autoconf libtool pkg-config libgmp-dev libreadline-dev libjansson-dev libev-dev cmake dnsutils libxtables-dev


# 9. Download and install libmnl
git clone git://git.netfilter.org/libmnl/
cd libmnl
sh autogen.sh
./configure
make
make install
cd ..


# 10. Download and install libnftnl
git clone git://git.netfilter.org/libnftnl
cd libnftnl
sh autogen.sh
./configure
make
make install
cd ..
ldconfig


# 11. Download and install nftables
git clone git://git.netfilter.org/nftables
cd nftables
sh autogen.sh
./configure
make
make install
cd ..
ldconfig


# 12. Download and install nftlb
git clone https://github.com/zevenet/nftlb
cd nftlb
autoreconf -fi
./configure
make
make install
cd ..
ldconfig


# 13. Start Minikube
minikube start --vm-driver=none --kubernetes-version="v1.12.0"