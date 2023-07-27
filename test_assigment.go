package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

// Data structure for request arguments
type RequestArguments struct {
	Title     string `json:"title"`
	Value     string `json:"value"`
	Status    string `json:"status"`
	Date      string `json:"add_time"`
	CloseDate string `json:"expected_close_date"`
}

// Data structure for deals
type Deal struct {
	ID                *int    `json:"id"`
	Title             *string `json:"title"`
	Value             *int    `json:"value"`
	Status            *string `json:"status"`
	AddTime           *string `json:"add_time"`
	ExpectedCloseDate *string `json:"expected_close_date"`
}

// Data structure for metrics
type Metrics struct {
	TotalCount    int
	TotalDuration time.Duration
	AvgDuration   time.Duration
	MaxDuration   time.Duration
	MinDuration   time.Duration
	TotalLatency  time.Duration
	AvgLatency    time.Duration
	MaxLatency    time.Duration
	MinLatency    time.Duration
}

// Constants for API URL and API Token
const apiURL = "https://zakhar-sandbox.pipedrive.com/api/v1/"
const apiToken = "d42a0ddf456139752b9e6c16fd952a62fecdc987"

// Global variable for metrics
var RequestMetrics *Metrics = new(Metrics)

// This function is used to display metrics
func handleMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Total requests: \t %d\n", RequestMetrics.TotalCount)
	fmt.Fprintf(w, "Total duration: \t %s\n", RequestMetrics.TotalDuration)
	fmt.Fprintf(w, "Avg duration: \t\t %s\n", RequestMetrics.AvgDuration)
	fmt.Fprintf(w, "Max duration: \t\t %s\n", RequestMetrics.MaxDuration)
	fmt.Fprintf(w, "Min duration: \t\t %s\n", RequestMetrics.MinDuration)
	fmt.Fprintf(w, "Total latency: \t\t %s\n", RequestMetrics.TotalLatency)
	fmt.Fprintf(w, "Avg Latency: \t\t %s\n", RequestMetrics.AvgLatency)
	fmt.Fprintf(w, "Max latency: \t\t %s\n", RequestMetrics.MaxLatency)
	fmt.Fprintf(w, "Min latency: \t\t %s\n", RequestMetrics.MinLatency)

}

// This function is used to write, calculate, display in the terminal
// and update the metrics
func WriteMetrics(duration time.Duration, latency time.Duration) {
	RequestMetrics.TotalCount++
	RequestMetrics.TotalDuration += duration
	RequestMetrics.TotalLatency += latency

	// Calculate average duration and average latency
	RequestMetrics.AvgDuration = RequestMetrics.TotalDuration / time.Duration(RequestMetrics.TotalCount)
	RequestMetrics.AvgLatency = RequestMetrics.TotalLatency / time.Duration(RequestMetrics.TotalCount)

	// Update Max and Min duration
	if RequestMetrics.TotalCount == 1 {
		RequestMetrics.MaxDuration = duration
		RequestMetrics.MinDuration = duration
	} else {
		if RequestMetrics.MaxDuration < duration {
			RequestMetrics.MaxDuration = duration
		}
		if RequestMetrics.MinDuration > duration {
			RequestMetrics.MinDuration = duration
		}
	}
	if RequestMetrics.TotalCount == 1 {
		RequestMetrics.MaxLatency = latency
		RequestMetrics.MinLatency = latency
	} else {
		if RequestMetrics.MaxLatency < latency {
			RequestMetrics.MaxLatency = latency
		}
		if RequestMetrics.MinLatency > latency {
			RequestMetrics.MinLatency = latency
		}
	}

	fmt.Println("Duration:\t", duration)
	fmt.Println("Latency:\t", latency)
	fmt.Println("Total requests:\t", RequestMetrics.TotalCount)
	fmt.Println("Total duration:\t", RequestMetrics.TotalDuration)
	fmt.Println("Avg duration:\t", RequestMetrics.AvgDuration)
	fmt.Println("Max duration:\t", RequestMetrics.MaxDuration)
	fmt.Println("Min duration:\t", RequestMetrics.MinDuration)
	fmt.Println("Total latency:\t", RequestMetrics.TotalLatency)
	fmt.Println("Avg Latency:\t", RequestMetrics.AvgLatency)
	fmt.Println("Max latency:\t", RequestMetrics.MaxLatency)
	fmt.Println("Min latency:\t", RequestMetrics.MinLatency)

}

