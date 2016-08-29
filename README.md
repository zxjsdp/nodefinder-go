NodeFinderGo
============

[![Build Status](https://travis-ci.org/zxjsdp/nodefinder-go.svg?branch=master)](https://travis-ci.org/zxjsdp/NodeFinderGo)

Go implementation of NodeFinder program.

- Please visit [NodeFinderGUI](https://github.com/zxjsdp/NodeFinderGUI) and [NodeFinder](https://github.com/zxjsdp/NodeFinder) for Python implementation of NodeFinder program.
- Please visit [NodeFinderC](https://github.com/zxjsdp/NodeFinderC) for C implementation of NodeFinder program.

Install
-------

    $ go install github.com/zxjsdp/nodefinder-go/cmd/nodefinder-go

Usage
-----

    $ nodefinder-go -input input.nwk -config config.txt -output output.nwk

Run Tests
---------

    $ cd $GOPATH/src/github.com/zxjsdp/nodefinder-go/ && go test -v ./...

Config Syntax
-------------

    # lines startswith # with be ignored
    # you can add one or multiple calibrations in single config file

    name_a, name_b, calibration_infomation_1
    name_c, name_d, calibration_infomation_2
    name_a, name_b, clade_label_information
    name, branch_label_information
    ..., ..., ...
