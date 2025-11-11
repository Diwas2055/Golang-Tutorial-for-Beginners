# Go Programming Exercises & Projects

Hands-on exercises and project ideas to practice your Go skills.

---

## Beginner Exercises

### Exercise 1: Calculator
Create a simple calculator that performs basic operations.

**Requirements**:
- Accept two numbers and an operator (+, -, *, /)
- Perform the calculation
- Handle division by zero
- Loop to allow multiple calculations

**Starter Code**:
```go
package main

import "fmt"

func main() {
    // Your code here
}
```

**Solution Hints**:
- Use `fmt.Scan()` for input
- Use `switch` statement for operators
- Return error for division by zero

---

### Exercise 2: Temperature Converter
Convert between Celsius and Fahrenheit.

**Requirements**:
- Menu to choose conversion direction
- Input temperature value
- Display converted result
- Formula: F = C × 9/5 + 32

**Example Output**:
```
Temperature Converter
1. Celsius to Fahrenheit
2. Fahrenheit to Celsius
Choose option: 1
Enter temperature in Celsius: 25
25°C = 77°F
```

---

### Exercise 3: Grade Calculator
Calculate letter grade based on score.

**Requirements**:
- Input student name and score
- Calculate letter grade (A: 90-100, B: 80-89, C: 70-79, D: 60-69, F: <60)
- Store multiple students
- Display all students with grades

**Data Structure**:
```go
type Student struct {
    Name  string
    Score int
    Grade string
}
```

---

### Exercise 4: Word Counter
Count words in a sentence.

**Requirements**:
- Input a sentence
- Count total words
- Count unique words
- Display word frequency

**Example**:
```
Input: "hello world hello"
Total words: 3
Unique words: 2
Frequency: hello=2, world=1
```

---

### Exercise 5: Number Guessing Game
Create a number guessing game.

**Requirements**:
- Generate random number (1-100)
- User guesses the number
- Provide hints (too high/too low)
- Count number of attempts
- Option to play again

---

## Intermediate Exercises

### Exercise 6: Todo List
Build a command-line todo list application.

**Requirements**:
- Add tasks
- List all tasks
- Mark tasks as complete
- Delete tasks
- Save/load from file

**Commands**:
```
1. Add task
2. List tasks
3. Complete task
4. Delete task
5. Exit
```

**Data Structure**:
```go
type Task struct {
    ID        int
    Title     string
    Completed bool
    CreatedAt time.Time
}
```

---

### Exercise 7: Contact Manager
Create a contact management system.

**Requirements**:
- Add contacts (name, phone, email)
- Search contacts by name
- Update contact information
- Delete contacts
- List all contacts
- Validate email format

**Data Structure**:
```go
type Contact struct {
    ID    int
    Name  string
    Phone string
    Email string
}
```

---

### Exercise 8: Bank Account System
Simulate a simple banking system.

**Requirements**:
- Create account with initial balance
- Deposit money
- Withdraw money (check sufficient balance)
- Check balance
- Transaction history
- Multiple accounts

**Data Structure**:
```go
type Account struct {
    AccountNumber string
    HolderName    string
    Balance       float64
    Transactions  []Transaction
}

type Transaction struct {
    Type      string  // "deposit" or "withdraw"
    Amount    float64
    Timestamp time.Time
}
```

---

### Exercise 9: File Statistics
Analyze text files and provide statistics.

**Requirements**:
- Read text file
- Count lines, words, characters
- Find most common words
- Calculate average word length
- Display statistics

**Example Output**:
```
File: document.txt
Lines: 150
Words: 1,234
Characters: 6,789
Most common words: the (45), and (32), is (28)
Average word length: 5.5
```

---

### Exercise 10: URL Shortener
Create a simple URL shortener.

**Requirements**:
- Generate short code for long URL
- Store URL mappings
- Retrieve original URL from short code
- Track click count
- List all URLs

**Data Structure**:
```go
type URLMapping struct {
    ShortCode   string
    OriginalURL string
    Clicks      int
    CreatedAt   time.Time
}
```

---

## Advanced Exercises

