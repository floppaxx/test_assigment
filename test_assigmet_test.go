package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// Testing the function HandleGetRequest
func TestHandleGetResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "/get-response", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handleGetResponse(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

}

// Testing the function HandlePostRequest
// Note that the testing creates a new deleted deal with the title "Test Deal"
func TestHandlePostRequest(t *testing.T) {
	form := url.Values{}
	form.Add("title", "Test Deal")
	form.Add("value", "0")
	form.Add("status", "deleted")
	form.Add("date", "")
	form.Add("exp_close_date", "")

	request, err := http.NewRequest("POST", "/post-request", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePostRequest)

	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "POST request was send"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}

}

// Testing the function HandlePutRequest
// Note that the id of the deal must exist (by default it is 2)
func TestHandlePutRequest(t *testing.T) {
	form := url.Values{}
	form.Add("id", "5")
	form.Add("title", "Test Deal")
	form.Add("value", "0")
	form.Add("status", "open")
	form.Add("date", "2023-07-26T19:11")
	form.Add("exp_close_date", "2023-08-26")

	request, err := http.NewRequest("PUT", "/put-request", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePutRequest)

	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "PUT request was send"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}

}

// Testing the function RetrieveDealsFromJSON
func TestRetrieveDealsFromJSON(t *testing.T) {
	jsonData := `{"success":true,"data":[{"id":1,"creator_user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"person_id":{"active_flag":true,"name":"Petro","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}],"owner_id":17918424,"value":1},"org_id":{"name":"comapny","people_count":1,"owner_id":17918424,"address":null,"active_flag":true,"cc_email":"John Doe-sandbox@pipedrivemail.com","owner_name":"John Doe","value":1},"stage_id":1,"title":"First ever deal","value":1000000,"currency":"USD","add_time":"2023-07-25 09:01:04","update_time":"2023-07-25 09:01:04","stage_change_time":null,"active":true,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":null,"pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"participants_count":1,"expected_close_date":"2023-07-28","last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"stage_order_nr":0,"person_name":"Petro","org_name":"comapny","next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$1,000,000","weighted_value":1000000,"formatted_weighted_value":"$1,000,000","weighted_value_currency":"USD","rotten_time":null,"owner_name":"John Doe","cc_email":"John Doe-sandbox+deal1@pipedrivemail.com","org_hidden":false,"person_hidden":false},{"id":2,"creator_user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"person_id":null,"org_id":null,"stage_id":1,"title":"jk","value":0,"currency":"USD","add_time":"2023-07-25 11:56:48","update_time":"2023-07-25 11:56:48","stage_change_time":null,"active":true,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":null,"pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"participants_count":0,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"stage_order_nr":0,"person_name":null,"org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$0","weighted_value":0,"formatted_weighted_value":"$0","weighted_value_currency":"USD","rotten_time":null,"owner_name":"John Doe","cc_email":"John Doe-sandbox+deal2@pipedrivemail.com","org_hidden":false,"person_hidden":false},{"id":3,"creator_user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"person_id":null,"org_id":null,"stage_id":1,"title":"testing post","value":250000,"currency":"USD","add_time":"2023-07-25 12:09:30","update_time":"2023-07-25 12:09:30","stage_change_time":null,"active":false,"deleted":false,"status":"won","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":"2023-07-25 12:09:30","pipeline_id":1,"won_time":"2023-07-25 12:09:30","first_won_time":"2023-07-25 12:09:30","lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"participants_count":0,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"stage_order_nr":0,"person_name":null,"org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$250,000","weighted_value":250000,"formatted_weighted_value":"$250,000","weighted_value_currency":"USD","rotten_time":null,"owner_name":"John Doe","cc_email":"John Doe-sandbox+deal3@pipedrivemail.com","org_hidden":false,"person_hidden":false},{"id":4,"creator_user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"person_id":null,"org_id":null,"stage_id":1,"title":"test","value":10000,"currency":"USD","add_time":"2023-07-26 15:53:00","update_time":"2023-07-26 12:54:02","stage_change_time":null,"active":false,"deleted":false,"status":"lost","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":"2023-07-26 12:54:02","pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":"2023-07-26 12:54:02","products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"participants_count":0,"expected_close_date":"2023-07-28","last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"stage_order_nr":0,"person_name":null,"org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$10,000","weighted_value":10000,"formatted_weighted_value":"$10,000","weighted_value_currency":"USD","rotten_time":null,"owner_name":"John Doe","cc_email":"John Doe-sandbox+deal4@pipedrivemail.com","org_hidden":false,"person_hidden":false},{"id":5,"creator_user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"user_id":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true,"value":17918424},"person_id":null,"org_id":null,"stage_id":1,"title":"wewe","value":0,"currency":"USD","add_time":"2023-07-25 12:15:41","update_time":"2023-07-25 15:47:12","stage_change_time":null,"active":false,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":"2023-07-25 15:47:12","pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"participants_count":0,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"stage_order_nr":0,"person_name":null,"org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$0","weighted_value":0,"formatted_weighted_value":"$0","weighted_value_currency":"USD","rotten_time":null,"owner_name":"John Doe","cc_email":"John Doe-sandbox+deal5@pipedrivemail.com","org_hidden":false,"person_hidden":false}],"additional_data":{"pagination":{"start":0,"limit":100,"more_items_in_collection":false}},"related_objects":{"user":{"17918424":{"id":17918424,"name":"John Doe","email":"example@example.com","has_pic":0,"pic_hash":null,"active_flag":true}},"organization":{"1":{"id":1,"name":"comapny","people_count":1,"owner_id":17918424,"address":null,"active_flag":true,"cc_email":"John Doe-sandbox@pipedrivemail.com","owner_name":"John Doe"}},"pipeline":{"1":{"id":1,"name":"Pipeline","url_title":"default","order_nr":1,"active":true,"deal_probability":false,"add_time":"2023-07-25 02:32:35","update_time":null}},"person":{"1":{"active_flag":true,"id":1,"name":"Petro","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}],"owner_id":17918424}},"stage":{"1":{"id":1,"company_id":12284304,"order_nr":0,"name":"Qualified","active_flag":true,"deal_probability":100,"pipeline_id":1,"rotten_flag":false,"rotten_days":null,"add_time":"2023-07-25 08:59:43","update_time":null,"pipeline_name":"Pipeline","pipeline_deal_probability":false}}}}`
	deals, err := retrieveDealsFromJSON(jsonData)
	if err != nil {
		t.Fatal(err)
	}

	if len(deals) != 5 {
		t.Errorf("Expected 5 deals, but got %d", len(deals))
	}
}

