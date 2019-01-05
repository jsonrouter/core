#!/bin/bash
echo "BUILDING 'http'"
(cd http && go build) || exit 10

echo "BUILDING 'validation'"
(cd ../validation && go build) || exit 10

echo "BUILDING 'config'"
(cd config && go build) || exit 10

echo "BUILDING 'tree'"
(cd tree && go build) || exit 10

echo "BUILDING 'openapi'"
(cd openapi && go build) || exit 10

echo "BUILDING 'core'"
go build || exit 10
