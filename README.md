NodeFinderGo
============

Go implementation of NodeFinder program.

Please visit [NodeFinderGUI](https://github.com/zxjsdp/NodeFinderGUI) and [NodeFinder](https://github.com/zxjsdp/NodeFinder) for Python implementation of NodeFinder program.

Please visit [NodeFinderC](https://github.com/zxjsdp/NodeFinderC) for C implementation of NodeFinder program.

Install
-------

    $ go install github.com/zxjsdp/nodefinder-go

Usage
-----

    $ nodefinder-go -input input.nwk -config config.txt -output output.nwk
    
If you just want to try it:

    $ go run github.com/zxjsdp/nodefinder-go/cmd/nodefinder-go/main.go -input input.nwk -config config.txt -output output.nwk

Run Tests
---------

    $ cd nodefindergo_test && go test -v -tags integration
