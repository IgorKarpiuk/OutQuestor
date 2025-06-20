# OutQuestor

OutQuestor is a small Go utility that prints information about outgoing network connections. It consists of a reusable `core` library and a command line interface built with Cobra.

## Repository layout

```
core/  - packet capturing library
cli/   - command line interface
```

## Building

Go 1.23 or newer is required. To build the CLI binary run:

```bash
cd cli
go build
```

This will create an executable named `cli` in the `cli` directory.

## Usage

```
./cli listen [--protocol tcp|upd] [--ipLayer v4|v6] [--httpOnly]
```

Flags:

- `--protocol` – filter by TCP or UDP traffic.
- `--ipLayer` – listen on IPv4 or IPv6 only.
- `--httpOnly` – restrict output to HTTP/HTTPS requests.

Example:

```bash
./cli listen --protocol tcp --httpOnly
```

The program prints a table with timestamp, destination host, port and protocol.

## License

This project is released under the MIT License.
