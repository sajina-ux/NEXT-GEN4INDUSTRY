import os
import csv
from flask import Flask, render_template, request, redirect, url_for, flash

app = Flask(__name__)
# Secret key enables the secure success pop-up notification system
app.secret_key = "next_gen_automation_secret_key"

# Pinpoints the exact folder path to save the spreadsheet file safely
BASE_DIR = os.path.abspath(os.path.dirname(__file__))
CSV_FILE_PATH = os.path.join(BASE_DIR, "leads.csv")

@app.route("/")
def home():
    return render_template("index.html")

# Route 1: Handle incoming customer form submissions
@app.route("/submit-inquiry", methods=["POST"])
def submit_inquiry():
    name = request.form.get("name", "").strip()
    email = request.form.get("email", "").strip()
    company = request.form.get("company", "").strip()
    message = request.form.get("message", "").strip()

    file_exists = os.path.exists(CSV_FILE_PATH)

    try:
        with open(CSV_FILE_PATH, mode="a", newline="", encoding="utf-8") as csv_file:
            writer = csv.writer(csv_file)
            if not file_exists:
                writer.writerow(["Name", "Email", "Company", "Message"]) # Headers
            
            writer.writerow([name, email, company, message]) # Data Entry row
            
        flash("Thank you! Your automation inquiry has been safely received.")
    except Exception as e:
        print(f"File writing error: {e}")
        flash("An error occurred while saving your inquiry.")

    return redirect(url_for('home') + "#contact")

# Route 2: Hidden Admin Dashboard to read and display your CSV leads file
@app.route("/admin/leads")
def view_leads():
    leads_list = []
    headers = ["Name", "Email", "Company", "Message"]
    
    if os.path.exists(CSV_FILE_PATH):
        try:
            with open(CSV_FILE_PATH, mode="r", encoding="utf-8") as csv_file:
                reader = csv.reader(csv_file)
                file_content = list(reader)
                if file_content:
                    headers = file_content[0] # Extract the headers row
                    leads_list = file_content[1:] # Extract all customer entries
        except Exception as e:
            print(f"Error reading CSV: {e}")

    # Directly outputs a clean HTML table inside your browser window
    table_rows = "".join([
        f"<tr>" + "".join([f"<td style='padding:12px; border:1px solid #cbd5e0;'>{cell}</td>" for cell in row]) + "</tr>"
        for row in leads_list
    ])
    
    header_cols = "".join([f"<th style='padding:12px; background-color:#1a365d; color:white; border:1px solid #cbd5e0; text-align:left;'>{h}</th>" for h in headers])

    return f"""
    <html>
    <head><title>NEXT-GEN4INDUSTRY | Admin Panel</title></head>
    <body style="font-family:'Segoe UI',Arial,sans-serif; background-color:#f9fbfd; margin:40px;">
        <h2 style="color:#1a365d;">Captured Client Inquiries (leads.csv)</h2>
        <p style="color:#718096; margin-bottom:20px;">Reviewing live pipeline customer queries processed by your Python backend server.</p>
        <table style="width:100%; border-collapse:collapse; background:white; box-shadow:0 4px 6px rgba(0,0,0,0.05); border-radius:8px; overflow:hidden;">
            <thead><tr>{header_cols}</tr></thead>
            <tbody>{table_rows if leads_list else "<tr><td colspan='4' style='padding:20px; text-align:center; color:#a0aec0;'>No customer entries recorded yet.</td></tr>"}</tbody>
        </table>
        <br><a href="/" style="color:#3182ce; font-weight:600; text-decoration:none;">&larr; Back to Live Website</a>
    </body>
    </html>
    """

if __name__ == "__main__":
    app.run(debug=True)
