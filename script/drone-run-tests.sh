#!/bin/sh
# remove old containers
docker-compose -p ci --file docker-compose.test.yml rm -f
# run tests
docker-compose -p ci --file docker-compose.test.yml up --build

echo "Inspecting exited containers:"
docker-compose -p ci --file docker-compose.test.yml ps
docker-compose -p ci --file docker-compose.test.yml ps -q | xargs docker inspect -f '{{ .State.ExitCode }}' | while read code; do
    if [ "$code" != "0" ]; then
       exit $code
    fi
done
