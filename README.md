# goxe

reduce large volumes of repetitive logs into compact, readable aggregates.

goxe is a high-performance log reduction tool written in go. it ingests logs (currently via syslog/udp),
normalizes and filters them, and aggregates repeated messages into a single-line format with occurrence counts.
the result is less noise, lower bandwidth usage, and cheaper storage without losing visibility into recurring issues.

goxe is designed to run continuously in the background as part of a logging pipeline or sidecar.

## requirements

* go 1.25.5 or higher (to build from source)

### aggregation behavior

goxe performs several transformations before aggregation:

* strips timestamps and date prefixes
* converts text to lowercase
* removes extra whitespace
* filters out configurable excluded words
* applies basic ascii beautification

after normalization, identical messages are grouped together and reported with repetition counts.

example input:

```
2025-01-01 12:00:01 error: connection failed
2025-01-01 12:00:02 error: connection failed
2025-01-01 12:00:03 error: connection failed
```

aggregated output:

```
error: connection failed (x3)
```

## architecture

goxe is built for concurrency and throughput:

* worker pool architecture for parallel log processing
* centralized, thread-safe aggregation state using `sync.mutex`
* periodic partial reporting using `time.ticker`
* streaming design with low memory overhead

the system is optimized to handle high log volumes with minimal latency.

## roadmap

### completed

* worker pool for parallel processing
* thread-safe state management
* automated partial reporting
* log normalization and filtering
* ascii beautification
* timestamp and date parsing
* syslog/udp network ingestion

### planned

* similarity clustering (group near-identical messages)
* graceful shutdown and signal handling
* configuration file support
* additional ingestion backends

## maintainers

* @dumbnoxx

## license

licensed under the apache license, version 2.0. see the [license file](license) for details.
