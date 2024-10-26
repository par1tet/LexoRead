if sudo docker ps | grep -- 'postgres'
then echo postgres is alredy started
else
echo starting postgres
systemctl start docker
sudo npx kill-port 5432
sudo docker compose -f ~/docker-compose.yml up -d
fi
python main.py