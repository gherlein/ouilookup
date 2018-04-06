OUI_DB     = oui.txt
OUI_URL    = http://standards-oui.ieee.org/oui.txt
OUI_DB_DIR = /usr/local/var/oui
BINARY     = ./ouilookup

${BINARY}: dependencies
	go build -ldflags "-X main.oui_db=${OUI_DB_DIR}/${OUI_DB}" -o ${BINARY}

dependencies:
	go get ./...


install: ${OUI_DB} 
	-mkdir -p $(OUI_DB_DIR)
	-cp $(OUI_DB) $(OUI_DB_DIR)
	-cp ${BINARY} /usr/local/bin

${OUI_DB}:
	-curl $(OUI_URL) > $(OUI_DB)


clean:
	-rm *~
	-rm ${BINARY}
