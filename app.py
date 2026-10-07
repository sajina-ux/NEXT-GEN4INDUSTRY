import os
import csv
from flask import Flask, render_template, request, redirect, url_for, flash

app = Flask(__name__)
app.secret_key = "next_gen_automation_secret_key"

# This forces the CSV to be saved in your exact website folder path
BASE_DIR = os.path.abspath(os.path.dirname(__file__))
CSV_FILE_PATH = os.path.join(BASE_DIR, "leads.csv")

@app.route("/")
def home():
    return render_template("index.html")

@app.route("/submit-inquiry", methods=["POST"])
def submit_inquiry():
    # 1. Safely extract values from the form inputs
    name = request.form.get("name", "").strip()
    email = request.form.get("email", "").strip()
    company = request.form.get("company", "").strip()
    message = request.form.get("message", "").strip()

    # 2. Check if the CSV exists to know if we need header rows
    file_exists = os.path.exists(CSV_FILE_PATH)

    try:
        # 3. Open the target spreadsheet in append mode
        with open(CSV_FILE_PATH, mode="a", newline="", encoding="utf-8") as csv_file:
            writer = csv.writer(csv_file)
            
            if not file_exists:
                # Write standard tracker headers if it is a brand new file
                writer.writerow(["Name", "Email", "Company", "Message"])
            
            # Write out your data entry fields row
            writer.writerow([name, email, company, message])
            
        flash("Thank you! Your automation inquiry has been safely received.")
    except Exception as e:
        # If a file layout block occurs, log it directly to your terminal
        print(f"Database error writing to CSV file: {e}")
        flash("An error occurred while handling your transmission.")

    return redirect(url_for('home') + "#contact")

if __name__ == "__main__":
    app.run(debug=True)
