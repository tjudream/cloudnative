from ubuntu:20.04
workdir /usr/local/bin/
copy main httpserver
entrypoint ./httpserver
