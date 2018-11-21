# Default env

Small lib for getting default value for ENV var

Order of reading var:
1. ENV
2. Yml file

Variables **reading only one time** when you call `GetInstance`

# Use and Example

```bash
    go get github.com/PxyUp/default-env
```

```go

    import (
        "github.com/PxyUp/default-env"
    )


    inst := default_env.GetInstance("env.yml") // ths is singleton and read only on time file and ENV var

    inst.Get("BACKEND_URL") // "http://main.go"
    inst.Get("FRONT_URL") // "http://frontend.go"
    inst.Get("SITE_URL") // "http://site.com"
```

```bash
    export SITE_URL=http://site.com
```

env.yml:

```yaml
    env_var:
        BACKEND_URL: http://main.go
        FRONT_URL: http://frontend.go
```