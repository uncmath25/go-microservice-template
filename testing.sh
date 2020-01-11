echo "Running basic tests against local golang http server..."

echo "Expecting negative feedback..."
curl localhost:8080/process_name/Joe

echo "Expecting positive feedback..."
curl localhost:8080/process_name/Colton
