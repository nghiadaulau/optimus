package web

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
}

// EnableCORS is a middleware function to add CORS headers to the response.
func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins for simplicity, you may want to restrict this in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request, respond with 200 OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		next(w, r)
	}
}

// Serve starts http web server.
func Serve(setups OrgSetup) {
	http.HandleFunc("/query", EnableCORS(setups.Query))
	http.HandleFunc("/invoke", EnableCORS(setups.Invoke))
	http.HandleFunc("/delete", EnableCORS(setups.Delete))
	http.HandleFunc("/search", EnableCORS(setups.Search))
	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
