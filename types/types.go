package types

import "time"

// User ...
type User struct {
	Username   string
	ChatID     int64
	Token      string
	Status     string
	Subscribed bool
}

// DVPNListResponse ...
type DVPNListResponse struct {
	Success bool      `json:"success"`
	Result  NodesList `json:"result"`
}

// List ...
type List struct {
	AccountAddr       string   `json:"account_addr"`
	IP                string   `json:"ip"`
	Latency           float64  `json:"latency"`
	VpnType           string   `json:"vpn_type"`
	Location          Location `json:"location"`
	NetSpeed          NetSpeed `json:"net_speed"`
	EncMethod         string   `json:"enc_method"`
	Version           string   `json:"version"`
	ActiveConnections int      `json:"active_connections"`
	Load              Load     `json:"load"`
	Rating            float64  `json:"rating,omitempty"`
	PricePerGB        float64  `json:"price_per_GB"`
	Moniker           string   `json:"moniker,omitempty"`
	Description       string   `json:"description,omitempty"`
}

type NodesList struct {
	nodes []Node `json:"nodes"`
}

// Location ...
type Location struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
}

// NetSpeed ...
type NetSpeed struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}

// Load ...
type Load struct {
	CPU    float64 `json:"cpu"`
	Memory int     `json:"memory"`
}

// NodeNetSpeed ...
type NodeNetSpeed struct {
	Download   float64    `json:"download"`
	BestServer BestServer `json:"best_server"`
	Upload     float64    `json:"upload"`
}

// BestServer ...
type BestServer struct {
	Latency float64 `json:"latency"`
	Host    string  `json:"host"`
}

// Bandwidth ... BW stats response
type Bandwidth struct {
	Success bool    `json:"success"`
	Units   string  `json:"units"`
	Stats   float64 `json:"stats"`
}

// TMExplorerResponce ...
type TMExplorerResponce struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      string                 `json:"id"`
	Result  map[string]interface{} `json:"result"`
}

type Node struct {
	Address string `json:"address"`
	Price   []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"price"`
	RemoteURL string    `json:"remote_url"`
	Status    int       `json:"status"`
	StatusAt  time.Time `json:"status_at"`
}
