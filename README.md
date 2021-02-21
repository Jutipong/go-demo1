## Install Gomon for run

```bash
go get -u github.com/TonPC64/gomon
```

## Initial

```bash
go test
```

## run

```bash
gomon .
```

## Postman collection

```bash
Go API Demo01.postman_collection.json
```

## SET IIS

```bash
To run on IIS, Create Application pool with No Managed Code
And config Physical path to web.config and our .exe file.
Then Start, and go to http://localhost:<your setup port>(for this setup just http://localhost).
```