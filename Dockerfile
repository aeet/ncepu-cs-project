# Use the official Golang image as the base image
FROM golang:1.16.5-alpine as build-stage

# Set environment variables
ENV APP_NAME ncepu-cs-project
ENV APP_HOME /app
ENV STATIC_DIR static

# Set the working directory for the application
WORKDIR ${APP_HOME}

# Copy the necessary files into the container
COPY . ${APP_HOME}

# Build the application
RUN go build -o ${APP_NAME}

# Start a new stage to create a lightweight image
FROM alpine:3.14 as production-stage

# Set environment variables
ENV APP_NAME ncepu-cs-project
ENV APP_HOME /app
ENV STATIC_DIR static

# Set the working directory for the application
WORKDIR ${APP_HOME}

# Copy the necessary files into the container
COPY --from=build-stage ${APP_HOME}/${APP_NAME} ${APP_HOME}/
COPY ${STATIC_DIR} ${APP_HOME}/${STATIC_DIR}/
COPY config.yaml ${APP_HOME}/

# Expose the port that the application runs on
EXPOSE 8080

# Start the application
CMD ["./ncepu-cs-project"]