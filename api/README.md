# FoodUnit API
> Das Server-Backend von FoodUnit.

## Lokale Installation

### Composer

Die API verwendet das [Slim-Framework](http://www.slimframework.com/) für das HTTP-Routing. Diese und alle weiteren Abhängikeiten können mit [Composer](https://getcomposer.org/) installiert und aktuell gehalten werden.

Nach der Installation von Composer wechselt man ins API-Verzeichnis und kann die Abhängigkeiten herunterladen:

```
$ cd api
$ composer install
```

Bei Fragen zu Composer steht [Ablont](https://github.com/ablont) zur Verfügung.

### Datenbank

Anschließend muss eine lokale Datenbank erstellt werden (empfohlener Name ist `foodunit2`). Diese Datenbank kann dann mit dem Skript `.docker/sql/setup.sql` eingerichtet werden.

### Projekt-Konfiguration

Zur Konfiguration muss die Datei `conf/conf.ini.example` kopiert und zu `conf.ini` umbenannt werden. Anschließend können die umgebungsspezifischen Konfigurationswerte, z. B. der Datenbank-Host, eingetragen werden.

Für den SendGrid-API-Key wird ein kostenloses SendGrid-Konto benötigt.

Fertig!

## Docker-Installation

[Docker](https://www.docker.com/) rules.

### Projekt-Konfiguration

Einfach die Datei `conf/conf.ini.docker` kopieren, in `conf.ini` umbenennen und einen gültigen `sg_api_key` sowie die `mail_from` eintragen.

Anschließend:

```
$ docker-compose up --build
```

Fertig!
