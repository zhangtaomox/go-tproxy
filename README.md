# go-tproxy
A simple and transparent proxy that supports any protocol.

Just for restrict network.

## Usage

``` bash
# redis proxy
./tproxy -l :8080 -r remoteip:6379
```

``` bash
# mongodb proxy
./tproxy -l :8080 -r remoteip:27017
```