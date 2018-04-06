# ouilookup

This is a simple command line tool to lookup the equipment maker associated with a MAC address.  The first three octets of a MAC address is the "Organizationally Unique Identifier" (OUI) and is registered with the IEEE.

## Usage

A typical use of this tool is to lookup what kind of device is assocaited with a MAC address on the network.  For example, suppose you use 'arp' and notice that you see a device on your network in the pool of addresses you reserve for DHCP addresses and you'd like to know what kind of device it is.  Example:

```
gherlein@jupiter:~$ ouilookup d0:e1:40:b4:0e:6a
d0:e1:40:b4:0e:6d   Apple, Inc.    
```

Nothing special - just a fast way to see what kind of device is associated with that address.


## Installation from Code

To compile and build:

```
make
sudo make install
```

If you look at the makefile you will see that it downloads the latest OUI database textfile and installs it on your system as part of the install.  This Makefile uses a trick to embed the installed location of that file into the code so that thre's no disconnect if you change the location.  This trick uses -ldflags to pass in a variable:

```
go build -ldflags "-X main.oui_db=${OUI_DB_DIR}/${OUI_DB}"
```

## Credits

Uses the excellent [OUI Library](https://github.com/klauspost/oui) by [Klaus Post](https://github.com/klauspost).


## References

[Organizationally Unique Identifier - Wikipedia](https://en.wikipedia.org/wiki/Organizationally_unique_identifier)
