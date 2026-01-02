package graph
import "github.com/yourusername/erebor-vault/internal/services"
type Resolver struct { vaultService *services.VaultService }
func NewResolver(vs *services.VaultService) *Resolver { return &Resolver{vaultService: vs} }