### Exercise 11: Concurrent Web Scraper
Scrape multiple websites concurrently.

**Requirements**:
- Accept list of URLs
- Fetch content concurrently using goroutines
- Extract titles and links
- Handle errors gracefully
- Display results

**Key Concepts**:
- Goroutines
- Channels
- WaitGroups
- HTTP requests

---

### Exercise 12: Chat Server
Build a simple TCP chat server.

**Requirements**:
- Multiple clients can connect
- Broadcast messages to all clients
- Private messages between users
- User nicknames
- Handle disconnections

**Key Concepts**:
- TCP sockets
- Goroutines for each client
- Channels for message passing
- Concurrent map access

---

### Exercise 13: REST API
Create a RESTful API for a blog.

**Requirements**:
- CRUD operations for posts
- JSON responses
- Error handling
- Middleware for logging
- In-memory storage

**Endpoints**:
```
GET    /posts       - List all posts
GET    /posts/:id   - Get single post
POST   /posts       - Create post
PUT    /posts/:id   - Update post
DELETE /posts/:id   - Delete post
```

---

### Exercise 14: File Sync Tool
Synchronize files between directories.

**Requirements**:
- Compare two directories
- Identify new, modified, deleted files
- Copy new/modified files
- Option for dry-run
- Progress reporting

**Key Concepts**:
- File I/O
- Directory traversal
- File hashing (MD5/SHA256)
- Concurrent file operations

---

### Exercise 15: Log Analyzer
Analyze server log files.

**Requirements**:
- Parse log files (Apache/Nginx format)
- Extract IP addresses, timestamps, status codes
- Generate statistics (requests per hour, status code distribution)
- Find most accessed URLs
- Detect potential security issues

**Example Log Line**:
```
192.168.1.1 - - [01/Jan/2024:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1234
```

---

## Project Ideas

### Project 1: CLI Task Manager
**Description**: Full-featured command-line task management tool

**Features**:
- Create, read, update, delete tasks
- Categories and tags
- Due dates and priorities
- Search and filter
- Export to JSON/CSV
- Persistent storage (JSON file)

**Technologies**:
- Cobra (CLI framework)
- JSON encoding/decoding
- File I/O

---

### Project 2: Weather CLI
**Description**: Command-line weather application

**Features**:
- Fetch weather data from API
- Display current weather
- 5-day forecast
- Multiple locations
- Save favorite locations
- Temperature unit conversion

**Technologies**:
- HTTP requests
- JSON parsing
- External API integration

---

### Project 3: Markdown to HTML Converter
**Description**: Convert Markdown files to HTML

**Features**:
- Parse Markdown syntax
- Generate HTML output
- Support headers, lists, links, images
- Code blocks with syntax highlighting
- Custom CSS themes
- Batch conversion

**Technologies**:
- String parsing
- Regular expressions
- File I/O
- HTML generation

---

### Project 4: Password Manager
**Description**: Secure password storage and generation

**Features**:
- Generate strong passwords
- Store passwords encrypted
- Master password protection
- Search passwords by service
- Copy to clipboard
- Password strength checker

**Technologies**:
- Encryption (AES)
- File I/O
- Password hashing (bcrypt)

---

### Project 5: Image Resizer
**Description**: Batch image processing tool

**Features**:
- Resize images
- Convert formats (JPEG, PNG, GIF)
- Maintain aspect ratio
- Batch processing
- Progress bar
- Concurrent processing

**Technologies**:
- Image processing library
- Goroutines
- File I/O

---

### Project 6: JSON API Server
**Description**: RESTful API with database

**Features**:
- User authentication (JWT)
- CRUD operations
- Database integration (SQLite/PostgreSQL)
- Middleware (logging, auth)
- API documentation
- Rate limiting

**Technologies**:
- Gin/Echo framework
- Database/SQL
- JWT authentication

---

### Project 7: Web Crawler
**Description**: Crawl websites and extract data

**Features**:
- Crawl multiple pages
- Extract specific data (prices, titles, etc.)
- Respect robots.txt
- Rate limiting
- Export to CSV/JSON
- Concurrent crawling

