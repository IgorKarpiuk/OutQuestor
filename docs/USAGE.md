# Usage

`OutQuestor` listens for outgoing network connections and prints them in a table.

```bash
outquestor listen [flags]
```

Available flags:

* `--protocol`, `-p` – `tcp`, `upd` or empty for both.
* `--ipLayer`, `-i` – `v4`, `v6` or empty for both.
* `--httpOnly` – show only HTTP and HTTPS requests.

## Output example

```
TIMESTAMP            DESTINATION                   PORT       PROTOCOL
2025-01-01 12:00:00  example.com                   443        HTTPS
```

The network interface used in `core` is currently hard coded to `en0`. Change it in `core/main.go` if your system uses a different device name.
