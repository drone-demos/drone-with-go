#!/bin/sh
# remove old containers
docker-compose -p ci --file docker-compose.test.yml rm -f
# run tests (and detatch)
docker-compose -p ci --file docker-compose.test.yml up --build -d


echo "Inspecting exited containers:"
docker-compose -p ci --file docker-compose.test.yml ps
COUNTER=0
docker-compose -p ci --file docker-compose.test.yml ps -q | xargs docker inspect -f '{{ .State.ExitCode }}' | while read code; do
    echo "Exit code is currently $code"
    if [ "$code" != "0" ]; then
       exit $code
    fi
    echo The counter is $COUNTER
    let COUNTER=COUNTER+1
    if [  $COUNTER -gt 50 ]; then
       exit 1
    fi
done
