package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	//"text/template"
	// Importing the general purpose Cosmos blockchain client
	// Importing the types package of your blog blockchain
	// Importing the general purpose Cosmos blockchain client

	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"ems/x/ems/types"
)

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/upcomingEvents", upcomingEventsHandler)
	http.HandleFunc("/register-event", registerEventHandler)
	http.HandleFunc("/get-event", getEventByIDHandler)
	http.HandleFunc("/purchase-ticket", PurchaseTicketHandler)
	http.HandleFunc("/list-tickets", ticketsHandler)
	http.HandleFunc("/resell-ticket", reSellTicketHandler)
	http.ListenAndServe(":8080", nil)
}

type PurchaseTicket struct {
	Amount    string `json:"amount"`
	EventID   string `json:"eventId"`
	Purchaser string `json:"purchaser"`
}

type ReSellTicket struct {
	Amount    string `json:"amount"`
	EventID   string `json:"eventId"`
	Purchaser string `json:"purchaser"`
	Id        string `json:"id"`
	Lender    string `json:"lender"`
}

// PageData holds the data to be passed to the template
type PageData struct {
	Title   string
	Content string
}

type EventId struct {
	EventID uint64 `json:"Id"`
}

// handler function to serve the HTML template
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Registration Application",
		Content: "This is a simple app developed for Hackathon.",
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Event struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Location     string `json:"location"`
	TicketsCount string `json:"ticketsCount"`
	Price        string `json:"price"`
	Organizer    string `json:"organizer"`
	Info         string `json:"info"`
}

type EventInfo struct {
	ID           uint64 `json:"id"`
	Name         string `json:"eventName"`
	Description  string `json:"eventDescription"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Location     string `json:"location"`
	TicketsCount uint64 `json:"numFtTickets"`
	TicketsLeft  uint64 `json:"ticketsLeft"`
	Price        uint64 `json:"ticketPrice:"`
	Organizer    string `json:"organizer"`
	Creator      string `json:"creator"`
}

func upcomingEventsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}
	queryClient := types.NewQueryClient(client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.ListEvent(ctx, &types.QueryListEventRequest{})
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResp)
}

func ticketsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}
	queryClient := types.NewQueryClient(client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.ListTicket(ctx, &types.QueryListTicketRequest{})
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResp)
}

func registerEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var event Event
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&event)

		ctx := context.Background()
		addressPrefix := "cosmos"

		// Create a Cosmos client instance
		client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
		if err != nil {
			log.Fatal(err)
		}

		// Account `alice` was initialized during `ignite chain serve`
		accountName := "alice"

		// Get account from the keyring
		account, err := client.Account(accountName)
		if err != nil {
			log.Fatal(err)
		}
		addr, err := account.Address(addressPrefix)
		if err != nil {
			log.Fatal(err)
		}
		tic, _ := strconv.ParseUint(event.TicketsCount, 10, 64)
		price, _ := strconv.ParseUint(event.Price, 10, 64)
		//tic, _ := strconv.ParseUint(event.TicketsCount, 10, 64)
		/*addr, err := account.Address(addressPrefix)
		if err != nil {
			log.Fatal(err)
		}*/
		// Define a message to create a post
		msg := &types.MsgCreateEvent{
			Creator:          addr,
			EventName:        event.Name,
			EventDescription: event.Description,
			StartDate:        event.StartDate,
			EndDate:          event.EndDate,
			Location:         event.Location,
			NumFtTickets:     tic,
			Organizer:        event.Organizer,
			TicketPrice:      price,
			TicketsLeft:      tic,
		}

		// Broadcast a transaction from account `alice` with the message
		// to create a post store response in txResp
		txResp, err := client.BroadcastTx(ctx, account, msg)
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process the event data (e.g., save to a database)
		// For now, let's just log it
		// Log the event data (for demonstration purposes)
		// In a real application, you might want to save this to a database
		log.Printf("Event registered: %+v\n", txResp)

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Event registered successfully!"})

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func PurchaseTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var purchase PurchaseTicket
	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	eventId, _ := strconv.ParseUint(purchase.EventID, 10, 64)
	//tic, _ := strconv.ParseUint(event.TicketsCount, 10, 64)
	/*addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}*/
	// Define a message to create a post
	msg := &types.MsgPurchaseTicket{
		Creator:   addr,
		Amount:    purchase.Amount + " token",
		EventId:   eventId,
		Purchaser: purchase.Purchaser,
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txResp)
}

func reSellTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var resellTicket ReSellTicket
	err := json.NewDecoder(r.Body).Decode(&resellTicket)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	eventId, _ := strconv.ParseUint(resellTicket.EventID, 10, 64)
	id, _ := strconv.ParseUint(resellTicket.Id, 10, 64)
	//tic, _ := strconv.ParseUint(event.TicketsCount, 10, 64)
	/*addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}*/
	// Define a message to create a post
	msg := &types.MsgResaleTicket{
		Creator:   addr,
		Amount:    resellTicket.Amount + " token",
		EventId:   eventId,
		Purchaser: resellTicket.Purchaser,
		Lender:    resellTicket.Lender,
		TicketId:  id,
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txResp)
}

// Handler to fetch event details by ID
func getEventByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var requestBody EventId
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)

		// Extract event ID from the URL
		ctx := context.Background()
		addressPrefix := "cosmos"

		// Create a Cosmos client instance
		client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
		if err != nil {
			log.Fatal(err)
		}
		queryClient := types.NewQueryClient(client.Context())
		//eventId, _ := strconv.ParseUint(requestBody.EventID, 10, 64)

		// Query the blockchain using the client's `PostAll` method
		// to get all posts store all posts in queryResp
		queryResp, err := queryClient.GetEvent(ctx, &types.QueryGetEventRequest{Id: requestBody.EventID})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Event registered: %+v\n", queryResp)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(queryResp)

	} //http.Error(w, "Event Not Found", http.StatusNotFound)
}

// serveTemplate function to render the HTML template
func serveTemplate(w http.ResponseWriter, templateName string, data PageData) {
	tmpl, err := template.ParseFiles(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
