# CHANGELOG

### [v0.2.7](https://github.com/speelynet/server/tree/v0.2.7)
- Upgrade [NGINX Unit](https://unit.nginx.org) version in [CI/CD workflow](https://github.com/speelynet/server/blob/v0.2.7/.github/workflows/CICD.yml)

### [v0.2.6](https://github.com/speelynet/server/tree/v0.2.6)
- Upgrade [NGINX Unit](https://unit.nginx.org) version in [CI/CD workflow](https://github.com/speelynet/server/blob/v0.2.6/.github/workflows/CICD.yml)

### [v0.2.5](https://github.com/speelynet/server/tree/v0.2.5)
- Update [components route](https://github.com/speelynet/server/blob/v0.2.5/components.go#L25)

### [v0.2.4](https://github.com/speelynet/server/tree/v0.2.4)
- Update logic for [defining subroutes of components route](https://github.com/speelynet/server/blob/v0.2.4/components.go#L21)
- Set 'Access-Control-Allow-Origin' header to '*' for [components](https://github.com/speelynet/server/blob/v0.2.4/components.go#L22) route

### [v0.2.3](https://github.com/speelynet/server/tree/v0.2.3)
- Modify [home](https://github.com/speelynet/server/blob/v0.2.3/home.go#L5) route to serve "home" directory instead of "static" directory

### [v0.2.2](https://github.com/speelynet/server/tree/v0.2.2)
- Disable strict-slash routing to fix redirection issue in [components](https://github.com/speelynet/server/blob/v0.2.2/components.go) route

### [v0.2.1](https://github.com/speelynet/server/tree/v0.2.1)
- Fix redirection bug in [components](https://github.com/speelynet/server/blob/v0.2.1/components.go) route

### [v0.2.0](https://github.com/speelynet/server/tree/v0.2.0)
- Start using [Godog](https://github.com/cucumber/godog#readme) for [Cucumber](https://cucumber.io/) tests
- Rewrite server to use a [gorilla/mux](https://github.com/gorilla/mux#readme) [router](https://github.com/speelynet/server/blob/v0.2.0/main.go#L9)
- Add [home](https://github.com/speelynet/server/blob/v0.2.0/home.go) and [components](https://github.com/speelynet/server/blob/v0.2.0/components.go) routes

### [v0.1.0](https://github.com/speelynet/server/tree/v0.1.0)
- Initial release