**Technologies**:
- HTTP requests
- HTML parsing (goquery)
- Goroutines and channels

---

### Project 8: System Monitor
**Description**: Monitor system resources

**Features**:
- CPU usage
- Memory usage
- Disk space
- Network traffic
- Process list
- Real-time updates
- Alerts for thresholds

**Technologies**:
- System calls
- Goroutines for monitoring
- Terminal UI (termui)

---

### Project 9: File Encryption Tool
**Description**: Encrypt and decrypt files

**Features**:
- AES encryption
- Password-based encryption
- Encrypt/decrypt files
- Batch processing
- Secure key derivation
- Progress reporting

**Technologies**:
- Crypto libraries
- File I/O
- Command-line interface

---

### Project 10: Microservice Template
**Description**: Production-ready microservice template

**Features**:
- REST API endpoints
- Database integration
- Configuration management
- Logging
- Health checks
- Docker support
- Unit tests

**Technologies**:
- Web framework
- Database
- Docker
- Testing

---

## Challenge Problems

### Challenge 1: Implement a Cache
Create an LRU (Least Recently Used) cache.

**Requirements**:
- Set(key, value) with capacity limit
- Get(key) returns value
- Evict least recently used when full
- Thread-safe operations

---

### Challenge 2: Build a Rate Limiter
Implement a token bucket rate limiter.

**Requirements**:
- Allow N requests per time window
- Reject excess requests
- Thread-safe
- Multiple users

---

### Challenge 3: Create a Job Queue
Build a concurrent job processing system.

**Requirements**:
- Add jobs to queue
- Multiple workers process jobs
- Retry failed jobs
- Job priorities
- Status tracking

---

### Challenge 4: Implement Pub/Sub
Create a publish-subscribe system.

**Requirements**:
- Topics and subscriptions
- Multiple publishers
- Multiple subscribers
- Message delivery guarantees
- Thread-safe

---

### Challenge 5: Build a Load Balancer
Simple load balancer for HTTP requests.

**Requirements**:
- Multiple backend servers
- Round-robin distribution
- Health checks
- Retry failed requests
- Connection pooling

---

## Testing Exercises

### Test Exercise 1: Unit Tests
Write comprehensive unit tests for a calculator package.

**Requirements**:
- Test all operations
- Test edge cases
- Test error conditions
- Achieve >90% coverage

---

### Test Exercise 2: Table-Driven Tests
Create table-driven tests for string validation functions.

**Functions to Test**:
- Email validation
- Phone number validation
- Password strength
- URL validation

---

### Test Exercise 3: Benchmark Tests
Write benchmark tests for different sorting algorithms.

**Algorithms**:
- Bubble sort
- Quick sort
- Merge sort

**Compare**:
- Execution time
- Memory usage

---

## Learning Path

### Week 1-2: Beginner Exercises
Complete exercises 1-5 to solidify fundamentals.

### Week 3-4: Intermediate Exercises
Work through exercises 6-10 to build practical applications.

### Week 5-6: Advanced Exercises
Tackle exercises 11-15 to master concurrency and advanced topics.

### Week 7-8: Projects
Choose 2-3 projects to build complete applications.

### Week 9-10: Challenges
Solve challenge problems to deepen understanding.

---

## Tips for Success

1. **Start Small**: Begin with simple exercises and gradually increase complexity
2. **Write Tests**: Practice test-driven development
3. **Read Code**: Study solutions from others
4. **Refactor**: Improve your code after it works
5. **Document**: Write clear comments and documentation
6. **Version Control**: Use Git for all projects
7. **Ask Questions**: Join Go communities (Reddit, Discord, Stack Overflow)
8. **Build Portfolio**: Showcase projects on GitHub

---

## Resources for Help

- **Go Playground**: [play.golang.org](https://play.golang.org)
- **Go by Example**: [gobyexample.com](https://gobyexample.com)
- **Exercism**: [exercism.org/tracks/go](https://exercism.org/tracks/go)
- **LeetCode**: Practice algorithms in Go
- **GitHub**: Search for Go projects to learn from

---

**Remember**: The best way to learn is by doing. Pick an exercise and start coding! 🚀
