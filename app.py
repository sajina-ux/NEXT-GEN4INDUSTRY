import os
import csv
from flask import Flask, render_template, request, redirect, url_for, flash

app = Flask(__name__)
app.secret_key = "next_gen_automation_secret_key"

BASE_DIR = os.path.abspath(os.path.dirname(__file__))
LEADS_CSV = os.path.join(BASE_DIR, "leads.csv")

# Mock database tracking active industrial Energy AI models
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

# Route 1: Legacy Lead Capturing Engine
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
        print(f"CSV Write Error: {e}")
    return redirect(url_for('home') + "#contact")

# Route 2: Central Energy AI Portfolio & MLOps Governance Workspace App
@app.route("/governance")
def governance_workspace():
    # Generate interactive model table blocks dynamically
    portfolio_rows = "".join([
        f"""<tr>
            <td style='padding:12px; border-bottom:1px solid #cbd5e0; font-weight:600;'>{model['name']}</td>
            <td style='padding:12px; border-bottom:1px solid #cbd5e0;'><span style='background:#e2e8f0; padding:4px 8px; border-radius:4px; font-size:0.85rem;'>{model['status']}</span></td>
            <td style='padding:12px; border-bottom:1px solid #cbd5e0; font-family:monospace; font-size:0.9rem;'>{model['dependencies']}</td>
            <td style='padding:12px; border-bottom:1px solid #cbd5e0; color:{"#e53e3e" if model['risk']=="High" else "#dd6b20" if model['risk']=="Medium" else "#38a169"}; font-weight:bold;'>{model['risk']}</td>
            <td style='padding:12px; border-bottom:1px solid #cbd5e0; color:#4a5568;'>{model['impact']}</td>
        </tr>""" for model in ENERGY_AI_PORTFOLIO
    ])

    return f"""
    <html>
    <head>
        <title>NEXT-GEN4INDUSTRY | Energy AI & MLOps Governance</title>
        <link rel="stylesheet" href="/static/style.css">
    </head>
    <body style="background-color:#f4f7f6; padding:0; margin:0;">
        <header class="navbar">
            <div class="logo">NEXT-GEN<span>4INDUSTRY</span></div>
            <nav><a href="/">Back to Core Site</a></nav>
        </header>
        
        <main style="padding:40px 8%; max-width:1200px; margin:0 auto;">
            <div style="background:#ffffff; padding:30px; border-radius:8px; box-shadow:0 4px 15px rgba(0,0,0,0.05); margin-bottom:40px;">
                <h1 style="color:#1a365d; margin-bottom:10px;">Central Energy AI Portfolio & MLOps Governance</h1>
                <p style="color:#718096; max-width:900px; margin-bottom:0;">
                    Standardizing development frameworks, agile methodologies, and scalable MLOps (Machine Learning Operations) practices to ensure reliable and highly secure deployment of AI models within heavy engineering and plant environments.
                </p>
            </div>

            <!-- Portfolio Framework Segment -->
            <h2 style="color:#1a365d; margin-bottom:20px;">Active AI Portfolio Registry</h2>
            <table style="width:100%; border-collapse:collapse; background:white; box-shadow:0 4px 6px rgba(0,0,0,0.02); border-radius:8px; overflow:hidden; margin-bottom:40px;">
                <thead>
                    <tr style="background-color:#1a365d; color:white; text-align:left;">
                        <th style="padding:12px;">Model Initiative</th>
                        <th style="padding:12px;">Deployment Status</th>
                        <th style="padding:12px;">SCADA / IoT Dependencies</th>
                        <th style="padding:12px;">Risk Index</th>
                        <th style="padding:12px;">Operational Business Value</th>
                    </tr>
                </thead>
                <tbody>
                    {portfolio_rows}
                </tbody>
            </table>

            <!-- MLOps Governance Core Guidelines Block -->
            <h2 style="color:#1a365d; margin-bottom:20px;">MLOps Infrastructure Standards</h2>
            <div style="display:grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap:25px;">
                <div style="background:white; padding:25px; border-radius:8px; border-left:4px solid #3182ce; box-shadow:0 2px 5px rgba(0,0,0,0.02);">
                    <h3 style="color:#2d3748; margin-bottom:10px;">Data Drift Protection</h3>
                    <p style="color:#4a5568; font-size:0.95rem;">Continuous out-of-bounds sensory tracking guarding production grid assets against data degradation and physical hardware drift errors.</p>
                </div>
                <div style="background:white; padding:25px; border-radius:8px; border-left:4px solid #3182ce; box-shadow:0 2px 5px rgba(0,0,0,0.02);">
                    <h3 style="color:#2d3748; margin-bottom:10px;">Plant Safety Containment</h3>
                    <p style="color:#4a5568; font-size:0.95rem;">Rigorous sandbox boundaries enforcing isolated testing spaces prior to injecting automated control hooks directly into live physical systems.</p>
                </div>
            </div>
        </main>
    </body>
    </html>
    """

if __name__ == "__main__":
    app.run(debug=True)
