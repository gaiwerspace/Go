package models
import "time"
type Keeper struct { ID int64; Name string }
type Vault struct { ID int64; KeeperID int64; Balance float64; CreatedAt time.Time }
type Transfer struct { ID int64; SourceVaultID int64; DestinationVaultID int64; Amount float64; CreatedAt time.Time }
