# Microfest
Manifest and configuration management for single page applications driven by micro frontend apps. Built to scale
horizontally across multiple production and non-production environments by default.

- [Background](#background)
- [Overview](#overview)
  * [`/manifest`](#--manifest-)
    + [`POST`](#-post-)
    + [`GET`](#-get-)
    + [`PUT`](#-put-)
  * [`/configuration`](#--configuration-)
    + [`POST`](#-post--1)
    + [`GET`](#-get--1)
  * [`/healthcheck`](#--healthcheck-)
    + [`GET`](#-get--2)
- [Javascript Example](#javascript-example)
- [Deploying on Kubernetes](#deploying-on-kubernetes)
- [Performance](#performance)

## Background
If you have adopted a micro frontend metaframework like [single-spa](https://github.com/CanopyTax/single-spa),
you will have noticed a bit of a disconnect between the development, build and deployment stories. While providing
more flexibility in development of new features as a SPA evolves with relatively little cost, the build and
deployment complexity increases with every new micro app added.

At a small scale, it possible to:
* Build all the micro app bundles at the same time
* Take the bundle hashes and update a manifest with them
* Run a container with everything built and the updated manifest inside

As more micro apps are added and the build time increases:
* It becomes ineffective to rebuild everything for changes to one micro app
* It make more sense to build bundles individually and serve them from a bucket or CDN

After the build process becomes independent:
* Updating the manifest with the newly hashed bundle names becomes trickier
* Either a new container can be rebuilt every time with the manifest generated from build args
* Or a new manifest can be mounted into the running container using a Docker volume or Kubernetes ConfigMap
* But then this requires killing the existing containers and spinning up new ones to pick up the changes

When the deployment process becomes independent:
* There needs to be a way to keep track of manifest versioning and history
* A `git` repo with automated commits and pushes on CI can handle this to a point
* But you need to either implement a semaphore or be sure to not be deploying two micro apps at the same time
* And then you also probably need a way to push more than one micro app at once in some cases

Finally, managing all of this across multiple production and non-production environments has the potential
to get very messy very quickly. Enter `microfest`.

## Overview
`microfest` provides a RESTful API for managing releases and manifests for micro frontend applications
which scales horizontally across multiple production and non-production environments with minimal
configuration and zero templating. Under the hood, `microfest` uses [bbolt](https://github.com/etcd-io/bbolt)
as a fast and reliable key/value store that is optimised for read-intensive workloads. As `bbolt` uses an
exclusive write lock on the database, this nicely handles the potential problem of multiple concurrent updates.

`microfest` exposes the following routes:
### `/manifest`
#### `POST`
Add a new, complete manifest for a hostname and set it as the manifest to be fetched when calling `GET`.

```json5
// full-manifest.json
{
  "navigation": "https://storage.googleapis.com/XXX/navigation.HASH.bundle.js",
  "settings": "https://storage.googleapis.com/XXX/settings.HASH.bundle.js",
  "help": "https://storage.googleapis.com/XXX/help.HASH.bundle.js"
}
```

```bash
curl -X POST 'http://localhost:8000/manifest?host=production.host' \
     -H 'Content-Type: application/json' -H 'X-API-KEY: XXX' \
     -d @full-manifest.json
     
# created manifest for host production.host
```

#### `GET`
```json5
// curl -X GET 'http://localhost:8000/manifest?host=production.host'

{
  "help": "https://storage.googleapis.com/XXX/help.HASH.bundle.js",
  "navigation": "https://storage.googleapis.com/XXX/navigation.HASH.bundle.js",
  "settings": "https://storage.googleapis.com/XXX/settings.HASH.bundle.js"
}
```

#### `PUT`
Create a new manifest by taking the current manifest for a hostname and patching the
diff from the payload.
```json5
// partial-manifest.json
{
  "settings": "https://storage.googleapis.com/XXX/settings.FIXED.bundle.js"
}
```

```bash
curl -X PUT 'http://localhost:8000/manifest?host=production.host' \
     -H 'Content-Type: application/json' -H 'X-API-KEY: XXX' \
     -d @partial-manifest.json
     
# created manifest for host production.host
```

```json5
// curl -X GET 'http://localhost:8000/manifest?host=production.host'
{
  "help": "https://storage.googleapis.com/XXX/help.HASH.bundle.js",
  "navigation": "https://storage.googleapis.com/XXX/navigation.HASH.bundle.js",
  "settings": "https://storage.googleapis.com/XXX/settings.FIXED.bundle.js"
}
```

### `/configuration`
#### `POST`
Add a new configuration version for a hostname and set it as the configuration to be fetched when calling `GET`.

```json5
// configuration.json
{
  "value1": "xxxxx",
  "value2": "xxxxx"
}
```

```bash
curl -X POST 'http://localhost:8000/configuration?host=production.host' \
     -H 'Content-Type: application/json' -H 'X-API-KEY: XXX' \
     -d @configuration.json
     
# created configuration for host production.host
```

#### `GET`
Get the configuration for a specific hostname.
```bash
curl -X GET 'http://localhost:8000/configuration?host=production.host'
```

```json
{
  "value1": "xxxxx",
  "value2": "xxxxx"
}
```

### `/healthcheck`
#### `GET`
A healthcheck endpoint for readiness and liveness probes.

```bash
curl -X GET 'http://localhost:8000/healthcheck'

# HTTP/1.1 200 OK
# Date: Sat, 13 Jul 2019 12:39:16 GMT
# Content-Length: 0
```

## Javascript Example
Once you have microfest deployed, it can be called to dynamically load the manifest for the correct
environment in a `single-spa` application like this:

```js
fetch(`https://mf.example.com/manifest?host=${window.location.hostname}`, { headers: {"X-API-KEY": apiKey }})
  .then(res => res.json())
  .then((manifest) => {
    window.manifest = manifest;

    navigation();
    settings();
    help();

    singleSpa.start();
  });
```

## Deploying on Kubernetes
In the `./kubernetes` a set of [kustomize](https://kustomize.io/) manifests are provided to help you get up and running
with `microfest` as quickly as possible on Kubernetes. Modify the values with the comments `# your own X here`:

```text
./kubernetes/deployment.yaml
16:                  - key: node-type # your own key here
19:                      - general # your own value here
53:          key: group # your own key here
55:          value: general # your own value here

./kubernetes/ingress.yaml
13:    - host: microfest.example.com # your own host here
21:        - microfest.example.com # your own host here
22:      secretName: tls # your own tls secret name here

./kubernetes/kustomization.yaml
7:namespace: microfest # your own namespace here
16:      - API_KEY=bla # your own API key here

./kubernetes/namespace.yaml
6:  name: microfest # your own namespace here
```

It is important ensure that Node Affinity and Taint Tolerations for your nodes are set when deploying `microfest`, to
ensure that any subsequent redeployments of `microfest` will always find the BoltDB data volume on the same node to
reattach.

Once you have modified the values, run `kustomize build ./kubernetes | kubectl apply -f -` to get `microfest` running
on your cluster.

## Performance
Report from running `vegeta attack -duration=60s` on the `GET /manifest` route:

```text
Requests      [total, rate]            3000, 50.02
Duration      [total, attack, wait]    59.982073627s, 59.980867273s, 1.206354ms
Latencies     [mean, 50, 95, 99, max]  1.556704ms, 1.185781ms, 1.99684ms, 4.028128ms, 234.552687ms
Bytes In      [total, mean]            3558000, 1186.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:3000
Error Set:
```
