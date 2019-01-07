dumpcerts
===

Call up a TLS server and dump the verified certificate chain.

Usage
---

    # dumpcerts [-v] [-p <proto>] <address>

    $ ./dumpcerts -v -p tcp darfk.net:443 > darfk.net.pem
    connecting to darfk.net:443
    established a tls connection to darfk.net:443
    1 verified certificate chain(s) for darfk.net:443
    writing pem to stdout
    write certificate for darfk.net issued by Let's Encrypt Authority X3
    write certificate for Let's Encrypt Authority X3 issued by DST Root CA X3
    write certificate for DST Root CA X3 issued by DST Root CA X3
    $ cat darfk.net.pem
    -----BEGIN CERTIFICATE-----
    ...

Switches
---

- To force dumpcerts to use a specific network use `-p proto` (default is `tcp`), documented here [go doc net Dial](https://golang.org/pkg/net/#Dial).

- To increase verbosity use `-v`

License
---

MIT