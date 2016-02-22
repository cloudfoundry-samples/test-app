# Lattice-App - a simple Go webapp

### Pushing the app to Cloud Foundry

```
cf push lattice-app
```

### Endpoints

`/`: a simple landing page displaying the index and uptime  
`/env`: displays environment variables  
`/exit`: instructs Lattice to exit with status code 1  
`/port`: returns the local port the request was received on

### Configure the app to listen on multiple ports

By providing a customer start command, you can configure the app to listen on multiple ports. The app responds the same way to each port.
```
cf push lattice-app -c "lattice-app --ports=7777,8888" 
```

### Pushing the app to Lattice as a Docker image

Lattice-App is packaged as a docker image at cloudfoundry/lattice-app

To push to [Lattice](https://github.com/cloudfoundry-incubator/lattice) using [ltc](https://github.com/cloudfoundry-incubator/lattice/ltc):

```bash
ltc create lattice-app cloudfoundry/lattice-app
```

### To rebuild the dockerimage:

```bash
./build.sh
```

Assumes you have the go toolchain (with the ability to cross-compile to different platforms) and docker installed and pointing at your docker daemon.
