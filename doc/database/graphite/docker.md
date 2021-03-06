# Docker

- https://github.com/graphite-project/docker-graphite-statsd
  - https://github.com/graphite-project/docker-graphite-statsd#mounted-volumes

Host | Container | Service
---- | --------- | -------------------------------------------------------------------------------------------------------------------
  80 |        80 | [nginx](https://www.nginx.com/resources/admin-guide/)
2003 |      2003 | [carbon receiver - plaintext](http://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-plaintext-protocol)
2004 |      2004 | [carbon receiver - pickle](http://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-pickle-protocol)
2023 |      2023 | [carbon aggregator - plaintext](http://graphite.readthedocs.io/en/latest/carbon-daemons.html#carbon-aggregator-py)
2024 |      2024 | [carbon aggregator - pickle](http://graphite.readthedocs.io/en/latest/carbon-daemons.html#carbon-aggregator-py)
8080 |      8080 | Graphite internal gunicorn port (without Nginx proxying).
8125 |      8125 | [statsd](https://github.com/etsy/statsd/blob/master/docs/server.md)
8126 |      8126 | [statsd admin](https://github.com/etsy/statsd/blob/v0.7.2/docs/admin_interface.md)