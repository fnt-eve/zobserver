package observer

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestQueryRedisQ tests queryRedisq with a mock server.
func TestQueryRedisQ(t *testing.T) {
	t.Parallel() // marks this test as capable of running in parallel with other tests

	testCases := []struct {
		name             string       // name of the test case
		statusCode       int          // status code to return from the mock server
		err              error        // expected error from the function
		expectedResponse ZkilResponse // expected response from the function
	}{
		{
			name:             "Success",
			statusCode:       http.StatusOK,
			err:              nil,
			expectedResponse: ZkilResponse{Package: ZkilPackage{KillID: 123}}, // replace with your actual response structure
		},
		{
			name:             "TooManyRequests",
			statusCode:       http.StatusTooManyRequests,
			err:              errTooManyRequests,
			expectedResponse: ZkilResponse{}, // irrelevant for 429 status code
		},
		{
			name:             "InvalidURL",
			statusCode:       0, // will cause net/http to fail when calling Do
			err:              errors.New("Get \"\": unsupported protocol scheme \"\""),
			expectedResponse: ZkilResponse{}, // irrelevant because we expect an error here
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create a mock server that returns the desired status code and body
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tc.statusCode)
				if tc.statusCode == http.StatusOK {
					_, _ = w.Write([]byte(`{"Package": {"KillID": 123}}`)) // replace with your actual response JSON
				}
			}))
			defer mockServer.Close()

			// override RedisQURL so it points to our mock server
			oldRedisQURL := RedisQURL
			RedisQURL = mockServer.URL
			defer func() { RedisQURL = oldRedisQURL }() // restore original URL after test

			// call queryRedisq
			resp, err := queryRedisq("testQueue", "0")

			if tc.err == nil {
				expectNoError(t, err, "Failed to fetch from RedisQ API")
			}

			if err == nil { // only need to compare response if no error occurs
				expectEqualResponses(t, *resp, tc.expectedResponse, "Unexpected Response from RedisQ API")
			}
		})
	}
}

func TestGenRand(t *testing.T) {
	testCases := []int{5, 10, 20, 30} // Different lengths to test

	for _, n := range testCases {
		t.Run(fmt.Sprintf("GeneratesRandomStringOfLength%d", n), func(t *testing.T) {
			got, err := GenRand(n)
			if err != nil {
				t.Errorf("Expected no errors, got %v", err)
			}

			if len(*got) > base64.StdEncoding.EncodedLen(n) {
				t.Errorf("Expected length <= %d, got %d", base64.StdEncoding.EncodedLen(n), len(*got))
			}
		})
	}
}

func expectNoError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("%s: Expected no error but got: %v", msg, err)
	}
}

func expectEqualResponses(t *testing.T, resp ZkilResponse, expected ZkilResponse, msg string) {
	if resp != expected {
		t.Errorf("%s: Expected response %v, got %v", msg, resp, expected)
	}
}
