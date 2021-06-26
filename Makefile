.PHONY: up down logs bench restart kataribe

gogo:
	sudo systemctl stop nginx
	sudo systemctl stop isu-go.service
	sudo truncate --size 0 /var/log/nginx/access.log
	sudo truncate --size 0 /var/log/nginx/error.log
	$(MAKE) build
	sudo systemctl start  isu-go.service
	sudo systemctl start nginx
	sleep 6
	$(MAKE) benchmark       

build:
	cd /home/isucon/isucon-practice-20210626/webapp/golang && go build -o app && cd /home/isucon/isucon-practice-20210626

benchmark:
	ssh ubuntu@172.31.27.253 "make bench"


kataribe:
	sudo cat /var/log/nginx/access.log | kataribe
