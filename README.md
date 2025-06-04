# ZObserver

`zobserver` is a Go-based application designed to monitor EVE Online killmail feeds, filter them based on specified criteria (characters, corporations, alliances), and relay relevant killmails to Discord webhooks. It leverages ZKillboard's RedisQ message queue for incoming killmails and the EVE Online ESI API for additional data.

## Features

* **EVE Online Killmail Monitoring**: Listens to a killmail feed (via RedisQ).
* **Configurable Filtering**: Filters killmails based on:
  * Character IDs
  * Corporation IDs
  * Alliance IDs
  * Option to capture all killmails.
* **Discord Notifications**: Sends formatted killmail notifications to configured Discord webhooks.
* **Containerized & Kubernetes Ready**: Designed to run as a Docker container and deployable via a Helm chart.

## Configuration

`zobserver` is configured via environment variables.

### Core Configuration

* `LOGGER_FORMAT`: Log output format (e.g., "prod", "dev").
* `LOGGER_LEVEL`: Logging level (e.g., "info", "debug").
* `QUEUE_NAME`: The name of the RedisQ queue to listen to.
* `TTW`: Time To Wait, in seconds. Default: `"10"`.
* `ESI_USER_AGENT`: User agent string for requests to the EVE Online ESI. Default: `"zobserver"`.
* `DESTINATIONS_FILE`: Path to a YAML file containing destination configurations (see format below).
* `DESTINATIONS`: Alternatively, a string containing the YAML configuration for destinations directly.

If neither `DESTINATIONS_FILE` nor `DESTINATIONS` is provided, the application will error as it requires at least one destination.

### Destinations Configuration (YAML Format)

This configuration can be provided via the `DESTINATIONS` environment variable directly or in a file specified by `DESTINATIONS_FILE`.

```yaml
- name: "Example Destination 1 (All Killmails)"
  all: true
  discord_webhooks:
    - id: "YOUR_DISCORD_WEBHOOK_ID_1"
      token: "YOUR_DISCORD_WEBHOOK_TOKEN_1"
- name: "Example Destination 2 (Specific Entities)"
  character_ids:
    - 90000001
    - 90000002
  corporation_ids:
    - 1000001
  alliance_ids:
    - 300001
  discord_webhooks:
    - id: "YOUR_DISCORD_WEBHOOK_ID_2"
      token: "YOUR_DISCORD_WEBHOOK_TOKEN_2"
    - id: "YOUR_DISCORD_WEBHOOK_ID_3"
      token: "YOUR_DISCORD_WEBHOOK_TOKEN_3"
```

*Refer to `internal/observer/config.go` for the structure definitions.*

## Getting Started

### Prerequisites

* Go (for building from source)
* Docker
* Kubernetes and Helm (for Kubernetes deployment)

### Run with Docker

A `Dockerfile` is present in the root of the project. The application image is also available on `ghcr.io/fnt-eve/zobserver`.

1. **Build the Docker image (if building locally):**
   ```bash
   docker build -t zobserver:latest .
   ```
2. **Run the Docker container:**
   You'll need to pass the necessary environment variables for configuration.
   ```bash
   docker run -d \
     -e QUEUE_NAME="your-redis-queue" \
     -e ESI_USER_AGENT="your-esi-agent/1.0 your@email.com" \
     -e DESTINATIONS_FILE="/app/destinations.yaml" \
     -v /path/to/your/destinations.yaml:/app/destinations.yaml \
     --name zobserver \
     ghcr.io/fnt-eve/zobserver:latest # Or your locally built tag
   ```


### Deploy to Kubernetes

The project includes a Helm chart located in the `k8s/zobserver/` directory.

1. Navigate to the chart directory:
   ```bash
   cd k8s/zobserver
   ```
2. Customize `values.yaml` as needed, especially the `env` section for configuration. The existing `k8s/zobserver/README.md` provides details on configurable values.
3. Deploy using Helm:
   ```bash
   helm install zobserver . -n <your-namespace>
   # Or
   helm upgrade --install zobserver . -n <your-namespace> -f my-values.yaml
   ```

## Project Structure

* `cmd/observer/`: Main application entry point.
* `internal/`: Contains the core logic of the observer.
  * `internal/logger/`: Logging setup.
  * `internal/observer/`: Core observer, Redis queue, routing, and ESI/Discord sending logic.
* `k8s/zobserver/`: Helm chart for Kubernetes deployment.
* `Dockerfile`: For building the Docker image.
* `go.mod`, `go.sum`: Go module files.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.
For major changes, please open an issue first to discuss what you would like to change.

## License

MIT