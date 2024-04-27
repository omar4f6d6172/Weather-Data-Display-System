package main

import (
    "context"
    "fmt"
    "html/template"
    "net/http"

    "github.com/hectormalot/omgo"
)

type City struct {
    Name       string
    State     string
    Latitude   float64
    Longitude  float64
    Temperature float64
}


var (
    client omgo.Client // Changed from *omgo.Client to omgo.Client
    capitals = []City{
        {"Stuttgart", "Baden-Württemberg", 48.7758, 9.1829, 0},
        {"Berlin", "Berlin", 52.5200, 13.4050, 0},
        {"Hamburg", "Hamburg", 53.5511, 9.9937, 0},
        {"Munich", "Bavaria", 48.1351, 11.5820, 0},
        {"Cologne", "North Rhine-Westphalia", 50.9375, 6.9603, 0},
        {"Frankfurt", "Hesse", 50.1109, 8.6821, 0},
        {"Düsseldorf", "North Rhine-Westphalia", 51.2277, 6.7735, 0},
        {"Dortmund", "North Rhine-Westphalia", 51.5136, 7.4653, 0},
        
        // Add other cities here...
    }
    templates *template.Template
)

func main() {
    var err error
    client, err = omgo.NewClient() // This now correctly matches the variable type
    if err != nil {
        fmt.Printf("Error creating omgo client: %v\n", err)
        return
    }

    templates = template.Must(template.ParseFiles("index.html"))

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/city/", cityHandler)

    fmt.Println("Starting server at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
    if err := templates.ExecuteTemplate(w, "index.html", capitals); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func cityHandler(w http.ResponseWriter, r *http.Request) {
    cityID := r.URL.Path[len("/city/"):]
    for _, city := range capitals {
        if city.Name == cityID {
            loc, _ := omgo.NewLocation(city.Latitude, city.Longitude)
            res, _ := client.CurrentWeather(context.Background(), loc, nil)
            city.Temperature = res.Temperature
            fmt.Fprintf(w, `<div class="card w-96 bg-base-100 shadow-xl">
                                <div class="card-body">
                                    <h2 class="card-title">%s</h2>
                                    <p>Temperature: %.2f°C</p>
                                </div>
                            </div>`, city.Name, city.Temperature)
            return
        }
    }
    http.NotFound(w, r)
}
