[Install Go](https://golang.org/dl/)

```bash
$ wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf ~/go1.8.linux-amd64.tar.gz
$ echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc && source ~/.zshrc
$ go version # go1.8 linux/amd64
$ rm go1.8.linux-amd64.tar.gz
```

Go cmd
```
$ go get
$ go get -d -t -v # download dependencies (verbose, including test deps)
$ go test
$ go build
$ go install
$ go test ./... -cover # recursive coverage
```
