package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am promi, I am a SW Engineer")
}

type Prodcut struct {
	ID          int
	Title       string
	Description string
	Price       float64
	Imgurl      string
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Plz give me GET request", 400)

	}

}
func main() {
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) //route
	mux.HandleFunc("/about", aboutHandler) //route
	mux.HandleFunc("/products", getProducts)
	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is sweet and juicy",
		Price:       100,
		Imgurl:      "https://images.othoba.com/images/thumbs/0040450_orange-1-kg.jpeg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Mango",
		Description: "Mango is juicy and delicious",
		Price:       300,
		Imgurl:      "https://dookan.com/cdn/shop/files/Dookan_Kesar_Mangoes_82bf0570-50bf-4f04-97b4-b1b415ebc733.png?v=1724714757&width=823",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "Red apple, crispy and sweet",
		Price:       200,
		Imgurl:      "https://upload.wikimedia.org/wikipedia/commons/1/15/Red_Apple.jpg",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Fresh yellow bananas, rich in potassium",
		Price:       50,
		Imgurl:      "https://upload.wikimedia.org/wikipedia/commons/8/8a/Banana-Single.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Grapes",
		Description: "Green seedless grapes, sweet and juicy",
		Price:       180,
		Imgurl:      "https://upload.wikimedia.org/wikipedia/commons/0/0b/Table_grapes_on_white.jpg",
	}
}
