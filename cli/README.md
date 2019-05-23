# FoodUnit CLI
> Das CLI-Tool für FoodUnit.

## Voraussetzungen

### Konfiguration

Die Konfigurationsdatei `conf.go.example` muss kopiert, in `conf.go` umbenannt und ausgefüllt werden.

### Go Toolchain

Das CLI-Tool ist in Go geschrieben. Um das Tool kompilieren zu können, muss die [Go Toolchain](https://golang.org/dl) installiert werden.

Falls noch nicht vorhanden, muss zunächst die Cobra-Library heruntergeladen werden.

```
$ go get github.com/spf13/cobra
```

Anschließend lässt sich eine ausführbare Datei erzeugen:

```
$ cd cli
$ go build -o foodunit.exe
```

Fertig!
