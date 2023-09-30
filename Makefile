name = com.github.davidkus.chrome-shortcuts
plist = ~/Library/LaunchAgents/$(name).plist

dev:
	air

build:
	go build -o out/

load:
	launchctl load $(plist)

unload:
	launchctl unload $(plist)

reload: unload load

status:
	launchctl list | grep $(name)

deploy:
	ln -hfs ~/Code/chrome-shortcuts/config.json /usr/local/etc/chrome-shortcuts/config.json
	ln -hfs ~/Code/chrome-shortcuts/info.plist $(plist)
	ln -hfs ~/Code/chrome-shortcuts/out/chrome-shortcuts /usr/local/bin/chrome-shortcuts

uninstall: unload
	rm -f $(plist)
	rm -f /usr/local/bin/chrome-shortcuts
	rm -f /usr/local/etc/chrome-shortcuts/config.json

logs:
	tail -f /usr/local/var/log/chrome-shortcuts.out.log
