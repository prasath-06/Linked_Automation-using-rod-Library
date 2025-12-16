# LinkedIn Automation using Go & Rod

## Project Overview

This project is a **Go-based LinkedIn automation proof-of-concept** built using the **Rod browser automation library**.
It demonstrates live browser automation with **human-like behavior simulation**, focusing on clean architecture and ethical automation practices.

The project is developed as a **technical demonstration**, not a production-scale bot.

---

## Project Objective

The objective of this project is to:

* Demonstrate advanced browser automation using **Go**
* Simulate **real human behavior** (typing delays, scrolling, navigation)
* Show understanding of **automation patterns**, **stealth awareness**, and **maintainable code structure**
* Provide a **proof-of-concept (POC)** that mimics authentic user interactions on LinkedIn

---

## Tech Stack

* **Language:** Go
* **Automation Library:** Rod
* **Browser:** Chromium (managed by Rod)
* **Platform:** Linux / macOS / Windows

---

## Project Structure

```
Linkedin_Automation/
│
├── main.go              # Entry point – controls overall automation flow
├── utils/  
│   └── browser.go       # Browser setup and human-like delay utilities
|   └── delay.go         # to delay like humanised
|   └── credentials.go   # contain username and password for linkedin
├── go.mod               # Go module definition
├── go.sum               # Dependency lock file
└── README.md            # Project documentation
```

---

## Automation Flow

1. Launch a visible Chromium browser using Rod
2. Navigate to LinkedIn login page
3. Enter login credentials with realistic typing delays
4. Submit login form and wait for feed to load
5. Navigate to LinkedIn people search page
6. Scroll search results in a human-like manner
7. Close browser gracefully after execution

---

## Human-like Behavior Simulation

To reduce bot-like patterns, the automation includes:

* Randomized delays between actions
* Gradual scrolling instead of instant jumps
* Visible (non-headless) browser execution
* Sequential navigation flow similar to real users

---

## How to Run the Project

### Install Go

Ensure Go is installed:

```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/your-username/Linkedin_Automation.git
cd Linkedin_Automation
```

### Initialize & Download Dependencies

```bash
go mod tidy
```

### Run the Automation

```bash
go run main.go
```

>  On first run, Rod will download a compatible Chromium browser automatically.

---

## Disclaimer

This project is created **strictly for educational and demonstration purposes**.

* It does NOT promote misuse of LinkedIn
* It does NOT scrape private data
* It does NOT send automated messages or spam connections

Users are responsible for complying with LinkedIn’s terms of service.

---

## Author

**Prasath Raja**
Final Year B.Tech – Artificial Intelligence & Data Science

---

