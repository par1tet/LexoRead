systemctl start docker
sudo docker compose -f docker-compose.yml up -d
python bot/main.py