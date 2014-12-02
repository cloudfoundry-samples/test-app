# Lattice-App - a simple Go webapp for demonig Lattice

Lattice-App is packaged as a docker image at cloudfoundry/lattice-app

To push to Lattice:

```bash
ltc start -i docker:///cloudfoundry/lattice-app -c /lattice-app
```

### Endpoints

`/`: a simple landing page
`/env`: displays environment variables
`/started-at`: the unix timestamp when Lattice-App was started
`/index`: the Instance Index of the Lattice-App
`/exit`: instructs Lattice to exit with status code 1

### To rebuild the dockerimage:

```bash
./build.sh
```

Assumes you have the go toolchain (with the ability to cross-compile to different platforms) and docker installed and pointing at your docker daemon.
