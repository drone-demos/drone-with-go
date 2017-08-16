sleep 5

ping -c 5 target

curl -X POST target:8080/api/id/1

if curl -X POST target:8080/api/id/1 | grep -q '{"op":"POST","id":1}'; then
  echo "Tests passed!"
  exit 0
else
  echo "Tests failed!"
  exit 1
fi