// This function is used to send, receive and display all the deals
// using GET request
func handleGetResponse(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	resp, err := http.Get(apiURL + "deals?api_token=" + apiToken)
	if err != nil {
		fmt.Println("Error sending API request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Add deal response status:", resp.Status)

	latency := time.Since(start)

	// Read the response body
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		return
	}

	// Retrieve the needed information from the response body
	deals, err := retrieveDealsFromJSON(string(response))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the retrieved deals
	for _, dealInfo := range deals {
		fmt.Fprintf(w, dealInfo+"\n")
		fmt.Println(dealInfo)
	}

	duration := time.Since(start)
	WriteMetrics(duration, latency)
}

// This function is used to send a POST request to add new deal
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Parse the form data
	r.ParseForm()
	title := r.FormValue("title")
	value := r.FormValue("value")
	status := r.FormValue("status")
	add_time := r.FormValue("date")
	expected_close_date := r.FormValue("exp_close_date")

	// Validate the input data
	result := ValidateInputs(title, value, status, add_time, expected_close_date)
	if result != nil {
		fmt.Fprintf(w, "Error: %s", result)
		return
	}

	// Create the request body
	RequestArguments := RequestArguments{
		Title:     title,
		Value:     value,
		Status:    status,
		Date:      add_time,
		CloseDate: expected_close_date,
	}

	//Creating JSON request body
	requestBody, err := json.Marshal(RequestArguments)
	if err != nil {
		fmt.Println(w, "Error creating JSON:", err)
		return
	}
	//fmt.Println(string(requestBody))
	resp, err := http.Post(apiURL+"deals?api_token="+apiToken, "application/json", bytes.NewBuffer(requestBody))

	latency := time.Since(start)

	if err != nil {
		fmt.Println(w, "Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Fprintf(w, "POST request failed with status: %d\n", resp.StatusCode)
		return
	}
	fmt.Fprintf(w, "POST request was send")

	duration := time.Since(start)
	WriteMetrics(duration, latency)
}

// This function is used to send a PUT request to update a deal
func handlePutRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Parse the form data
	r.ParseForm()
	id := r.FormValue("id")
	title := r.FormValue("title")
	value := r.FormValue("value")
	status := r.FormValue("status")
	add_time := r.FormValue("date")
	expected_close_date := r.FormValue("exp_close_date")

	//Displaying the id of the deal to be updated
	//And validating the input data
	fmt.Println("ID: " + id)
	result := ValidateInputs(title, value, status, add_time, expected_close_date)

	if result != nil {
		fmt.Fprintf(w, "Error: %s", result)
		return
	} else if id == "" || !regexp.MustCompile(`^[0-9]+$`).MatchString(id) {
		fmt.Fprintf(w, "Error: id is not valid")
		return
	}

	// Create the request body
	RequestArguments := RequestArguments{
		Title:     title,
		Value:     value,
		Status:    status,
		Date:      add_time,
		CloseDate: expected_close_date,
	}

	requestBody, err := json.Marshal(RequestArguments)
	if err != nil {
		fmt.Println(w, "Error creating JSON:", err)
		return
	}

	// Create a PUT request
	req, err := http.NewRequest(http.MethodPut, apiURL+"deals/"+id+"?api_token="+apiToken, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the PUT request using the default http.Client
	client := &http.Client{}
	resp, err := client.Do(req)

	latency := time.Since(start)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("PUT request failed with status: %d\nCheck the deal id", resp.StatusCode)
		fmt.Fprintf(w, "Error sending PUT request: %d\nCheck the deal id", resp.StatusCode)
		return
	}

	fmt.Fprintf(w, "PUT request was send")

	duration := time.Since(start)
	WriteMetrics(duration, latency)
}

