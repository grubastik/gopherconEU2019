http://workshop.grahovac.pro


- go mod - dependency management
- goproxy - speedup download and update
- testing - testify lib but need to think about more idiomatic code
- profiling - simple to use. enable by defult in prod app

Observability:
- log tracing, metrics, etc.
- use panic wisely
- opensensu -> datadog, eger, prometeus,...

Matrics:
- bussines - understanding usage
- tracing - understanding data flow

Operability:
- healthy
- readiness
- split servers for diagnostics and bussines

Handle os flags correctly

Use version during compilation time using ldFlags

Don't hesitate to use Makefile

Prepare "Cloud native Hello world" service template (caldera)

Keel
k3s
k3d