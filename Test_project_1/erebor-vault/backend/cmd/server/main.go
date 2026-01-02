package main
import ("log"; "net/http"; "os"; "github.com/99designs/gqlgen/graphql/handler"; "github.com/99designs/gqlgen/graphql/playground"; "github.com/yourusername/erebor-vault/graph"; "github.com/yourusername/erebor-vault/internal/database"; "github.com/yourusername/erebor-vault/internal/services")
func main() {
	port := "8080"
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" { dbURL = "postgres://postgres@localhost:5432/erebor?sslmode=disable" }
	db, err := database.NewDB(dbURL)
	if err != nil { log.Fatal(err) }
	defer db.Close()
	if err := database.RunMigrations(db); err != nil { log.Fatal(err) }
	svc := services.NewVaultService(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(svc)}))
	cors := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" { w.WriteHeader(200); return }
			h.ServeHTTP(w, r)
		})
	}
	http.Handle("/", playground.Handler("Erebor", "/query"))
	http.Handle("/query", cors(srv))
	log.Printf("🏔️ Server: http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
