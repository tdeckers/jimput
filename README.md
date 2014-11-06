# Overview

At times, there's a need to provide a docker container with input that can't be shared through the environment variables.  This is the case when the input is large or binary data. _jimput_ was trying to solve this in a complex way, leveraging code from [confd](https://github.com/kelseyhightower/confd).

As it turns out, this can be solved in a much simpler fashion by just piping tarred data into the containers stdin.

* Create input

	mkdir tmp
	echo "Some text" > tmp/datafile
	tar -cf input.tar tmp/datafile

* Pipe the input into the container

	ID=$(cat input.tar | docker run -i -a stdin ubuntu /bin/bash -c "tar -x -C / && cat /tmp/datafile")
	docker logs $ID

This mechanism allows you to inject multiple data files, config files, etc.. into a container.  In a practical example, you'll want to replace the `cat /tmp/datafile` with the actual command you want your container to run.  Very likely you'll want to use `exec` to ensure the container properly responds to signals (e.g. to stop the container process)
