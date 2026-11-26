```
docker run -it \
  --name container \
  --privileged \
  -v "$(pwd):/workspace" \
  -w /workspace \
  golang:1.25-bookworm \
  bash
```

inside the docker container :
```
apt-get update
apt-get install -y debootstrap

rm -rf /container-fs
mkdir -p /container-fs
debootstrap --arch=arm64 jammy /container-fs http://ports.ubuntu.com/ubuntu-ports

cd /workspace
go run main.go run /bin/bash

ls -la
hostname
ps
cat /etc/os-release
```