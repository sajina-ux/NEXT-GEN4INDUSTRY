import csv
import os
from flask import Flask, flash, redirect, render_template, request, url_for

app = Flask(__name__)
# A secret key is required by Flask to handle success messages safely
app.secret_key = "next_gen_automation_secret_key"

CSV_FILE_PATH = "leads.csv"


# Route for the main landing page
@app.route("/")
def home():
    return render_template("index.html")


# Route to handle form submissions
@app.route("/submit-inquiry", methods=["POST"])
def submit_inquiry():
    # 1. Capture form fields from the incoming request
    client_name = request.form.get("name")
    client_email = request.form.get("email")
    company_name = request.form.get("company")
    client_message = request.form.get("message")

    # 2. Check if the CSV file exists; if not, create it and write headers
    file_exists = os.path.isfile(CSV_FILE_PATH)

    with open(CSV_FILE_PATH, mode="a", newline="", encoding="utf-8") as file:
        writer = csv.writer(file)
        if not file_exists:
            writer.writerow(["Name", "Email", "Company", "Message"])  # Headers

        # 3. Append the client inquiry row
        writer.writerow([client_name, client_email, company_name, client_message])

    # 4. Display a friendly confirmation banner on the page
    flash("Thank you! Your automation inquiry has been safely received.")
    return redirect(url_for("home") + "#contact")


if __name__ == "__main__":
    app.run(debug=True)
