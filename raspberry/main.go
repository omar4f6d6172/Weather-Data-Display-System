package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"

	"github.com/hectormalot/omgo"
)

type City struct {
	Name        string
	State       string
	Latitude    float64
	Longitude   float64
	Temperature float64
}

var (
	client   omgo.Client
	capitals = []City{
		{"Stuttgart", "Baden-Württemberg", 48.7758, 9.1829, 0},
        {"Berlin", "Berlin", 52.5200, 13.4050, 0},
        {"Hamburg", "Hamburg", 53.5511, 9.9937, 0},
        {"Munich", "Bavaria", 48.1351, 11.5820, 0},
        {"Cologne", "North Rhine-Westphalia", 50.9375, 6.9603, 0},
        {"Frankfurt", "Hesse", 50.1109, 8.6821, 0},
        {"Düsseldorf", "North Rhine-Westphalia", 51.2277, 6.7735, 0},
        {"Dortmund", "North Rhine-Westphalia", 51.5136, 7.4653, 0},
        {"Essen", "North Rhine-Westphalia", 51.4556, 7.0116, 0},
        {"Leipzig", "Saxony", 51.3397, 12.3731, 0},
        {"Bremen", "Bremen", 53.0793, 8.8017, 0},
        {"Dresden", "Saxony", 51.0504, 13.7373, 0},
        {"Hanover", "Lower Saxony", 52.3759, 9.7320, 0},
        {"Nuremberg", "Bavaria", 49.4521, 11.0768, 0},
        {"Duisburg", "North Rhine-Westphalia", 51.4344, 6.7623, 0},
        {"Bochum", "North Rhine-Westphalia", 51.4818, 7.2195, 0},
        {"Wuppertal", "North Rhine-Westphalia", 51.2562, 7.1507, 0},
        {"Bielefeld", "North Rhine-Westphalia", 52.0302, 8.5325, 0},
		// Add other cities...
	}
	templates *template.Template
	spiPort   spi.PortCloser
)

func main() {
	if _, err := host.Init(); err != nil {
		fmt.Printf("Failed to initialize hardware: %v\n", err)
		return
	}

	var err error
	spiPort, err = spireg.Open("/dev/spidev0.0")
	if err != nil {
		fmt.Printf("Error opening SPI port: %v\n", err)
		return
	}
	defer spiPort.Close()

	spiConn, err := spiPort.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		fmt.Printf("Failed to connect to SPI device: %v\n", err)
		return
	}

	client, err = omgo.NewClient()
	if err != nil {
		fmt.Printf("Error creating omgo client: %v\n", err)
		return
	}

	templates, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("Error parsing template files: %v\n", err)
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/city/", func(w http.ResponseWriter, r *http.Request) {
		cityHandler(w, r, spiConn) // Pass spiConn to the handler
	})

    fmt.Println("Now you can access the web interface at https://omar.svenjarene.team")
	fmt.Println("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", capitals); err != nil {
		http.Error(w, "Failed to execute template: "+err.Error(), http.StatusInternalServerError)
	}
}

func cityHandler(w http.ResponseWriter, r *http.Request, conn spi.Conn) {
	cityID := r.URL.Path[len("/city/"):]
	for _, city := range capitals {
		if city.Name == cityID {
			loc, err := omgo.NewLocation(city.Latitude, city.Longitude)
			if err != nil {
				http.Error(w, "Invalid location data", http.StatusBadRequest)
				return
			}
			res, err := client.CurrentWeather(context.Background(), loc, nil)
			if err != nil {
				http.Error(w, "Failed to fetch current weather: "+err.Error(), http.StatusInternalServerError)
				return
			}
			city.Temperature = res.Temperature
			if err := sendSPI(conn, city.Name, city.Temperature); err != nil {
				http.Error(w, "Failed to send data via SPI: "+err.Error(), http.StatusInternalServerError)
				return
			}
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

func sendSPI(conn spi.Conn, cityName string, temperature float64) error {
	tempStr := fmt.Sprintf("%.2f", temperature)
	data := []byte(cityName + ":" + tempStr)
	return conn.Tx(data, nil)
}