// Test the function ValidateInputs
func TestValidateInputs(t *testing.T) {
	// Test case 1: Valid input
	err := ValidateInputs("Test Deal", "100", "open", "2023-07-25T12:34", "2023-08-25")
	if err != nil {
		t.Errorf("Test case 1: expected no error, but got %v", err)
	}

	// Test case 2: Empty title
	err = ValidateInputs("", "100", "open", "2023-07-25T12:34", "2023-08-25")
	if err == nil {
		t.Errorf("Test case 2: expected error for empty title, but got no error")
	}

	// Test case 3: Invalid value (non-numeric)
	err = ValidateInputs("Test Deal", "abc", "open", "2023-07-25T12:34", "2023-08-25")
	if err == nil {
		t.Errorf("Test case 3: expected error for non-numeric value, but got no error")
	}

	// Test case 4: Invalid status
	err = ValidateInputs("Test Deal", "100", "invalid", "2023-07-25T12:34", "2023-08-25")
	if err == nil {
		t.Errorf("Test case 5: expected error for invalid status, but got no error")
	}

	// Test case 5: Invalid date format
	err = ValidateInputs("Test Deal", "100", "open", "2023-07-27T12:34:99", "2023-08-25")
	if err == nil {
		t.Errorf("Test case 6: expected error for invalid date format, but got no error")
	}

	// Test case 6: Invalid expected close date format
	err = ValidateInputs("Test Deal", "100", "open", "2023-07-25T12:34", "2023-08-32")
	if err == nil {
		t.Errorf("Test case 8: expected error for invalid expected close date format, but got no error")
	}
}

