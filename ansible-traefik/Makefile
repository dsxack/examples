ANSIBLE := ansible-playbook main.yml \
			-i hosts.ini

.PHONY: all
all:
	$(ANSIBLE)

.PHONY: traefik
traefik:
	$(ANSIBLE) --tags traefik

.PHONY: whoami
whoami:
	$(ANSIBLE) --tags whoami