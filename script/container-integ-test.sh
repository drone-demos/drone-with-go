# Check if it is alive
ping -c 5 target

EXIT_CODE=0
# RUN THE TEST
if curl -X POST target:8080/api/id/1 | grep -q '{"op":"POST","id":1}'; then
  echo "Tests passed!"
else
  echo "Tests failed!"
  let EXIT_CODE=1
fi
# GOOD API DESIGN OF A SERVICE INCLUDES A LIFECYCLE MGMT INTERFACE
# WHICH WE CAN LEVERAGE FOR TESTING... INVOKE IT TO STOP THE SERVICE
curl -X POST target:8080/.service/lifecycle/stop
sleep 1
exit $EXIT_CODE
