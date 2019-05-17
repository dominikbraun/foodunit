# FoodUnit API
> Das Server-Backend von FoodUnit

## Voraussetzungen

### Composer

Die API verwendet das [Slim-Framework](http://www.slimframework.com/) für das HTTP-Routing. Diese und alle weiteren Abhängikeiten können mit [Composer](https://getcomposer.org/) installiert und aktuell gehalten werden.

Nach der Installation vom Composer wechselt man ins API-Verzeichnis und kann die Abhängigkeiten herunterladen:

```
$ cd api
$ composer install
```

Bei Fragen zu Composer steht [Ablont](https://github.com/ablont) zur Verfügung.

### Datenbank

Die Datenbank lässt sich mit dem Skript `database/setup.sql` einfach einrichten.

### Projekt-Konfiguration

Umgebungsspezifische Konfigurationen, z. B. der Datenbank-Nutzername, müssen in `conf/conf.ini.example` eingetragen werden. Anschließend kann die Datei kopiert und in `conf.ini` umbenannt werden.

Fertig!