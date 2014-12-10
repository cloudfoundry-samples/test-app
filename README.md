# Lattice-App - a simple Go webapp for demoing Lattice

Lattice-App is packaged as a docker image at cloudfoundry/lattice-app

To push to Lattice:

```bash
ltc start lattice-app -i docker:///cloudfoundry/lattice-app -c /lattice-app
```

### Endpoints

`/`: a simple landing page displaying the index and uptime
`/env`: displays environment variables
`/exit`: instructs Lattice to exit with status code 1

### To rebuild the dockerimage:

```bash
./build.sh
```

Assumes you have the go toolchain (with the ability to cross-compile to different platforms) and docker installed and pointing at your docker daemon.