// This function is used to send a retrieve needed deals from the API response
func retrieveDealsFromJSON(jsonData string) ([]string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON data: %w", err)
	}

	var deals_arr []string
	deals := data["data"].([]interface{})
	for _, deal := range deals {
		dealMap := deal.(map[string]interface{})

		// Handle nullable fields using pointers
		var id *int
		if dealMap["id"] != nil {
			idValue := int(dealMap["id"].(float64))
			id = &idValue
		}

		var title *string
		if dealMap["title"] != nil {
			titleValue := dealMap["title"].(string)
			title = &titleValue
		}

		var value *int
		if dealMap["value"] != nil {
			valueValue := int(dealMap["value"].(float64))
			value = &valueValue
		}

		var status *string
		if dealMap["status"] != nil {
			statusValue := dealMap["status"].(string)
			status = &statusValue
		}

		var addTime *string
		if dealMap["add_time"] != nil {
			addTimeValue := dealMap["add_time"].(string)
			addTime = &addTimeValue
		}

		var expectedCloseDate *string
		if dealMap["expected_close_date"] != nil {
			expectedCloseDateValue := dealMap["expected_close_date"].(string)
			expectedCloseDate = &expectedCloseDateValue
		}

		// Create the Deal struct with the nullable fields
		dealInfo := Deal{
			ID:                id,
			Title:             title,
			Value:             value,
			Status:            status,
			AddTime:           addTime,
			ExpectedCloseDate: expectedCloseDate,
		}

		// Format the deal information and append it to the deals_arr
		deals_arr = append(deals_arr, formatDealInfo(dealInfo))
	}

	return deals_arr, nil
}

// This function is used to format the deal information to be displayed in the terminal
func formatDealInfo(deal Deal) string {
	var result string

	// Deal ID
	if deal.ID != nil {
		result += fmt.Sprintf("Deal Number: %-10d", *deal.ID)
	}

	// Deal Title
	if deal.Title != nil {
		result += fmt.Sprintf(" Deal Title: %-15s", *deal.Title)
	}

	// Deal Value
	if deal.Value != nil {
		result += fmt.Sprintf(" Deal Value: %-10d", *deal.Value)
	}

	// Deal Status
	if deal.Status != nil {
		result += fmt.Sprintf(" Deal Status: %-15s", *deal.Status)
	}

	// Deal Add Time
	if deal.AddTime != nil {
		result += fmt.Sprintf(" Deal Add Time: %-15s", *deal.AddTime)
	}

	// Deal Expected Close Date
	if deal.ExpectedCloseDate != nil {
		result += fmt.Sprintf(" Deal Expected Close Date: %-15s", *deal.ExpectedCloseDate)
	}

	return result
}

// This function is used to validate the inputs provided by the user
func ValidateInputs(title string, value string, status string, date string, exp_close_date string) error {
	if title == "" {
		fmt.Println("title cannot be empty")
		return fmt.Errorf("title cannot be empty")
	} else {
		fmt.Println("Title validated:", title)
	}

	match, err := regexp.MatchString("^[0-9]*$", value)
	if match && err == nil {
		fmt.Println("Value validated:", match)
	} else if value == "" {
		fmt.Println("No value provided")
	} else {
		fmt.Println("Value error. Match:", match)
		return fmt.Errorf("value cannot be empty")
	}

	if status == "" {
		fmt.Println("No status provided")
	} else if status != "open" && status != "won" && status != "lost" && status != "deleted" {
		return fmt.Errorf("status cannot be empty or be invalid")
	} else {
		fmt.Println("Status validated:", status)
	}

	dateFormat := "2006-01-02T15:04"
	_, err = time.Parse(dateFormat, date)
	if err != nil {
		if date == "" {
			fmt.Println("No date provided")

		} else {
			fmt.Println("Error parsing date:", err)
			return err
		}
	} else {
		fmt.Println("Date validated:", date)
	}

	expDateFormat := "2006-01-02"
	_, err = time.Parse(expDateFormat, exp_close_date)
	if err != nil {
		if date == "" {
			fmt.Println("No expected close date provided")

		} else {
			fmt.Println("Error parsing expected close date:", err)
			return err
		}

	} else {
		fmt.Println("Expected close date validated:", exp_close_date)
	}

	return nil
}

// Main function to start the server
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/get-response", handleGetResponse)
	http.HandleFunc("/post-request", handlePostRequest)
	http.HandleFunc("/put-request", handlePutRequest)
	http.HandleFunc("/metrics", handleMetrics)

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
