package odoo

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
	"time"

	"github.com/ybbus/jsonrpc/v2"
)

// Client is ...
type Client struct {
	url       string
	db        string
	username  string
	password  string
	uid       int
	rpcClient jsonrpc.RPCClient
}

// Dial creates a client for an RPC endpoint.
func Dial(url string) *Client {
	c := Client{
		url:       url,
		rpcClient: jsonrpc.NewClient(url),
	}
	return &c
}

// Authenticate connects to the endpoint with username and password.
func (c *Client) Authenticate(db string, username string, password string) error {
	response, err := c.rpcClient.Call("call", &Request{
		Service: "common",
		Method:  "login",
		Args:    []interface{}{db, username, password},
	})

	if err != nil {
		return err
	}

	c.db = db
	c.username = username
	c.password = password

	//if n, ok := response.Result.(json.Number); !ok {
	if uid, err := response.GetInt(); err != nil {
		return errors.New("error getting UID")
	} else {
		c.uid = int(uid)
	}
	return nil
}

type Timesheet struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	CreateDate  string       `json:"create_date"`
	User        IDLabelField `json:"user_id"`
	Tier        IDLabelField `json:"tier_id"`
	UnitAmount  float32      `json:"unit_amount"`
	DisplayName string       `json:"display_name"`
	Project     IDLabelField `json:"project_id"`
	//Amount      float32      `json:"amount"`
	TokenAmount float32      `json:"token_amount"`
	TaskID      IDLabelField `json:"task_id"`
	Date        string       `json:"date"`
	Tokenized   bool         `json:"tokenized"`
}

type TokenizedLine struct {
	Tokenized bool `json:"tokenized"`
}

// GetTimesheets returns all timesheets
func (c *Client) GetTimesheets(i Interval) ([]Timesheet, error) {
	q := c.NewSearch(
		"account.analytic.line",
		NewFields(StructToTags(Timesheet{})...),
		NewWhere("end", ">=", i.Start).
			And("end", "<", i.End).
			And("task_stage_id", "ilike", "approved").
			And("tokenized", "=", false),
	)
	var r []Timesheet
	err := c.rpcClient.CallFor(&r, "call", q.ToRequest())
	if err != nil {
		return nil, err
	}
	return r, nil
}

// User represents a user in the Odoo domain.
type User struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Email           String  `json:"email"`
	EthereumAddress Address `json:"ethereum_address"`
}

// GetUsers returns all users in Odoo.
func (c *Client) GetUsers() ([]User, error) {
	q := c.NewSearch(
		"res.users",
		NewFields(StructToTags(User{})...),
		NewWhere("active", "=", true),
	)
	var r []User
	err := c.rpcClient.CallFor(&r, "call", q.ToRequest())
	if err != nil {
		return nil, err
	}
	return r, nil
}

// TokenAllocation holds information about the token allocation for a user.
type TokenAllocation struct {
	User        `json:"user"`
	HoursAmount float32     `json:"hours_amount"`
	TokenAmount float32     `json:"token_amount"`
	Timesheets  []Timesheet `json:"timesheets"`
}

// CalculateTokenAllocations takes users and timesheets and calculates the new
// token allocation.
func CalculateTokenAllocations(u []User, t []Timesheet) []TokenAllocation {
	userMap := make(map[int]*User)
	for i, user := range u {
		userMap[user.ID] = &u[i]
	}
	allocationMap := make(map[int]*TokenAllocation)
	for _, timesheet := range t {
		a := allocationMap[timesheet.User.ID]
		if a == nil {
			a = &TokenAllocation{}
			allocationMap[timesheet.User.ID] = a
		}
		a.User = *userMap[timesheet.User.ID]
		a.Timesheets = append(a.Timesheets, timesheet)
		a.HoursAmount += timesheet.UnitAmount
		a.TokenAmount += timesheet.TokenAmount
	}
	v := make([]TokenAllocation, len(allocationMap))
	i := 0
	for _, value := range allocationMap {
		v[i] = *value
		i++
	}
	return v
}

type Rewards struct {
	Interval         `json:"interval"`
	TokenAllocations []TokenAllocation `json:"token_allocations"`
}

func (c *Client) GetRewards() (*Rewards, error) {
	i := NewInterval(time.Now(), MonthlySchedule)
	p := i.Decrement()
	users, err := c.GetUsers()
	if err != nil {
		return nil, err
	}

	timesheets, err := c.GetTimesheets(p)
	if err != nil {
		return nil, err
	}

	tokenAllocations := CalculateTokenAllocations(users, timesheets)

	rewards := &Rewards{
		Interval:         p,
		TokenAllocations: tokenAllocations,
	}

	return rewards, nil
}

func (c *Client) SetTimesheetTokenization(id []int, b bool) error {
	r := c.NewWriteRequest("account.analytic.line", id, TokenizedLine{b})

	if response, err := c.rpcClient.Call("call", r.ToRequest()); err != nil {
		return err
	} else if ok, err := response.GetBool(); !ok || err != nil {
		return err
	}
	return nil
}

func (p *Rewards) PrintTable() {
	fmt.Printf("Rewards from %s to %s\n", p.Start.Format("Jan 2, 2006"), p.End.Format("Jan 2, 2006"))
	for _, a := range p.TokenAllocations {
		fmt.Printf("%s <%s>, %s will receive %.2f Tokens\n", a.Name, a.Email, a.EthereumAddress, a.TokenAmount)
		for _, t := range a.Timesheets {
			fmt.Printf("%d %.2fh\t%s\n", t.ID, t.UnitAmount, t.DisplayName)
		}
		fmt.Println()
	}
}

type Resolution struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type RewardsResolution struct {
	Rewards    `json:"rewards"`
	Resolution Resolution `json:"resolution"`
}

func NewRewardsResolution(p *Rewards) *RewardsResolution {
	title := fmt.Sprintf("Rewarding Contributors for %s", p.Interval.Start.Format("January 2006"))
	//t, err := template.ParseFiles("cli/templates/resolution-payments.md")
	t, err := template.New("content").Parse(`1. Object the confirmation of the contributed time by Contributors to teledisko DAO from {{.Interval.Start.Format "02.01.2006"}} to {{.Interval.GetEndInclude.Format "02.01.2006"}} or the minting of the corresponding number of tokens to them in the following manner:{{range $a := .TokenAllocations}}
    1. {{$a.Name}}, {{printf "%.4f" $a.HoursAmount}} hours, {{printf "%.2f" $a.TokenAmount}} tokens;{{end}}`)
	if err != nil {
		panic(err)
	}
	b := bytes.Buffer{}
	err = t.Execute(&b, p)
	if err != nil {
		panic(err)
	}
	return &RewardsResolution{
		Rewards:    *p,
		Resolution: Resolution{title, b.String()},
	}
}
