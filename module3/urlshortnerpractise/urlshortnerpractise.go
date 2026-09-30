package main

import (
	"fmt"
	"sync"
)

// safe access of map so no 2 grs access it same time.
func main() {
	urls := make(map[string]string)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			code := fmt.Sprintf("code%d", i)
			url := fmt.Sprintf("https://example%d.com", i)
			mu.Lock()
			urls[code] = url
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(urls)
}

// 1. go value / struct to json
// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }
// func main() {
// 	user := User{
// 		Name: "Shri Ram",
// 		Age:  22,
// 	}
// 	data, err := json.Marshal(user)// gives you []bytes
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 2. json to go value / struct
// 	data := []byte(`{"name":"Shri Ram","age":33}`)
// 	var user User
// 	err := json.Unmarshal(data, &user)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(user.Name)
// 	fmt.Println(user.Age)
// 	fmt.Println(string(data))

// 3. encoder - when sending go data to json stream (when sending response)
// type User struct {
// 	Name        string `json:"name"`
// 	Age         int    `json:"age"`
// 	Nationality string `json:"nationality"`
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	user := User{
// 		Name:        "Aniket",
// 		Age:         25,
// 		Nationality: "Indian",
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// func main() {
// 	http.HandleFunc("/user", userHandler)
// 	http.ListenAndServe(":8080", nil)
// }

//4. decoder(when receiving http request)
// type User struct {
// 	Name        string `json:"name"`
// 	Age         int    `json:"age"`
// 	Nationality string `json:"nationality"`
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Name:", user.Name)
// 	fmt.Println("Age:", user.Age)
// }

// func main() {
// 	http.HandleFunc("/user", userHandler)

// 	http.ListenAndServe(":8080", nil)
// }

//5. create random codes
// const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// func generateCode(length int) string {
// 	code := make([]byte, length)

// 	for i := range code {
// 		code[i] = characters[rand.IntN(len(characters))]
// 	}
// 	return string(code)
// }
// func main() {
// 	usedCodes := make(map[string]bool)

// 	for i := 0; i < 10; i++ {

// 		code := generateCode(6)

// 		if usedCodes[code] {
// 			fmt.Println("Duplicate:", code)
// 			continue
// 		}

// 		usedCodes[code] = true

// 		fmt.Println("Generated:", code)
// 	}
// }

//6 . redirect
// func googleHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "https://www.google.com", http.StatusFound)
// }

// func main() {
// 	http.HandleFunc("/google", googleHandler)
// 	http.ListenAndServe(":8080", nil)
// }

// 	http.HandleFunc("/hello", helloHandler)
// 	http.HandleFunc("/about", aboutHandler)
// 	http.HandleFunc("/greet", greetHandler)
// 	http.HandleFunc("/user/", userHandler)
// 	fmt.Println("Server running on port 8080")
// 	http.ListenAndServe(":8080", nil)
// func greetHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("surname")
// 	fmt.Fprintln(w, "Hello", name)
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path

// 	parts := strings.Split(path, "/")

// 	userID := parts[2]

// 	fmt.Fprintln(w, "User ID:", userID)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Method:", r.Method)
// 	fmt.Println("Path:", r.URL.Path)
// 	fmt.Fprintln(w, "Hello Aniket")
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Method:", r.Method)
// 	fmt.Println("Path:", r.URL.Path)
// 	fmt.Fprintln(w, "This is the about page")
// }
