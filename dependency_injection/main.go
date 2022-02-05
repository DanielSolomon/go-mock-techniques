package di

import "time"

func main() {
	TestableCreateUser("foo", 123, time.Now)
}
