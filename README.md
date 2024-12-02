# Website Monitoring Tool

This is a simple Go application that monitors the status of a list of websites. It performs multiple checks on the specified websites and reports their HTTP status codes, helping you identify if any site is down or encountering issues.

## Features

- Monitors a list of websites.
- Performs multiple checks with a delay between them.
- Reports HTTP status codes to indicate the health of the websites.

## Prerequisites

To run this project, you need:

- Go installed on your system (version 1.17 or later).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/guusebumps/website-monitoring-tool.git
   ```

2. Navigate to the project directory:
   ```bash
   cd website-monitoring-tool
   ```

## Usage

1. Run the program:
   ```bash
   go run main.go
   ```

2. The application will:
  * Test the predefined list of websites.
  * Report their status codes.
  * Repeat the process multiple times as defined in the monitoringTimes constant.

## Customizing the Application
You can modify the list of websites to monitor by updating the sites slice in the startMonitor function:
   ```bash
   sites := []string{"https://example.com", "https://another-example.com"}
   ```
You can also change the number of monitoring iterations (monitoringTimes) or the delay between checks (delay):
  ```bash
  const monitoringTimes = 4
  const delay = 2 // seconds
  ```

## Example Output

  ```bash
  ----------------------------------
  
  *** Testing 1 /3 times. ***
  
  Position: 0 // Site: https://httpstat.us/random/200,201,500-504
  Website loaded successfully (StatusCode): 200
  ----------------------------------
  Position: 1 // Site: https://www.alura.com.br
  Website loaded successfully (StatusCode): 200
  ----------------------------------
  Position: 2 // Site: https://www.caelum.com.br
  Website error (StatusCode): 503
  ----------------------------------
  ```

Happy monitoring! 🎉
