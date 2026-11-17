import csv
import os
import random
import time
from flask import Flask, render_template, request, redirect, url_for, jsonify

app = Flask(__name__)

# ==========================================================
# 1. WEBSITE & GOVERNANCE WORKSPACE ROUTES
# ==========================================================
@app.route("/")
def home():
    return render_template("index.html")

@app.route("/governance")
def governance():
    return render_template("governance.html")

@app.route("/submit-inquiry", methods=["POST"])
def submit_inquiry():
    name = request.form.get("name", "")
    email = request.form.get("email", "")
    company = request.form.get("company", "")
    message = request.form.get("message", "")

    file_exists = os.path.isfile("leads.csv")
    with open("leads.csv", mode="a", newline="", encoding="utf-8") as f:
        writer = csv.writer(f)
        if not file_exists:
            writer.writerow(["Timestamp", "Name", "Email", "Company", "Message"])
        writer.writerow([time.strftime("%Y-%m-%d %H:%M:%S"), name, email, company, message])

    return redirect(url_for("home"))

# ==========================================================
# 2. AI4INDUSTRY.BAA — EMBEDDED EDGE AI, DIGITAL TWIN & RAG API
# ==========================================================
OEM_MANUAL_SOPS = [
    {
        "sop_id": "SOP-CNC-101",
        "title": "Spindle Overheating & Coolant Flow Recovery",
        "keywords": ["temp", "temperature", "heat", "overheat", "coolant", "chiller", "75"],
        "action": "Execute Emergency Feed Hold (Code M00). Verify Chiller_Pump_CP20 pressure (3.2–4.0 bar) and clean the 50-micron inline coolant strainer."
    },
    {
        "sop_id": "SOP-CNC-204",
        "title": "High-Frequency Spindle Vibration & Bearing Inspection",
        "keywords": ["vibration", "vib", "bearing", "b7008", "belt", "3.5", "shake"],
        "action": "Isolate main breaker (LOTO #L-09). Inspect Ceramic_Bearings_B7008 for axial play (>0.015 mm requires replacement) and check Drive_Belt_HTD5M tension (145 Hz)."
    },
    {
        "sop_id": "SOP-CNC-309",
        "title": "Combined Thermal-Vibration Runaway Protocol",
        "keywords": ["both", "runaway", "seizure", "critical", "emergency", "80", "4.5"],
        "action": "Immediate spindle seizure risk. Do NOT run cool-down rotation. Cut power immediately, inspect ISO VG 32 lubrication mist unit, and replace B7008 bearing cartridge."
    }
]

@app.route("/api/ai4industry/telemetry/<machine_id>", methods=["GET"])
def ai4industry_telemetry(machine_id):
    is_fault = random.random() < 0.15
    temp = round(random.uniform(86.0, 98.5), 2) if is_fault else round(random.uniform(48.0, 54.5), 2)
    vib = round(random.uniform(4.8, 7.6), 2) if is_fault else round(random.uniform(0.8, 1.4), 2)
    rpm = random.randint(2100, 2400) if is_fault else random.randint(2950, 3050)

    if is_fault:
        diagnosis = (
            f"[Graph Trace: {machine_id} -> Spindle_Assembly -> Ceramic_Bearings_B7008] "
            f"Combined thermal/vibration anomaly detected. Execute SOP-CNC-309 immediately."
        )
    else:
        diagnosis = "Nominal operation verified against Asset Knowledge Graph & IEC 62443 OT-IT boundary rules."

    return jsonify({
        "machine_id": machine_id,
        "temperature_c": temp,
        "vibration_mms": vib,
        "rpm": rpm,
        "status": "CRITICAL" if is_fault else "HEALTHY",
        "ai_diagnosis": diagnosis,
        "timestamp": time.strftime("%H:%M:%S")
    })

@app.route("/api/ai4industry/copilot", methods=["POST"])
def ai4industry_copilot():
    data = request.get_json(silent=True) or {}
    question = data.get("question", "").lower()

    best_sop = OEM_MANUAL_SOPS[0]
    best_score = -1
    for sop in OEM_MANUAL_SOPS:
        score = sum(1 for kw in sop["keywords"] if kw in question)
        if score > best_score:
            best_score = score
            best_sop = sop

    return jsonify({
        "sop_id": best_sop["sop_id"],
        "title": best_sop["title"],
        "answer": f"[{best_sop['sop_id']} — {best_sop['title']}] {best_sop['action']}"
    })

if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    app.run(host="0.0.0.0", port=port, debug=False)
