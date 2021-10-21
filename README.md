# Responder
This is a very simple web-app which simply always returns HTTP status code 200.
It will also wait for an amount of time which can be set in URL parameters.

It's intended use is as a simulated upstream when evaluating network components in distributed systems.
My use case for example was to evaluate the performance profile and resource consumption of load balancers.

## Parametrization
`/static?delay=<delay in ms>` e.g. `/static?delay=25`
`/uniform?from=<from delay in ms>&to=<to delay in ms>` e.g. `/uniform?from=10&to=50`
`/lognorm?mean=<mean ms>&stdev=<standard deviation ms>` e.g. `/lognorm?mean=25&stdev=25`

You can also always use those without parameters. In this case the following defaults are used:

| `/static`  | `time=25`        |
|------------|------------------|
| `/uniform` | `from=10&to=30`  |
| `/lognorm` | `mean=25&stdev=10` |

## Request structure
Apart from the URL and POST request type, you can use any request you like, with any other parameters you like.
Nothing will happen.

## Response structure
Response will always be `HTTP 200 OK`, with a placeholder JSON response body:
```json
{
  "first": 0.12345,
  "second": 0.12345,
  "third": 0.12345,
  "forth": 0.12345,
  "fifth": 0.12345,
  "sixth": 0.12345
}
```

Feel free the change this in code should your use case require something else.
## Building
Generally you will build this just like you would any other go app.
I won't explain how to build a go app here.

### Docker
For debugging or simple local execution you can run the default `Dockerfile` in the root directory.

Depending on what your goals are you may need to build multi-arch docker images to make everything run on ARM systems as well.
To this end use the convenience scripts like so `scripts/docker-build-multiarch.sh <your version tag>`.

To push you can use a script too: `scripts/docker-push-all.sh <your version tag>`.

Please note that in the scripts the image name is set as `jjnp/responder`.
Feel free to change this to push images to whatever repository you need them in.

*Hint*: In case there are build errors with the ARM images you might need
to run the following command first, as described [here](https://github.com/multiarch/qemu-user-static)

`docker run --rm --privileged multiarch/qemu-user-static --reset -p yes`
