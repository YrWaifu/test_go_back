from locust import HttpUser, task, between, stats


stats.PERCENTILES_TO_STATISTICS = [0.5, 0.8, 0.9, 0.99]
stats.PERCENTILES_TO_CHART = [0.5, 0.8, 0.9, 0.99]
stats.PERCENTILES_TO_REPORT = [0.5, 0.8, 0.9, 0.99]

class MyUser(HttpUser):
    host = "http://localhost:8080"
    wait_time = between(0.1, 0.2)

    def on_start(self):
        self.token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM5NzE4OTg1LCJpYXQiOjE3Mzk3MTUzODV9.m_q4TBMZF04bSNTBEXdKXPqlSQxe5wI1jLzg6xNXArQ"

    @task
    def get_info(self):
        headers = {
            "Authorization": f"Bearer {self.token}"
        }
        self.client.get("/api/info", headers=headers)