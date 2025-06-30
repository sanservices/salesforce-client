package salesforceclient

import (
	"context"
	"log"
	"os"

	Grpc "github.com/sanservices/salesforce-client/grpc"
	"gopkg.in/yaml.v2"
)

const (
	fPath = "settings.yaml"
)

type Settings struct {
	Sfmc Grpc.SettingsSfmc `yaml:"SfmcConfig"`
}

type Person struct {
	First_Name string
	Last_Name  string
	Email      string
	ID         string
}

func (c *Person) ToValue() interface{} {
	return map[string]interface{}{
		"FIRST_NAME":    c.First_Name,
		"LAST_NAME":     c.Last_Name,
		"PERSON_ID":     c.ID,
		"EmailAddress":  c.Email,
		"SubscriberKey": c.Email,
	}
}

func Test1() {
	s, err := New(context.Background())
	if err != nil {
		log.Fatalf("ERR: %v", err)
	}

	c, err := Grpc.SfmcNewClient(&s.Sfmc)

	if err != nil {
		log.Fatalf("ERR: %v", err)
	}
	data := []Person{
		{
			First_Name: "Ana",
			Last_Name:  "Mejia",
			Email:      "gabriela.diaz@sanservices.hn",
			ID:         "0502200005098",
		},
	}
	// Convert []Person to []interface{}
	dataInterface := make([]interface{}, len(data))
	for i, v := range data {
		dataInterface[i] = v.ToValue()
	}
	r, err := c.InsertDataRows(context.Background(), dataInterface, "beaches", "PersonTesting")
	if err != nil {
		log.Fatalf("ERR: %v", err)
	}
	log.Print(r.RequestId)

}

func New(ctx context.Context) (*Settings, error) {
	settings := &Settings{}

	//Read settings file
	cf, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(cf, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
