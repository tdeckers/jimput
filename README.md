# Overview

_jimput_ takes json as stdin and applies transformations to create (configuration) files.  It's created as a way to provide complex data input structures into a docker container, but other uses might exist.
The concept is very similar to [confd](https://github.com/kelseyhightower/confd) but with json at the stdin instead of pulling data from a remote service.

# Use

_work in progress:_
1. json in
2. use go templates to parse input and write to files
3. launch delegate application using exec

_jimput_ will return non-zero if it encountered an issue parsing the input json or writing the output files.
