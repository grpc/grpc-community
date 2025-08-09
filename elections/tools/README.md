# gRPC Steering Committee Elections Tooling

This directory contains tooling used to conduct the annual gRPC Steering
Committee elections.

First, write a config file using [the example config](examples/config.yaml) as a
guide. If you have not changed the ballot format since the last election, the
only change you will need to make is the list of candidates.

To generate a results markdown file given a CSV of ballots, use the following command:

```
go run ./cmd/tally.go [VOTES_CSV] -c [CONFIG_FILE]
```

This will generate a markdown file at `results.md` with the results of the
election.
