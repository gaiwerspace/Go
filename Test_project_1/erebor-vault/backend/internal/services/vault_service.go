package services
import ("database/sql"; "errors"; "github.com/yourusername/erebor-vault/internal/models")
var (ErrInsufficientFunds = errors.New("insufficient funds"); ErrVaultNotFound = errors.New("vault not found"); ErrKeeperNotFound = errors.New("keeper not found"); ErrInvalidAmount = errors.New("invalid amount"); ErrSameVault = errors.New("same vault"))
type VaultService struct { db *sql.DB }
func NewVaultService(db *sql.DB) *VaultService { return &VaultService{db: db} }
func (s *VaultService) CreateVault(keeperID int64, deposit float64) (*models.Vault, error) {
	if deposit < 0 { return nil, ErrInvalidAmount }
	var exists bool
	s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM keepers WHERE id=$1)", keeperID).Scan(&exists)
	if !exists { return nil, ErrKeeperNotFound }
	var v models.Vault
	err := s.db.QueryRow("INSERT INTO vaults (keeper_id, balance, created_at) VALUES ($1,$2,NOW()) RETURNING id,keeper_id,balance,created_at", keeperID, deposit).Scan(&v.ID, &v.KeeperID, &v.Balance, &v.CreatedAt)
	return &v, err
}
func (s *VaultService) TransferGold(srcID, dstID int64, amt float64) (*models.Transfer, error) {
	if amt <= 0 { return nil, ErrInvalidAmount }
	if srcID == dstID { return nil, ErrSameVault }
	tx, _ := s.db.Begin()
	defer tx.Rollback()
	var srcBal float64
	err := tx.QueryRow("SELECT balance FROM vaults WHERE id=$1 FOR UPDATE", srcID).Scan(&srcBal)
	if err == sql.ErrNoRows { return nil, ErrVaultNotFound }
	if srcBal < amt { return nil, ErrInsufficientFunds }
	var dstBal float64
	tx.QueryRow("SELECT balance FROM vaults WHERE id=$1 FOR UPDATE", dstID).Scan(&dstBal)
	tx.Exec("UPDATE vaults SET balance=balance-$1 WHERE id=$2", amt, srcID)
	tx.Exec("UPDATE vaults SET balance=balance+$1 WHERE id=$2", amt, dstID)
	var t models.Transfer
	tx.QueryRow("INSERT INTO transfers (source_vault_id,destination_vault_id,amount,created_at) VALUES ($1,$2,$3,NOW()) RETURNING id,source_vault_id,destination_vault_id,amount,created_at", srcID, dstID, amt).Scan(&t.ID, &t.SourceVaultID, &t.DestinationVaultID, &t.Amount, &t.CreatedAt)
	tx.Commit()
	return &t, nil
}
func (s *VaultService) GetVault(id int64) (*models.Vault, error) {
	var v models.Vault
	err := s.db.QueryRow("SELECT id,keeper_id,balance,created_at FROM vaults WHERE id=$1", id).Scan(&v.ID, &v.KeeperID, &v.Balance, &v.CreatedAt)
	if err == sql.ErrNoRows { return nil, ErrVaultNotFound }
	return &v, err
}
func (s *VaultService) GetVaultHistory(vaultID int64) ([]*models.Transfer, error) {
	rows, _ := s.db.Query("SELECT id,source_vault_id,destination_vault_id,amount,created_at FROM transfers WHERE source_vault_id=$1 OR destination_vault_id=$1 ORDER BY created_at DESC", vaultID)
	defer rows.Close()
	var ts []*models.Transfer
	for rows.Next() {
		var t models.Transfer
		rows.Scan(&t.ID, &t.SourceVaultID, &t.DestinationVaultID, &t.Amount, &t.CreatedAt)
		ts = append(ts, &t)
	}
	return ts, nil
}
func (s *VaultService) GetKeepers() ([]*models.Keeper, error) {
	rows, _ := s.db.Query("SELECT id,name FROM keepers ORDER BY id")
	defer rows.Close()
	var ks []*models.Keeper
	for rows.Next() { var k models.Keeper; rows.Scan(&k.ID, &k.Name); ks = append(ks, &k) }
	return ks, nil
}
func (s *VaultService) GetKeeper(id int64) (*models.Keeper, error) {
	var k models.Keeper
	err := s.db.QueryRow("SELECT id,name FROM keepers WHERE id=$1", id).Scan(&k.ID, &k.Name)
	if err == sql.ErrNoRows { return nil, ErrKeeperNotFound }
	return &k, err
}
func (s *VaultService) GetVaultsByKeeper(keeperID int64) ([]*models.Vault, error) {
	rows, _ := s.db.Query("SELECT id,keeper_id,balance,created_at FROM vaults WHERE keeper_id=$1 ORDER BY created_at DESC", keeperID)
	defer rows.Close()
	var vs []*models.Vault
	for rows.Next() { var v models.Vault; rows.Scan(&v.ID, &v.KeeperID, &v.Balance, &v.CreatedAt); vs = append(vs, &v) }
	return vs, nil
}
func (s *VaultService) GetTotalBalanceByKeeper(keeperID int64) (float64, error) {
	var total float64
	s.db.QueryRow("SELECT COALESCE(SUM(balance),0) FROM vaults WHERE keeper_id=$1", keeperID).Scan(&total)
	return total, nil
}
