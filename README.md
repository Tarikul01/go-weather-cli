## 🌤️ Go Weather CLI  

A simple and lightweight CLI tool built with Golang to fetch and display real-time weather data using WeatherAPI.  

---

### 🚀 Features  
✅ Fetches real-time weather data  
✅ Displays temperature, wind speed, humidity, and condition  
✅ Simple and easy-to-use CLI tool  

---

### 📝 Installation  

#### 1⃣ Clone the Repository  
```sh
git clone https://github.com/Tarikul01/go-weather-cli.git
cd go-weather-cli
```

#### 2⃣ Build the Project  
```sh
go build -o sun
```

#### 3⃣ Run the CLI Tool  
```sh
./sun
```

---

### 🛠 Configuration  

To use this tool, you need an API key from **WeatherAPI**.  

1. Sign up at [WeatherAPI](https://www.weatherapi.com/) and get your API key.  
2. Replace the API key in your `main.go` file:  
   ```go
   res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=YOUR_API_KEY&q=Dhaka")
   ```
3. Rebuild the project:  
   ```sh
   go build -o sun
   ```

---

### 🛠 Usage  

Run the CLI tool with a city name:  
```sh
./sun Dhaka
```
Sample Output:  
```
Weather in Dhaka, Bangladesh:
Temperature: 25.0°C (Feels like 26.5°C)
Condition: Clear
Wind Speed: 13.0 km/h
Humidity: 68%
```

---

### 🌟 License  
This project is licensed under the **MIT License**.  

---

### 🤝 Contributing  
Feel free to fork this repository and submit a pull request! Contributions are always welcome.  

---


