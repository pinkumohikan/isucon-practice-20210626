.PHONY: up down logs bench restart kataribe

gogo:
	sudo systemctl stop nginx
	sudo systemctl stop isu-go.service
	ssh isucon-app2 "sudo systemctl stop isu-go.service"
	ssh isucon-app3 "sudo systemctl stop mysql"
	sudo truncate --size 0 /var/log/nginx/access.log
	sudo truncate --size 0 /var/log/nginx/error.log
	ssh isucon-app3 "sudo truncate --size 0 /var/log/mysql/mysql-slow.log"
	$(MAKE) build
	scp webapp/golang/app isucon-app2:private_isu/webapp/golang/app
	ssh isucon-app3 "sudo systemctl start mysql"
	sleep 3
	sudo systemctl start isu-go.service
	ssh isucon-app2 "sudo systemctl start isu-go.service"
	sudo systemctl start nginx
	sleep 3
	$(MAKE) benchmark       

build:
	cd /home/isucon/isucon-practice-20210626/webapp/golang && go build -o app && cd /home/isucon/isucon-practice-20210626

benchmark:
	ssh ubuntu@172.31.27.253 "make bench"


kataribe:
	sudo cat /var/log/nginx/access.log | kataribe
