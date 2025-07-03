###  **Use Case: Worker Pool Pattern**

You want to process a list of jobs (say, resizing images or scraping URLs) using a fixed number of worker goroutines.

---

###  **Why use a channel here?**

* Channels are used to send jobs to workers (goroutines).
* Channels let you **safely share data** between goroutines without race conditions.
* Closing the channel signals all workers: "No more work."

---

###  **What's going on?**

* You spin up 3 workers (`go worker(...)`).
* Each worker blocks on `jobs <-chan int`. When a job arrives, it processes it.
* After sending 5 jobs, we `close(jobs)` so workers exit their loop when the channel is drained.
* The `results` channel collects the output from each worker.

---

###  **Why is this useful?**

* This structure **limits concurrency** — only 3 jobs will be worked on at the same time.
* Channels make **safe communication** between goroutines effortless.
* You **don’t need mutexes** — channels handle the synchronization.
* **Backpressure** is natural. If all workers are busy, the main goroutine blocks on `jobs <-`.
