#!/bin/bash

(cd appengine && go test -v)
(cd fasthttp && go test -v)
