# ActivityTrackerAPI
CSUSM Activity Tracker API

Install Go
```
$ sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
$ sudo apt-get update
$ sudo apt-get install golang
```

Set $GOPATH in your $HOME/.bashrc
```
export PATH=${PATH}:/usr/local/go/bin
export GOPATH=$HOME/path/to/go
export GOBIN=$GOPATH/bin
export PATH=$GOPATH:$GOBIN:$PATH
```

Install dependencies
`$ godep restore`

documentation: https://docs.google.com/document/d/1XWebh2Yba7p5x-PzjQ9Q718E6VY--y_VTITsJubt79Q/edit?usp=sharing
