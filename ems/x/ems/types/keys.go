package types

const (
	// ModuleName defines the module name
	ModuleName = "ems"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ems"

	MemCountKey = "cnt_ems"

	TicketStoreKey = "mem_tic_ems"

	TicketCountKey = "cnt_tic_ems"
)

var (
	ParamsKey = []byte("p_ems")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
