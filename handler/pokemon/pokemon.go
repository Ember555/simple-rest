package pokemon

import (
	"encoding/json"
	"net/http"
	"simple-rest/models"
	"simple-rest/repository"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control")
}

type Handler struct {
	pokemonRepos repository.PokemonRepository
	userRepos    repository.UserRepository
}

func NewHandler(pokemonRepos repository.PokemonRepository, userRepos repository.UserRepository) (*Handler, error) {
	return &Handler{pokemonRepos, userRepos}, nil
}

func (h *Handler) Handle(r *mux.Router) {
	r.HandleFunc("/search", h.SearchAll).Methods("GET")
	r.HandleFunc("/search/id/{id}", h.SearchByName).Methods("GET")
	r.HandleFunc("/create", h.Create).Methods("POST")
	r.HandleFunc("/update", h.Update).Methods("POST")
	r.HandleFunc("/delete/id/{id}", h.Delete).Methods("POST")

	r.HandleFunc("/signup", h.Signup).Methods("POST")
	r.HandleFunc("/signin", h.Signin).Methods("POST")
}

func Authen(w http.ResponseWriter, r *http.Request) error {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	return nil
}

func (h *Handler) SearchAll(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	err := Authen(w, r)
	if err != nil {
		http.Error(w, "Authen Fail", http.StatusUnauthorized)
		return
	}

	result, err := h.pokemonRepos.QueryAll()
	if err != nil {
		http.Error(w, "Search Fail", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func (h *Handler) SearchByName(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	err := Authen(w, r)
	if err != nil {
		http.Error(w, "Authen Fail", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	result, err := h.pokemonRepos.Query(params["id"])
	if err != nil {
		http.Error(w, "Search Fail", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	err := Authen(w, r)
	if err != nil {
		http.Error(w, "Authen Fail", http.StatusUnauthorized)
		return
	}

	var data models.PokemonModel
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = h.pokemonRepos.Insert(&data)
	if err != nil {
		http.Error(w, "Create Fail", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Create Success")
	return

	http.Error(w, "Not Found", http.StatusNotFound)
	return
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	err := Authen(w, r)
	if err != nil {
		http.Error(w, "Authen Fail", http.StatusUnauthorized)
		return
	}

	var data models.PokemonModel
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = h.pokemonRepos.Update(&data)
	if err != nil {
		http.Error(w, "Update Fail", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Update Success")
	return

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	err := Authen(w, r)
	if err != nil {
		http.Error(w, "Authen Fail", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	err = h.pokemonRepos.Delete(params["id"])
	if err != nil {
		http.Error(w, "Delete Fail", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode("Delete Success")
	return

}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)

	var data models.UserModel
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := h.userRepos.SignUp(&data)
	if err != nil {
		http.Error(w, "Sign up Fail", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Sign up Success")
	return

	http.Error(w, "Not Found", http.StatusNotFound)
	return
}

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	var creds models.UserModel

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Request Payload Invalid", http.StatusBadRequest)
		return
	}

	err = h.userRepos.Signin(creds.Username, creds.Password)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: "",
		})
		http.Error(w, "Log in Fail", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	json.NewEncoder(w).Encode("Log in Success")
}
