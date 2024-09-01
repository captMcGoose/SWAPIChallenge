#!/bin/bash

# URL of the service endpoint we need to test
# (This assumes the service is running on Minikube and accessible via NodePort)
URL="http://$(minikube ip):30001/people"

echo "Starting load test..."

# Loop to perform 1000 parallel HTTP GET requests to the service endpoint to simulate a concurrent load
for i in {1..1000}
do
    # Perform the HTTP GET request and discard the output to /dev/null
    curl $URL > /dev/null &
done

echo "Load test completed."
