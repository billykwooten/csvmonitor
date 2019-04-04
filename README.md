# CSV Monitor
Monitors CSV for stuff

## Requirements

* Linux / MacOS / Windows
* [docker](https://www.docker.com)

## Development

* [go](https://golang.org/dl)

## Configuration

Dark sky exporter can be controlled by both ENV or CLI flags as described below.

| Environment        	     | CLI (`--flag`)              | Default                 	 | Description                                                                                                      |
|----------------------------|-----------------------------|---------------------------- |------------------------------------------------------------------------------------------------------------------|
| `CSVLOC`                   | `csv`                       | `<REQUIRED>`                | File path of the CSV file |
| `SLACKHOOK`                | `slackhook`                 | `<REQUIRED>`                | Webhook URL for Slacks    |

## Usage

```
# Export weather metrics from Seattle using binary
export LOG=*; ./csvmonitor --csv test.csv --slackhook https://hooks.slack.com/services/JIGJ42V99/BHDTFBH9P/K489FDQyPjlHVjU5492AjTWfQ

# Export weather metrics from Seattle using docker
docker run -d --restart on-failure --env LOG=* --name=csvmonitor billykwooten/csvmonitor --csv test.csv --slackhook https://hooks.slack.com/services/JIGJ42V99/BHDTFBH9P/K489FDQyPjlHVjU5492AjTWfQ
```