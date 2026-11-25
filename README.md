```docker run -it \
  --name container \
  --privileged \
  -v "$(pwd):/workspace" \
  -w /workspace \
  golang:1.25-bookworm \
  bash
```