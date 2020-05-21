# Puma Cloud Native Buildpack

The Puma CNB sets the start command for a given ruby application that runs on a puma server.

## Integration

It provides `puma` as a dependency, but currently there's no scenario
we can imagine that you would use a downstream buildpack to require this
dependency. If a user likes to include some other functionality, it can be done
independent of the Puma CNB without requiring a dependency of it.

To package this buildpack for consumption:
```
$ ./scripts/package.sh
```
This builds the buildpack's source using GOOS=linux by default. You can supply another value as the first argument to package.sh.

## `buildpack.yml` Configurations

There are no extra configurations for this buildpack based on `buildpack.yml`.
