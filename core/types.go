package core

// TODO: create enums

type ListenArgs struct {
	Protocol string // tcp | upd | empty - tcp and upd
	IpLayer  string // v4 | v6 - empty to display both
	HttpOnly bool   // Display only http and https requests
}
