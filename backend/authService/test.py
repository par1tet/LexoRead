import requests

req = requests.post("http://127.0.0.1:8000/log", {"username": "test", "password": "test"})

print(req.text)