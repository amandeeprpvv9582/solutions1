#include <iostream>
#include <vector>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <chrono>
#include <queue>

std::queue<int> sharedQueue;
std::mutex mtx;
std::condition_variable cv;

void producer(int id, int numIterations) {
    for (int i = 0; i < numIterations; ++i) {
        std::this_thread::sleep_for(std::chrono::milliseconds(500));  // Simulate some work

        std::unique_lock<std::mutex> lock(mtx);
        sharedQueue.push(i);
        lock.unlock();

        cv.notify_all();  // Notify consumers that new data is available
    }
}

void consumer(int id) {
    while (true) {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, []{ return !sharedQueue.empty(); });  // Wait for new data

        int data = sharedQueue.front();
        sharedQueue.pop();
        lock.unlock();

        // Process the data
        std::cout << "Consumer " << id << " processed data: " << data << std::endl;

        std::this_thread::sleep_for(std::chrono::milliseconds(1000));  // Simulate some work
    }
}

int main() {
    const int numProducers = 2;
    const int numConsumers = 2;
    const int numIterations = 5;

    std::vector<std::thread> producers;
    std::vector<std::thread> consumers;

    for (int i = 0; i < numProducers; ++i) {
        producers.emplace_back(producer, i, numIterations);
    }

    for (int i = 0; i < numConsumers; ++i) {
        consumers.emplace_back(consumer, i);
    }

    for (auto& thread : producers) {
        thread.join();
    }

    cv.notify_all();  // Notify consumers to exit

    for (auto& thread : consumers) {
        thread.join();
    }

    return 0;
}
