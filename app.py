import os
from flask import Flask, render_template

app = Flask(__name__)
app.secret_key = "next_gen_automation_secret_key"

ENERGY_AI_PORTFOLIO = [
    {
        "name": "Smart Grid Load Forecaster",
        "status": "Production",
        "dependencies": "SCADA, Substation IoT Streams",
        "risk": "Low",
        "impact": "12% Peak Demand Reduction"
    },
    {
        "name": "Power Transformer Thermal Optimizer",
        "status": "Staging",
        "dependencies": "DGA Sensor Logs, Modbus Telemetry",
        "risk": "Medium",
        "impact": "Predictive Failure Mitigation"
    },
    {
        "name": "Renewable Generation Wind Mesh",
        "status": "In Development",
        "dependencies": "NVIDIA DGX, Weather API",
        "risk": "High",
        "impact": "Grid Balancing Optimization"
    }
]

@app.route("/")
def home():
    return render_template("index.html")

@app.route("/governance")
def governance_workspace():
    return render_template("governance.html", portfolio=ENERGY_AI_PORTFOLIO)

if __name__ == "__main__":
    app.run(debug=True)
