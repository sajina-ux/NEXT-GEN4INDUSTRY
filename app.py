import os
import csv
from flask import Flask, render_template, request, redirect, url_for, flash

app = Flask(__name__)
app.secret_key = "next_gen_automation_secret_key"

BASE_DIR = os.path.abspath(os.path.dirname(__file__))
LEADS_CSV = os.path.join(BASE_DIR, "leads.csv")

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

@app.route("/submit-inquiry", methods=["POST"])
def submit_inquiry():
    name = request.form.get("name", "").strip()
    email = request.form.get("email", "").strip()
    company = request.form.get("company", "").strip()
    message = request.form.get("message", "").strip()
    file_exists = os.path.exists(LEADS_CSV)
    try:
        with open(LEADS_CSV, mode="a", newline="", encoding="utf-8") as f:
            writer = csv.writer(f)
            if not file_exists:
                writer.writerow(["Name", "Email", "Company", "Message"])
            writer.writerow([name, email, company, message])
        flash("Thank you! Your inquiry has been safely processed.")
    except Exception as e:
        print(f"CSV Error: {e}")
    return redirect(url_for('home') + "#contact")

# Corrected layout parsing route pipeline
@app.route("/governance")
def governance_workspace():
    # Pass portfolio array safely to the true standalone layout canvas
    return render_template("governance.html", portfolio=ENERGY_AI_PORTFOLIO)

if __name__ == "__main__":
    app.run(debug=True)
