#!/usr/bin/env bash
$GOPATH/bin/godepgraph -s -i gopkg.in/mgo.v2,gopkg.in/mgo.v2/bson,gopkg.in/mgo.v2/internal.json,github.com/spf13/viper,github.com/jolestar/go-commons-pool,github.com/asaskevich/EventBus,github.com/labstack/echo,github.com/labstack/echo/middleware github.com/open-gtd/server  | dot -Tpng -o godepgraph.png
