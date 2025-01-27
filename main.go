package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
    // "os"
    "strings"
)

type Person struct {
    Name     string  `json:"name"`
    Language string  `json:"language"`
    ID       string  `json:"id"`
    Bio      string  `json:"bio"`
    Version  float64 `json:"version"`
}

var dataFile = "data.json"

// Fetch all persons
func getPersons(c *gin.Context) {
    data, err := ioutil.ReadFile(dataFile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read data"})
        return
    }

    var persons map[string]Person
    if err := json.Unmarshal(data, &persons); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse data"})
        return
    }

    c.JSON(http.StatusOK, persons)
}

// Add a new person
func addPerson(c *gin.Context) {
    var newPerson Person
    if err := c.ShouldBindJSON(&newPerson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Read existing data
    data, err := ioutil.ReadFile(dataFile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read data"})
        return
    }

    var persons map[string]Person
    if err := json.Unmarshal(data, &persons); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse data"})
        return
    }

    // Create a key from the name
    key := strings.ReplaceAll(newPerson.Name, " ", "_")
    
    // Add new person to the map
    persons[key] = newPerson

    // Write updated data back to file
    updatedData, err := json.MarshalIndent(persons, "", "  ")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not marshal updated data"})
        return
    }

    if err := ioutil.WriteFile(dataFile, updatedData, 0644); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not write data"})
        return
    }

    c.JSON(http.StatusCreated, newPerson)
}

// Fetch person by ID
func fetchByID(c *gin.Context) {
    id := c.Param("id")

    // Read existing data
    data, err := ioutil.ReadFile(dataFile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read data"})
        return
    }

    var persons map[string]Person
    if err := json.Unmarshal(data, &persons); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse data"})
        return
    }

    // Search for the person by ID in all entries
    for _, person := range persons {
        if person.ID == id {
            c.JSON(http.StatusOK, person)
            return
        }
    }

    // If not found, return error message
    c.JSON(http.StatusNotFound, gin.H{"err": "an object with id not found"})
}

// Fetch person by name (using query parameter)
func fetchByName(c *gin.Context) {
    name := c.Query("name") // Get the name from query parameters

    // Read existing data
    data, err := ioutil.ReadFile(dataFile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read data"})
        return
    }

    var persons map[string]Person
    if err := json.Unmarshal(data, &persons); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse data"})
        return
    }

    // Check if the name exists in the map
    person, exists := persons[name]
    
	// If found, return the person object; otherwise return error message.
	if exists {
		c.JSON(http.StatusOK, person)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"err": "an object with name not found", "query": name})
	}
}

func main() {
  	router := gin.Default()

  	router.GET("/persons", getPersons)
  	router.POST("/persons", addPerson)
		router.GET("/fetchid/:id", fetchByID)
		router.GET("/fetchname", fetchByName)
    router.Run(":1651")
}

