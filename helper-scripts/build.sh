#!/bin/sh

echo "Stopping go_linkr service..."
service go_linkr stop

go mod tidy
go build -o go_linkr ../main.go

mv go_linkr /usr/local/bin

echo "Executable built and moved to /usr/local/bin/go_linkr..."
echo "Starting service go_linkr..."

service go_linkr start

echo "Done!"
