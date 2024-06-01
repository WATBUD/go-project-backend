# Deployment

# Running with Docker
docker pull openpolicyagent/opa
docker run -p 8181:8181 openpolicyagent/opa run --server --log-level debug
docker ps