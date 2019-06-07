#!/bin/bash
brew install graphviz
cd vendor/github.com/TrueFurby/go-callvis
make
cd ../../../..
go-callvis  -focus "candidate_service/handlers" -group pkg -nointer -nostd ./main
