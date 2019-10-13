hostname=flux-scanner

help:
	@grep '^[a-z].*:$$' Makefile|tr -d ':'

run:
	vagrant up

provision:
	vagrant provision

config-ssh:
	vagrant ssh-config|sed -e '1 s/default/$(hostname)/'|tee ~/.ssh/config.d/scanner
