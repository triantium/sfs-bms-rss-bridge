# SFS-Feed-Generator

Parst und generiert Feeds aus der übersicht der Freien Lehrgangsplätze von
https://lega.sfs-bayern.de/

## In

## API

```bash

curl localhost:8080/feed/atom
curl localhost:8080/feed/rss
curl localhost:8080/feed/json
```

## Verwendete Frameworks

https://github.com/PuerkitoBio/goquery

https://github.com/gorilla/feeds

https://github.com/gin-gonic/gin