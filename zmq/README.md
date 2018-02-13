ZMQ install

[ZeroMQ](http://zeromq.org/)

```bash
$ sudo apt-get install libtool-bin
$ wget https://github.com/zeromq/zeromq3-x/releases/download/v3.2.5/zeromq-3.2.5.tar.gz
$ tar -xzf zeromq-3.2.5.tar.gz && rm zeromq-3.2.5.tar.gz && cd zeromq-3.2.5
$ ./autogen.sh && ./configure && make && sudo make install && sudo ldconfig
```
