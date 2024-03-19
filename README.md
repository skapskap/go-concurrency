# Synchronous vs Asynchronous Fetching in Go 🚀

This repository contains a Go program that demonstrates the difference between synchronous and asynchronous fetching of data. It showcases how asynchronous fetching using goroutines and channels can lead to faster execution times compared to synchronous fetching. 🌐⏰

## Overview 📋

The program consists of two main parts:

1. **Synchronous Fetching**: Fetches mock user data (username, likes, and match) sequentially using synchronous function calls.
2. **Asynchronous Fetching**: Fetches mock user data concurrently using goroutines and channels, allowing for parallel execution.

The program measures the execution time for both approaches and displays the fetched data along with the time taken.

## Results 📊

Here are the typical results obtained from running the program:

![Synchronous vs Asynchronous Results](/assets/sync_async_results.png)

As you can see, asynchronous fetching generally takes less time compared to synchronous fetching, as it allows for concurrent execution of the data fetching tasks. 🚀💨

## Benchmarking 🏋️‍♂️

The repository also includes a `main_test.go` file that contains benchmark functions to measure the performance of synchronous and asynchronous fetching. To run the benchmarks, use the following command:

```bash
  go test -bench=.
```

The benchmark results will provide insights into the performance difference between the two approaches, including the number of iterations, average time per iteration, and memory allocations. 📈

![Benchmark Results](/assets/benchmark_results.png)

## How to Run 🏃‍♂️

To run the program, follow these steps:

1. Make sure you have Go installed on your system. 🐹
2. Clone this repository: `git clone https://github.com/skapskap/go-concurrency`. 📥
3. Navigate to the project directory: `cd go-concurrency`. 📂
4. Run the program: `go run main.go`. ▶️

You should see the fetched user data and the execution times for both synchronous and asynchronous fetching. ⏱️

## Contributions 🤝

Contributions to this repository are welcome! If you have any ideas, improvements, or bug fixes, feel free to open an issue or submit a pull request. Let's make this program even better together! 😄

## License 📄

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as per the terms of the license. 📜

---

Happy coding! 💻✨
