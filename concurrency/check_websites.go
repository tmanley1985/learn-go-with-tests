package concurrency

type WebsiteChecker func(string) bool

// Notice the key value pairs don't have a name?
// This is nice when we don't know exactly what we'd name a field.
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Two points here:
	// 1.) We use an IFFE for the goroutine and pass in any state. Why?
	// Well, remember, the loop is defining the url variable and the closure captures it
	// so if the url variable is reassigned, the closures reference to it will change to!
	//2.) We have to avoid concurrent writes to the map. Maps REALLY don't like it when you do that.
	// The answer here is to use CHANNELS. Just like in CSP (communicating sequential processes).
	for _, url := range urls {
		go func(u string) {

			// The SEND statement has a channel on the left and a value on the right.
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// The RECEIVE statement has a channel on the right and a value on the left.
		// This action is BLOCKING. It's almost like saying: yield, or await.
		r := <-resultChannel
		// Once we have a value, THEN we can put them onto the map.
		results[r.string] = r.bool
	}

	return results
}