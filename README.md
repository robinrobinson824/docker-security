# docker-security
Securing Docker Containers

Simple application displaying "Hello, World!" out on the screen written in GO.
Dockerfile using multi-stage build for better efficiency and security.

Builder stage: 
Uses golang:1.22-alpine image pulled from Docker Hub.
Copies application to the container, runs updates, initializes new module and downloads the dependencies.

Final image stage:
Copies artifact from builder stage to the final stage.
Creates user myuser and group mygroup, assigns myuser to mygroup. Assigns 755 permissions to app folder.
Switches to myuser for security purposes - don't want to leave root in use.
Runs ./app when container starts.