// Test the function WriteMetrics
func TestWriteMetrics(t *testing.T) {
	RequestMetrics = &Metrics{}

	// Test case 1: Single duration and latency
	duration1 := time.Second * 2
	latency1 := time.Millisecond * 500
	WriteMetrics(duration1, latency1)

	expectedAvgDuration1 := duration1
	expectedAvgLatency1 := latency1

	if RequestMetrics.TotalCount != 1 {
		t.Errorf("TotalCount: got %d, want %d", RequestMetrics.TotalCount, 1)
	}
	if RequestMetrics.TotalDuration != duration1 {
		t.Errorf("TotalDuration: got %s, want %s", RequestMetrics.TotalDuration, duration1)
	}
	if RequestMetrics.TotalLatency != latency1 {
		t.Errorf("TotalLatency: got %s, want %s", RequestMetrics.TotalLatency, latency1)
	}
	if RequestMetrics.AvgDuration != expectedAvgDuration1 {
		t.Errorf("AvgDuration: got %s, want %s", RequestMetrics.AvgDuration, expectedAvgDuration1)
	}
	if RequestMetrics.AvgLatency != expectedAvgLatency1 {
		t.Errorf("AvgLatency: got %s, want %s", RequestMetrics.AvgLatency, expectedAvgLatency1)
	}
	if RequestMetrics.MaxDuration != duration1 {
		t.Errorf("MaxDuration: got %s, want %s", RequestMetrics.MaxDuration, duration1)
	}
	if RequestMetrics.MinDuration != duration1 {
		t.Errorf("MinDuration: got %s, want %s", RequestMetrics.MinDuration, duration1)
	}
	if RequestMetrics.MaxLatency != latency1 {
		t.Errorf("MaxLatency: got %s, want %s", RequestMetrics.MaxLatency, latency1)
	}
	if RequestMetrics.MinLatency != latency1 {
		t.Errorf("MinLatency: got %s, want %s", RequestMetrics.MinLatency, latency1)
	}

	// Test case 2: Multiple durations and latencies
	duration2 := time.Second * 3
	latency2 := time.Millisecond * 800
	WriteMetrics(duration2, latency2)

	expectedAvgDuration2 := (duration1 + duration2) / 2
	expectedAvgLatency2 := (latency1 + latency2) / 2

	if RequestMetrics.TotalCount != 2 {
		t.Errorf("TotalCount: got %d, want %d", RequestMetrics.TotalCount, 2)
	}
	if RequestMetrics.TotalDuration != (duration1 + duration2) {
		t.Errorf("TotalDuration: got %s, want %s", RequestMetrics.TotalDuration, (duration1 + duration2))
	}
	if RequestMetrics.TotalLatency != (latency1 + latency2) {
		t.Errorf("TotalLatency: got %s, want %s", RequestMetrics.TotalLatency, (latency1 + latency2))
	}
	if RequestMetrics.AvgDuration != expectedAvgDuration2 {
		t.Errorf("AvgDuration: got %s, want %s", RequestMetrics.AvgDuration, expectedAvgDuration2)
	}
	if RequestMetrics.AvgLatency != expectedAvgLatency2 {
		t.Errorf("AvgLatency: got %s, want %s", RequestMetrics.AvgLatency, expectedAvgLatency2)
	}
	if RequestMetrics.MaxDuration != duration2 {
		t.Errorf("MaxDuration: got %s, want %s", RequestMetrics.MaxDuration, duration2)
	}
	if RequestMetrics.MinDuration != duration1 {
		t.Errorf("MinDuration: got %s, want %s", RequestMetrics.MinDuration, duration1)
	}
	if RequestMetrics.MaxLatency != latency2 {
		t.Errorf("MaxLatency: got %s, want %s", RequestMetrics.MaxLatency, latency2)
	}
	if RequestMetrics.MinLatency != latency1 {
		t.Errorf("MinLatency: got %s, want %s", RequestMetrics.MinLatency, latency1)
	}
}
