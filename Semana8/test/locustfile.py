from locust import HttpUser, task, between
import random

class MilitaryReportUser(HttpUser):
    wait_time = between(0.1, 0.5)

    @task
    def send_report(self):
        countries =["USA", "RUS", "CHN", "ESP", "GMT"]
        payload = {
            "country": random.choice(countries),
            "warplanes_in_air": random.randint(0, 50),
            "warships_in_water": random.randint(0, 30),
            "timestamp": "2026-03-12T20:15:30Z"
        }
        self.client.post("/grpc-2026", json=payload)