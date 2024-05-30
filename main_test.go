package main

import (
	"testing"
    "time"
    "errors"
)

func TestCheck(t *testing.T) {
	//Arrange
	testTable := []struct {
		wait int
        p ping
		expected ping
	} {
		{
            wait: 1000,
            p: ping{url: "https://wikipedia.org"}, 
			expected: ping{
                url: "https://wikipedia.org",
                err1: nil,
                err2: nil,
                err3: nil,
                status: "200 OK",
                duration: time.Duration(1 * time.Second), 
            },
		},
        {
            wait: 1,
            p: ping{url: "https://wikipedia.org"}, 
			expected: ping{
                url: "https://wikipedia.org",
                err1: errors.New("any error"),
                err2: nil,
                err3: nil,
                status: "",
                duration: time.Duration(1 * time.Second),
            },
		},
        {
            wait: 1000,
            p: ping{url: "https://fake.fake"}, 
			expected: ping{
                url: "https://fake.fake",
                err1: errors.New("any error"),
                err2: nil,
                err3: nil,
                status: "",
                duration: time.Duration(1 * time.Second), 
            },
		},
        {
            wait: 1000,
            p: ping{url: "http://cnn.com"}, 
			expected: ping{
                url: "http://cnn.com",
                err1: nil,
                err2: errors.New("any error"),
                err3: errors.New("any error"),
                status: "200 OK",
                duration: time.Duration(1 * time.Second), 
            },
		},
        {
            wait: 1000,
            p: ping{url: "https://example.com/fake"}, 
			expected: ping{
                url: "https://example.com/fake",
                err1: nil,
                err2: nil,
                err3: nil,
                status: "500 Internal Server Error",
                duration: time.Duration(1 * time.Second), 
            },
		},
	}
	
	//Act
	for _, testCase := range testTable {
        Wait = testCase.wait
        testCase.p.check()
        result := testCase.p
        
        if result.url != testCase.expected.url {
			t.Errorf("Incorrect result. Expect url [%s], got [%s]", testCase.expected.url, result.url)
            continue
		}
		if (result.err1 != nil && testCase.expected.err1 == nil) || 
                (result.err1 == nil && testCase.expected.err1 != nil) {
			t.Errorf("Incorrect result. Expect err1 [%v], got [%v]", testCase.expected.err1, result.err1)
            continue
		}
		if (result.err2 != nil && testCase.expected.err2 == nil) || 
                (result.err2 == nil && testCase.expected.err2 != nil) {
			t.Errorf("Incorrect result. Expect err2 [%v], got [%v]", testCase.expected.err2, result.err2)
            continue
		}
		if (result.err3 != nil && testCase.expected.err3 == nil) || 
                (result.err3 == nil && testCase.expected.err3 != nil) {
			t.Errorf("Incorrect result. Expect err3 [%v], got [%v]", testCase.expected.err3, result.err3)
            continue
		}
        if result.status != testCase.expected.status {
			t.Errorf("Incorrect result. Expect status [%s], got [%s]", testCase.expected.status, result.status)
            continue
		}
        if result.duration > testCase.expected.duration {
			t.Errorf("Incorrect result. Expect duration [%v] > [%v]", testCase.expected.duration, result.duration)
            continue
		}
	}
}
